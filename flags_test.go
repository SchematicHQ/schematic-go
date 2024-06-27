package schematic

import "testing"

func TestFlagCheckCacheKey(t *testing.T) {
	tests := []struct {
		name     string
		evalCtx  *CheckFlagRequestBody
		flagKey  string
		expected string
	}{
		{
			name: "Empty context and flag key",
			evalCtx: &CheckFlagRequestBody{
				Company: map[string]string{},
				User:    map[string]string{},
			},
			flagKey:  "",
			expected: "f:",
		},
		{
			name: "Context with company and user data",
			evalCtx: &CheckFlagRequestBody{
				Company: map[string]string{
					"id":   "123",
					"name": "ACME Inc.",
				},
				User: map[string]string{
					"id":    "456",
					"email": "john@example.com",
				},
			},
			flagKey:  "feature_flag_1",
			expected: "f:feature_flag_1;c:id:123;c:name:ACME Inc.;u:email:john@example.com;u:id:456",
		},
		{
			name: "Context with only company data",
			evalCtx: &CheckFlagRequestBody{
				Company: map[string]string{
					"id":   "789",
					"name": "XYZ Corp.",
				},
				User: map[string]string{},
			},
			flagKey:  "feature_flag_2",
			expected: "f:feature_flag_2;c:id:789;c:name:XYZ Corp.",
		},
		{
			name: "Context with only user data",
			evalCtx: &CheckFlagRequestBody{
				Company: map[string]string{},
				User: map[string]string{
					"id":    "abc",
					"email": "jane@example.com",
				},
			},
			flagKey:  "feature_flag_3",
			expected: "f:feature_flag_3;u:email:jane@example.com;u:id:abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := flagCheckCacheKey(tt.evalCtx, tt.flagKey)
			if result != tt.expected {
				t.Errorf("flagCheckCacheKey() = %v, want %v", result, tt.expected)
			}
		})
	}
}
