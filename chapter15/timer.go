package main

import (
	"fmt"
	"time"
)

func main() {
	var timer *time.Timer = time.NewTimer(3 * time.Second)
	defer timer.Stop()

	fmt.Println(time.Now())
	fmt.Println(<- timer.C)

	timer.Reset(3 * time.Second)
	fmt.Println(<- timer.C)

	timer.Reset(time.Second)

	for {
		select {
		case now := <- timer.C:
			fmt.Println(now)
			timer.Reset(time.Second * 2)
		}
	}

}
