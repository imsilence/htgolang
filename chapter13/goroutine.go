package main

import (
	"fmt"
	"runtime"
)

func printChars(prefix string) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s:%c\n", prefix, ch)
		// 让出CPU
		// time.Sleep(1 * time.Microsecond)
		runtime.Gosched()
	}
}

func main() {
	// 创建例程
	go printChars("gochars01")
	go printChars("gochars02")

	printChars("gochars03")

	// //休眠3s 等待例程运行结束
	// time.Sleep(3000 * time.Millisecond)
}
