package datastream

import (
	"testing"
	"time"

	"github.com/schematichq/rulesengine"
	"github.com/schematichq/schematic-go/cache"
	"github.com/stretchr/testify/assert"
)

func TestCompanyIDCacheKeyFormat(t *testing.T) {
	client := &DataStreamClient{}

	tests := []struct {
		name     string
		id       string
		expected string
	}{
		{
			name:     "Standard ID",
			id:       "comp_abc123",
			expected: "schematic:company:comp_abc123",
		},
		{
			name:     "UUID-style ID",
			id:       "550e8400-e29b-41d4-a716-446655440000",
			expected: "schematic:company:550e8400-e29b-41d4-a716-446655440000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := client.companyIDCacheKey(tt.id)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetCompanyFromCache_NewFormat(t *testing.T) {
	// Test the two-step lookup: lookup key -> company ID -> company at ID key
	companyCache := cache.NewLocalCache[*rulesengine.Company](100, time.Minute)
	lookupCache := cache.NewLocalCache[string](100, time.Minute)

	client := &DataStreamClient{
		companyCacheProvider:       companyCache,
		companyLookupCacheProvider: lookupCache,
	}

	company := &rulesengine.Company{
		ID:            "comp_abc123",
		AccountID:     "acct_001",
		EnvironmentID: "env_001",
		Keys: map[string]string{
			"domain":     "test.com",
			"company_id": "acme-corp",
		},
	}

	// Populate caches in the new format
	idKey := client.companyIDCacheKey(company.ID)
	companyCache.Set(nil, idKey, company, nil)

	for key, value := range company.Keys {
		lookupKey := client.resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)
		lookupCache.Set(nil, lookupKey, company.ID, nil)
	}

	// Resolve via each key
	result := client.getCompanyFromCache(map[string]string{"domain": "test.com"})
	assert.NotNil(t, result)
	assert.Equal(t, "comp_abc123", result.ID)

	result = client.getCompanyFromCache(map[string]string{"company_id": "acme-corp"})
	assert.NotNil(t, result)
	assert.Equal(t, "comp_abc123", result.ID)
}

func TestGetCompanyFromCache_OldFormat(t *testing.T) {
	// Test backwards compatibility: lookup key -> full company directly
	companyCache := cache.NewLocalCache[*rulesengine.Company](100, time.Minute)
	lookupCache := cache.NewLocalCache[string](100, time.Minute)

	client := &DataStreamClient{
		companyCacheProvider:       companyCache,
		companyLookupCacheProvider: lookupCache,
	}

	company := &rulesengine.Company{
		ID:            "comp_old_format",
		AccountID:     "acct_001",
		EnvironmentID: "env_001",
		Keys: map[string]string{
			"domain": "legacy.com",
		},
	}

	// Populate cache in the OLD format (full company at lookup key, no ID key)
	lookupKey := client.resourceKeyToCacheKey(cacheKeyPrefixCompany, "domain", "legacy.com")
	companyCache.Set(nil, lookupKey, company, nil)
	// Note: lookup cache is empty — no ID string stored there

	// Should still resolve via the old format fallback
	result := client.getCompanyFromCache(map[string]string{"domain": "legacy.com"})
	assert.NotNil(t, result, "Should resolve company via old format fallback")
	assert.Equal(t, "comp_old_format", result.ID)
}

func TestGetCompanyFromCache_NewFormatMissingIDKey(t *testing.T) {
	// Test fallback when lookup cache has the ID but the ID key is missing
	companyCache := cache.NewLocalCache[*rulesengine.Company](100, time.Minute)
	lookupCache := cache.NewLocalCache[string](100, time.Minute)

	client := &DataStreamClient{
		companyCacheProvider:       companyCache,
		companyLookupCacheProvider: lookupCache,
	}

	// Populate lookup cache with company ID, but don't populate the ID key
	lookupKey := client.resourceKeyToCacheKey(cacheKeyPrefixCompany, "domain", "orphaned.com")
	lookupCache.Set(nil, lookupKey, "comp_missing", nil)

	// Should return nil since ID key doesn't exist and old format doesn't exist either
	result := client.getCompanyFromCache(map[string]string{"domain": "orphaned.com"})
	assert.Nil(t, result, "Should return nil when ID key is missing and no old format fallback")
}

func TestGetCompanyFromCache_NewFormatPreferredOverOld(t *testing.T) {
	// If both old and new format data exist, new format should be used
	companyCache := cache.NewLocalCache[*rulesengine.Company](100, time.Minute)
	lookupCache := cache.NewLocalCache[string](100, time.Minute)

	client := &DataStreamClient{
		companyCacheProvider:       companyCache,
		companyLookupCacheProvider: lookupCache,
	}

	oldCompany := &rulesengine.Company{
		ID:            "comp_old",
		EnvironmentID: "env_old",
	}

	newCompany := &rulesengine.Company{
		ID:            "comp_new",
		EnvironmentID: "env_new",
	}

	lookupKey := client.resourceKeyToCacheKey(cacheKeyPrefixCompany, "domain", "dual.com")

	// Old format: full company at lookup key
	companyCache.Set(nil, lookupKey, oldCompany, nil)

	// New format: lookup key -> ID, ID key -> company
	lookupCache.Set(nil, lookupKey, newCompany.ID, nil)
	idKey := client.companyIDCacheKey(newCompany.ID)
	companyCache.Set(nil, idKey, newCompany, nil)

	// New format should win
	result := client.getCompanyFromCache(map[string]string{"domain": "dual.com"})
	assert.NotNil(t, result)
	assert.Equal(t, "comp_new", result.ID, "New format should be preferred over old format")
}

func TestGetCompanyFromCache_NotFound(t *testing.T) {
	companyCache := cache.NewLocalCache[*rulesengine.Company](100, time.Minute)
	lookupCache := cache.NewLocalCache[string](100, time.Minute)

	client := &DataStreamClient{
		companyCacheProvider:       companyCache,
		companyLookupCacheProvider: lookupCache,
	}

	result := client.getCompanyFromCache(map[string]string{"domain": "nonexistent.com"})
	assert.Nil(t, result)
}

func TestGetCompanyFromCache_MultipleKeys(t *testing.T) {
	// Test that any matching key resolves the company
	companyCache := cache.NewLocalCache[*rulesengine.Company](100, time.Minute)
	lookupCache := cache.NewLocalCache[string](100, time.Minute)

	client := &DataStreamClient{
		companyCacheProvider:       companyCache,
		companyLookupCacheProvider: lookupCache,
	}

	company := &rulesengine.Company{
		ID: "comp_multi",
		Keys: map[string]string{
			"domain":     "multi.com",
			"company_id": "multi-corp",
		},
	}

	// Only cache one lookup key (domain), not company_id
	idKey := client.companyIDCacheKey(company.ID)
	companyCache.Set(nil, idKey, company, nil)

	lookupKey := client.resourceKeyToCacheKey(cacheKeyPrefixCompany, "domain", "multi.com")
	lookupCache.Set(nil, lookupKey, company.ID, nil)

	// Looking up by domain should work
	result := client.getCompanyFromCache(map[string]string{"domain": "multi.com"})
	assert.NotNil(t, result)
	assert.Equal(t, "comp_multi", result.ID)

	// Looking up by company_id (not cached) should still work if domain is also in the map
	// because Go map iteration will eventually hit the cached key
	// But looking up ONLY by the uncached key should return nil
	result = client.getCompanyFromCache(map[string]string{"company_id": "multi-corp"})
	assert.Nil(t, result, "Should return nil when only the uncached key is provided")
}
