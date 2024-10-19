package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"go-ethereum/contracts/store" // for demo
)

// 连接到以太坊网络：使用ethclient.Dial方法连接到以太坊网络。
// 加载私钥：使用crypto.HexToECDSA方法将十六进制私钥加载为ECDSA私钥。
// 获取公钥：通过私钥生成公钥，并断言公钥类型为*ecdsa.PublicKey。
// 获取发送地址：使用公钥生成以太坊地址。
// 获取nonce：通过客户端获取发送地址的nonce。
// 获取建议的gas价格：通过客户端获取网络建议的gas价格。
// 设置链ID：创建一个大整数表示链ID。
// 创建交易对象：使用bind.NewKeyedTransactorWithChainID方法创建一个具有链ID的交易对象。
// 部署合约：使用交易对象和客户端部署一个合约，并获取合约地址、交易和合约实例。
// 打印结果：打印合约地址和交易哈希。
func main() {
	// 加载 .env 文件
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 读取环境变量
	PRIVATE_KEY := os.Getenv("PRIVATE_KEY")

	client, err := ethclient.Dial("https://ethereum-sepolia.blockpi.network/v1/rpc/public")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID := big.NewInt(11155111)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0xA4c46d29370884D58A2a1e6890628C8F616B761b
	fmt.Println(tx.Hash().Hex()) // 0x46119f95564dc0e5747e87f7cd4a9abae9b6fcca804d28bf16544e0297a8b449

	_ = instance
}
