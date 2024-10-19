package main

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

func main() {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Error("failed GenerateKey with %s.", err)
	}

	fmt.Println("private key have 0x   \n", hexutil.Encode(crypto.FromECDSA(key)))
	fmt.Println("private key no 0x \n", hex.EncodeToString(crypto.FromECDSA(key)))

	if err := crypto.SaveECDSA("privatekey", key); err != nil {
		log.Error(fmt.Sprintf("Failed to persist node key: %v", err))
	}

	fmt.Println("public key have 0x   \n", hexutil.Encode(crypto.FromECDSAPub(&key.PublicKey)))
	fmt.Println("public key no 0x \n", hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey)))

	// 由私钥字符串转换私钥
	acc1Key, _ := crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	address1 := crypto.PubkeyToAddress(acc1Key.PublicKey)
	fmt.Println("address ", address1.String())

	dummyAddr := common.HexToAddress("9b2055d370f73ec7d8a03e965129118dc8f5bf83")
	fmt.Println("dummyAddr", dummyAddr.String())

	// 字节转地址
	addr3 := common.BytesToAddress([]byte("ethereum"))
	fmt.Println("address ", addr3.String())

	// 字节转hash
	hash1 := common.BytesToHash([]byte("topic1"))
	fmt.Println("hash ", hash1.String())

	testAddrHex := "970e8128ab834e8eac17ab8e3812f010678cf791"
	addrtest := common.HexToAddress(testAddrHex)

	testPrivHex := "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
	key1, _ := crypto.HexToECDSA(testPrivHex)

	// 计算给定数据的 Keccak-256 哈希值
	msg := crypto.Keccak256([]byte("foo"))
	sig, err := crypto.Sign(msg, key1)
	if err != nil {
		log.Error("failed Sign with %s.", err)
	}

	// 推出公钥字节
	recoveredPub, err := crypto.Ecrecover(msg, sig)
	if err != nil {
		log.Error("failed Ecrecover with %s.", err)
	}
	// 字节转 公钥
	pubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	// 摘要和签名推出公钥
	recoveredPub2, _ := crypto.SigToPub(msg, sig)
	recoveredAddr2 := crypto.PubkeyToAddress(*recoveredPub2)

	fmt.Println("addrtest ", addrtest.String())
	fmt.Println("recoveredAddr ", recoveredAddr.String())
	fmt.Println("recoveredAddr2 ", recoveredAddr2.String())
}
