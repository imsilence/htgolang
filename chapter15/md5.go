package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {

	// 计算字节切片md5散列值
	fmt.Printf("%x\n", md5.Sum([]byte("Hi,kk")))

	// 创建md5 hash对象
	hash := md5.New()

	// 批量写入字符串计算散列
	io.WriteString(hash, "Hi,")
	io.WriteString(hash, "kk")

	// 计算MD5散列
	fmt.Printf("%x\n", hash.Sum(nil))

}
