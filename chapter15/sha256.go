package main

import (
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {

	// 计算字节切片sha256散列值
	fmt.Printf("%x\n", sha256.Sum256([]byte("Hi,kk")))

	// 创建sha256 hash对象
	hash := sha256.New()

	// 批量写入字符串计算散列
	io.WriteString(hash, "Hi,")
	io.WriteString(hash, "kk")

	// 计算sha256散列
	fmt.Printf("%x\n", hash.Sum(nil))

}
