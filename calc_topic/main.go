package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
)

// 定义一个transferFnSignature字节数组，包含字符串"Transfer(address,address,uint256)"，用于表示一个函数签名。
// 创建一个sha3 LegacyKeccak256哈希对象。
// 将transferFnSignature写入哈希对象中。
// 打印出哈希对象的十六进制编码结果。
// 将哈希结果的前4个字节赋值给methodID。
// 打印出methodID的十六进制编码结果，即0xddf252ad。
func main() {
	// 计算ERC20转账事件topic的示例代码
	transferFnSignature := []byte("Transfer(address,address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	fmt.Println(hexutil.Encode(hash.Sum(nil)))
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xddf252ad
}
