# gozec
![Gozec_logo](docs/assets/gozec.png)

An idiomatic Go wrapper around a Rust FFI for interacting with a Zcash wallet.

## Overview
- **Project:** `gozec` — a small Go library that calls into a Rust FFI (in `ffi/`) to manage a Zcash wallet.
- **Use case:** programmatic wallet creation, address retrieval, syncing, balance checks and sending transactions from Go.

## Requirements
- **Go:** `go 1.25.4` (see `go.mod`).
- **Rust toolchain:** to build the FFI library in `ffi/` (the Go code links against `-lrust_ffi_go`).
- **CGo:** enabled (bundled into `go build` when CGo is available on your system).

## Installation / Build

1. Prepare FFI bindings. you can either

- **Build from source** [zcash rust ffi library](https://github.com/steel-feel/zcash_rust_ffi) and put inside the `ffi/` folder at home directory or where ever you like.

#### OR

- **Download from here** - [libzcash_rust_ffi.dylib](https://ljmu-my.sharepoint.com/:u:/g/personal/uplhjai2_ljmu_ac_uk/IQB7-_q_wP5sSJlke0UcYGtnAefQTVMzDHQXI9vaAvZ4MrQ)


2. Set below enviorment variables

```bash
export GOZEC_FFI=<path/to/folder/of/rust_ffi_go[.dylib/so]>
export CGO_LDFLAGS="-L$GOZEC_FFI -lrust_ffi_go"
```

3. Build or test the Go package from the repository root:

```bash
go test ./...
go build ./...
```

Note: cgo will use the `#cgo LDFLAGS` specified in `gozec.go` to link the Rust library from `ffi/`.

**Quick Usage Example**

```go
package main

import (
	"fmt"
	"github.com/steel-feel/gozec"
)

func main() {
	// Initialize or create a wallet at `./walletdir` on Mainnet
	w, err := gozec.Init("./walletdir", gozec.Mainnet)
	if err != nil {
		panic(err)
	}

	// Sync the wallet (calls into the ffi)
	w.Sync()

	// Get addresses
	addr := w.GetAddress()
	fmt.Println("t-address:", addr.tAddress)
	fmt.Println("u-address:", addr.uAddress)

	// Get balances
	bal := w.GetBalance()
	fmt.Printf("height=%s total=%d shielded=%d unshielded=%d\n", bal.height, bal.total, bal.shielded, bal.unshielded)

	// Send transaction (example)
	// txid := w.SendTransaction("u1...", 1000)
	// fmt.Println("txid:", txid)
}
```

**API Reference (high level)**
- `Init(walletDir string, networkType NetworkType) (*GozecWallet, error)`: Initialize or create a wallet directory and return a `*GozecWallet`.
- `(*GozecWallet).GetAddress() ZecAddress`: Returns `tAddress` (transparent) and `uAddress` (unified) addresses for the account.
- `(*GozecWallet).Sync()`: Triggers a sync operation via the FFI.
- `(*GozecWallet).GetBalance() ZecBalance`: Returns a `ZecBalance` struct with `height`, `shielded`, `unshielded`, and `total` fields.
- `(*GozecWallet).SendTransaction(to string, value uint64) string`: Sends `value` to `to` (address) and returns a string (FFI-returned result, e.g. tx id or error message).

Types
- `NetworkType` constants: `Testnet`, `Mainnet`.
- `ZecAccount`, `ZecAddress`, `ZecBalance`, and `GozecWallet` are defined in `gozec.go` for use by callers.

**Testing**
- Unit tests in `gozec_test.go` expect a wallet path such as `../hello` (adjust as needed).
- Some tests require the FFI and/or a prepared wallet directory and cannot run in CI without the FFI and network access.

**Notes & Caveats**
- This package wraps a Rust FFI. You must ensure the Rust library is built and discoverable in the `ffi/` folder before running or building the Go package.
- `SendTransaction` requires on-chain interaction and a funded wallet; test coverage for sending is limited in the included tests.
- The FFI function signatures (in `gozec.go`) allocate C strings and structures — the Go wrapper frees them where appropriate, but be careful when changing FFI signatures.

**Contributing**
- Fork the repo, make your changes, and open a pull request with a clear description.
- If you modify the Rust side, include build instructions and ensure library artifacts are produced under `ffi/` or that `#cgo` paths are updated.

**License**
- This repository includes a `LICENSE` file — follow its terms.

**Contact / Author**
- Repository: `github.com/steel-feel/gozec`

If you'd like, I can also add a small example program under `examples/` or add a CI job that builds the Rust FFI and runs the Go tests — tell me which you'd prefer.
