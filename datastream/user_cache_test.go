package datastream

import (
	"testing"
	"time"

	"github.com/schematichq/rulesengine"
	"github.com/schematichq/schematic-go/cache"
	"github.com/stretchr/testify/assert"
)

func TestUserIDCacheKeyFormat(t *testing.T) {
	client := &DataStreamClient{}

	tests := []struct {
		name     string
		id       string
		expected string
	}{
		{
			name:     "Standard ID",
			id:       "user_abc123",
			expected: "schematic:user:" + rulesengine.VersionKey + ":user_abc123",
		},
		{
			name:     "UUID-style ID",
			id:       "550e8400-e29b-41d4-a716-446655440000",
			expected: "schematic:user:" + rulesengine.VersionKey + ":550e8400-e29b-41d4-a716-446655440000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := client.userIDCacheKey(tt.id)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetUserFromCache_NewFormat(t *testing.T) {
	// Test the two-step lookup: lookup key -> user ID -> user at ID key
	userCache := cache.NewLocalCache[*rulesengine.User](100, time.Minute)
	lookupCache := cache.NewLocalCache[string](100, time.Minute)

	client := &DataStreamClient{
		userCacheProvider:       userCache,
		userLookupCacheProvider: lookupCache,
	}

	user := &rulesengine.User{
		ID:            "user_abc123",
		AccountID:     "acct_001",
		EnvironmentID: "env_001",
		Keys: map[string]string{
			"email":   "alice@test.com",
			"user_id": "usr-123",
		},
	}

	// Populate caches in the new format
	idKey := client.userIDCacheKey(user.ID)
	userCache.Set(nil, idKey, user, nil)

	for key, value := range user.Keys {
		lookupKey := client.resourceKeyToCacheKey(cacheKeyPrefixUser, key, value)
		lookupCache.Set(nil, lookupKey, user.ID, nil)
	}

	// Resolve via each key
	result := client.getUserFromCache(map[string]string{"email": "alice@test.com"})
	assert.NotNil(t, result)
	assert.Equal(t, "user_abc123", result.ID)

	result = client.getUserFromCache(map[string]string{"user_id": "usr-123"})
	assert.NotNil(t, result)
	assert.Equal(t, "user_abc123", result.ID)
}

func TestGetUserFromCache_NotFound(t *testing.T) {
	userCache := cache.NewLocalCache[*rulesengine.User](100, time.Minute)
	lookupCache := cache.NewLocalCache[string](100, time.Minute)

	client := &DataStreamClient{
		userCacheProvider:       userCache,
		userLookupCacheProvider: lookupCache,
	}

	result := client.getUserFromCache(map[string]string{"email": "nonexistent@test.com"})
	assert.Nil(t, result)
}

func TestGetUserFromCache_MultipleKeys(t *testing.T) {
	// Test that any matching key resolves the user
	userCache := cache.NewLocalCache[*rulesengine.User](100, time.Minute)
	lookupCache := cache.NewLocalCache[string](100, time.Minute)

	client := &DataStreamClient{
		userCacheProvider:       userCache,
		userLookupCacheProvider: lookupCache,
	}

	user := &rulesengine.User{
		ID: "user_multi",
		Keys: map[string]string{
			"email":   "multi@test.com",
			"user_id": "multi-user",
		},
	}

	// Only cache one lookup key (email), not user_id
	idKey := client.userIDCacheKey(user.ID)
	userCache.Set(nil, idKey, user, nil)

	lookupKey := client.resourceKeyToCacheKey(cacheKeyPrefixUser, "email", "multi@test.com")
	lookupCache.Set(nil, lookupKey, user.ID, nil)

	// Looking up by email should work
	result := client.getUserFromCache(map[string]string{"email": "multi@test.com"})
	assert.NotNil(t, result)
	assert.Equal(t, "user_multi", result.ID)

	// Looking up ONLY by the uncached key should return nil
	result = client.getUserFromCache(map[string]string{"user_id": "multi-user"})
	assert.Nil(t, result, "Should return nil when only the uncached key is provided")
}

func TestGetUserFromCache_MissingIDKey(t *testing.T) {
	// Test that nil is returned when lookup cache has the ID but the ID key is missing
	userCache := cache.NewLocalCache[*rulesengine.User](100, time.Minute)
	lookupCache := cache.NewLocalCache[string](100, time.Minute)

	client := &DataStreamClient{
		userCacheProvider:       userCache,
		userLookupCacheProvider: lookupCache,
	}

	// Populate lookup cache with user ID, but don't populate the ID key
	lookupKey := client.resourceKeyToCacheKey(cacheKeyPrefixUser, "email", "orphaned@test.com")
	lookupCache.Set(nil, lookupKey, "user_missing", nil)

	// Should return nil since ID key doesn't exist
	result := client.getUserFromCache(map[string]string{"email": "orphaned@test.com"})
	assert.Nil(t, result, "Should return nil when ID key is missing")
}
