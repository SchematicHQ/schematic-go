package rulesengine_test

import (
	"context"
	"fmt"
	"runtime"
	"testing"

	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/require"
)

// BenchmarkCheckFlagParallel measures how evaluation scales with pool size.
// The other SDKs serialize every check behind a single lock because the module
// uses shared linear memory; this exists to show what that would cost Go, which
// is replacing a lock-free in-process engine.
func BenchmarkCheckFlagParallel(b *testing.B) {
	for _, size := range []int{1, 2, 4, runtime.GOMAXPROCS(0)} {
		b.Run(fmt.Sprintf("pool=%d", size), func(b *testing.B) {
			ctx := context.Background()
			e, err := rulesengine.NewEngine(ctx, rulesengine.WithPoolSize(size))
			require.NoError(b, err)
			defer func() { _ = e.Close(ctx) }()

			flag := testFlag("my-flag", true)
			company := testCompany("comp-1")

			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if _, err := e.CheckFlag(ctx, company, nil, flag); err != nil {
						b.Fatal(err)
					}
				}
			})
		})
	}
}

// BenchmarkCheckFlagSerial measures single-threaded per-check cost.
func BenchmarkCheckFlagSerial(b *testing.B) {
	ctx := context.Background()
	e, err := rulesengine.NewEngine(ctx, rulesengine.WithPoolSize(1))
	require.NoError(b, err)
	defer func() { _ = e.Close(ctx) }()

	flag := testFlag("my-flag", true)
	company := testCompany("comp-1")

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := e.CheckFlag(ctx, company, nil, flag); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkNewEngine measures construction, which compiles the module and is
// the reason an Engine should be created once and reused.
func BenchmarkNewEngine(b *testing.B) {
	ctx := context.Background()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		e, err := rulesengine.NewEngine(ctx, rulesengine.WithPoolSize(1))
		require.NoError(b, err)
		_ = e.Close(ctx)
	}
}
