// Package wasm holds the compiled Schematic rules engine and embeds it into the
// binary at build time.
//
// The .wasm file is built from the Rust rules engine in the schematic-api repo
// (rulesengine-rust) and shared, byte for byte, with the C#, Python, Java, and
// Ruby SDKs. It is pinned by the WASM_VERSION file at the repo root and
// refreshed by scripts/download-wasm.sh.
package wasm

import _ "embed"

// Binary is the compiled rules engine module, loaded by the rulesengine package.
//
// This is committed to git rather than downloaded at build time, which is what
// every other Schematic SDK does. It has to be: `go get` serves a module's
// source tree directly from its git tag, so there is no packaging step in which
// a gitignored artifact could be added, and go:embed fails to compile outright
// if the file is absent. CI re-downloads the pinned asset and diffs it against
// this file, so it is verified rather than trusted.
//
//go:embed rulesengine.wasm
var Binary []byte
