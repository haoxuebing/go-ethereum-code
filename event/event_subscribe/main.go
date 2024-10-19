package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 网络:sepolia
// 代币：USDC

// 使用ethclient.Dial连接到以太坊Sepolia测试网络的RPC节点。
// 定义需要监控的智能合约地址和查询过滤器。
// 创建一个接收日志的通道logs。
// 通过client.SubscribeFilterLogs方法订阅过滤日志，并将结果发送到logs通道中。
// 使用select语句监听错误和日志，若有错误则使用log.Fatal记录并终止程序，若接收到日志则打印日志内容。
func main() {
	client, err := ethclient.Dial("wss://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
