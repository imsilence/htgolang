package main

import "fmt"

func main() {
	// 初始化管道
	channel := make(chan rune)

	go func() {
		for c := 'A'; c <= 'Z'; c++ {
			channel <- c
		}
	}()
	for c := range channel {
		fmt.Println(c)
	}
}
