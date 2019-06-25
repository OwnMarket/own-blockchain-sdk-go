# own-blockchain-sdk-go

Own Blockchain SDK for Go

## Quick Start

```bash
$ git clone https://github.com/OwnMarket/own-blockchain-sdk-go.git
$ cd own-blockchain-sdk-go
```

Run tests:
```bash
$ go test -v
```

## Usage

Fetch Own Blockchain SDK for Go package

```bash
$ go get github.com/OwnMarket/own-blockchain-sdk-go
```

Use the package in Go code

```bash
package main

import (
	"fmt"

	ownSdk "github.com/OwnMarket/own-blockchain-sdk-go"
)

func main() {
	// Create a new wallet
	wallet := ownSdk.GenerateWallet()

	// Compose a transaction with nonce = 1 and actionFee = 0.01
	tx := ownSdk.CreateTx(wallet.Address, 1, 0.01, 0)
	tx.AddTransferChxAction("CHxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", 100) // Transfer 100 CHX to CHxxx... address.

	// Look at the raw transaction in JSON format
	fmt.Println(tx.ToJson(true))

	// Sign the transaction for submission to node API on TestNet
	networkCode := []byte("OWN_PUBLIC_BLOCKCHAIN_TESTNET")
	signedTx := tx.Sign(networkCode, wallet.PrivateKey)
	fmt.Println(signedTx.Signature)
}
```
