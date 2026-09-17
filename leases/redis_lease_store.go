package leases

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	// DefaultKeyPrefix namespaces every lease and reservation key.
	DefaultKeyPrefix = "schematic:"

	leaseKeyNamespace = "credit-lease:"
	// leaseTTLGraceMs is how long after the declared expiry the row survives
	// before Redis evicts it. It gives the sweeper a window to refund expired
	// reservations before the lease state underneath them disappears.
	leaseTTLGraceMs = 60_000
)

// Every Lua script below touches exactly ONE key (the lease hash), keeping
// them safe under Redis Cluster (multi-key scripts spanning slots raise
// CROSSSLOT). Only the lease hash needs atomic mutation; cross-key
// bookkeeping uses ordinary single-key commands.
//
// Expiry is decided against the *Redis server's* clock (`redis.call('TIME')`),
// not the calling pod's: with many pods sharing one lease, local clock skew
// would let pods disagree on whether the lease is live. The `leaseNowMs`
// snippet converts TIME to integer milliseconds (matching the stored
// `expiresAt`); `redis.replicate_commands()` first, so the non-deterministic
// TIME read is allowed alongside writes on Redis 5/6.
//
// The key layout, hash fields and script bodies are identical to the Node and
// Python SDKs', which is what lets all three fleets share one Redis. Try-reserve
// returns the leaseId it debited: reading it in the same atomic step is how a
// hold gets pinned to the lease the credits actually came from.
const leaseNowMs = `
redis.replicate_commands()
local t = redis.call('TIME')
local now = (tonumber(t[1]) * 1000) + math.floor(tonumber(t[2]) / 1000)
`

// Atomic `replace`. Writes the lease hash only when the slot is empty or the
// existing lease has expired. Returns 1 on write, 0 if a *live* lease already
// occupies the slot, even one with a different leaseId, e.g. installed by a
// sibling instance that raced this acquire. An expired row with the SAME
// leaseId is reconciled like an extend instead of rewritten, which would reset
// the balance and erase debits whose reservations are still open.
const replaceScriptSource = leaseNowMs + `
local existing_id = redis.call('HGET', KEYS[1], 'leaseId')
local existing_expiry = tonumber(redis.call('HGET', KEYS[1], 'expiresAt') or '0')
local new_id = ARGV[1]
local new_granted = ARGV[2]
local new_expiry = tonumber(ARGV[3])
local grace = tonumber(ARGV[4])

if existing_id and existing_expiry > now then
    return 0
end

if existing_id == new_id then
    local granted = tonumber(redis.call('HGET', KEYS[1], 'grantedAmount') or '0')
    local add = tonumber(new_granted) - granted
    if add > 0 then
        local remaining = tonumber(redis.call('HGET', KEYS[1], 'localRemainingCredits') or '0')
        redis.call('HSET', KEYS[1],
            'grantedAmount', new_granted,
            'localRemainingCredits', tostring(remaining + add))
    end
    if new_expiry > existing_expiry then
        redis.call('HSET', KEYS[1], 'expiresAt', ARGV[3])
        redis.call('PEXPIREAT', KEYS[1], new_expiry + grace)
    end
    return 0
end

redis.call('DEL', KEYS[1])
redis.call('HSET', KEYS[1],
    'leaseId', new_id,
    'companyId', ARGV[5],
    'creditTypeId', ARGV[6],
    'grantedAmount', new_granted,
    'localRemainingCredits', new_granted,
    'expiresAt', ARGV[3])
redis.call('PEXPIREAT', KEYS[1], new_expiry + grace)
return 1
`

// Atomic check-and-decrement on `localRemainingCredits`. Returns the post-debit
// balance as a string (a Lua number reply truncates to integer, which would
// corrupt fractional credit costs) alongside the leaseId it came out of; nil if
// there is no lease, the lease has expired, or there is insufficient remaining.
// A hash without a leaseId is refused rather than debited: nothing could pin a
// reservation to it, so the debit would be one no refund could ever reach.
// The expiry guard compares against the Redis server clock, so a reserve
// against an expired-but-not-yet-evicted row during the TTL grace window is
// rejected.
const tryReserveScriptSource = leaseNowMs + `
local raw = redis.call('HGET', KEYS[1], 'localRemainingCredits')
if not raw then return false end
local lease_id = redis.call('HGET', KEYS[1], 'leaseId')
if not lease_id then return false end
local expiry = tonumber(redis.call('HGET', KEYS[1], 'expiresAt') or '0')
if expiry <= now then return false end
local remaining = tonumber(raw)
local requested = tonumber(ARGV[1])
if remaining < requested then return false end
local new_remaining = remaining - requested
redis.call('HSET', KEYS[1], 'localRemainingCredits', tostring(new_remaining))
return { tostring(new_remaining), lease_id }
`

// Refund credits, clamped at `grantedAmount`. ARGV[2], when non-empty, pins the
// refund to a specific leaseId: if the slot now holds a different lease, the
// refund is dropped, because the expired lease's unspent remainder was already
// returned to the company balance server-side, so crediting the successor would
// mint phantom credits.
const refundScriptSource = `
local raw_remaining = redis.call('HGET', KEYS[1], 'localRemainingCredits')
if not raw_remaining then return 0 end
local required_lease = ARGV[2]
if required_lease and required_lease ~= '' then
    local current_lease = redis.call('HGET', KEYS[1], 'leaseId')
    if current_lease ~= required_lease then return 0 end
end
local remaining = tonumber(raw_remaining)
local granted = tonumber(redis.call('HGET', KEYS[1], 'grantedAmount') or '0')
local refund = tonumber(ARGV[1])
local new_balance = remaining + refund
if new_balance > granted then new_balance = granted end
redis.call('HSET', KEYS[1], 'localRemainingCredits', tostring(new_balance))
return 1
`

// Reconcile the lease to the server-authoritative grantedAmount total
// (ARGV[1]), crediting the difference to localRemainingCredits. The delta is
// computed HERE, atomically against the hash's current total, never by the
// caller from a pre-wire-call read: two pods extending the same shared lease
// concurrently would each apply a delta against the same stale read and mint
// phantom credits. Reconciling to the absolute total converges regardless of
// arrival order. Expiry only ever moves forward. ARGV[4], when non-empty, pins
// the extend to a leaseId, mirroring the refund script.
const extendScriptSource = `
local raw_granted = redis.call('HGET', KEYS[1], 'grantedAmount')
if not raw_granted then return 0 end
local required_lease = ARGV[4]
if required_lease and required_lease ~= '' then
    local current_lease = redis.call('HGET', KEYS[1], 'leaseId')
    if current_lease ~= required_lease then return 0 end
end
local granted = tonumber(raw_granted)
local target = tonumber(ARGV[1])
local add = target - granted
if add > 0 then
    local remaining = tonumber(redis.call('HGET', KEYS[1], 'localRemainingCredits') or '0')
    redis.call('HSET', KEYS[1],
        'grantedAmount', tostring(target),
        'localRemainingCredits', tostring(remaining + add))
end
local new_expiry = tonumber(ARGV[2])
local grace = tonumber(ARGV[3])
local current_expiry = tonumber(redis.call('HGET', KEYS[1], 'expiresAt') or '0')
if new_expiry > current_expiry then
    redis.call('HSET', KEYS[1], 'expiresAt', ARGV[2])
    redis.call('PEXPIREAT', KEYS[1], new_expiry + grace)
end
return 1
`

var (
	replaceScript    = redis.NewScript(replaceScriptSource)
	tryReserveScript = redis.NewScript(tryReserveScriptSource)
	refundScript     = redis.NewScript(refundScriptSource)
	extendScript     = redis.NewScript(extendScriptSource)
)

// RedisLeaseStore keeps lease slots in Redis, one hash per slot, mutated by
// single-key Lua so it is correct on standalone and clustered Redis alike.
// Balances are stored as strings so fractional credit amounts survive the round
// trip.
type RedisLeaseStore struct {
	client               redis.UniversalClient
	keyPrefix            string
	defaultLeaseDuration time.Duration
	clock                Clock
}

// RedisLeaseStoreOptions configures a RedisLeaseStore. The zero value is the
// production configuration.
type RedisLeaseStoreOptions struct {
	// KeyPrefix defaults to DefaultKeyPrefix. Pass the same prefix to the
	// reservation store.
	KeyPrefix string
	// DefaultLeaseDuration is only reached when a direct caller extends without
	// an expiry; the lease manager always passes one. Defaults to
	// DefaultLeaseDuration.
	DefaultLeaseDuration time.Duration
	// Clock defaults to time.Now. Note that liveness is decided by the Redis
	// server's clock, not this one; it only sizes the fallback extend expiry.
	Clock Clock
}

func NewRedisLeaseStore(client redis.UniversalClient, opts RedisLeaseStoreOptions) *RedisLeaseStore {
	prefix := opts.KeyPrefix
	if prefix == "" {
		prefix = DefaultKeyPrefix
	}
	duration := opts.DefaultLeaseDuration
	if duration <= 0 {
		duration = DefaultLeaseDuration
	}
	return &RedisLeaseStore{
		client:               client,
		keyPrefix:            prefix,
		defaultLeaseDuration: duration,
		clock:                orNow(opts.Clock),
	}
}

func (s *RedisLeaseStore) hashKey(companyID, creditTypeID string) string {
	return s.keyPrefix + leaseKeyNamespace + LeaseKey(companyID, creditTypeID)
}

func (s *RedisLeaseStore) Get(ctx context.Context, companyID, creditTypeID string) (*LeaseState, error) {
	raw, err := s.client.HGetAll(ctx, s.hashKey(companyID, creditTypeID)).Result()
	if err != nil {
		return nil, err
	}
	if raw["leaseId"] == "" {
		return nil, nil
	}
	return &LeaseState{
		LeaseID:               raw["leaseId"],
		CompanyID:             fallbackString(raw["companyId"], companyID),
		CreditTypeID:          fallbackString(raw["creditTypeId"], creditTypeID),
		GrantedAmount:         parseAmount(raw["grantedAmount"]),
		LocalRemainingCredits: parseAmount(raw["localRemainingCredits"]),
		ExpiresAt:             fromEpochMs(int64(parseAmount(raw["expiresAt"]))),
	}, nil
}

func (s *RedisLeaseStore) Replace(ctx context.Context, grant LeaseGrant) (bool, error) {
	result, err := replaceScript.Run(
		ctx,
		s.client,
		[]string{s.hashKey(grant.CompanyID, grant.CreditTypeID)},
		// No client clock here: the script reads `now` from the Redis server
		// via TIME, so every pod agrees on expiry.
		grant.LeaseID,
		formatAmount(grant.GrantedAmount),
		strconv.FormatInt(toEpochMs(grant.ExpiresAt), 10),
		strconv.Itoa(leaseTTLGraceMs),
		grant.CompanyID,
		grant.CreditTypeID,
	).Int64()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

func (s *RedisLeaseStore) TryReserve(ctx context.Context, companyID, creditTypeID string, credits float64) (float64, string, bool, error) {
	// Reject non-finite/negative debits before they reach the script: the
	// string form of NaN parses back to a Lua nan, slips through the `<`
	// comparison, and would poison the SHARED balance for every pod.
	if !IsValidQuantity(credits) {
		return 0, "", false, nil
	}
	reply, err := tryReserveScript.Run(
		ctx,
		s.client,
		[]string{s.hashKey(companyID, creditTypeID)},
		formatAmount(credits),
	).Slice()
	if errors.Is(err, redis.Nil) {
		return 0, "", false, nil
	}
	if err != nil {
		return 0, "", false, err
	}
	if len(reply) != 2 {
		return 0, "", false, fmt.Errorf("try-reserve returned %d fields, want the balance and the lease id", len(reply))
	}
	rawBalance, _ := reply[0].(string)
	balance, err := strconv.ParseFloat(rawBalance, 64)
	if err != nil {
		return 0, "", false, fmt.Errorf("try-reserve returned an unparseable balance %q: %w", rawBalance, err)
	}
	leaseID, _ := reply[1].(string)
	return balance, leaseID, true, nil
}

func (s *RedisLeaseStore) Refund(ctx context.Context, companyID, creditTypeID string, credits float64, pinLeaseID string) error {
	if !IsValidQuantity(credits) || credits <= 0 {
		return nil
	}
	return refundScript.Run(
		ctx,
		s.client,
		[]string{s.hashKey(companyID, creditTypeID)},
		formatAmount(credits),
		// An empty string disables the lease pin (Lua has no nil ARGV).
		pinLeaseID,
	).Err()
}

func (s *RedisLeaseStore) Extend(ctx context.Context, companyID, creditTypeID string, grantedTotal float64, newExpiresAt *time.Time, pinLeaseID string) error {
	expiry := s.clock().Add(s.defaultLeaseDuration)
	if newExpiresAt != nil {
		expiry = *newExpiresAt
	}
	return extendScript.Run(
		ctx,
		s.client,
		[]string{s.hashKey(companyID, creditTypeID)},
		// grantedTotal is the server-authoritative TOTAL; the script computes
		// the delta against the stored total.
		formatAmount(grantedTotal),
		strconv.FormatInt(toEpochMs(expiry), 10),
		strconv.Itoa(leaseTTLGraceMs),
		pinLeaseID,
	).Err()
}

func (s *RedisLeaseStore) Drop(ctx context.Context, companyID, creditTypeID string) error {
	// A plain single-key delete: no secondary index to keep in sync.
	return s.client.Del(ctx, s.hashKey(companyID, creditTypeID)).Err()
}

// toEpochMs writes an instant the way the Node SDK does, so a shared row reads
// alike from every fleet.
func toEpochMs(instant time.Time) int64 {
	return instant.UnixMilli()
}

func fromEpochMs(ms int64) time.Time {
	return time.UnixMilli(ms).UTC()
}

// formatAmount writes a credit amount without exponent notation, so Lua's
// tonumber and every SDK's parser read back the same figure.
func formatAmount(value float64) string {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return "0"
	}
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func parseAmount(raw string) float64 {
	value, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0
	}
	return value
}

func fallbackString(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
