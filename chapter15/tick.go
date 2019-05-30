package main

import (
	"fmt"
	"time"
)

func main() {
	var ticker <-chan time.Time = time.Tick(3 * time.Second)
	fmt.Println(time.Now())
	for now := range ticker {
		fmt.Println(now)
	}
}
