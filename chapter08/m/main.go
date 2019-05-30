package main

import (
	"fmt"

	"github.com/imsilence/chapter08/modtest/pkgname/pkg01"
	"github.com/imsilence/chapter08/modtest/pkgname/pkg02"
)

func main() {
	fmt.Println("module main")
	fmt.Println(pkg01.Name)
	fmt.Println(pkg02.Name)
}
