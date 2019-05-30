package main

import (
	"fmt"
	"time"
)

func main() {
	var timeout <-chan time.Time = time.After(3 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timeout)
}
