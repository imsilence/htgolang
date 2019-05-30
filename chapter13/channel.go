package main

import (
	"fmt"
	"runtime"
)

func printChars(prefix string, channel chan string) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s:%c\n", prefix, ch)
		runtime.Gosched()
	}
	channel <- prefix //像管道写入数据
}

func main() {
	// 初始化管道
	channel := make(chan string)
	n := 10
	// 创建例程
	for i := 0; i < n; i++ {
		go printChars(fmt.Sprintf("gochars%02d", i), channel)
	}

	for i := 0; i < n; i++ {
		tag := <-channel //从管道读取数据
		fmt.Println("over: ", tag)
	}
	fmt.Println("over")
}
