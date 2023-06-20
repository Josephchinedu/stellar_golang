package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
)

func createAccount() {
	pair, err := keypair.Random()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(pair.Seed())
	// SAV76USXIJOBMEQXPANUOQM6F5LIOTLPDIDVRJBFFE2MDJXG24TAPUU7
	log.Println(pair.Address())
	// GCFXHS4GXL6BVUCXBWXGTITROWLVYXQKQLF4YH5O5JT3YZXCYPAFBJZB
}

func fundAccount(address string) {

	resp, err := http.Get("https://friendbot.stellar.org/?addr=" + address)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func getAccountBalance(address string) {

	request := horizonclient.AccountRequest{AccountID: address}

	account, err := horizonclient.DefaultTestNetClient.AccountDetail(request)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Balances for account:", address)

	for _, balance := range account.Balances {
		log.Println(balance)
	}
}

func main() {
	// fundAccount("GB4XYPCKCX5FKKNFYTG5DEQPNQ354A6WERXDYBWIV4OJSHM4ZU2FRRXI")

	getAccountBalance("GB4XYPCKCX5FKKNFYTG5DEQPNQ354A6WERXDYBWIV4OJSHM4ZU2FRRXI")

	// SDICFI722MCPAZAN42HNDBUF3DYQKKP5D46MPBARYNK4N5HG72QUDFYK seed

	// GDJ4F3Z6VYPAGBJUHH2IMN6DHO2SARFMT3FWTYAN2WG5FWE3GRHNIFUO

	// FIRST WALLET

	// SECOND WALLET
	// createAccount()

	// SCA32MUQWCMDHUL27533CHT4CO4EDLEWDQRC5B6QIHSSU2XOMVLQYKIS seed

	// GB4XYPCKCX5FKKNFYTG5DEQPNQ354A6WERXDYBWIV4OJSHM4ZU2FRRXI address

}
