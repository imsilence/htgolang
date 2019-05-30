package main

import (
	"fmt"
)

func main() {
	// 定义两个带缓冲区的管道
	ch01 := make(chan int, 1)
	ch02 := make(chan int, 1)

	// 启用例程像两个管道中写入数据
	go func() {
		ch01 <- 0
	}()

	go func() {
		ch02 <- 0
	}()

	// 使用select从两个管道中读取数据
	select {
	case <-ch01:
		fmt.Println("ch01")
	case <-ch02:
		fmt.Println("ch02")
	default:
		fmt.Println("default")
	}
}
