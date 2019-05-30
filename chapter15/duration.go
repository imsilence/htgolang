package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	time.Sleep(3)
	fmt.Println(time.Since(start))
	fmt.Println(time.Until(start))

	duration, _ := time.ParseDuration("1.5h20m30s100ms")
	fmt.Println(duration.Nanoseconds())
	fmt.Println(duration.Seconds())
	fmt.Println(duration.Minutes())
	fmt.Println(duration.Hours())
	fmt.Println(duration.String())


}
