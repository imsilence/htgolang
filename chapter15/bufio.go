package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 缓冲输入流
	sreader := strings.NewReader("Hi, kk\nHi, Silence")
	breader := bufio.NewReader(sreader)
	fmt.Println(breader.ReadLine())
	fmt.Println(breader.ReadLine())
	fmt.Println(breader.ReadLine())

	sreader.Seek(0, os.SEEK_SET)
	breader.WriteTo(os.Stdout)
	fmt.Println()

	// 定义扫描器
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// 缓冲输入流
	bbuffer := bytes.NewBufferString("")
	bwriter := bufio.NewWriter(bbuffer)
	bwriter.WriteString("Hi, kk\n")
	bwriter.Write([]byte("Hi, Silence"))
	bwriter.Flush()

	bbuffer.WriteTo(os.Stdout)
}
