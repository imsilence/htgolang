package main

import (
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {

	// 计算字节切片sha512散列值
	fmt.Printf("%x\n", sha512.Sum512([]byte("Hi,kk")))

	// 创建sha256 hash对象
	hash := sha512.New()

	// 批量写入字符串计算散列1
	io.WriteString(hash, "Hi,")
	io.WriteString(hash, "kk")

	// 计算sha512散列
	fmt.Printf("%x\n", hash.Sum(nil))

}
