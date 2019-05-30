package main

import (
	"fmt"
	"time"
)

func main() {
	var ticker *time.Ticker = time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	fmt.Println(time.Now())
	fmt.Println(<- ticker.C)
	fmt.Println(<- ticker.C)

	for {
		select {
		case now := <- ticker.C:
			fmt.Println(now)
		}
	}
}
