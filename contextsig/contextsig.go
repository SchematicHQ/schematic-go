// Package contextsig signs the context a browser SDK sends with a publishable
// key, for environments that require context signatures.
//
// A publishable key is public, so the API can't trust the company and user a
// browser names with it. When the environment requires signatures, your
// backend, which holds the secret API key, signs the logged-in user's own
// company and user, and the browser SDK sends that signature in the
// X-Schematic-Context-Sig header. Sign the context from your own authenticated
// session, never from what the browser asks for.
//
// The signature is `iat.exp.sig`, where `sig` is the unpadded base64url
// HMAC-SHA256 of `canonical(context).iat.exp` keyed with the secret API key,
// and `canonical` is the context's JSON with keys sorted and no whitespace.
package contextsig

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// Header is the request header the browser SDK sends the signature in.
const Header = "X-Schematic-Context-Sig"

// DefaultLifetime is how long a signature is valid for unless WithLifetime
// says otherwise. It matches how long the browser SDK reuses one.
const DefaultLifetime = 10 * time.Minute

// MaxLifetime is the longest the API accepts a signature as valid for.
const MaxLifetime = time.Hour

// ErrLifetimeInvalid is returned for a lifetime under a second or over
// MaxLifetime.
var ErrLifetimeInvalid = errors.New("contextsig: lifetime must be between one second and one hour")

// ErrSecretMissing is returned when no secret API key is given.
var ErrSecretMissing = errors.New("contextsig: secret API key is required")

// Context is the company and user a browser SDK checks flags and sends events
// for, by their keys. A nil map is left out of the signature; an empty one is
// signed as empty. It must match the context the browser SDK is given exactly.
type Context struct {
	Company map[string]string
	User    map[string]string
}

// Option configures Sign.
type Option func(*options)

type options struct {
	lifetime time.Duration
}

// WithLifetime sets how long the signature is valid for, from one second up to
// MaxLifetime. It's counted in whole seconds.
func WithLifetime(lifetime time.Duration) Option {
	return func(o *options) {
		o.lifetime = lifetime
	}
}

// Sign returns the X-Schematic-Context-Sig value for ctx, valid from now for
// DefaultLifetime or the lifetime WithLifetime gives. secret is the secret API
// key for the environment the browser SDK's publishable key belongs to.
func Sign(secret string, ctx Context, opts ...Option) (string, error) {
	o := options{lifetime: DefaultLifetime}
	for _, opt := range opts {
		opt(&o)
	}
	if o.lifetime < time.Second || o.lifetime > MaxLifetime {
		return "", ErrLifetimeInvalid
	}

	iat := time.Now().Unix()
	return sign(secret, ctx, iat, iat+int64(o.lifetime/time.Second))
}

func sign(secret string, ctx Context, iat, exp int64) (string, error) {
	if secret == "" {
		return "", ErrSecretMissing
	}

	canonical, err := canonicalize(ctx)
	if err != nil {
		return "", err
	}

	payload := make([]byte, 0, len(canonical)+32)
	payload = append(payload, canonical...)
	payload = append(payload, '.')
	payload = strconv.AppendInt(payload, iat, 10)
	payload = append(payload, '.')
	payload = strconv.AppendInt(payload, exp, 10)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	sig := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return strconv.FormatInt(iat, 10) + "." + strconv.FormatInt(exp, 10) + "." + sig, nil
}

// canonicalize encodes ctx the way the API does before verifying: maps encode
// with their keys in byte order, and HTML escaping is off so `<`, `>` and `&`
// stay literal.
func canonicalize(ctx Context) ([]byte, error) {
	body := map[string]map[string]string{}
	if ctx.Company != nil {
		body["company"] = ctx.Company
	}
	if ctx.User != nil {
		body["user"] = ctx.User
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(body); err != nil {
		return nil, err
	}

	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
