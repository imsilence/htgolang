package main

import (
	"fmt"
)

func main() {
	var channel chan int
	fmt.Printf("%T %v", channel, channel)

	channel = make(chan int)

	channel02 := make(chan string, 10)
	fmt.Printf("%T %v %d\n", channel, channel, len(channel))
	fmt.Printf("%T %v %d\n", channel02, channel02, len(channel02))

	channel02 <- "1"
	channel02 <- "2"
	fmt.Printf("%T %v %d\n", channel02, channel02, len(channel02))

	fmt.Println(<-channel02)
	fmt.Printf("%T %v %d\n", channel02, channel02, len(channel02))

	close(channel02)
	for {
		e, ok := <-channel02
		if !ok {
			fmt.Println("over")
			break
		} else {
			fmt.Println(e)
		}
	}

	channel03 := make(chan int)

	go func() {
		for e := range channel03 {
			fmt.Println(e)
		}
		channel <- 0
	}()

	go func() {
		for i := 0; i < 5; i++ {
			channel03 <- i
		}
		close(channel03)
	}()

	<-channel

	channel04 := make(chan rune, 10)

	// 声明例程管道只写
	go func(cl chan<- rune) {
		for c := 'A'; c <= 'G'; c++ {
			cl <- c
		}
		close(cl)

	}(channel04)

	// 声明例程管道只读
	go func(cl <-chan rune) {
		for c := range cl {
			fmt.Printf("%c\n", c)
		}
		channel <- 0
	}(channel04)

	<-channel

}
