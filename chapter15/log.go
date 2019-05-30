package main

import (
	"log"
	"os"
	"bufio"
)

func main() {
	log.Println("Line 1")

	log.SetPrefix("log:")
	log.Println("Line 2")

	log.Println(log.Flags())
	log.SetFlags(log.Flags() | log.Lshortfile)
	log.Println("Line 3")

	file, _ := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE, os.ModePerm)
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	log.SetOutput(writer)
	log.Println("Line 4")
	log.Println("Line 5")
	log.Println("Line 6")

	log.Output(1, "Line 7")
	log.Output(2, "Line 8")
	log.Output(3, "Line 9")

	file2, _ := os.OpenFile("log2.log", os.O_APPEND|os.O_CREATE, os.ModePerm)
	defer file2.Close()
	writer2 := bufio.NewWriter(file2)
	defer writer2.Flush()

	
	logger := log.New(writer2, "log2:", log.LstdFlags|log.Lshortfile)

	logger.Println("Line 1")
	logger.Println("Line 2")
	logger.Println("Line 3")

}
