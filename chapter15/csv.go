package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {

	// 写入csv文件
	file, err := os.Create("user.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		writer := csv.NewWriter(file)
		writer.Write([]string{"name", "addr"})
		writer.Write([]string{"kk", "西安"})
		writer.Write([]string{"woniu", "北京"})
		writer.Flush()
		file.Close()
	}

	// 读取csv文件
	file, err = os.Open("user.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		reader := csv.NewReader(file)
		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(line)
		}
		file.Close()
	}
}
