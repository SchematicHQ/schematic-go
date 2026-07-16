package rulesengine_test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestEngine(t *testing.T, opts ...rulesengine.EngineOption) *rulesengine.Engine {
	t.Helper()
	ctx := context.Background()
	e, err := rulesengine.NewEngine(ctx, opts...)
	require.NoError(t, err)
	t.Cleanup(func() { _ = e.Close(context.Background()) })
	return e
}

func testFlag(key string, defaultValue bool) *rulesengine.Flag {
	return &rulesengine.Flag{
		ID:            "flag-" + key,
		AccountID:     "acct-1",
		EnvironmentID: "env-1",
		Key:           key,
		DefaultValue:  defaultValue,
	}
}

func testCompany(id string) *rulesengine.Company {
	return &rulesengine.Company{
		ID:            id,
		AccountID:     "acct-1",
		EnvironmentID: "env-1",
	}
}

func TestNewEngine(t *testing.T) {
	e := newTestEngine(t)
	assert.NotEmpty(t, e.VersionKey())
	assert.NotEqual(t, "1", e.VersionKey(), "should read the real version key, not the fallback")
}

func TestCloseIsIdempotent(t *testing.T) {
	ctx := context.Background()
	e, err := rulesengine.NewEngine(ctx)
	require.NoError(t, err)
	require.NoError(t, e.Close(ctx))
	require.NoError(t, e.Close(ctx))
}

func TestCheckFlagDefaultValue(t *testing.T) {
	e := newTestEngine(t)
	ctx := context.Background()

	for _, defaultValue := range []bool{true, false} {
		res, err := e.CheckFlag(ctx, nil, nil, testFlag("my-flag", defaultValue))
		require.NoError(t, err)
		assert.Equal(t, defaultValue, res.Value)
		assert.Equal(t, rulesengine.ReasonNoRulesMatched, res.Reason)
	}
}

// TestCheckFlagPopulatesIdentifiers is the regression test for the camelCase
// mismatch: the engine emits flagKey/flagId/companyId while the generated types
// are tagged snake_case. Decoding naively leaves these empty while value and
// reason still look correct.
func TestCheckFlagPopulatesIdentifiers(t *testing.T) {
	e := newTestEngine(t)
	ctx := context.Background()

	res, err := e.CheckFlag(ctx, testCompany("comp-1"), nil, testFlag("my-flag", true))
	require.NoError(t, err)

	assert.Equal(t, "my-flag", res.FlagKey, "flagKey must survive the camelCase boundary")
	require.NotNil(t, res.FlagID)
	assert.Equal(t, "flag-my-flag", *res.FlagID)
	require.NotNil(t, res.CompanyID)
	assert.Equal(t, "comp-1", *res.CompanyID)
}

// TestCheckFlagWithNilSlices is the regression test for null-stripping. The
// generated types have no omitempty on their slice fields, so a nil Rules
// marshals to `"rules": null`, which the engine rejects outright.
func TestCheckFlagWithNilSlices(t *testing.T) {
	e := newTestEngine(t)
	ctx := context.Background()

	company := testCompany("comp-1") // Rules, Traits, Metrics all nil
	user := &rulesengine.User{ID: "user-1", AccountID: "acct-1", EnvironmentID: "env-1"}
	flag := testFlag("my-flag", true) // Rules nil

	res, err := e.CheckFlag(ctx, company, user, flag)
	require.NoError(t, err, "nil slices must not be sent to the engine as JSON null")
	assert.True(t, res.Value)
}

func TestCheckFlagNilFlag(t *testing.T) {
	e := newTestEngine(t)

	res, err := e.CheckFlag(context.Background(), nil, nil, nil)
	require.NoError(t, err)
	assert.False(t, res.Value)
	assert.Equal(t, rulesengine.ReasonFlagNotFound, res.Reason)
	assert.ErrorIs(t, res.Err, rulesengine.ErrorFlagNotFound)
}

func TestCheckFlagNegativePreflight(t *testing.T) {
	e := newTestEngine(t)
	ctx := context.Background()
	flag := testFlag("my-flag", true)

	t.Run("negative usage", func(t *testing.T) {
		_, err := e.CheckFlag(ctx, nil, nil, flag, rulesengine.WithUsage(-1))
		assert.ErrorIs(t, err, rulesengine.ErrorNegativePreflightUsage)
	})

	t.Run("negative event usage", func(t *testing.T) {
		_, err := e.CheckFlag(ctx, nil, nil, flag, rulesengine.WithEventUsage("sub", -1))
		assert.ErrorIs(t, err, rulesengine.ErrorNegativePreflightUsage)
	})

	t.Run("negative credit cost", func(t *testing.T) {
		_, err := e.CheckFlag(ctx, nil, nil, flag, rulesengine.WithCreditCost("credit-1", -1))
		assert.ErrorIs(t, err, rulesengine.ErrorNegativePreflightCreditCost)
	})

	t.Run("zero usage is accepted", func(t *testing.T) {
		_, err := e.CheckFlag(ctx, nil, nil, flag, rulesengine.WithUsage(0))
		assert.NoError(t, err)
	})
}

// TestCheckFlagConcurrent exercises the pool. The module uses shared linear
// memory, so an instance serving two checks at once would corrupt both. Every
// result carries its own flag key, which makes cross-talk visible.
func TestCheckFlagConcurrent(t *testing.T) {
	e := newTestEngine(t, rulesengine.WithPoolSize(4))
	ctx := context.Background()

	const goroutines = 32
	const iterations = 25

	var wg sync.WaitGroup
	errs := make(chan error, goroutines*iterations)

	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			key := fmt.Sprintf("flag-%d", g)
			want := g%2 == 0
			for i := 0; i < iterations; i++ {
				res, err := e.CheckFlag(ctx, testCompany(fmt.Sprintf("comp-%d", g)), nil, testFlag(key, want))
				if err != nil {
					errs <- err
					return
				}
				if res.FlagKey != key {
					errs <- fmt.Errorf("cross-talk: got flag key %q, want %q", res.FlagKey, key)
					return
				}
				if res.Value != want {
					errs <- fmt.Errorf("cross-talk: got value %v for %q, want %v", res.Value, key, want)
					return
				}
				if res.CompanyID == nil || *res.CompanyID != fmt.Sprintf("comp-%d", g) {
					errs <- fmt.Errorf("cross-talk: wrong company on %q", key)
					return
				}
			}
		}(g)
	}

	wg.Wait()
	close(errs)
	for err := range errs {
		t.Fatal(err)
	}
}

func TestPoolSizeOne(t *testing.T) {
	e := newTestEngine(t, rulesengine.WithPoolSize(1))
	ctx := context.Background()

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := e.CheckFlag(ctx, nil, nil, testFlag("my-flag", true))
			assert.NoError(t, err)
			assert.True(t, res.Value)
		}()
	}
	wg.Wait()
}
