package cache

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// Benchmark for Get operation
func BenchmarkLocalCache_Get(b *testing.B) {
	ctx := context.Background()
	cache := NewLocalCache[string](1000, 0) // No TTL to avoid cleanup overhead

	// Pre-populate cache with some data
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key-%d", i)
		cache.Set(ctx, key, fmt.Sprintf("value-%d", i), nil)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Get random key (80% hits, 20% misses)
			var key string
			if rand.Float32() < 0.8 {
				// Hit
				key = fmt.Sprintf("key-%d", rand.Intn(1000))
			} else {
				// Miss
				key = fmt.Sprintf("missing-%d", rand.Intn(1000))
			}
			cache.Get(ctx, key)
		}
	})
}

// Benchmark for Set operation
func BenchmarkLocalCache_Set(b *testing.B) {
	ctx := context.Background()
	cache := NewLocalCache[string](1000, 0) // No TTL to avoid cleanup overhead

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := fmt.Sprintf("key-%d-%d", i, rand.Intn(1000))
			cache.Set(ctx, key, fmt.Sprintf("value-%d", i), nil)
			i++
		}
	})
}

// Benchmark for Delete operation
func BenchmarkLocalCache_Delete(b *testing.B) {
	ctx := context.Background()
	cache := NewLocalCache[string](1000, 0) // No TTL to avoid cleanup overhead

	// Pre-populate cache with some data
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key-%d", i)
		cache.Set(ctx, key, fmt.Sprintf("value-%d", i), nil)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Delete random key (80% hits, 20% misses)
			var key string
			if rand.Float32() < 0.8 {
				// Hit
				key = fmt.Sprintf("key-%d", rand.Intn(1000))
			} else {
				// Miss
				key = fmt.Sprintf("missing-%d", rand.Intn(1000))
			}
			cache.Delete(ctx, key)
		}
	})
}

// Benchmark for a mixed workload
func BenchmarkLocalCache_MixedWorkload(b *testing.B) {
	ctx := context.Background()
	cache := NewLocalCache[string](1000, 0) // No TTL to avoid cleanup overhead

	// Pre-populate cache with some data
	for i := 0; i < 500; i++ {
		key := fmt.Sprintf("key-%d", i)
		cache.Set(ctx, key, fmt.Sprintf("value-%d", i), nil)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			// Mix of operations: 70% Get, 20% Set, 10% Delete
			op := rand.Float32()

			var key string
			if op < 0.7 {
				// Get
				if rand.Float32() < 0.8 {
					// Hit
					key = fmt.Sprintf("key-%d", rand.Intn(500))
				} else {
					// Miss
					key = fmt.Sprintf("missing-%d", rand.Intn(1000))
				}
				cache.Get(ctx, key)
			} else if op < 0.9 {
				// Set
				key = fmt.Sprintf("key-%d-%d", i, rand.Intn(1000))
				cache.Set(ctx, key, fmt.Sprintf("value-%d", i), nil)
				i++
			} else {
				// Delete
				if rand.Float32() < 0.8 {
					// Hit
					key = fmt.Sprintf("key-%d", rand.Intn(500))
				} else {
					// Miss
					key = fmt.Sprintf("missing-%d", rand.Intn(1000))
				}
				cache.Delete(ctx, key)
			}
		}
	})
}

// Benchmark for cleanup routine
func BenchmarkLocalCache_CleanupRoutine(b *testing.B) {
	ctx := context.Background()

	// Run this benchmark multiple times with different cache sizes
	for _, size := range []int{100, 1000, 10000} {
		b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
			// Create cache with very short TTL
			cache := NewLocalCache[string](size, 10*time.Millisecond)
			defer cache.Stop()

			// Pre-populate with items that will expire quickly
			for i := 0; i < size; i++ {
				key := fmt.Sprintf("key-%d", i)
				cache.Set(ctx, key, fmt.Sprintf("value-%d", i), nil)
			}

			b.ResetTimer()

			// Let the cleanup routine do its work multiple times
			for i := 0; i < b.N; i++ {
				// Wait for items to expire and cleanup to run
				time.Sleep(20 * time.Millisecond)

				// Add new items
				for j := 0; j < size/10; j++ {
					key := fmt.Sprintf("key-%d-%d", i, j)
					cache.Set(ctx, key, fmt.Sprintf("value-%d-%d", i, j), nil)
				}
			}
		})
	}
}

// Benchmark LRU eviction
func BenchmarkLocalCache_LRUEviction(b *testing.B) {
	ctx := context.Background()

	// Create a small cache to trigger eviction
	cache := NewLocalCache[string](100, 0) // No TTL to isolate LRU performance

	b.ResetTimer()

	// Continuously add new items to trigger LRU eviction
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		cache.Set(ctx, key, fmt.Sprintf("value-%d", i), nil)
	}
}

// Benchmark different value types
func BenchmarkLocalCache_ValueTypes(b *testing.B) {
	// String values (small)
	b.Run("SmallStrings", func(b *testing.B) {
		ctx := context.Background()
		cache := NewLocalCache[string](1000, 0)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := strconv.Itoa(i % 1000)
			cache.Set(ctx, key, "small value", nil)
		}
	})

	// String values (large)
	b.Run("LargeStrings", func(b *testing.B) {
		ctx := context.Background()
		cache := NewLocalCache[string](1000, 0)

		// Create a 10KB string
		largeValue := string(make([]byte, 10*1024))

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := strconv.Itoa(i % 1000)
			cache.Set(ctx, key, largeValue, nil)
		}
	})

	// Struct values
	b.Run("Structs", func(b *testing.B) {
		type TestStruct struct {
			ID       int
			Name     string
			Values   []int
			Metadata map[string]string
		}

		ctx := context.Background()
		cache := NewLocalCache[TestStruct](1000, 0)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := strconv.Itoa(i % 1000)
			value := TestStruct{
				ID:     i,
				Name:   fmt.Sprintf("Test-%d", i),
				Values: []int{1, 2, 3, 4, 5},
				Metadata: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			}
			cache.Set(ctx, key, value, nil)
		}
	})
}
