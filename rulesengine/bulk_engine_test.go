package rulesengine_test

import (
	"context"
	"testing"

	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/require"
)

// End-to-end checks that the bulk envelope this package builds is one the
// engine accepts. The serialization invariants are pinned as text in
// bulk_envelope_test.go; these run the payload through wasm, which is the only
// place a wire mistake actually shows up.

func newBulkTestEngine(t *testing.T) *rulesengine.Engine {
	t.Helper()

	ctx := context.Background()
	e, err := rulesengine.NewEngine(ctx, rulesengine.WithPoolSize(1))
	require.NoError(t, err)
	t.Cleanup(func() { _ = e.Close(ctx) })

	return e
}

// TestBulkMatchesIndividual guards the benchmark: the two paths it compares
// must answer identically, or the comparison is measuring different work.
func TestBulkMatchesIndividual(t *testing.T) {
	ctx := context.Background()
	e, err := rulesengine.NewEngine(ctx, rulesengine.WithPoolSize(1))
	require.NoError(t, err)
	defer func() { _ = e.Close(ctx) }()

	flags := benchFlags(10)
	company := benchCompany(flags, 5)

	bulk, err := e.CheckFlags(ctx, company, nil, flags)
	require.NoError(t, err)
	require.Len(t, bulk, len(flags))

	for i, flag := range flags {
		single, err := e.CheckFlag(ctx, company, nil, flag)
		require.NoError(t, err)

		require.Equal(t, single.Value, bulk[i].Value, "flag %s", flag.Key)
		require.Equal(t, single.FlagKey, bulk[i].FlagKey)
		require.Equal(t, single.Reason, bulk[i].Reason)
	}
}

// TestCheckFlagsNilCollectionsRoundTrip is the worst case the wire types are
// meant to absorb: a company and flags whose every collection is left nil. If
// any of them reached the engine as null, serde would reject the envelope and
// the whole batch would fail rather than answer.
func TestCheckFlagsNilCollectionsRoundTrip(t *testing.T) {
	ctx := context.Background()
	e := newBulkTestEngine(t)

	flags := []*rulesengine.Flag{
		{ID: "flag-1", AccountID: "a", EnvironmentID: "e", Key: "k1", DefaultValue: true},
		{ID: "flag-2", AccountID: "a", EnvironmentID: "e", Key: "k2", DefaultValue: false},
	}
	company := &rulesengine.Company{ID: "comp-1", AccountID: "a", EnvironmentID: "e"}

	results, err := e.CheckFlags(ctx, company, nil, flags)
	require.NoError(t, err)
	require.Len(t, results, len(flags))

	for i, flag := range flags {
		require.Nil(t, results[i].Err, "flag %s", flag.Key)
		require.Equal(t, flag.Key, results[i].FlagKey)
		require.Equal(t, flag.DefaultValue, results[i].Value, "a flag with no rules reports its default")
	}
}

// TestCheckFlagsNilUserRoundTrip covers the other optional member: a nil user
// must be omitted from the envelope rather than sent as null.
func TestCheckFlagsNilUserRoundTrip(t *testing.T) {
	ctx := context.Background()
	e := newBulkTestEngine(t)

	flags := benchFlags(3)
	company := benchCompany(flags, 2)

	results, err := e.CheckFlags(ctx, company, nil, flags)
	require.NoError(t, err)
	require.Len(t, results, len(flags))
	for _, result := range results {
		require.Nil(t, result.Err)
	}
}

// TestCheckFlagsNilEntry pins that a nil flag never reaches the engine -- the
// engine would reject the envelope outright -- and that its result still lands
// in the position the caller passed it.
//
// The result is the one CheckFlag gives for a nil flag: {Reason: "Flag not
// found", Err: ErrorFlagNotFound} and a nil error, with FlagKey, FlagID and
// Value left at their zero values, so a caller assembling flags from a lookup
// that can miss gets a result per position and a batch that still succeeds.
func TestCheckFlagsNilEntry(t *testing.T) {
	ctx := context.Background()
	e := newBulkTestEngine(t)

	flags := []*rulesengine.Flag{
		{ID: "flag-1", AccountID: "a", EnvironmentID: "e", Key: "k1", DefaultValue: true},
		nil,
		{ID: "flag-2", AccountID: "a", EnvironmentID: "e", Key: "k2", DefaultValue: true},
	}

	results, err := e.CheckFlags(ctx, nil, nil, flags)
	require.NoError(t, err, "a nil flag is a result, not a failed batch")
	require.Len(t, results, 3)

	require.Equal(t, "k1", results[0].FlagKey)
	require.Equal(t, "k2", results[2].FlagKey)

	notFound := results[1]
	require.Equal(t, rulesengine.ReasonFlagNotFound, notFound.Reason)
	require.Equal(t, rulesengine.ErrorFlagNotFound, notFound.Err)
	require.False(t, notFound.Value, "no flag, so no default to report")
	require.Empty(t, notFound.FlagKey)
	require.Nil(t, notFound.FlagID)
	require.Nil(t, notFound.Entitlement)
}

// TestCheckFlagsNilEntryMatchesCheckFlag ties the above to the single-flag
// path, which is the behaviour the pre-Rust engine set and schematic-go's
// CheckFlag still carries. Bulk must not become a second opinion on it.
func TestCheckFlagsNilEntryMatchesCheckFlag(t *testing.T) {
	ctx := context.Background()
	e := newBulkTestEngine(t)

	single, err := e.CheckFlag(ctx, nil, nil, nil)
	require.NoError(t, err)

	bulk, err := e.CheckFlags(ctx, nil, nil, []*rulesengine.Flag{nil})
	require.NoError(t, err)
	require.Len(t, bulk, 1)

	require.Equal(t, single, bulk[0])
}

// TestCheckFlagsAllNilEntries covers the envelope never being built at all:
// every flag nil means no engine call, and a result per position regardless.
func TestCheckFlagsAllNilEntries(t *testing.T) {
	ctx := context.Background()
	e := newBulkTestEngine(t)

	results, err := e.CheckFlags(ctx, nil, nil, []*rulesengine.Flag{nil, nil})
	require.NoError(t, err)
	require.Len(t, results, 2)
	for _, result := range results {
		require.Equal(t, rulesengine.ReasonFlagNotFound, result.Reason)
	}
}

// TestCheckFlagsEmpty pins the empty case: no engine call, an empty result.
func TestCheckFlagsEmpty(t *testing.T) {
	ctx := context.Background()
	e := newBulkTestEngine(t)

	results, err := e.CheckFlags(ctx, nil, nil, nil)
	require.NoError(t, err)
	require.Empty(t, results)
}
