# Credit lease conformance suite

`SPEC.md` and `vectors/*.json` are copied verbatim from `conformance/` on `main`
in [schematic-node](https://github.com/SchematicHQ/schematic-node), the
reference implementation. Do not edit them here: change them there, then copy
the new versions across, so every SDK runs the same contract.

The runner is the only language-specific piece. This SDK's lives in
`leases/conformance_test.go`.
