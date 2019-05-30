package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	sreader := strings.NewReader("Hi, kk\n")
	io.Copy(os.Stdout, sreader)

	breader := bytes.NewReader([]byte("Hi, kk\n"))
	io.Copy(os.Stdout, breader)

	bbuffer := bytes.NewBufferString("Hi, kk\n")
	io.CopyN(os.Stdout, bbuffer, 2)
	fmt.Println()
	io.CopyN(os.Stdout, bbuffer, 2)
	fmt.Println()
	io.CopyN(os.Stdout, bbuffer, 2)
	fmt.Println()

	io.WriteString(os.Stdout, "Hi, kk\n")
}
