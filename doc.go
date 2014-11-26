// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

/*
Package rddscript implements Reddcoin transaction scripts.

A complete description of the script language used by Reddcoin can be found at
https://en.Reddcoin.it/wiki/Script.  The following only serves as a quick
overview to provide information on how to use the package.

This package provides data structures and functions to parse and execute
Reddcoin transaction scripts.

Script Overview

Bitcoin transaction scripts are written in a stack-base, FORTH-like language.

The Reddcoin script language consists of a number of opcodes which fall into
several categories such pushing and popping data to and from the stack,
performing basic and bitwise arithmetic, conditional branching, comparing
hashes, and checking cryptographic signatures.  Scripts are processed from left
to right and intentionally do not provide loops.

The vast majority of Bitcoin scripts at the time of this writing are of several
standard forms which consist of a spender providing a public key and a signature
which proves the spender owns the associated private key.  This information
is used to prove the the spender is authorized to perform the transaction.

One benefit of using a scripting language is added flexibility in specifying
what conditions must be met in order to spend Reddcoins.

Errors

Errors returned by this package are of the form rddscript.ErrStackX where X
indicates the specific error.  See Variables in the package documentation for a
full list.
*/
package rddscript
