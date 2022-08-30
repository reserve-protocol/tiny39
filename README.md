# Tiny39: Very Tiny BIP39 Mnemonic Generator

This is a tiny go script that just generates a single, 256-bit bip39 mnemonic from a cryptographic randomness source. It's a tiny executable wrapper around the `github.com/tyler-smith/go-bip39` package.

If you're looking for a CLI tool with more than one feature, you probably want the bip39 tool over here: https://github.com/kubetrail/bip39

# Installation and usage
Installation and usage are pretty simple though. Assuming no familiarity with standard module handling in `go`:

1. Have the [go compiler toolchain](https://go.dev/dl/) installed. You can check this by running `go version`.
2. Run `go install github.com/reserve-protocol/tiny39@latest`.
3. Run `~/go/bin/tiny39`  (or, if you've tinkered with your go configuration before, `$(go env GOPATH)/bin/tiny39`)


