package datastream

import (
	"testing"
)

func TestGetRulesEngineVersionKey(t *testing.T) {
	t.Run("Generates consistent version key", func(t *testing.T) {
		key1 := GetRulesEngineVersionKey()
		key2 := GetRulesEngineVersionKey()

		if key1 != key2 {
			t.Errorf("Expected version keys to be consistent, got %s and %s", key1, key2)
		}

		if len(key1) != 8 {
			t.Errorf("Expected version key to be 8 characters, got %d", len(key1))
		}
	})

	t.Run("Version key is not empty", func(t *testing.T) {
		key := GetRulesEngineVersionKey()

		if key == "" {
			t.Error("Expected version key to not be empty")
		}
	})

	t.Run("Global variable matches function call", func(t *testing.T) {
		if RulesEngineVersionKey != GetRulesEngineVersionKey() {
			t.Errorf("Expected global variable to match function call, got %s vs %s",
				RulesEngineVersionKey, GetRulesEngineVersionKey())
		}
	})
}
