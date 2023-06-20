package main

import (
	"fmt"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
)

func main() {
	source := "SDICFI722MCPAZAN42HNDBUF3DYQKKP5D46MPBARYNK4N5HG72QUDFYK"
	destination := "GB4XYPCKCX5FKKNFYTG5DEQPNQ354A6WERXDYBWIV4OJSHM4ZU2FRRXI"
	client := horizonclient.DefaultTestNetClient

	// Make sure destination account exists
	destAccountRequest := horizonclient.AccountRequest{AccountID: destination}
	destinationAccount, err := client.AccountDetail(destAccountRequest)
	if err != nil {
		panic(err)
	}

	fmt.Println("Destination Account", destinationAccount)

	// Load the source account
	sourceKP := keypair.MustParseFull(source)
	sourceAccountRequest := horizonclient.AccountRequest{AccountID: sourceKP.Address()}
	sourceAccount, err := client.AccountDetail(sourceAccountRequest)
	if err != nil {
		panic(err)
	}

	// Build transaction
	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &sourceAccount,
			IncrementSequenceNum: true,
			BaseFee:              txnbuild.MinBaseFee,
			Preconditions: txnbuild.Preconditions{
				TimeBounds: txnbuild.NewInfiniteTimeout(), // Use a real timeout in production!
			},
			Operations: []txnbuild.Operation{
				&txnbuild.Payment{
					Destination: destination,
					Amount:      "10",
					Asset:       txnbuild.NativeAsset{},
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}

	// Sign the transaction to prove you are actually the person sending it.
	tx, err = tx.Sign(network.TestNetworkPassphrase, sourceKP)
	if err != nil {
		panic(err)
	}

	// And finally, send it off to Stellar!
	resp, err := horizonclient.DefaultTestNetClient.SubmitTransaction(tx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successful Transaction:")
	fmt.Println("Ledger:", resp.Ledger)
	fmt.Println("Hash:", resp.Hash)
}
