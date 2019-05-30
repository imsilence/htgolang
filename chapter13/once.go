package main

import (
	"fmt"
	"sync"
)

func main() {
	once := &sync.Once{}

	for i := 0; i <= 10; i++ {
		func() {
			fmt.Println("func:", i)
		}()

		once.Do(func() {
			fmt.Println("once:", i)
		})

	}

}
