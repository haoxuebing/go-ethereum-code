package main

import (
	"fmt"
	"log"

	"github.com/miguelmota/go-ethereum-hdwallet"
)

/*
使用助记词生成私钥和地址
*/

/*
CLI
go install github.com/miguelmota/go-ethereum-hdwallet/cmd/geth-hdwallet@latest

$ geth-hdwallet -mnemonic "tag volcano eight thank tide danger coast health above argue embrace heavy" -path "m/44'/60'/0'/0/0"
*/


func main() {
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	// Derive the first account of the wallet
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947

	// Derive the second account of the wallet
	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559
}