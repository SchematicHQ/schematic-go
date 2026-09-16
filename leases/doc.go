// Package leases implements client-mode credit leases: the SDK draws a tranche
// of credits from the server per (company, credit type), carves reservations
// out of it locally at check time, and settles them at track time.
//
// The semantics are shared across the Schematic SDKs and pinned by the
// language-agnostic vectors in conformance/ at the repo root. The Redis
// backends use the same key layout, hash fields and Lua scripts as the Node and
// Python SDKs, so pods from all three can share one Redis.
package leases
