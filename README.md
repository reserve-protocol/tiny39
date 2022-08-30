# Tiny39: Very Tiny BIP39 Mnemonic Generator

This is a tiny go script that just generates a single, 256-bit bip39 mnemonic from a cryptographic randomness source. It's a tiny executable wrapper around the `github.com/tyler-smith/go-bip39` package.

If you're looking for a CLI tool with more than one feature, you probably want the bip39 tool over here: https://github.com/kubetrail/bip39

# Installation
Installation and usage are pretty simple though. Assuming no familiarity with standard module handling in `go`:

1. Have the [go compiler toolchain](https://go.dev/dl/) installed. You can check this by running `go version`.
2. Run `go install github.com/reserve-protocol/tiny39@latest`.

# Usage

Just `~/go/bin/tiny39`. It takes no options.

Or, if you've tinkered with your go defaults before and you no longer know where your binary , run `$(go env GOPATH)/bin/tiny39`.

Output will look something like:

    MNEMONIC="pizza cereal violin genuine prosper donor word marine desk upgrade slam rebel shop tribe wide robust tragic motor mixture system plug naive gain exile"


