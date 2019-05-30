package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer 01")
	}()
	defer func() {
		fmt.Println("defer 02")
	}()
	defer func() {
		fmt.Println("defer 03")
	}()

	fmt.Println("over")
}
