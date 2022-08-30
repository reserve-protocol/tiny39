# Simple BIP-39 Mnemonic Generator

This is just a tiny go script that just generates a single, 256-bit bip39 mnemonic from cryptographic randomness. There's very little here and that's on purpose. All the real work is being done by the `github.com/tyler-smith/go-bip39` package.

Installation and usage, though, are dead simple, and just the result of standard module handling in `go`:

1. Have the [go compiler toolchain](https://go.dev/dl/) installed. Check this by running `go version`.
2. From this directory, run `go run .`

Or, alternately, you can do `go install` and then run `tiny39` out of your go binary path. If you haven't modified `GOPATH` or `GOBIN`, that'll be `~/go/bin/tiny39`. However, if you're installing something for long-term use, you probably instead want the fully-featured, properly packaged, bip39 command-line tool, over here: https://github.com/kubetrail/bip39

