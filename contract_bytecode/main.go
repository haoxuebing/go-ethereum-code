package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// https://sepolia.etherscan.io/address/0xA4c46d29370884D58A2a1e6890628C8F616B761b#code
func main() {
	client, err := ethclient.Dial("https://rpc.sepolia.org")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xA4c46d29370884D58A2a1e6890628C8F616B761b")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode)) // 60806...a0033
}
