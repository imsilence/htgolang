package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	// 读取文件所有内容
	if file, err := os.Open("test01.txt"); err == nil {
		defer file.Close()
		content, _ := ioutil.ReadAll(file)
		fmt.Println(string(content))
	}

	// 读取文件目录
	if ffs, err := ioutil.ReadDir("."); err == nil {
		for _, fs := range ffs {
			fmt.Println(fs.Name(), fs.IsDir(), fs.Size(), fs.ModTime(), fs.Mode())
		}
	}

	// 读取文件内容
	if content, err := ioutil.ReadFile("test01.txt"); err == nil {
		fmt.Println(string(content))
	}

	// 写入文件
	ioutil.WriteFile("test01.txt", []byte("吾日三省吾身。为人谋而不忠乎？与朋友交而不信乎？传不习乎？"), 0600)

	// 创建临时文件夹
	fmt.Println(ioutil.TempDir("", "test"))

	// 创建临时文件
	tempfile, _ := ioutil.TempFile("", "test")
	defer tempfile.Close()
	fmt.Println(tempfile.Name())
	io.WriteString(tempfile, "吾日三省吾身。为人谋而不忠乎？与朋友交而不信乎？传不习乎？")

}
