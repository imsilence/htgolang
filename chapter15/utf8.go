package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Println(utf8.RuneCountInString("abcdef我"))
	fmt.Println(utf8.RuneCountInString("abcdef"))
	fmt.Println(len("abcdef我"))
	fmt.Println(len("abcdef"))
}
