// Package rulesengine evaluates Schematic feature flags locally.
//
// Evaluation runs the shared Schematic rules engine, a WebAssembly module built
// from the Rust implementation in schematic-api and used byte for byte by the
// C#, Python, Java, and Ruby SDKs. Sharing one engine means rule semantics
// cannot drift between languages.
//
// The module is hosted by wazero, a pure-Go WebAssembly runtime. wazero needs
// no cgo, so cross-compilation and static binaries keep working -- which is why
// it is used here rather than wasmtime, the runtime the other SDKs embed.
package rulesengine

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"runtime"
	"sync"
	"time"

	"github.com/schematichq/schematic-go/rulesengine/wasm"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

// defaultVersionKey matches the fallback the other SDKs use when the module
// predates the version-key export.
const defaultVersionKey = "1"

// Engine evaluates flags against the WebAssembly rules engine. It is safe for
// concurrent use. Create one with NewEngine and reuse it; construction compiles
// the module, which is far more expensive than a check.
type Engine struct {
	runtime    wazero.Runtime
	pool       chan *instance
	versionKey string
	closeOnce  sync.Once
	closeErr   error
}

// EngineOption configures an Engine.
type EngineOption func(*engineConfig)

type engineConfig struct {
	poolSize int
}

// WithPoolSize sets how many module instances the engine keeps. Each instance
// serves one concurrent check and owns its linear memory, so this caps
// concurrent evaluation. Defaults to GOMAXPROCS. Values below 1 are ignored.
func WithPoolSize(n int) EngineOption {
	return func(c *engineConfig) {
		if n > 0 {
			c.poolSize = n
		}
	}
}

// instance is a single module instance plus its resolved exports. The engine's
// linear memory is shared across every call into one instance, so an instance
// must only ever serve one check at a time; the pool enforces that by handing
// each out to a single caller.
type instance struct {
	mod     api.Module
	mem     api.Memory
	alloc   api.Function
	dealloc api.Function
	check   api.Function
	getRes  api.Function
	getLen  api.Function
	setTime api.Function // optional; absent on modules predating the export

	// Metric-period boundary exports. Optional for the same reason as setTime:
	// a module predating them still serves checks, callers just get nil.
	curPeriodCalendar   api.Function
	curPeriodSubscript  api.Function
	nextPeriodCalendar  api.Function
	nextPeriodSubscript api.Function
}

// NewEngine compiles the rules engine and prepares it for evaluation.
func NewEngine(ctx context.Context, opts ...EngineOption) (*Engine, error) {
	cfg := &engineConfig{poolSize: runtime.GOMAXPROCS(0)}
	for _, opt := range opts {
		opt(cfg)
	}
	if cfg.poolSize < 1 {
		cfg.poolSize = 1
	}

	rt := wazero.NewRuntime(ctx)

	// The module barely uses WASI -- it reaches for fd_write only when the Rust
	// side hits an eprintln! error path -- but it imports the interface, so the
	// host must provide it or instantiation fails.
	if _, err := wasi_snapshot_preview1.Instantiate(ctx, rt); err != nil {
		_ = rt.Close(ctx)
		return nil, fmt.Errorf("instantiate wasi: %w", err)
	}

	// Compile once, instantiate many. Compilation is the expensive step; each
	// instance is comparatively cheap and gets its own linear memory.
	compiled, err := rt.CompileModule(ctx, wasm.Binary)
	if err != nil {
		_ = rt.Close(ctx)
		return nil, fmt.Errorf("compile rules engine: %w", err)
	}

	e := &Engine{
		runtime:    rt,
		pool:       make(chan *instance, cfg.poolSize),
		versionKey: defaultVersionKey,
	}

	for i := 0; i < cfg.poolSize; i++ {
		inst, err := newInstance(ctx, rt, compiled)
		if err != nil {
			_ = rt.Close(ctx)
			return nil, err
		}
		if i == 0 {
			// Read the version key from a pristine instance. The export returns
			// a pointer into linear memory that later calls will overwrite, so
			// it has to be copied out before this instance serves any check.
			e.versionKey = readVersionKey(ctx, inst)
		}
		e.pool <- inst
	}

	return e, nil
}

func newInstance(ctx context.Context, rt wazero.Runtime, compiled wazero.CompiledModule) (*instance, error) {
	// An empty name instantiates anonymously, which is what lets one compiled
	// module back many instances -- named modules must be unique per runtime.
	// Stderr is discarded: the module writes to it only on Rust error paths,
	// which are already surfaced through the -1 return and the result's err.
	cfg := wazero.NewModuleConfig().WithName("").WithStderr(io.Discard).WithStdout(io.Discard)

	mod, err := rt.InstantiateModule(ctx, compiled, cfg)
	if err != nil {
		return nil, fmt.Errorf("instantiate rules engine: %w", err)
	}

	inst := &instance{
		mod:     mod,
		mem:     mod.Memory(),
		alloc:   mod.ExportedFunction("alloc"),
		dealloc: mod.ExportedFunction("dealloc"),
		check:   mod.ExportedFunction("checkFlagCombined"),
		getRes:  mod.ExportedFunction("getResultJson"),
		getLen:  mod.ExportedFunction("getResultJsonLength"),
		setTime: mod.ExportedFunction("setCurrentTimeMillis"),

		curPeriodCalendar:   mod.ExportedFunction("getCurrentMetricPeriodStartForCalendarMetricPeriod"),
		curPeriodSubscript:  mod.ExportedFunction("getCurrentMetricPeriodStartForCompanyBillingSubscription"),
		nextPeriodCalendar:  mod.ExportedFunction("getNextMetricPeriodStartForCalendarMetricPeriod"),
		nextPeriodSubscript: mod.ExportedFunction("getNextMetricPeriodStartForCompanyBillingSubscription"),
	}

	if inst.mem == nil {
		return nil, fmt.Errorf("rules engine exports no memory")
	}
	for name, fn := range map[string]api.Function{
		"alloc":               inst.alloc,
		"dealloc":             inst.dealloc,
		"checkFlagCombined":   inst.check,
		"getResultJson":       inst.getRes,
		"getResultJsonLength": inst.getLen,
	} {
		if fn == nil {
			return nil, fmt.Errorf("rules engine missing required export %q", name)
		}
	}

	return inst, nil
}

// readVersionKey reads the NUL-terminated version key, falling back to
// defaultVersionKey when the export is absent, as the other SDKs do.
func readVersionKey(ctx context.Context, inst *instance) string {
	fn := inst.mod.ExportedFunction("get_version_key_wasm")
	if fn == nil {
		return defaultVersionKey
	}
	res, err := fn.Call(ctx)
	if err != nil || len(res) == 0 {
		return defaultVersionKey
	}

	ptr := api.DecodeU32(res[0])
	var out []byte
	for i := ptr; ; i++ {
		b, ok := inst.mem.ReadByte(i)
		if !ok || b == 0 {
			break
		}
		out = append(out, b)
	}
	if len(out) == 0 {
		return defaultVersionKey
	}
	return string(out)
}

// VersionKey identifies the shape of the engine's models. It changes whenever
// those models change, which makes it suitable for namespacing caches of
// engine data so a model change invalidates stale entries.
func (e *Engine) VersionKey() string {
	return e.versionKey
}

// Close releases the engine's resources. It is safe to call more than once.
func (e *Engine) Close(ctx context.Context) error {
	e.closeOnce.Do(func() {
		e.closeErr = e.runtime.Close(ctx)
	})
	return e.closeErr
}

// CheckFlag evaluates flag against the given company and user, either of which
// may be nil. Preflight options simulate pending usage without mutating state.
func (e *Engine) CheckFlag(
	ctx context.Context,
	company *Company,
	user *User,
	flag *Flag,
	opts ...CheckFlagOption,
) (*CheckFlagResult, error) {
	options := newCheckFlagOptions()
	for _, opt := range opts {
		opt(options)
	}

	if err := options.validate(); err != nil {
		return &CheckFlagResult{Reason: ReasonNoRulesMatched, Err: err}, err
	}

	// Handled here rather than in the engine, whose checkFlagCombined requires a
	// flag object and reports a missing one only by its empty id.
	if flag == nil {
		return &CheckFlagResult{Reason: ReasonFlagNotFound, Err: ErrorFlagNotFound}, nil
	}

	env := &checkFlagEnvelope{Flag: flag, Company: company, User: user}
	if !options.isZero() {
		env.Options = options
	}

	input, err := marshalEnvelope(env)
	if err != nil {
		return nil, err
	}

	return e.CheckFlagJSON(ctx, input)
}

// CheckFlagJSON evaluates a pre-marshaled envelope.
//
// It exists for callers that already hold the wire types in their own packages
// -- the Schematic API owns these models and publishes them as the OpenAPI
// Rulesengine* schemas -- and so cannot pass this package's structs to
// CheckFlag. Marshaling their own envelope and calling this avoids a
// per-check conversion between two field-identical mirrors.
//
// The envelope is the object CheckFlag builds: {"flag":..., "company":...,
// "user":..., "options":...}, where absent members are omitted. Collection
// fields must serialize as [] / {} rather than null; see marshalEnvelope.
func (e *Engine) CheckFlagJSON(ctx context.Context, envelope []byte) (*CheckFlagResult, error) {
	// Wait for a free instance, but honor cancellation: if every instance is
	// busy and ctx is already done, return rather than block indefinitely.
	var inst *instance
	select {
	case inst = <-e.pool:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	defer func() { e.pool <- inst }()

	out, err := inst.checkFlag(ctx, envelope)
	if err != nil {
		return nil, err
	}

	var result CheckFlagResult
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, fmt.Errorf("decode rules engine result: %w", err)
	}

	return &result, nil
}

// checkFlag runs one evaluation. The caller must hold exclusive use of inst.
func (i *instance) checkFlag(ctx context.Context, input []byte) ([]byte, error) {
	res, err := i.alloc.Call(ctx, uint64(len(input)))
	if err != nil {
		return nil, fmt.Errorf("rules engine alloc: %w", err)
	}
	ptr := api.DecodeU32(res[0])

	// The input buffer is ours to free; the result buffer is not (it lives in a
	// thread-local owned by the module).
	defer func() { _, _ = i.dealloc.Call(ctx, uint64(ptr), uint64(len(input))) }()

	if !i.mem.Write(ptr, input) {
		return nil, fmt.Errorf("rules engine: writing %d bytes at %d exceeds memory", len(input), ptr)
	}

	// The raw wasm32-unknown-unknown build has no system clock. Without this the
	// engine still evaluates correctly, but metric-period reset timestamps are
	// silently omitted from results.
	if i.setTime != nil {
		if _, err := i.setTime.Call(ctx, api.EncodeI64(time.Now().UnixMilli())); err != nil {
			return nil, fmt.Errorf("rules engine setCurrentTimeMillis: %w", err)
		}
	}

	checked, err := i.check.Call(ctx, uint64(ptr), uint64(len(input)))
	if err != nil {
		return nil, fmt.Errorf("rules engine checkFlagCombined: %w", err)
	}
	if api.DecodeI32(checked[0]) < 0 {
		return nil, ErrorUnexpected
	}

	resPtr, err := i.getRes.Call(ctx)
	if err != nil {
		return nil, fmt.Errorf("rules engine getResultJson: %w", err)
	}
	resLen, err := i.getLen.Call(ctx)
	if err != nil {
		return nil, fmt.Errorf("rules engine getResultJsonLength: %w", err)
	}

	view, ok := i.mem.Read(api.DecodeU32(resPtr[0]), api.DecodeU32(resLen[0]))
	if !ok {
		return nil, fmt.Errorf("rules engine: result read out of range")
	}

	// Read hands back a view into linear memory, which the next call overwrites.
	out := make([]byte, len(view))
	copy(out, view)
	return out, nil
}
