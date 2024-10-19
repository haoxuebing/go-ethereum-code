package main

import (
	"log"
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/miguelmota/go-ethereum-hdwallet"
)

/*
创建一个加密钱包并进行交易签名
1.使用给定的助记词（"tag volcano eight thank tide danger coast health above argue embrace heavy"）创建一个HD钱包。
2.解析派生路径为"m/44'/60'/0'/0/0"，用于派生钱包账户。
3.从HD钱包中派生指定路径的账户。
4.创建一个交易，设置交易的必要参数，如nonce（序号）、value（价值）、toAddress（接收地址）、gasLimit（燃气限制）、gasPrice（燃气价格）和data（附加数据）。
5.使用HD钱包和账户签署交易。
6.使用spew.Dump输出签名的交易信息。
*/
func main() {
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Fatal(err)
	}

	nonce := uint64(0)
	value := big.NewInt(1000000000000000000)
	toAddress := common.HexToAddress("0x0")
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(21000000000)
	var data []byte

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := wallet.SignTx(account, tx, nil)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(signedTx)
}