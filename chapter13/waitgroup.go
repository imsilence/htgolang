package main

import (
	"fmt"
	"runtime"
	"sync"
)

func printChars(prefix string, group *sync.WaitGroup) {
	defer group.Done() // 通过信号量通知执行结束
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s:%c\n", prefix, ch)
		runtime.Gosched()
	}
}

func main() {
	// 定义信号量
	group := &sync.WaitGroup{}
	n := 10

	group.Add(n)

	// 创建例程
	for i := 0; i < n; i++ {
		go printChars(fmt.Sprintf("gochars%02d", i), group)
	}

	group.Wait() // 等待所有例程执行结束
	fmt.Println("over")
}
