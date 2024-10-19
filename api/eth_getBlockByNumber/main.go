package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://arbitrum-sepolia.blockpi.network/v1/rpc/public")
	if err != nil {
		log.Printf("Failed to connect: %v", err)
		return
	}

	blockNumber := big.NewInt(69914586)
	blockMap, err := getArbitrumBlock(client, blockNumber)
	if err != nil {
		log.Printf("Failed to get block: %v", err)
		return
	}

	txs, ok := blockMap["transactions"].([]interface{})
	if !ok {
		log.Println("Transactions is not a slice of interfaces")
		return
	}

	targetTxHash := "0xe8a514a9f530d4250725db8691e2f1568b11fe9bee737faa7489002a2a61059f"

	for _, tx := range txs {
		if tx == nil {
			continue
		}
		txm, ok := tx.(map[string]interface{})
		if !ok {
			log.Println("Transaction is not a map")
			continue
		}

		if txm["hash"] != targetTxHash {
			continue
		}

		fmt.Println("BlockHash:", blockMap["hash"])

		fmt.Println("TxHash:", txm["hash"])
		fmt.Println("From:", txm["from"])
		fmt.Println("To:", txm["to"])
		fmt.Println("BlockNum:", hexStringToInt64(blockMap["number"]))
		fmt.Println("BlockTime:", hexStringToInt64(blockMap["timestamp"]))
		fmt.Println("Gas:", hexStringToInt64(txm["gas"]))
		fmt.Println("GasPrice:", hexStringToInt64(txm["gasPrice"]))
		fmt.Println("Value:", hexStringToInt64(txm["value"]))
		fmt.Println("Op:", txm["type"])
		fmt.Println("ChainId:", hexStringToInt64(txm["chainId"]))
	}
}

// getArbitrumBlock gets the block at the given number
func getArbitrumBlock(client *ethclient.Client, number *big.Int) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := client.Client().CallContext(context.Background(), &result, "eth_getBlockByNumber", fmt.Sprintf("0x%x", number.Int64()), true)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// hexStringToInt64 converts a hex string to an int64
func hexStringToInt64(hexInput interface{}) int64 {
	hexStr := hexInput.(string)
	intVal, err := strconv.ParseInt(hexStr[2:], 16, 64)
	if err != nil {
		return 0
	}
	return intVal
}
