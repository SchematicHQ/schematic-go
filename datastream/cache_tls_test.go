package datastream

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/schematichq/schematic-go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// startTLSRedisStub runs a TLS listener that speaks just enough RESP to answer
// go-redis' connection handshake and a PING. It returns the address and a cert
// pool that trusts it. A plaintext client cannot talk to it, so a successful
// PING proves the TLS config reached the dialer.
func startTLSRedisStub(t *testing.T) (addr string, roots *x509.CertPool) {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "localhost"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:     []string{"localhost"},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}
	der, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	require.NoError(t, err)

	keyDER, err := x509.MarshalECPrivateKey(key)
	require.NoError(t, err)

	cert, err := tls.X509KeyPair(
		pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}),
		pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER}),
	)
	require.NoError(t, err)

	listener, err := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	})
	require.NoError(t, err)
	t.Cleanup(func() { _ = listener.Close() })

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go serveRESP(conn, listener.Addr().String())
		}
	}()

	parsed, err := x509.ParseCertificate(der)
	require.NoError(t, err)
	roots = x509.NewCertPool()
	roots.AddCert(parsed)

	return listener.Addr().String(), roots
}

// serveRESP answers PING with +PONG, CLUSTER SLOTS with a single-node topology, and
// everything else with a generic error, which is all go-redis needs to consider the
// connection usable. own is the listener's address, which the slot map has to name so
// that a cluster client dials this same stub for the node it is told about.
func serveRESP(conn net.Conn, own string) {
	defer func() { _ = conn.Close() }()

	reader := bufio.NewReader(conn)
	for {
		args, err := readRESPCommand(reader)
		if err != nil {
			return
		}
		if len(args) == 0 {
			continue
		}

		reply := fmt.Sprintf("-ERR unknown command '%s'\r\n", args[0])
		switch {
		case strings.EqualFold(args[0], "PING"):
			reply = "+PONG\r\n"
		case strings.EqualFold(args[0], "CLUSTER") && len(args) > 1 && strings.EqualFold(args[1], "SLOTS"):
			reply = clusterSlotsReply(own)
		}
		if _, err := conn.Write([]byte(reply)); err != nil {
			return
		}
	}
}

// clusterSlotsReply hands back a topology of one node owning every slot, pointing at
// the stub itself. A cluster client will not route a command until it has one.
func clusterSlotsReply(own string) string {
	host, port, err := net.SplitHostPort(own)
	if err != nil {
		return "-ERR bad stub address\r\n"
	}

	const nodeID = "0123456789abcdef0123456789abcdef01234567"

	return fmt.Sprintf("*1\r\n*3\r\n:0\r\n:16383\r\n*3\r\n$%d\r\n%s\r\n:%s\r\n$%d\r\n%s\r\n",
		len(host), host, port, len(nodeID), nodeID)
}

func readRESPCommand(reader *bufio.Reader) ([]string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	if len(line) < 2 || line[0] != '*' {
		return nil, fmt.Errorf("expected array, got %q", line)
	}

	count, err := strconv.Atoi(line[1 : len(line)-2])
	if err != nil {
		return nil, err
	}

	args := make([]string, 0, count)
	for i := 0; i < count; i++ {
		header, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		if len(header) < 2 || header[0] != '$' {
			return nil, fmt.Errorf("expected bulk string, got %q", header)
		}

		length, err := strconv.Atoi(header[1 : len(header)-2])
		if err != nil {
			return nil, err
		}

		buf := make([]byte, length+2) // payload plus trailing CRLF
		if _, err := io.ReadFull(reader, buf); err != nil {
			return nil, err
		}
		args = append(args, string(buf[:length]))
	}

	return args, nil
}

// TestBuildRedisClientTLSHandshake exercises the whole path an encrypted
// ElastiCache/Valkey cluster takes: a TLSConfig on the cache config has to
// survive buildRedisClient and reach the dialer, or the handshake never happens.
func TestBuildRedisClientTLSHandshake(t *testing.T) {
	addr, roots := startTLSRedisStub(t)

	t.Run("TLSConfig connects", func(t *testing.T) {
		client := buildRedisClient(&core.DatastreamOptions{
			CacheConfig: &core.RedisCacheConfig{
				Addr:        addr,
				DialTimeout: 5 * time.Second,
				ReadTimeout: 5 * time.Second,
				TLSConfig:   &tls.Config{RootCAs: roots, MinVersion: tls.VersionTLS12},
			},
		}, nil)
		require.NotNil(t, client)
		t.Cleanup(func() { _ = client.Close() })

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		pong, err := client.Ping(ctx).Result()
		require.NoError(t, err)
		assert.Equal(t, "PONG", pong)
	})

	t.Run("rediss URL connects without an explicit TLSConfig", func(t *testing.T) {
		// The URL supplies TLS; RootCAs still has to be threaded in for the
		// self-signed stub, but ServerName comes from the URL host.
		opts := ToRedisOptions(&core.RedisCacheConfig{
			Addr:        "rediss://" + addr,
			DialTimeout: 5 * time.Second,
			ReadTimeout: 5 * time.Second,
		})
		require.NotNil(t, opts.TLSConfig, "rediss scheme should imply TLS")
		opts.TLSConfig.RootCAs = roots
		opts.TLSConfig.ServerName = "localhost"

		client := redis.NewClient(opts)
		t.Cleanup(func() { _ = client.Close() })

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		pong, err := client.Ping(ctx).Result()
		require.NoError(t, err)
		assert.Equal(t, "PONG", pong)
	})

	t.Run("caller-supplied client keeps its own TLS", func(t *testing.T) {
		// The point of WithRedisClient: the SDK maps a fixed set of go-redis
		// fields, but a client handed over whole arrives with whatever its owner
		// configured. Nothing here is expressible through RedisCacheConfig.
		own := redis.NewClient(&redis.Options{
			Addr:        addr,
			DialTimeout: 5 * time.Second,
			ReadTimeout: 5 * time.Second,
			TLSConfig:   &tls.Config{RootCAs: roots, ServerName: "localhost", MinVersion: tls.VersionTLS13},
		})
		t.Cleanup(func() { _ = own.Close() })

		client := buildRedisClient(&core.DatastreamOptions{
			CacheConfig: core.WithRedisClient(own),
		}, nil)
		require.Same(t, own, client, "the datastream should share the client, not rebuild it")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		pong, err := client.Ping(ctx).Result()
		require.NoError(t, err)
		assert.Equal(t, "PONG", pong)
	})

	t.Run("cluster TLSConfig connects", func(t *testing.T) {
		// The case this whole change exists for: an encrypted ElastiCache/Valkey
		// cluster. Mapping the option is not proof it reaches the dialer.
		client := buildRedisClient(&core.DatastreamOptions{
			CacheConfig: &core.RedisCacheClusterConfig{
				Addrs:       []string{addr},
				DialTimeout: 5 * time.Second,
				ReadTimeout: 5 * time.Second,
				TLSConfig:   &tls.Config{RootCAs: roots, ServerName: "localhost", MinVersion: tls.VersionTLS12},
			},
		}, nil)
		require.NotNil(t, client)
		t.Cleanup(func() { _ = client.Close() })

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		pong, err := client.Ping(ctx).Result()
		require.NoError(t, err)
		assert.Equal(t, "PONG", pong)
	})

	t.Run("cluster rediss URL connects without an explicit TLSConfig", func(t *testing.T) {
		// ParseClusterURL takes ServerName from the URL host, which is the stub's
		// 127.0.0.1; RootCAs still has to be threaded in for the self-signed cert.
		opts, err := toRedisClusterOptions(&core.RedisCacheClusterConfig{
			Addrs:       []string{"rediss://" + addr},
			DialTimeout: 5 * time.Second,
			ReadTimeout: 5 * time.Second,
		})
		require.NoError(t, err)
		require.NotNil(t, opts.TLSConfig, "rediss scheme should imply TLS")
		assert.Equal(t, []string{addr}, opts.Addrs, "the scheme should be stripped off the address")
		opts.TLSConfig.RootCAs = roots

		client := redis.NewClusterClient(opts)
		t.Cleanup(func() { _ = client.Close() })

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		pong, err := client.Ping(ctx).Result()
		require.NoError(t, err)
		assert.Equal(t, "PONG", pong)
	})

	t.Run("nil TLSConfig cannot talk to an encrypted server", func(t *testing.T) {
		client := buildRedisClient(&core.DatastreamOptions{
			CacheConfig: &core.RedisCacheConfig{
				Addr:        addr,
				DialTimeout: 2 * time.Second,
				ReadTimeout: 2 * time.Second,
				MaxRetries:  -1,
			},
		}, nil)
		require.NotNil(t, client)
		t.Cleanup(func() { _ = client.Close() })

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err := client.Ping(ctx).Result()
		assert.Error(t, err, "a plaintext client should not reach a TLS-only server")
	})
}
