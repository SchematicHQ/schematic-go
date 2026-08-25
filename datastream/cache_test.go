package datastream

import (
	"context"
	"crypto/tls"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/schematichq/schematic-go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// recordingLogger captures warnings so the malformed-URL path can be asserted on.
type recordingLogger struct {
	warnings []string
}

func (l *recordingLogger) Debug(context.Context, string, ...any) {}
func (l *recordingLogger) Info(context.Context, string, ...any)  {}
func (l *recordingLogger) Error(context.Context, string, ...any) {}
func (l *recordingLogger) Warn(_ context.Context, message string, _ ...any) {
	l.warnings = append(l.warnings, message)
}

func TestToRedisOptionsBareAddress(t *testing.T) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	config := &core.RedisCacheConfig{
		Network:               "tcp",
		Addr:                  "localhost:6379",
		ClientName:            "schematic",
		Protocol:              3,
		Username:              "user",
		Password:              "secret",
		DB:                    4,
		MaxRetries:            3,
		MinRetryBackoff:       time.Millisecond,
		MaxRetryBackoff:       2 * time.Millisecond,
		DialTimeout:           5 * time.Second,
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          4 * time.Second,
		ContextTimeoutEnabled: true,
		PoolFIFO:              true,
		PoolSize:              10,
		PoolTimeout:           6 * time.Second,
		MinIdleConns:          1,
		MaxIdleConns:          2,
		MaxActiveConns:        20,
		ConnMaxIdleTime:       7 * time.Second,
		ConnMaxLifetime:       8 * time.Second,
		DisableIndentity:      true,
		DisableIdentity:       true,
		IdentitySuffix:        "suffix",
		UnstableResp3:         true,
		TLSConfig:             tlsConfig,
	}

	opts, err := toRedisOptions(config)
	require.NoError(t, err)

	assert.Equal(t, "tcp", opts.Network)
	assert.Equal(t, "localhost:6379", opts.Addr)
	assert.Equal(t, "schematic", opts.ClientName)
	assert.Equal(t, 3, opts.Protocol)
	assert.Equal(t, "user", opts.Username)
	assert.Equal(t, "secret", opts.Password)
	assert.Equal(t, 4, opts.DB)
	assert.Equal(t, 3, opts.MaxRetries)
	assert.Equal(t, time.Millisecond, opts.MinRetryBackoff)
	assert.Equal(t, 2*time.Millisecond, opts.MaxRetryBackoff)
	assert.Equal(t, 5*time.Second, opts.DialTimeout)
	assert.Equal(t, 3*time.Second, opts.ReadTimeout)
	assert.Equal(t, 4*time.Second, opts.WriteTimeout)
	assert.True(t, opts.ContextTimeoutEnabled)
	assert.True(t, opts.PoolFIFO)
	assert.Equal(t, 10, opts.PoolSize)
	assert.Equal(t, 6*time.Second, opts.PoolTimeout)
	assert.Equal(t, 1, opts.MinIdleConns)
	assert.Equal(t, 2, opts.MaxIdleConns)
	assert.Equal(t, 20, opts.MaxActiveConns)
	assert.Equal(t, 7*time.Second, opts.ConnMaxIdleTime)
	assert.Equal(t, 8*time.Second, opts.ConnMaxLifetime)
	assert.True(t, opts.DisableIndentity) //nolint:staticcheck // deprecated in go-redis, still mapped
	assert.True(t, opts.DisableIdentity)
	assert.Equal(t, "suffix", opts.IdentitySuffix)
	assert.True(t, opts.UnstableResp3) //nolint:staticcheck // deprecated in go-redis, still mapped
	assert.Same(t, tlsConfig, opts.TLSConfig)
}

func TestToRedisOptionsDefaultsToPlaintext(t *testing.T) {
	opts, err := toRedisOptions(&core.RedisCacheConfig{Addr: "localhost:6379"})

	require.NoError(t, err)
	assert.Nil(t, opts.TLSConfig)
}

func TestToRedisOptionsParsesSchemeInAddr(t *testing.T) {
	opts, err := toRedisOptions(&core.RedisCacheConfig{Addr: "redis://localhost:6379/2"})

	require.NoError(t, err)
	assert.Equal(t, "localhost:6379", opts.Addr)
	assert.Equal(t, 2, opts.DB)
	assert.Nil(t, opts.TLSConfig)
}

func TestToRedisOptionsParsesRedissAddr(t *testing.T) {
	opts, err := toRedisOptions(&core.RedisCacheConfig{Addr: "rediss://user:secret@cache.example.com:6379"})

	require.NoError(t, err)
	assert.Equal(t, "cache.example.com:6379", opts.Addr)
	assert.Equal(t, "user", opts.Username)
	assert.Equal(t, "secret", opts.Password)
	require.NotNil(t, opts.TLSConfig)
	assert.Equal(t, "cache.example.com", opts.TLSConfig.ServerName)
}

func TestToRedisOptionsURLFieldWinsOverAddr(t *testing.T) {
	opts, err := toRedisOptions(&core.RedisCacheConfig{
		URL:  "rediss://cache.example.com:6379",
		Addr: "localhost:6379",
	})

	require.NoError(t, err)
	assert.Equal(t, "cache.example.com:6379", opts.Addr)
	assert.NotNil(t, opts.TLSConfig)
}

func TestToRedisOptionsExplicitFieldsOverrideURL(t *testing.T) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS13}

	opts, err := toRedisOptions(&core.RedisCacheConfig{
		URL:         "rediss://urluser:urlsecret@cache.example.com:6379/1",
		Username:    "configuser",
		Password:    "configsecret",
		DB:          7,
		TLSConfig:   tlsConfig,
		DialTimeout: 5 * time.Second,
	})

	require.NoError(t, err)
	assert.Equal(t, "configuser", opts.Username)
	assert.Equal(t, "configsecret", opts.Password)
	assert.Equal(t, 7, opts.DB)
	assert.Same(t, tlsConfig, opts.TLSConfig)
	assert.Equal(t, 5*time.Second, opts.DialTimeout)
}

func TestToRedisOptionsKeepsURLDerivedValuesWhenConfigIsZero(t *testing.T) {
	opts, err := toRedisOptions(&core.RedisCacheConfig{
		URL: "redis://urluser:urlsecret@cache.example.com:6379/1?dial_timeout=9s",
	})

	require.NoError(t, err)
	assert.Equal(t, "urluser", opts.Username)
	assert.Equal(t, "urlsecret", opts.Password)
	assert.Equal(t, 1, opts.DB)
	assert.Equal(t, 9*time.Second, opts.DialTimeout)
}

func TestToRedisOptionsMalformedURLFallsBack(t *testing.T) {
	opts, err := toRedisOptions(&core.RedisCacheConfig{
		Addr:     "http://localhost:6379",
		PoolSize: 12,
	})

	require.Error(t, err)
	assert.Equal(t, "http://localhost:6379", opts.Addr)
	assert.Equal(t, 12, opts.PoolSize)
}

func TestNewRedisClientWarnsOnMalformedURL(t *testing.T) {
	logger := &recordingLogger{}

	// A nil client is the signal buildCacheProvidersFromConfig reads as "use the
	// local cache".
	client := newRedisClient(&core.RedisCacheConfig{Addr: "http://localhost:6379"}, logger)

	assert.True(t, client == nil, "expected an untyped nil, got %#v", client)
	require.Len(t, logger.warnings, 1)
	assert.Contains(t, logger.warnings[0], "http://localhost:6379")
	assert.Contains(t, logger.warnings[0], "local cache")
}

// With the URL in the URL field, Addr is empty, and go-redis turns an empty Addr into
// localhost:6379 — so a typo would dial whatever Redis runs beside the application.
func TestNewRedisClientMalformedURLDoesNotDialLocalhost(t *testing.T) {
	logger := &recordingLogger{}

	client := newRedisClient(&core.RedisCacheConfig{URL: "cache.example.com:6379"}, logger)

	assert.True(t, client == nil, "expected an untyped nil, got %#v", client)
	require.Len(t, logger.warnings, 1)
	assert.Contains(t, logger.warnings[0], "cache.example.com:6379")
	assert.NotContains(t, logger.warnings[0], "localhost")
}

// The URL that failed to parse may carry a password, and is about to reach the logs.
func TestNewRedisClientRedactsCredentials(t *testing.T) {
	logger := &recordingLogger{}

	client := newRedisClient(&core.RedisCacheConfig{
		URL: "rediss://user:hunter2@cache.example.com:6379/not-a-number",
	}, logger)

	assert.True(t, client == nil, "expected an untyped nil, got %#v", client)
	require.Len(t, logger.warnings, 1)
	assert.NotContains(t, logger.warnings[0], "hunter2")
	assert.NotContains(t, logger.warnings[0], "user:")
	assert.Contains(t, logger.warnings[0], "rediss://redacted@cache.example.com:6379/not-a-number")
}

func TestRedactRedisURL(t *testing.T) {
	tests := map[string]struct{ redacted string }{
		"rediss://user:secret@cache.example.com:6379/0": {"rediss://redacted@cache.example.com:6379/0"},
		"redis://user:secret@cache.example.com:6379":    {"redis://redacted@cache.example.com:6379"},
		"user:secret@cache.example.com:6379":            {"redacted@cache.example.com:6379"},
		"redis://cache.example.com:6379/0":              {"redis://cache.example.com:6379/0"},
		"cache.example.com:6379":                        {"cache.example.com:6379"},
		"":                                              {""},
		// A '/' or '?' in the password hides the userinfo '@'; the stray '@' beyond the
		// guessed authority is the tell, and the whole URL goes.
		"redis://user:pa/ss@cache.example.com:6379":  {"redis://redacted"},
		"rediss://user:se?cret@cache.example.com:63": {"rediss://redacted"},
		// The same tell fires on a query-value '@', costing the host. That is the trade.
		"redis://cache.example.com:6379/0?client_name=a@b": {"redis://redacted"},
	}

	for input, expected := range tests {
		assert.Equal(t, expected.redacted, redactRedisURL(input), "input %q", input)
	}
}

// The parse error is the harder half of the warning: net/url quotes the URL into it
// %q-escaped, and quotes fragments besides — "invalid port" names whatever sat where the
// port should have been, which for a password holding a '/' is the password.
func TestNewRedisClientRedactsCredentialsInErrorText(t *testing.T) {
	for _, connURL := range []string{
		// A control character is what makes url.Parse fail *and* what makes %q escape.
		"redis://user:hunter2@cache.example.com:6379/\x7f",
		"redis://user:hun\x7fter2@cache.example.com:6379",
		// The '/' leaves the password where net/url expects a port.
		"redis://user:hunter2/x@cache.example.com:6379",
		"redis://user:hunter2?x@cache.example.com:6379",
	} {
		t.Run(connURL, func(t *testing.T) {
			logger := &recordingLogger{}

			client := newRedisClient(&core.RedisCacheConfig{URL: connURL}, logger)

			assert.True(t, client == nil, "expected an untyped nil, got %#v", client)
			require.Len(t, logger.warnings, 1)
			assert.NotContains(t, logger.warnings[0], "hunter2")
			assert.NotContains(t, logger.warnings[0], "hun")
			assert.NotContains(t, logger.warnings[0], "user")
		})
	}
}

// A URL with no userinfo has nothing to leak, so it keeps the diagnostic in full.
func TestNewRedisClientKeepsErrorDetailWithoutCredentials(t *testing.T) {
	logger := &recordingLogger{}

	newRedisClient(&core.RedisCacheConfig{URL: "http://cache.example.com:6379"}, logger)

	require.Len(t, logger.warnings, 1)
	assert.Contains(t, logger.warnings[0], "invalid URL scheme")
}

// The other way a URL reaches localhost:6379: redis.ParseURL accepts a hostless URL and
// defaults it, so a half-written one parses cleanly and dials the wrong server.
func TestNewRedisClientRejectsHostlessURL(t *testing.T) {
	for _, connURL := range []string{"redis://", "redis:///0", "rediss://"} {
		t.Run(connURL, func(t *testing.T) {
			logger := &recordingLogger{}

			client := newRedisClient(&core.RedisCacheConfig{URL: connURL}, logger)

			assert.True(t, client == nil, "expected an untyped nil, got %#v", client)
			require.Len(t, logger.warnings, 1)
			assert.Contains(t, logger.warnings[0], "local cache")
			assert.NotContains(t, logger.warnings[0], "localhost")
		})
	}
}

// A unix:// URL has no host by design; the hostless check must not swallow it.
func TestToRedisOptionsAcceptsUnixSocketURL(t *testing.T) {
	opts, err := toRedisOptions(&core.RedisCacheConfig{URL: "unix:///var/run/redis.sock"})

	require.NoError(t, err)
	assert.Equal(t, "unix", opts.Network)
	assert.Equal(t, "/var/run/redis.sock", opts.Addr)
}

func TestNewRedisClientToleratesNilLogger(t *testing.T) {
	assert.NotPanics(t, func() {
		newRedisClient(&core.RedisCacheConfig{Addr: "http://localhost:6379"}, nil)
	})
}

func TestNewRedisClientDoesNotWarnOnValidConfig(t *testing.T) {
	logger := &recordingLogger{}

	client := newRedisClient(&core.RedisCacheConfig{Addr: "rediss://cache.example.com:6379"}, logger)
	require.NotNil(t, client)
	t.Cleanup(func() { _ = client.Close() })

	assert.Empty(t, logger.warnings)
}

func TestToRedisClusterOptionsBareAddrs(t *testing.T) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	config := &core.RedisCacheClusterConfig{
		Addrs:                 []string{"one.example.com:6379", "two.example.com:6379"},
		MaxRedirects:          5,
		RouteByLatency:        true,
		RouteRandomly:         true,
		Protocol:              3,
		Username:              "user",
		Password:              "secret",
		MaxRetries:            3,
		MinRetryBackoff:       time.Millisecond,
		MaxRetryBackoff:       2 * time.Millisecond,
		DialTimeout:           5 * time.Second,
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          4 * time.Second,
		ContextTimeoutEnabled: true,
		PoolFIFO:              true,
		PoolSize:              10,
		PoolTimeout:           6 * time.Second,
		MinIdleConns:          1,
		MaxIdleConns:          2,
		MaxActiveConns:        20,
		ConnMaxIdleTime:       7 * time.Second,
		ConnMaxLifetime:       8 * time.Second,
		DisableIndentity:      true,
		DisableIdentity:       true,
		IdentitySuffix:        "suffix",
		UnstableResp3:         true,
		TLSConfig:             tlsConfig,
	}

	opts, err := toRedisClusterOptions(config)
	require.NoError(t, err)

	assert.Equal(t, []string{"one.example.com:6379", "two.example.com:6379"}, opts.Addrs)
	assert.Equal(t, 5, opts.MaxRedirects)
	assert.True(t, opts.RouteByLatency)
	assert.True(t, opts.RouteRandomly)
	assert.Equal(t, 3, opts.Protocol)
	assert.Equal(t, "user", opts.Username)
	assert.Equal(t, "secret", opts.Password)
	assert.Equal(t, 3, opts.MaxRetries)
	assert.Equal(t, time.Millisecond, opts.MinRetryBackoff)
	assert.Equal(t, 2*time.Millisecond, opts.MaxRetryBackoff)
	assert.Equal(t, 5*time.Second, opts.DialTimeout)
	assert.Equal(t, 3*time.Second, opts.ReadTimeout)
	assert.Equal(t, 4*time.Second, opts.WriteTimeout)
	assert.True(t, opts.ContextTimeoutEnabled)
	assert.True(t, opts.PoolFIFO)
	assert.Equal(t, 10, opts.PoolSize)
	assert.Equal(t, 6*time.Second, opts.PoolTimeout)
	assert.Equal(t, 1, opts.MinIdleConns)
	assert.Equal(t, 2, opts.MaxIdleConns)
	assert.Equal(t, 20, opts.MaxActiveConns)
	assert.Equal(t, 7*time.Second, opts.ConnMaxIdleTime)
	assert.Equal(t, 8*time.Second, opts.ConnMaxLifetime)
	assert.True(t, opts.DisableIndentity) //nolint:staticcheck // deprecated in go-redis, still mapped
	assert.True(t, opts.DisableIdentity)
	assert.Equal(t, "suffix", opts.IdentitySuffix)
	assert.True(t, opts.UnstableResp3) //nolint:staticcheck // deprecated in go-redis, still mapped
	assert.Same(t, tlsConfig, opts.TLSConfig)
}

func TestToRedisClusterOptionsParsesSchemeInAddrs(t *testing.T) {
	opts, err := toRedisClusterOptions(&core.RedisCacheClusterConfig{
		Addrs: []string{"rediss://user:secret@clustercfg.example.com:6379"},
	})

	require.NoError(t, err)
	assert.Equal(t, []string{"clustercfg.example.com:6379"}, opts.Addrs)
	assert.Equal(t, "user", opts.Username)
	assert.Equal(t, "secret", opts.Password)
	require.NotNil(t, opts.TLSConfig)
	assert.Equal(t, "clustercfg.example.com", opts.TLSConfig.ServerName)
}

// A URL beside bare entries names one node; the bare ones name the rest.
func TestToRedisClusterOptionsKeepsBareAddrsBesideURL(t *testing.T) {
	opts, err := toRedisClusterOptions(&core.RedisCacheClusterConfig{
		Addrs: []string{"redis://one.example.com:6379", "two.example.com:6379"},
	})

	require.NoError(t, err)
	assert.Equal(t, []string{"one.example.com:6379", "two.example.com:6379"}, opts.Addrs)
}

// go-redis' own way of naming extra nodes on a cluster URL.
func TestToRedisClusterOptionsHonoursAddrQueryParam(t *testing.T) {
	opts, err := toRedisClusterOptions(&core.RedisCacheClusterConfig{
		URL: "redis://one.example.com:6379?addr=two.example.com:6379",
	})

	require.NoError(t, err)
	assert.Equal(t, []string{"one.example.com:6379", "two.example.com:6379"}, opts.Addrs)
}

func TestToRedisClusterOptionsURLFieldWinsOverAddrs(t *testing.T) {
	opts, err := toRedisClusterOptions(&core.RedisCacheClusterConfig{
		URL:   "rediss://clustercfg.example.com:6379",
		Addrs: []string{"ignored.example.com:6379"},
	})

	require.NoError(t, err)
	assert.Equal(t, []string{"clustercfg.example.com:6379"}, opts.Addrs)
	assert.NotNil(t, opts.TLSConfig)
}

func TestToRedisClusterOptionsExplicitFieldsOverrideURL(t *testing.T) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS13}

	opts, err := toRedisClusterOptions(&core.RedisCacheClusterConfig{
		URL:         "rediss://urluser:urlsecret@clustercfg.example.com:6379?max_retries=2",
		Username:    "configuser",
		Password:    "configsecret",
		MaxRetries:  5,
		TLSConfig:   tlsConfig,
		DialTimeout: 5 * time.Second,
	})

	require.NoError(t, err)
	assert.Equal(t, "configuser", opts.Username)
	assert.Equal(t, "configsecret", opts.Password)
	assert.Equal(t, 5, opts.MaxRetries)
	assert.Same(t, tlsConfig, opts.TLSConfig)
	assert.Equal(t, 5*time.Second, opts.DialTimeout)
}

func TestToRedisClusterOptionsKeepsURLDerivedValuesWhenConfigIsZero(t *testing.T) {
	opts, err := toRedisClusterOptions(&core.RedisCacheClusterConfig{
		URL: "redis://urluser:urlsecret@clustercfg.example.com:6379?max_retries=2&dial_timeout=9s",
	})

	require.NoError(t, err)
	assert.Equal(t, "urluser", opts.Username)
	assert.Equal(t, "urlsecret", opts.Password)
	assert.Equal(t, 2, opts.MaxRetries)
	assert.Equal(t, 9*time.Second, opts.DialTimeout)
}

// Two URLs carry two sets of credentials and TLS, so the config is refused outright.
func TestToRedisClusterOptionsRejectsTwoURLs(t *testing.T) {
	_, err := toRedisClusterOptions(&core.RedisCacheClusterConfig{
		Addrs: []string{"rediss://one.example.com:6379", "rediss://two.example.com:6379"},
	})

	require.ErrorIs(t, err, errRedisMultipleClusterURLs)
}

func TestNewRedisClusterClientFallsBack(t *testing.T) {
	tests := map[string]struct {
		config  *core.RedisCacheClusterConfig
		warning string
	}{
		"malformed URL": {
			config:  &core.RedisCacheClusterConfig{URL: "http://clustercfg.example.com:6379"},
			warning: "invalid URL scheme",
		},
		"hostless URL": {
			config:  &core.RedisCacheClusterConfig{URL: "redis://"},
			warning: "no host",
		},
		"scheme in Addrs": {
			config:  &core.RedisCacheClusterConfig{Addrs: []string{"http://clustercfg.example.com:6379"}},
			warning: "invalid URL scheme",
		},
		"two URLs": {
			config: &core.RedisCacheClusterConfig{Addrs: []string{
				"rediss://one.example.com:6379",
				"rediss://two.example.com:6379",
			}},
			warning: "more than one connection URL",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			logger := &recordingLogger{}

			client := newRedisClusterClient(test.config, logger)

			assert.True(t, client == nil, "expected an untyped nil, got %#v", client)
			require.Len(t, logger.warnings, 1)
			assert.Contains(t, logger.warnings[0], test.warning)
			assert.Contains(t, logger.warnings[0], "local cache")
			assert.NotContains(t, logger.warnings[0], "localhost")
		})
	}
}

func TestNewRedisClusterClientRedactsCredentials(t *testing.T) {
	logger := &recordingLogger{}

	// ParseClusterURL rejects an unknown query parameter; it ignores the path outright,
	// there being no DB to select on a cluster.
	client := newRedisClusterClient(&core.RedisCacheClusterConfig{
		URL: "rediss://user:hunter2@clustercfg.example.com:6379?not_a_param=1",
	}, logger)

	assert.True(t, client == nil, "expected an untyped nil, got %#v", client)
	require.Len(t, logger.warnings, 1)
	assert.NotContains(t, logger.warnings[0], "hunter2")
	assert.NotContains(t, logger.warnings[0], "user")
	assert.Contains(t, logger.warnings[0], "rediss://redacted@clustercfg.example.com:6379")
}

func TestBuildRedisClientClusterURL(t *testing.T) {
	t.Run("builds from a URL", func(t *testing.T) {
		client := buildRedisClient(&core.DatastreamOptions{
			CacheConfig: core.RedisCacheClusterConfig{URL: "rediss://clustercfg.example.com:6379"},
		}, nil)

		require.NotNil(t, client)
		t.Cleanup(func() { _ = client.Close() })
		assert.IsType(t, &redis.ClusterClient{}, client)
	})

	t.Run("declines an unusable URL", func(t *testing.T) {
		client := buildRedisClient(&core.DatastreamOptions{
			CacheConfig: &core.RedisCacheClusterConfig{URL: "http://clustercfg.example.com:6379"},
		}, nil)

		assert.True(t, client == nil, "expected an untyped nil, got %#v", client)
	})
}

func TestToRedisClusterOptionsTLSConfig(t *testing.T) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	withTLS := ToRedisClusterOptions(&core.RedisCacheClusterConfig{
		Addrs:     []string{"cache.example.com:6379"},
		TLSConfig: tlsConfig,
	})
	assert.Same(t, tlsConfig, withTLS.TLSConfig)

	withoutTLS := ToRedisClusterOptions(&core.RedisCacheClusterConfig{
		Addrs: []string{"cache.example.com:6379"},
	})
	assert.Nil(t, withoutTLS.TLSConfig)
}

// The WithRedisClient path: the client is shared rather than rebuilt.
func TestBuildRedisClientCallerSupplied(t *testing.T) {
	t.Run("shares the client", func(t *testing.T) {
		own := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
		t.Cleanup(func() { _ = own.Close() })

		assert.Same(t, own, buildRedisClient(&core.DatastreamOptions{
			CacheConfig: core.WithRedisClient(own),
		}, nil))
	})

	t.Run("accepts the config by pointer", func(t *testing.T) {
		own := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
		t.Cleanup(func() { _ = own.Close() })

		assert.Same(t, own, buildRedisClient(&core.DatastreamOptions{
			CacheConfig: &core.RedisClientConfig{Client: own},
		}, nil))
	})

	t.Run("nil client falls back to the local cache", func(t *testing.T) {
		got := buildRedisClient(&core.DatastreamOptions{
			CacheConfig: core.WithRedisClient(nil),
		}, nil)

		// Compared with ==, not assert.Nil, which reports a typed nil pointer as nil
		// too and so cannot tell the two apart.
		assert.True(t, got == nil, "expected an untyped nil, got %#v", got)
	})

	t.Run("a non-pointer implementation is used as-is", func(t *testing.T) {
		// reflect.Value.IsNil panics on a kind that cannot be nil, so a client
		// implementing the interface by value must be recognised without it.
		own := valueClient{redis.NewClient(&redis.Options{Addr: "localhost:6379"})}
		t.Cleanup(func() { _ = own.Close() })

		var got redis.UniversalClient
		require.NotPanics(t, func() {
			got = buildRedisClient(&core.DatastreamOptions{
				CacheConfig: core.WithRedisClient(own),
			}, nil)
		})
		assert.Equal(t, own, got)
	})

	t.Run("typed-nil client falls back to the local cache", func(t *testing.T) {
		// A (*redis.Client)(nil) in the interface is not == nil, so without the
		// reflect guard this reaches the caches and panics on first command.
		var typedNil *redis.Client

		got := buildRedisClient(&core.DatastreamOptions{
			CacheConfig: core.WithRedisClient(typedNil),
		}, nil)

		assert.True(t, got == nil, "expected an untyped nil, got %#v", got)
	})
}

// valueClient implements redis.UniversalClient with a struct rather than a pointer,
// which is what makes an unguarded reflect.Value.IsNil panic.
type valueClient struct {
	redis.UniversalClient
}
