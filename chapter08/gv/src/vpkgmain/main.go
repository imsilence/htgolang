package main

import (
	"fmt"

	"vpkgname/pkg01"
	"vpkgname/pkg02"
)

func main() {
	fmt.Println("vpkgmain")
	fmt.Println(pkg01.Name)
	fmt.Println(pkg02.Name)
}
