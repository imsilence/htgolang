package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	// 定义超时管道和执行结果管道
	timeout := make(chan int)
	result := make(chan int)

	// 启用超时例程
	go func() {
		time.Sleep(5 * time.Second) //等待5s结束
		timeout <- 0
	}()

	// 启动工作例程
	go func() {
		interval := time.Duration(rand.Int31n(10)) * time.Second
		fmt.Println("sleep:", interval)
		time.Sleep(interval)
		result <- 0
	}()

	// 使用select从两个管道中读取数据
	select {
	case <-result:
		fmt.Println("ok")
	case <-timeout:
		fmt.Println("timeout")
	}
}
