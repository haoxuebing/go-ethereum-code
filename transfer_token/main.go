package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/sha3"
)

// 初始化RPC客户端：通过ethclient.Dial方法连接到Arbitrum的RPC节点。
// 加载私钥：使用crypto.HexToECDSA方法加载一个以太坊私钥。
// 获取公钥：通过privateKey.Public()方法从私钥生成公钥，并断言类型为*ecdsa.PublicKey。
// 计算地址：使用crypto.PubkeyToAddress方法从公钥计算以太坊地址。
// 获取nonce：通过client.PendingNonceAt方法获取下一个可用的交易nonce。
// 设置交易价值：使用big.NewInt(0)设置交易金额为0。
// 获取燃气价格：通过client.SuggestGasPrice方法获取当前建议的燃气价格。
// 设置接收地址和代币地址：分别使用common.HexToAddress方法解析接收地址和代币地址。
// 构建交易数据：根据ERC20代币的transfer函数签名和参数构建交易数据。
// 估算燃气限制：通过client.EstimateGas方法估算交易所需的燃气量。
// 创建和发送交易：使用types.NewTransaction创建交易对象，通过types.SignTx方法签名交易，并通过client.SendTransaction方法发送交易。
// 最终，代码将打印出发送的交易哈希。

func main() {
	// 连接到Arbitrum Sepolia的RPC节点
	// 0xeF20B71f4C9a9036450E01D5361917D6C97262c7 转1个USDC到 0x4C34899060680556B0B20Fafa9D433be25615b15
	// Arbitrum Sepolia的USDC合约地址是 0x75faf114eafb1bdbe2f0316df893fd58ce46aa4d

	// 加载 .env 文件
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 读取环境变量
	PRIVATE_KEY := os.Getenv("PRIVATE_KEY")

	client, err := ethclient.Dial("https://sepolia-rollup.arbitrum.io/rpc")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// 打印公钥信息
	fmt.Println(hexutil.Encode(crypto.FromECDSAPub(&privateKey.PublicKey)))

	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4C34899060680556B0B20Fafa9D433be25615b15")
	tokenAddress := common.HexToAddress("0x75faf114eafb1bdbe2f0316df893fd58ce46aa4d")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amount := new(big.Int)
	amount.SetString("1000000", 10) // sets the value to 1000 tokens, in the token denomination

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 估算所需的燃气量
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}
