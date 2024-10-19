package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"go-ethereum/contracts/store" // for demo
)

// 使用 ethclient.Dial 连接到以太坊 Sepolia 测试网的 RPC 节点。
// 如果连接失败，则记录错误并终止程序。
// 通过合约地址和以太坊客户端创建一个 store.Store 实例，代表智能合约 Store。
// 如果创建实例失败，则记录错误并终止程序。
// 输出 "contract is loaded" 表示合约已成功加载，并未使用 instance 变量 (_ = instance)。

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

	fmt.Println("contract is loaded")
	_ = instance
}
