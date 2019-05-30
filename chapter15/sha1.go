package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {

	// 计算字节切片sha1散列值
	fmt.Printf("%x\n", sha1.Sum([]byte("Hi,kk")))

	// 创建sha1 hash对象
	hash := sha1.New()

	// 批量写入字符串计算散列
	io.WriteString(hash, "Hi,")
	io.WriteString(hash, "kk")

	// 计算sha1散列
	fmt.Printf("%x\n", hash.Sum(nil))
}
