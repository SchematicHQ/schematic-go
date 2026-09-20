package leases

import (
	"context"
	"errors"
	"math"
	"sync"
	"testing"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Unit coverage for the paths the conformance vectors cannot reach: entity
// resolution, store and engine failures, and the flag_check event the lease path
// owes the analytics pipeline. The vectors pin the happy path and the failure
// modes that are expressible as JSON.

const (
	testCreditID = "ct_1"
	testSubtype  = "inference_tokens"
)

// fakeDataStream answers from fixtures and can be made to fail at any seam.
type fakeDataStream struct {
	flag        *rulesengine.Flag
	flagMissing bool
	company     *rulesengine.Company
	companyErr  error
	user        *rulesengine.User
	userErr     error

	// results are handed out in call order; an exhausted script errors.
	results []*rulesengine.CheckFlagResult
	errs    []error
	calls   []engineCall
}

func (d *fakeDataStream) GetFlag(context.Context, string) (*rulesengine.Flag, bool) {
	if d.flagMissing {
		return nil, false
	}
	return d.flag, true
}

func (d *fakeDataStream) GetCompany(context.Context, map[string]string) (*rulesengine.Company, error) {
	if d.companyErr != nil {
		return nil, d.companyErr
	}
	return d.company, nil
}

func (d *fakeDataStream) GetUser(context.Context, map[string]string) (*rulesengine.User, error) {
	if d.userErr != nil {
		return nil, d.userErr
	}
	return d.user, nil
}

func (d *fakeDataStream) EvaluateFlag(
	_ context.Context,
	_ *rulesengine.Flag,
	company *rulesengine.Company,
	_ *rulesengine.User,
	opts *EvalOptions,
) (*rulesengine.CheckFlagResult, error) {
	balances := map[string]float64{}
	if company != nil {
		balances = company.CreditBalances
	}
	d.calls = append(d.calls, engineCall{creditBalances: balances, options: opts})

	if len(d.errs) > 0 {
		err := d.errs[0]
		d.errs = d.errs[1:]
		if err != nil {
			if len(d.results) > 0 {
				d.results = d.results[1:]
			}
			return nil, err
		}
	}
	if len(d.results) == 0 {
		return nil, errors.New("unscripted engine call")
	}
	result := d.results[0]
	d.results = d.results[1:]
	return result, nil
}

func creditEntitlementResult(value bool, reason string) *rulesengine.CheckFlagResult {
	creditID := testCreditID
	subtype := testSubtype
	rate := 10.0
	return &rulesengine.CheckFlagResult{
		Value:   value,
		Reason:  reason,
		FlagKey: "inference",
		Entitlement: &rulesengine.FeatureEntitlement{
			FeatureID:       "feat_1",
			FeatureKey:      "inference",
			ValueType:       rulesengine.EntitlementValueTypeCredit,
			CreditID:        &creditID,
			ConsumptionRate: &rate,
			EventSubtype:    &subtype,
		},
	}
}

func plainResult(value bool, reason string) *rulesengine.CheckFlagResult {
	return &rulesengine.CheckFlagResult{Value: value, Reason: reason, FlagKey: "inference"}
}

// checkFixture is one flow under test, with its stores and its recorded events.
type checkFixture struct {
	datastream CheckDataStream
	leases     *InMemoryLeaseStore
	store      ReservationStore
	manager    *LeaseManager
	wire       *scriptedWireClient
	clock      *virtualClock
	events     []*schematicgo.EventBodyFlagCheck
	fellBack   bool
}

func newCheckFixture(t *testing.T, datastream CheckDataStream) *checkFixture {
	t.Helper()
	clock := newVirtualClock()
	leaseStore := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	reservations := NewInMemoryReservationStore(leaseStore, InMemoryReservationStoreOptions{Clock: clock.Now})
	wire := &scriptedWireClient{clock: clock}
	manager := NewLeaseManager(wire, leaseStore, LeaseManagerOptions{
		Reservations: reservations,
		Config:       ResolvedLeaseConfig{LeaseSize: 1000, LeaseDuration: 5 * time.Minute, ReservationTTL: time.Minute},
		Clock:        clock.Now,
	})
	t.Cleanup(manager.Stop)
	return &checkFixture{
		datastream: datastream,
		leases:     leaseStore,
		store:      reservations,
		manager:    manager,
		wire:       wire,
		clock:      clock,
	}
}

func (f *checkFixture) deps() CheckDeps {
	return CheckDeps{
		DataStream:   f.datastream,
		Leases:       f.leases,
		Reservations: f.store,
		Manager:      f.manager,
		Clock:        f.clock.Now,
		EmitFlagCheck: func(_ context.Context, body *schematicgo.EventBodyFlagCheck) {
			f.events = append(f.events, body)
		},
	}
}

func (f *checkFixture) run(t *testing.T, req CheckRequest) *CheckOutcome {
	t.Helper()
	outcome := CheckWithLease(context.Background(), f.deps(), req, func(context.Context) *CheckOutcome {
		f.fellBack = true
		return &CheckOutcome{Allowed: true, Value: true, Reason: "fallback", FlagKey: req.FlagKey}
	})
	f.manager.drainBackground()
	return outcome
}

// installTestLease gives the fixture a live lease to draw on.
func (f *checkFixture) installTestLease(t *testing.T, companyID string, granted float64) {
	t.Helper()
	installLease(t, f.leases, f.clock, "lse_1", companyID, testCreditID, granted, 300000)
}

func baseRequest() CheckRequest {
	return CheckRequest{
		FlagKey:      "inference",
		Company:      map[string]string{"id": "co_1"},
		Usage:        10,
		EventSubtype: testSubtype,
	}
}

func testCompany() *rulesengine.Company {
	return &rulesengine.Company{ID: "co_1", CreditBalances: map[string]float64{testCreditID: 5000}}
}

func TestCheckWithLeaseResolvesUserBeforeEvaluating(t *testing.T) {
	datastream := &fakeDataStream{
		flag:    &rulesengine.Flag{ID: "flag_1", Key: "inference"},
		company: testCompany(),
		user:    &rulesengine.User{ID: "user_1"},
		results: []*rulesengine.CheckFlagResult{creditEntitlementResult(true, "probe"), plainResult(true, "ok")},
	}
	fixture := newCheckFixture(t, datastream)
	fixture.installTestLease(t, "co_1", 1000)

	request := baseRequest()
	request.User = map[string]string{"id": "user_1"}
	outcome := fixture.run(t, request)

	require.True(t, outcome.Allowed)
	require.NotNil(t, outcome.Reservation)
	require.Len(t, fixture.events, 1)
	require.NotNil(t, fixture.events[0].UserID)
	assert.Equal(t, "user_1", *fixture.events[0].UserID)
}

func TestCheckWithLeaseFallsBackWhenAnEntityCannotBeResolved(t *testing.T) {
	t.Run("company", func(t *testing.T) {
		datastream := &fakeDataStream{
			flag:       &rulesengine.Flag{ID: "flag_1", Key: "inference"},
			companyErr: errors.New("datastream not connected"),
		}
		fixture := newCheckFixture(t, datastream)

		outcome := fixture.run(t, baseRequest())

		assert.True(t, fixture.fellBack)
		assert.Equal(t, "fallback", outcome.Reason)
		assert.Empty(t, fixture.events, "the plain check reports its own flag_check")
	})

	t.Run("user", func(t *testing.T) {
		datastream := &fakeDataStream{
			flag:    &rulesengine.Flag{ID: "flag_1", Key: "inference"},
			company: testCompany(),
			userErr: errors.New("timeout while waiting for user data"),
		}
		fixture := newCheckFixture(t, datastream)

		request := baseRequest()
		request.User = map[string]string{"id": "user_1"}
		outcome := fixture.run(t, request)

		assert.True(t, fixture.fellBack)
		assert.Equal(t, "fallback", outcome.Reason)
	})

	t.Run("flag", func(t *testing.T) {
		fixture := newCheckFixture(t, &fakeDataStream{flagMissing: true})

		outcome := fixture.run(t, baseRequest())

		assert.True(t, fixture.fellBack)
		assert.Equal(t, "fallback", outcome.Reason)
	})
}

func TestCheckWithLeaseFallsBackWhenTheProbeErrors(t *testing.T) {
	datastream := &fakeDataStream{
		flag:    &rulesengine.Flag{ID: "flag_1", Key: "inference"},
		company: testCompany(),
		results: []*rulesengine.CheckFlagResult{nil},
		errs:    []error{errors.New("wasm unavailable")},
	}
	fixture := newCheckFixture(t, datastream)

	outcome := fixture.run(t, baseRequest())

	assert.True(t, fixture.fellBack)
	assert.Equal(t, "fallback", outcome.Reason)
	assert.Nil(t, outcome.Reservation)
}

func TestCheckWithLeaseFallsBackWithoutAnEventSubtype(t *testing.T) {
	probe := creditEntitlementResult(true, "probe")
	probe.Entitlement.EventSubtype = nil
	datastream := &fakeDataStream{
		flag:    &rulesengine.Flag{ID: "flag_1", Key: "inference"},
		company: testCompany(),
		results: []*rulesengine.CheckFlagResult{probe},
	}
	fixture := newCheckFixture(t, datastream)
	fixture.installTestLease(t, "co_1", 1000)

	request := baseRequest()
	request.EventSubtype = ""
	outcome := fixture.run(t, request)

	assert.True(t, fixture.fellBack)
	assert.Nil(t, outcome.Reservation)
	entry, err := fixture.leases.Get(context.Background(), "co_1", testCreditID)
	require.NoError(t, err)
	assert.Equal(t, 1000.0, entry.LocalRemainingCredits, "no hold is taken on a fallback")
}

func TestCheckWithLeaseRejectsNonFiniteUsageWithoutTouchingTheStores(t *testing.T) {
	datastream := &fakeDataStream{
		flag:    &rulesengine.Flag{ID: "flag_1", Key: "inference"},
		company: testCompany(),
	}
	fixture := newCheckFixture(t, datastream)
	fixture.installTestLease(t, "co_1", 1000)

	request := baseRequest()
	request.Usage = math.NaN()
	outcome := fixture.run(t, request)

	assert.False(t, outcome.Allowed)
	assert.Equal(t, "invalid_usage", outcome.Reason)
	assert.Equal(t, "invalid_usage", outcome.Error)
	assert.False(t, fixture.fellBack)
	assert.Empty(t, datastream.calls, "a malformed usage never reaches the engine")
	entry, err := fixture.leases.Get(context.Background(), "co_1", testCreditID)
	require.NoError(t, err)
	assert.Equal(t, 1000.0, entry.LocalRemainingCredits)
	require.Len(t, fixture.events, 1, "the lease path reports the check it resolved itself")
}

func TestCheckWithLeaseFailOpenBlanketAllowsWhenTheFailOpenEvaluationErrors(t *testing.T) {
	datastream := &fakeDataStream{
		flag:    &rulesengine.Flag{ID: "flag_1", Key: "inference"},
		company: testCompany(),
		results: []*rulesengine.CheckFlagResult{creditEntitlementResult(true, "probe"), nil},
		errs:    []error{nil, errors.New("wasm unavailable")},
	}
	fixture := newCheckFixture(t, datastream)
	// No lease installed and no scripted acquire, so the acquire fails.

	request := baseRequest()
	request.FailOpen = true
	outcome := fixture.run(t, request)

	assert.True(t, outcome.Allowed)
	assert.Equal(t, "lease_acquire_failed_fail_open", outcome.Reason)
	assert.Equal(t, "lease_acquire_failed", outcome.Error)
	assert.Nil(t, outcome.Reservation)
}

// failingReservationStore fails whichever call the test arms, so the flow's
// containment paths can be driven without an unreachable Redis.
type failingReservationStore struct {
	ReservationStore
	addErr     error
	consumeErr error
}

func (s *failingReservationStore) Add(ctx context.Context, reservation ReservationRecord) error {
	if s.addErr != nil {
		return s.addErr
	}
	return s.ReservationStore.Add(ctx, reservation)
}

func (s *failingReservationStore) Consume(ctx context.Context, id string, credits float64) (float64, bool, error) {
	if s.consumeErr != nil {
		return 0, false, s.consumeErr
	}
	return s.ReservationStore.Consume(ctx, id, credits)
}

func TestCheckWithLeaseUndoesTheDebitWhenTheHoldCannotBePersisted(t *testing.T) {
	datastream := &fakeDataStream{
		flag:    &rulesengine.Flag{ID: "flag_1", Key: "inference"},
		company: testCompany(),
		results: []*rulesengine.CheckFlagResult{creditEntitlementResult(true, "probe")},
	}
	fixture := newCheckFixture(t, datastream)
	fixture.installTestLease(t, "co_1", 1000)
	fixture.store = &failingReservationStore{ReservationStore: fixture.store, addErr: errors.New("redis down")}

	outcome := fixture.run(t, baseRequest())

	assert.False(t, outcome.Allowed)
	assert.Equal(t, "lease_store_error", outcome.Reason)
	assert.Nil(t, outcome.Reservation)
	entry, err := fixture.leases.Get(context.Background(), "co_1", testCreditID)
	require.NoError(t, err)
	assert.Equal(t, 1000.0, entry.LocalRemainingCredits, "the debit is returned rather than stranded")
}

func TestCheckWithLeaseReportsTheEngineVerdictAsAFlagCheckEvent(t *testing.T) {
	ruleID := "rule_1"
	denied := plainResult(false, "denied_by_targeting")
	denied.RuleID = &ruleID
	datastream := &fakeDataStream{
		flag:    &rulesengine.Flag{ID: "flag_1", Key: "inference"},
		company: testCompany(),
		results: []*rulesengine.CheckFlagResult{creditEntitlementResult(true, "probe"), denied},
	}
	fixture := newCheckFixture(t, datastream)
	fixture.installTestLease(t, "co_1", 1000)

	outcome := fixture.run(t, baseRequest())

	assert.False(t, outcome.Allowed)
	require.Len(t, fixture.events, 1)
	event := fixture.events[0]
	assert.Equal(t, "inference", event.FlagKey)
	assert.False(t, event.Value)
	assert.Equal(t, "denied_by_targeting", event.Reason)
	require.NotNil(t, event.RuleID)
	assert.Equal(t, ruleID, *event.RuleID)
	assert.Equal(t, map[string]string{"id": "co_1"}, event.ReqCompany)
}

func TestSettleReservationRefundsTheUnspentSliceAndBillsTheActual(t *testing.T) {
	datastream := &fakeDataStream{
		flag:    &rulesengine.Flag{ID: "flag_1", Key: "inference"},
		company: testCompany(),
		results: []*rulesengine.CheckFlagResult{creditEntitlementResult(true, "probe"), plainResult(true, "ok")},
	}
	fixture := newCheckFixture(t, datastream)
	fixture.installTestLease(t, "co_1", 1000)

	outcome := fixture.run(t, baseRequest())
	require.NotNil(t, outcome.Reservation)

	settled := SettleReservation(context.Background(), fixture.store, *outcome.Reservation, 4)

	require.NoError(t, settled.Err)
	assert.True(t, settled.SettledLocally)
	assert.Equal(t, testSubtype, settled.Track.Event)
	require.NotNil(t, settled.Track.Quantity)
	assert.Equal(t, int64(4), *settled.Track.Quantity)
	require.NotNil(t, settled.Track.LeaseID)
	assert.Equal(t, "lse_1", *settled.Track.LeaseID)
	assert.Nil(t, settled.Track.ReservationID, "a client-mode settle never names a server hold")

	entry, err := fixture.leases.Get(context.Background(), "co_1", testCreditID)
	require.NoError(t, err)
	assert.Equal(t, 960.0, entry.LocalRemainingCredits)
}

func TestSettleReservationStillBillsWhenTheLocalSettleFails(t *testing.T) {
	store := &failingReservationStore{
		ReservationStore: NewInMemoryReservationStore(NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{}), InMemoryReservationStoreOptions{}),
		consumeErr:       errors.New("redis down"),
	}
	record := ReservationRecord{
		ID:              "res_1",
		LeaseID:         "lse_1",
		CompanyID:       "co_1",
		CreditTypeID:    testCreditID,
		EventSubtype:    testSubtype,
		CreditsReserved: 100,
		ConsumptionRate: 10,
		Company:         map[string]string{"id": "co_1"},
	}

	settled := SettleReservation(context.Background(), store, record, 7)

	require.Error(t, settled.Err)
	assert.False(t, settled.SettledLocally)
	require.NotNil(t, settled.Track.Quantity)
	assert.Equal(t, int64(7), *settled.Track.Quantity)
}

func TestSettleQuantityRoundsAPartialUnitUp(t *testing.T) {
	assert.Equal(t, int64(4), SettleQuantity(4))
	assert.Equal(t, int64(5), SettleQuantity(4.1))
	assert.Equal(t, int64(0), SettleQuantity(0))
}

// The API rejects a non-integer quantity while processing the event, so a
// fractional settle billed as-is would be dropped server-side while the local
// ledger had already debited it. Rounded up, the direction the hold takes, so
// the hold, the local debit and the billed quantity all agree.
func TestSettleReservationBillsAPartialUnitAsAWholeOne(t *testing.T) {
	leaseStore := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{})
	_, err := leaseStore.Replace(context.Background(), LeaseGrant{
		LeaseID:       "lse_1",
		CompanyID:     "co_1",
		CreditTypeID:  testCreditID,
		GrantedAmount: 1000,
		ExpiresAt:     time.Now().Add(5 * time.Minute),
	})
	require.NoError(t, err)
	store := NewInMemoryReservationStore(leaseStore, InMemoryReservationStoreOptions{})
	record := ReservationRecord{
		ID:               "res_1",
		LeaseID:          "lse_1",
		CompanyID:        "co_1",
		CreditTypeID:     testCreditID,
		EventSubtype:     testSubtype,
		QuantityReserved: 0.5,
		CreditsReserved:  10,
		ConsumptionRate:  10,
		ExpiresAt:        time.Now().Add(time.Minute),
		Company:          map[string]string{"id": "co_1"},
	}
	require.NoError(t, store.Add(context.Background(), record))
	_, _, reserved, err := leaseStore.TryReserve(context.Background(), "co_1", testCreditID, 10)
	require.NoError(t, err)
	require.True(t, reserved)

	settled := SettleReservation(context.Background(), store, record, 0.5)

	require.NoError(t, settled.Err)
	assert.True(t, settled.SettledLocally)
	require.NotNil(t, settled.Track.Quantity)
	assert.Equal(t, int64(1), *settled.Track.Quantity)

	// The debit moved the lease by what the event bills: one whole unit.
	entry, err := leaseStore.Get(context.Background(), "co_1", testCreditID)
	require.NoError(t, err)
	assert.Equal(t, 990.0, entry.LocalRemainingCredits)
}

// swapBeforeReserve replaces the slot's lease just before the debit lands,
// which is the window a check's acquire and its reserve straddle.
type swapBeforeReserve struct {
	LeaseStore
	once sync.Once
	swap func()
}

func (s *swapBeforeReserve) TryReserve(ctx context.Context, companyID, creditTypeID string, credits float64) (float64, string, bool, error) {
	s.once.Do(s.swap)
	return s.LeaseStore.TryReserve(ctx, companyID, creditTypeID, credits)
}

// The window between a check's acquire and its debit spans a network call, so
// the slot can be carrying a successor by the time the debit lands. The hold
// has to name the lease the credits actually came out of: pinned to the lease
// the acquire returned, its refund would be dropped and the settling track
// event would bill a lease that never held the usage.
func TestCheckWithLeasePinsTheHoldToTheLeaseItDebited(t *testing.T) {
	for _, factory := range backendFactories {
		t.Run(factory.name, func(t *testing.T) {
			ctx := context.Background()
			b := factory.make(t)
			installLease(t, b.leases, b.clock, "lse_1", "co_1", testCreditID, 1000, 300_000)

			swapped := &swapBeforeReserve{LeaseStore: b.leases, swap: func() {
				// lse_1 expires and a successor takes the slot, as it would
				// while the check sat in its extend call.
				b.clock.advance(400 * time.Second)
				installLease(t, b.leases, b.clock, "lse_2", "co_1", testCreditID, 1000, 800_000)
			}}

			manager := NewLeaseManager(&scriptedWireClient{clock: b.clock}, b.leases, LeaseManagerOptions{
				Reservations: b.reservations,
				Config:       ResolvedLeaseConfig{LeaseSize: 1000, LeaseDuration: 5 * time.Minute, ReservationTTL: time.Minute},
				Clock:        b.clock.Now,
			})
			t.Cleanup(manager.Stop)

			deps := CheckDeps{
				DataStream: &fakeDataStream{
					flag:    &rulesengine.Flag{ID: "flag_1", Key: "inference"},
					company: testCompany(),
					results: []*rulesengine.CheckFlagResult{creditEntitlementResult(true, "probe"), plainResult(true, "ok")},
				},
				Leases:       swapped,
				Reservations: b.reservations,
				Manager:      manager,
				Clock:        b.clock.Now,
			}
			outcome := CheckWithLease(ctx, deps, baseRequest(), func(context.Context) *CheckOutcome {
				t.Error("the check should have gated, not fallen back")
				return &CheckOutcome{}
			})
			manager.drainBackground()

			require.NotNil(t, outcome.Reservation)
			assert.Equal(t, "lse_2", outcome.Reservation.LeaseID, "the hold names the lease the debit came out of")
			entry, err := b.leases.Get(ctx, "co_1", testCreditID)
			require.NoError(t, err)
			assert.Equal(t, 900.0, entry.LocalRemainingCredits, "and that is the lease that paid")

			// The refund is pinned to the record's lease, so a record naming
			// the wrong one silently drops it.
			_, claimed, err := b.reservations.Consume(ctx, outcome.Reservation.ID, 0)
			require.NoError(t, err)
			require.True(t, claimed)
			entry, err = b.leases.Get(ctx, "co_1", testCreditID)
			require.NoError(t, err)
			assert.Equal(t, 1000.0, entry.LocalRemainingCredits, "the refund reaches the lease that was debited")
		})
	}
}
