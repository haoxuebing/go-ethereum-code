package main

import (
	"fmt"
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

// 通过助记词生成seed 或 生成一个新的随机种子 seed。
// 从种子创建一个 HD (Hierarchical Deterministic) 钱包实例 wallet。
// 根据 BIP-44 路径 m/44'/60'/0'/0/0 衍生第一个账户，并打印其地址（十六进制格式）。
// 根据 BIP-44 路径 m/44'/60'/0'/0/1 衍生第二个账户，并打印其地址（十六进制格式）。

func main() {
	// 通过助记词生成seed
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	seed, err := hdwallet.NewSeedFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	// 生成随机seed
	// seed, _ := hdwallet.NewSeed()

	// 通过seed生成钱包
	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559
}
