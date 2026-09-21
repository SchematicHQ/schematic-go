package tracing

import (
	"sort"
	"strings"

	"go.opentelemetry.io/otel/attribute"
)

// Attribute keys for the spans the SDK emits. They are namespaced under
// "schematic." so they cannot collide with an application's own attributes or
// with OpenTelemetry's semantic conventions.
const (
	// AttrFlagKey is the flag a check was asked about.
	AttrFlagKey = attribute.Key("schematic.flag.key")
	// AttrFlagKeys is the set of flags a bulk check was asked about.
	AttrFlagKeys = attribute.Key("schematic.flag.keys")
	// AttrFlagValue is the boolean the flag evaluated to.
	AttrFlagValue = attribute.Key("schematic.flag.value")
	// AttrFlagReason is the SDK's or the server's explanation of the verdict.
	AttrFlagReason = attribute.Key("schematic.flag.reason")
	// AttrFlagID is the flag's Schematic ID, when the check resolved one.
	AttrFlagID = attribute.Key("schematic.flag.id")

	// AttrCheckSource says what answered a check: the local cache, DataStream,
	// the API, or a configured default. It is the attribute to reach for first
	// when a check returns something unexpected.
	AttrCheckSource = attribute.Key("schematic.check.source")
	// AttrCheckAllowed is Check's verdict, which differs from the flag value on
	// the credit paths.
	AttrCheckAllowed = attribute.Key("schematic.check.allowed")
	// AttrCheckFailOpen records that the caller asked to be allowed through
	// when the check itself could not be completed.
	AttrCheckFailOpen = attribute.Key("schematic.check.fail_open")

	// AttrCompanyID is the company the check resolved to.
	AttrCompanyID = attribute.Key("schematic.company.id")
	// AttrUserID is the user the check resolved to.
	AttrUserID = attribute.Key("schematic.user.id")
	// AttrCompanyKeyNames and AttrUserKeyNames name the lookup dimensions the
	// caller passed. Only the names: the values are the caller's own
	// identifiers and may carry personal data, so they are never recorded.
	AttrCompanyKeyNames = attribute.Key("schematic.company.key_names")
	AttrUserKeyNames    = attribute.Key("schematic.user.key_names")

	// AttrEventType is "identify" or "track".
	AttrEventType = attribute.Key("schematic.event.type")
	// AttrEventSubtype is the caller's name for a tracked event.
	AttrEventSubtype = attribute.Key("schematic.event.subtype")

	// AttrUsageQuantity is the units of a feature a check was asked to hold
	// credits for.
	AttrUsageQuantity = attribute.Key("schematic.usage.quantity")
	// AttrLeaseMode is where the credit hold lives: "client" or "server".
	AttrLeaseMode = attribute.Key("schematic.lease.mode")
	// AttrReservationID is the credit hold a check took, when it took one.
	AttrReservationID = attribute.Key("schematic.reservation.id")
	// AttrCreditTypeID is the credit type a hold draws down.
	AttrCreditTypeID  = attribute.Key("schematic.credit.type_id")
	AttrCreditTypeIDs = attribute.Key("schematic.credit.type_ids")
	// AttrCreditsReserved is the credits a hold covers.
	AttrCreditsReserved = attribute.Key("schematic.credit.reserved")
	// AttrActualQuantity is the usage a reservation was settled at.
	AttrActualQuantity = attribute.Key("schematic.usage.actual_quantity")
	// AttrLeaseID is the lease a wire call acted on.
	AttrLeaseID = attribute.Key("schematic.lease.id")
	// AttrLeaseRequested is the credit amount an acquire or extend asked for.
	AttrLeaseRequested = attribute.Key("schematic.lease.requested")
	// AttrLeaseGranted is the lease total the server came back with, which an
	// extend reports as the new total rather than the increment.
	AttrLeaseGranted = attribute.Key("schematic.lease.granted")

	// AttrOffline records that the client is in offline mode, which answers
	// every check from defaults without reaching the API.
	AttrOffline = attribute.Key("schematic.offline")
)

// Values for AttrCheckSource.
const (
	SourceOffline    = "offline"
	SourceCache      = "cache"
	SourceDataStream = "datastream"
	SourceAPI        = "api"
	SourceDefault    = "default"
)

// KeyNames returns the sorted names of a lookup key map, for recording which
// dimensions a caller identified a company or user by.
//
// Only the names. The values are the caller's own identifiers — an email, an
// internal account ID — so recording them would put personal data in a trace
// that the caller never asked to send anywhere. The resolved Schematic ID goes
// on the span instead, which identifies the company just as well for debugging
// without carrying anything of the caller's.
func KeyNames(keys map[string]string) string {
	if len(keys) == 0 {
		return ""
	}
	names := make([]string, 0, len(keys))
	for name := range keys {
		names = append(names, name)
	}
	sort.Strings(names)
	return strings.Join(names, ",")
}
