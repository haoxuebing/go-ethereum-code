package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// 一个生成AK/SK 的函数，用于web 服务端，jwt 接口校验

// 生成随机字符串作为 Access Key
func generateAccessKey(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	ak := make([]byte, length)
	for i := range ak {
		ak[i] = charset[rand.Intn(len(charset))]
	}
	return string(ak)
}

// 生成 HMAC-SHA256 哈希作为 Secret Key
func generateSecretKey(accessKey string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(accessKey))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	accessKey := generateAccessKey(32) // 生成长度为 16 的随机字符串作为 Access Key
	secret := "wisdom-uae"             // 使用的密钥种子，可以是一个固定的字符串或随机生成
	secretKey := generateSecretKey(accessKey, secret)

	fmt.Println("Access Key:", accessKey)
	fmt.Println("Secret Key:", secretKey)
}
