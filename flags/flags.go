package flags

import (
	"fmt"
	"sort"
	"strings"

	schematicgo "github.com/schematichq/schematic-go"
)

// TODO: Use separators that are not valid characters in key values
func FlagCheckCacheKey(evalCtx *schematicgo.CheckFlagRequestBody, flagKey string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("f:%s;", flagKey))

	companyKeys := make([]string, 0, len(evalCtx.Company))
	for k := range evalCtx.Company {
		companyKeys = append(companyKeys, k)
	}
	sort.Strings(companyKeys)
	for _, k := range companyKeys {
		sb.WriteString(fmt.Sprintf("c:%s:%s;", k, *evalCtx.Company[k]))
	}

	userKeys := make([]string, 0, len(evalCtx.User))
	for k := range evalCtx.User {
		userKeys = append(userKeys, k)
	}
	sort.Strings(userKeys)
	for _, k := range userKeys {
		sb.WriteString(fmt.Sprintf("u:%s:%s;", k, *evalCtx.User[k]))
	}

	return strings.TrimRight(sb.String(), ";")
}
