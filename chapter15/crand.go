package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	rand_bytes := make([]byte, 20)
	// 生成随机字符串
	n, err := rand.Read(rand_bytes)
	fmt.Println(n, err)
	// base64编码
	fmt.Println(base64.StdEncoding.EncodeToString(rand_bytes))
}
