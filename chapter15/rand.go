package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(101))
	fmt.Println(rand.Float64())

	letters := "0123456789abcdefhijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	buffer := bytes.Buffer{}
	for i := 0; i < 10; i++ {
		buffer.WriteByte(letters[rand.Intn(len(letters))])
	}
	fmt.Println(buffer.String())
}
