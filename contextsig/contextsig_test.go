package contextsig

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// signingVectors is testdata/context-signing.json, copied verbatim from
// api/lib/contextsig/testdata in the Schematic API, which generates it from
// the verifier and checks it in its own tests. Don't edit it here.
type signingVectors struct {
	MaxLifetimeSeconds int64 `json:"max_lifetime_seconds"`
	Vectors            []struct {
		Canonical string `json:"canonical"`
		Context   struct {
			Company map[string]string `json:"company"`
			User    map[string]string `json:"user"`
		} `json:"context"`
		Exp       int64  `json:"exp"`
		Iat       int64  `json:"iat"`
		Name      string `json:"name"`
		Secret    string `json:"secret"`
		Signature string `json:"signature"`
	} `json:"vectors"`
}

func TestSigningVectors(t *testing.T) {
	t.Parallel()

	raw, err := os.ReadFile(filepath.Join("testdata", "context-signing.json"))
	require.NoError(t, err)
	var vectors signingVectors
	require.NoError(t, json.Unmarshal(raw, &vectors))
	require.NotEmpty(t, vectors.Vectors)
	assert.Equal(t, int64(MaxLifetime/time.Second), vectors.MaxLifetimeSeconds)

	for _, v := range vectors.Vectors {
		t.Run(v.Name, func(t *testing.T) {
			t.Parallel()
			ctx := Context{Company: v.Context.Company, User: v.Context.User}

			canonical, err := canonicalize(ctx)
			require.NoError(t, err)
			assert.Equal(t, v.Canonical, string(canonical))

			signature, err := sign(v.Secret, ctx, v.Iat, v.Exp)
			require.NoError(t, err)
			assert.Equal(t, v.Signature, signature)
		})
	}
}

func TestSign(t *testing.T) {
	t.Parallel()
	ctx := Context{Company: map[string]string{"id": "comp_1"}}

	t.Run("signs for the default lifetime from now", func(t *testing.T) {
		t.Parallel()
		before := time.Now().Unix()
		signature, err := Sign("secret", ctx)
		require.NoError(t, err)

		iat, exp := parseWindow(t, signature)
		assert.GreaterOrEqual(t, iat, before)
		assert.LessOrEqual(t, iat, time.Now().Unix())
		assert.Equal(t, int64(DefaultLifetime/time.Second), exp-iat)
	})

	t.Run("signs for the lifetime given", func(t *testing.T) {
		t.Parallel()
		signature, err := Sign("secret", ctx, WithLifetime(MaxLifetime))
		require.NoError(t, err)

		iat, exp := parseWindow(t, signature)
		assert.Equal(t, int64(MaxLifetime/time.Second), exp-iat)
	})

	t.Run("signs from the time given", func(t *testing.T) {
		t.Parallel()
		issuedAt := time.Unix(1_700_000_000, 0)
		signature, err := Sign("secret", ctx, WithIssuedAt(issuedAt), WithLifetime(5*time.Minute))
		require.NoError(t, err)

		iat, exp := parseWindow(t, signature)
		assert.Equal(t, issuedAt.Unix(), iat)
		assert.Equal(t, issuedAt.Add(5*time.Minute).Unix(), exp)

		expected, err := sign("secret", ctx, iat, exp)
		require.NoError(t, err)
		assert.Equal(t, expected, signature)
	})

	t.Run("refuses a lifetime the API would reject", func(t *testing.T) {
		t.Parallel()
		for _, lifetime := range []time.Duration{0, -time.Minute, 500 * time.Millisecond, MaxLifetime + time.Second, 24 * time.Hour} {
			_, err := Sign("secret", ctx, WithLifetime(lifetime))
			assert.ErrorIs(t, err, ErrLifetimeInvalid, lifetime.String())
		}
	})

	t.Run("refuses a missing secret", func(t *testing.T) {
		t.Parallel()
		_, err := Sign("", ctx)
		assert.ErrorIs(t, err, ErrSecretMissing)
	})
}

func parseWindow(t *testing.T, signature string) (int64, int64) {
	t.Helper()
	parts := strings.Split(signature, ".")
	require.Len(t, parts, 3)
	iat, err := strconv.ParseInt(parts[0], 10, 64)
	require.NoError(t, err)
	exp, err := strconv.ParseInt(parts[1], 10, 64)
	require.NoError(t, err)
	return iat, exp
}
