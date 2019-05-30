package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {

	key := []byte("1234567890abcdefg")
	// 创建hmac hash对象
	hash := hmac.New(sha256.New, key)

	// 写入字符串计算散列
	io.WriteString(hash, "Hi,kk")

	// 计算hmac散列
	fmt.Printf("%x\n", hash.Sum(nil))

	hash2 := hmac.New(sha256.New, key)
	hash2.Write([]byte("Hi,kk"))
	fmt.Println(hmac.Equal(hash2.Sum(nil), hash.Sum(nil)))

}
