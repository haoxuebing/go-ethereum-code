package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// 将字符串形式的大数 r 和 s 转换为 big.Int 类型，并转换为十六进制字符串
// 拼接 r、s 的十六进制字符串及 v，形成签名
// 从给定十六进制字符串获取哈希值
// 使用 Ecrecover 函数从签名和哈希值恢复公钥
// 输出恢复后的公钥和地址
// 移除签名中的恢复标识符
// 验证签名是否正确

// 如果r 和 s 的长度不是64位，前面补充0
func main() {
	r := "93706950132453901348378792405885520616079580121126560497173953355583730590215"
	s := "19267692409323608135360436114343820840749991697725077116478667028585874280021"
	v := "01"
	bigR := new(big.Int)
	bigR.SetString(r, 10)
	rhash := bigR.Text(16)
	fmt.Println("rhash:", rhash)

	bigS := new(big.Int)
	bigS.SetString(s, 10)
	shash := bigS.Text(16)
	fmt.Println("shash:", shash)

	signature := common.FromHex(fmt.Sprintf("0x%s%s%s", rhash, shash, v))

	hash := common.FromHex("0xfac794d2c35fdcf5ac77bc4d7a77af7d3e0278b221cfc04039edd93e2189f963")

	sigPublicKey, err := crypto.Ecrecover(hash, signature)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("recoverePubkey: ", hexutil.Encode(sigPublicKey)) // common.Bytes2Hex(sigPublicKey)

	pubKey, _ := crypto.UnmarshalPubkey(sigPublicKey)
	fmt.Println("recoveredPubkey: ", pubKey)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	fmt.Println("recoveredAddr ", recoveredAddr.String())

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	// VerifySignature函数，该函数接收字节格式的公钥、哈希值和原始数据的签名
	verified := crypto.VerifySignature(sigPublicKey, hash, signatureNoRecoverID)
	fmt.Println(verified) // true
}
