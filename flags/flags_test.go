package flags

import (
	"testing"

	schematicgo "github.com/schematichq/schematic-go"
)

func TestFlagCheckCacheKey(t *testing.T) {
	id1 := "123"
	acme := "ACME Inc."

	id2 := "456"
	emailJohn := "john@example.com"

	id3 := "789"
	xyz := "XYZ Corp."

	id4 := "abc"
	emailJane := "jane@example.com"
	tests := []struct {
		name     string
		evalCtx  *schematicgo.CheckFlagRequestBody
		flagKey  string
		expected string
	}{
		{
			name: "Empty context and flag key",
			evalCtx: &schematicgo.CheckFlagRequestBody{
				Company: map[string]string{},
				User:    map[string]string{},
			},
			flagKey:  "",
			expected: "f:",
		},
		{
			name: "Context with company and user data",
			evalCtx: &schematicgo.CheckFlagRequestBody{
				Company: map[string]string{
					"id":   id1,
					"name": acme,
				},
				User: map[string]string{
					"id":    id2,
					"email": emailJohn,
				},
			},
			flagKey:  "feature_flag_1",
			expected: "f:feature_flag_1;c:id:123;c:name:ACME Inc.;u:email:john@example.com;u:id:456",
		},
		{
			name: "Context with only company data",
			evalCtx: &schematicgo.CheckFlagRequestBody{
				Company: map[string]string{
					"id":   id3,
					"name": xyz,
				},
				User: map[string]string{},
			},
			flagKey:  "feature_flag_2",
			expected: "f:feature_flag_2;c:id:789;c:name:XYZ Corp.",
		},
		{
			name: "Context with only user data",
			evalCtx: &schematicgo.CheckFlagRequestBody{
				Company: map[string]string{},
				User: map[string]string{
					"id":    id4,
					"email": emailJane,
				},
			},
			flagKey:  "feature_flag_3",
			expected: "f:feature_flag_3;u:email:jane@example.com;u:id:abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FlagCheckCacheKey(tt.evalCtx, tt.flagKey)
			if result != tt.expected {
				t.Errorf("flagCheckCacheKey() = %v, want %v", result, tt.expected)
			}
		})
	}
}
