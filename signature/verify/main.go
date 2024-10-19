package main

import (
    "bytes"
    "crypto/ecdsa"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
)

// 将十六进制私钥转换为ECDSA私钥。
// 提取与ECDSA私钥对应的公钥。
// 将公钥转换为字节切片。
// 对字符串"hello"进行Keccak-256哈希运算，并打印哈希值（十六进制格式）。
// 使用ECDSA私钥对哈希进行签名，并打印签名值（十六进制格式）。
// 使用哈希和签名恢复公钥，并检查它与原始公钥是否匹配。
// 使用哈希、签名和公钥进行验证，确保签名是有效的。
// 移除签名中的恢复ID，并再次验证签名的有效性。
func main() {
    privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

    data := []byte("hello")
    hash := crypto.Keccak256Hash(data)
    fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

    signature, err := crypto.Sign(hash.Bytes(), privateKey)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301

	// 调用Ecrecover（椭圆曲线签名恢复）来检索签名者的公钥
    sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
    if err != nil {
        log.Fatal(err)
    }

    matches := bytes.Equal(sigPublicKey, publicKeyBytes)
    fmt.Println(matches) // true

    sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
    if err != nil {
        log.Fatal(err)
    }

    sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
    matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
    fmt.Println(matches) // true

    signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	// VerifySignature函数，该函数接收字节格式的公钥、哈希值和原始数据的签名
    verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
    fmt.Println(verified) // true
}