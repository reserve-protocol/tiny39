package main

import (
	"fmt"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, err := bip39.NewMnemonic(entropy)

	// Display mnemonic or error
	if err == nil {
		fmt.Printf("MNEMONIC=\"%v\"\n", mnemonic)
	} else {
		fmt.Println("building mnemonic:", err)
	}
}
