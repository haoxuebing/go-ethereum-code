package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"go-ethereum/contracts/store"
)

func main() {
	client, err := ethclient.Dial("https://rpc.sepolia.org")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xA4c46d29370884D58A2a1e6890628C8F616B761b")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}
