package leases

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

const (
	reservationKeyNamespace = "credit-reservation:"
	// reservationIndexKey names the sorted set scoring open reservations by
	// expiry so the sweeper can pop expired entries in O(log n). Members encode
	// the full (company, credit, id) tuple.
	reservationIndexKey = "credit-reservations:byExpiry"
	// reservationByCreditNamespace names the per-(company, credit) index of
	// open holds, one hash of id -> creditsReserved, so ReservedCredits reads a
	// tenant's holds with one HGETALL. The hash is also the source of truth for
	// that sum: a field exists exactly while its reservation is open and
	// unrefunded.
	reservationByCreditNamespace = "credit-reservations:byCredit:"
	// reservationTTLGraceMs buffers past expiry before Redis evicts the row, so
	// the sweeper has a window to refund.
	reservationTTLGraceMs = 30_000

	// sweepBatchSize is the page size for the sweeper's ZRANGEBYSCORE. Without
	// a limit, a backlog of expired holds (after a Redis outage or a long pod
	// pause) would come back as one giant reply on every pod's next tick;
	// paging bounds the reply while the per-member ZREM keeps offset 0
	// advancing through the backlog.
	sweepBatchSize = 256
	// maxSweepBatches bounds one sweep's work and guards against an endless
	// loop if ZREM persistently fails. Anything left over is picked up next
	// tick.
	maxSweepBatches = 16

	// memberDelimiter is absent from Schematic ids and from the reservation id.
	memberDelimiter = "|"
)

// Atomic claim: read the reservation hash and delete it in one step, returning
// its fields (or nil if it was already gone). Touches a single key. The atomic
// read-then-delete is what makes consume exactly-once: of two racing callers (a
// normal settle and a sweeper, say) only one gets the fields back and proceeds
// to refund. The refund to the lease hash is a separate single-key op; a crash
// in the gap leaves the unspent slice held on the lease until the lease itself
// expires, never double-refunded.
const claimScriptSource = `
local raw = redis.call('HGETALL', KEYS[1])
if #raw == 0 then return nil end
redis.call('DEL', KEYS[1])
return raw
`

var claimScript = redis.NewScript(claimScriptSource)

// RedisReservationStore keeps the reservation table in Redis and refunds
// through the lease store it is handed.
//
// Every mutation is a single-key operation (or single-key Lua), so the store is
// correct on standalone and clustered Redis alike: the unspent-slice refund is
// delegated to the lease store rather than reaching across to the lease hash
// inside a multi-key script.
type RedisReservationStore struct {
	client    redis.UniversalClient
	leases    ReservationRefunder
	keyPrefix string
	clock     Clock
}

// RedisReservationStoreOptions configures a RedisReservationStore. The zero
// value is the production configuration.
type RedisReservationStoreOptions struct {
	// KeyPrefix defaults to DefaultKeyPrefix. Pass the same prefix as the lease
	// store.
	KeyPrefix string
	// Clock defaults to time.Now. It decides the sweep cutoff.
	Clock Clock
}

func NewRedisReservationStore(client redis.UniversalClient, leases ReservationRefunder, opts RedisReservationStoreOptions) *RedisReservationStore {
	prefix := opts.KeyPrefix
	if prefix == "" {
		prefix = DefaultKeyPrefix
	}
	return &RedisReservationStore{
		client:    client,
		leases:    leases,
		keyPrefix: prefix,
		clock:     orNow(opts.Clock),
	}
}

func (s *RedisReservationStore) hashKey(id string) string {
	return s.keyPrefix + reservationKeyNamespace + id
}

func (s *RedisReservationStore) indexKey() string {
	return s.keyPrefix + reservationIndexKey
}

func (s *RedisReservationStore) byCreditKey(companyID, creditTypeID string) string {
	return s.keyPrefix + reservationByCreditNamespace + companyID + ":" + creditTypeID
}

func (s *RedisReservationStore) Add(ctx context.Context, reservation ReservationRecord) error {
	expiresMs := toEpochMs(reservation.ExpiresAt)
	hashKey := s.hashKey(reservation.ID)
	// The hash goes out first so the reservation exists before anything
	// references it. These are independent single-key ops rather than one
	// multi-key script: a partial failure at worst leaves an un-indexed
	// reservation that the TTL reaps, never a double-spend.
	if err := s.client.HSet(ctx, hashKey, map[string]any{
		"id":               reservation.ID,
		"leaseId":          reservation.LeaseID,
		"companyId":        reservation.CompanyID,
		"creditTypeId":     reservation.CreditTypeID,
		"eventSubtype":     reservation.EventSubtype,
		"quantityReserved": formatAmount(reservation.QuantityReserved),
		"creditsReserved":  formatAmount(reservation.CreditsReserved),
		"consumptionRate":  formatAmount(reservation.ConsumptionRate),
		"expiresAt":        strconv.FormatInt(expiresMs, 10),
		"evalCtx":          encodeEvalCtx(reservation),
	}).Err(); err != nil {
		return err
	}
	if err := s.client.PExpireAt(ctx, hashKey, fromEpochMs(expiresMs+reservationTTLGraceMs)).Err(); err != nil {
		return err
	}
	member := encodeMember(reservation.CompanyID, reservation.CreditTypeID, reservation.ID)
	if err := s.client.ZAdd(ctx, s.indexKey(), redis.Z{Score: float64(expiresMs), Member: member}).Err(); err != nil {
		return err
	}
	return s.client.HSet(
		ctx,
		s.byCreditKey(reservation.CompanyID, reservation.CreditTypeID),
		reservation.ID,
		formatAmount(reservation.CreditsReserved),
	).Err()
}

func (s *RedisReservationStore) Get(ctx context.Context, id string) (*ReservationRecord, error) {
	raw, err := s.client.HGetAll(ctx, s.hashKey(id)).Result()
	if err != nil {
		return nil, err
	}
	if raw["id"] == "" {
		return nil, nil
	}
	return decodeReservation(raw), nil
}

func (s *RedisReservationStore) Consume(ctx context.Context, id string, creditsConsumed float64) (float64, bool, error) {
	claimed, err := claimScript.Run(ctx, s.client, []string{s.hashKey(id)}).Slice()
	if errors.Is(err, redis.Nil) {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, err
	}
	raw := decodeFlat(claimed)
	if raw["id"] == "" {
		return 0, false, nil
	}

	companyID := raw["companyId"]
	creditTypeID := raw["creditTypeId"]
	reserved := parseAmount(raw["creditsReserved"])

	// Index cleanup, single-key ops. The per-tenant hash loses the slice BEFORE
	// the refund below, so the lease (local remaining plus this hash) never
	// transiently double-counts it. Both are best-effort: a failed cleanup must
	// not abort the settle.
	member := encodeMember(companyID, creditTypeID, id)
	_ = s.client.ZRem(ctx, s.indexKey(), member).Err()
	_ = s.client.HDel(ctx, s.byCreditKey(companyID, creditTypeID), id).Err()

	consumed := clampConsumption(creditsConsumed, reserved)
	if refund := reserved - consumed; refund > 0 {
		// The lease store owns the lease hash, which keeps this cross-key write
		// out of a single Lua script. Pinned to the reservation's lease so a
		// hold carved out of an expired lease cannot inflate a successor's
		// balance.
		if err := s.leases.Refund(ctx, companyID, creditTypeID, refund, raw["leaseId"]); err != nil {
			return 0, false, err
		}
	}
	return consumed, true, nil
}

func (s *RedisReservationStore) ReservedCredits(ctx context.Context, companyID, creditTypeID string) (float64, error) {
	raw, err := s.client.HGetAll(ctx, s.byCreditKey(companyID, creditTypeID)).Result()
	if err != nil {
		return 0, err
	}
	total := 0.0
	for _, value := range raw {
		total += parseAmount(value)
	}
	return total, nil
}

func (s *RedisReservationStore) SweepExpired(ctx context.Context) (int, error) {
	cutoff := toEpochMs(s.clock())
	swept := 0
	// Page through expired members rather than fetching them all at once. Each
	// processed member is removed below, so re-reading at offset 0 advances
	// through the backlog.
	for batch := 0; batch < maxSweepBatches; batch++ {
		expired, err := s.client.ZRangeByScore(ctx, s.indexKey(), &redis.ZRangeBy{
			Min:    "0",
			Max:    strconv.FormatInt(cutoff, 10),
			Offset: 0,
			Count:  sweepBatchSize,
		}).Result()
		if err != nil {
			return swept, err
		}
		if len(expired) == 0 {
			return swept, nil
		}
		for _, member := range expired {
			companyID, creditTypeID, reservationID, ok := decodeMember(member)
			if !ok {
				// Nothing but Add writes members, so this is belt-and-braces:
				// drop it rather than let it wedge the sweeper.
				_ = s.client.ZRem(ctx, s.indexKey(), member).Err()
				continue
			}
			_, claimed, err := s.Consume(ctx, reservationID, 0)
			if err != nil {
				return swept, err
			}
			// Always drop the member just read. On the success path Consume
			// already removed it, so this is idempotent; it also covers the
			// hash-evicted path below.
			_ = s.client.ZRem(ctx, s.indexKey(), member).Err()
			if claimed {
				swept++
				continue
			}
			// No reservation hash: either a racing settle consumed it (and
			// reconciled the byCredit field, making this a no-op) or the hash
			// TTL-evicted before the sweeper reached it, orphaning the field.
			// Reconcile so ReservedCredits stops summing an evicted hold.
			// Deliberately no refund: without the hash, exactly-once cannot be
			// arbitrated across racing sweepers, so the slice waits for the
			// lease to expire server-side.
			_ = s.client.HDel(ctx, s.byCreditKey(companyID, creditTypeID), reservationID).Err()
		}
		if len(expired) < sweepBatchSize {
			return swept, nil
		}
	}
	return swept, nil
}

func (s *RedisReservationStore) Count(ctx context.Context) (int, error) {
	count, err := s.client.ZCard(ctx, s.indexKey()).Result()
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

// encodeMember packs the whole tuple into an expiry-index member. The sweeper
// needs company and credit to clean the per-tenant hash even after the
// reservation hash has TTL-evicted, at which point the claim returns nil and
// cannot report them; otherwise the orphaned field would inflate
// ReservedCredits forever.
func encodeMember(companyID, creditTypeID, reservationID string) string {
	return companyID + memberDelimiter + creditTypeID + memberDelimiter + reservationID
}

func decodeMember(member string) (companyID, creditTypeID, reservationID string, ok bool) {
	parts := strings.Split(member, memberDelimiter)
	if len(parts) != 3 {
		return "", "", "", false
	}
	return parts[0], parts[1], parts[2], true
}

func encodeEvalCtx(reservation ReservationRecord) string {
	ctx := make(map[string]map[string]string, 2)
	if reservation.Company != nil {
		ctx["company"] = reservation.Company
	}
	if reservation.User != nil {
		ctx["user"] = reservation.User
	}
	encoded, err := json.Marshal(ctx)
	if err != nil {
		return "{}"
	}
	return string(encoded)
}

// decodeFlat decodes the flat [field, value, ...] reply the claim script
// returns.
func decodeFlat(raw []any) map[string]string {
	out := make(map[string]string, len(raw)/2)
	for i := 0; i+1 < len(raw); i += 2 {
		key, keyOK := raw[i].(string)
		value, valueOK := raw[i+1].(string)
		if keyOK && valueOK {
			out[key] = value
		}
	}
	return out
}

func decodeReservation(raw map[string]string) *ReservationRecord {
	var ctx struct {
		Company map[string]string `json:"company"`
		User    map[string]string `json:"user"`
	}
	if encoded := raw["evalCtx"]; encoded != "" {
		_ = json.Unmarshal([]byte(encoded), &ctx)
	}
	return &ReservationRecord{
		ID:               raw["id"],
		LeaseID:          raw["leaseId"],
		CompanyID:        raw["companyId"],
		CreditTypeID:     raw["creditTypeId"],
		EventSubtype:     raw["eventSubtype"],
		QuantityReserved: parseAmount(raw["quantityReserved"]),
		CreditsReserved:  parseAmount(raw["creditsReserved"]),
		ConsumptionRate:  parseAmount(raw["consumptionRate"]),
		ExpiresAt:        fromEpochMs(int64(parseAmount(raw["expiresAt"]))),
		Company:          ctx.Company,
		User:             ctx.User,
	}
}
