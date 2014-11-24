// Copyright (c) 2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rddscript_test

import (
	"encoding/hex"
	"fmt"

	"github.com/reddcoin-project/rddnet"
	"github.com/reddcoin-project/rddscript"
	"github.com/reddcoin-project/rddutil"
)

// This example demonstrates creating a script which pays to a Reddcoin address.
// It also prints the created script hex and uses the DisasmString function to
// display the disassembled script.
func ExamplePayToAddrScript() {
	// Parse the address to send the coins to into a rddutil.Address
	// which is useful to ensure the accuracy of the address and determine
	// the address type.  It is also required for the upcoming call to
	// PayToAddrScript.
	addressStr := "RkQDYcqiv7mzQfNYMc8FfYv3dtQ8wuSGoM"
	address, err := rddutil.DecodeAddress(addressStr, &rddnet.MainNetParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a public key script that pays to the address.
	script, err := rddscript.PayToAddrScript(address)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Script Hex: %x\n", script)

	disasm, err := rddscript.DisasmString(script)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Script Disassembly:", disasm)

	// Output:
	// Script Hex: 76a914814089fb909f05918d54e530f0ad8e339a4edffe88ac
	// Script Disassembly: OP_DUP OP_HASH160 814089fb909f05918d54e530f0ad8e339a4edffe OP_EQUALVERIFY OP_CHECKSIG
}

// This example demonstrates extracting information from a standard public key
// script.
func ExampleExtractPkScriptAddrs() {
	// Start with a standard pay-to-pubkey-hash script.
	scriptHex := "76a914814089fb909f05918d54e530f0ad8e339a4edffe88ac"
	script, err := hex.DecodeString(scriptHex)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Extract and print details from the script.
	scriptClass, addresses, reqSigs, err := rddscript.ExtractPkScriptAddrs(
		script, &rddnet.MainNetParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Script Class:", scriptClass)
	fmt.Println("Addresses:", addresses)
	fmt.Println("Required Signatures:", reqSigs)

	// Output:
	// Script Class: pubkeyhash
	// Addresses: [RkQDYcqiv7mzQfNYMc8FfYv3dtQ8wuSGoM]
	// Required Signatures: 1
}
