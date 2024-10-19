package main

import (
	"fmt"
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

/*
初始化一个助记词 mnemonic。
通过助记词创建一个新的 HD 钱包 wallet。
解析一个导出路径 path 用于生成特定的钱包账户。
根据路径从钱包中派生出一个账户 account。
打印账户的地址。
获取并打印账户的私钥（十六进制格式）。
获取并打印账户的公钥（十六进制格式）。
*/
func main() {
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Account address: %s\n", account.Address.Hex())

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Private key in hex: %s\n", privateKey)

	publicKey, _ := wallet.PublicKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Public key in hex: %s\n", publicKey)
}
