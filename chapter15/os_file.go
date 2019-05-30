package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Mkdir("test/test01/test02", os.ModePerm))
	fmt.Println(os.Mkdir("test02", os.ModePerm))
	fmt.Println(os.MkdirAll("test/test01/test02", os.ModePerm))

	fmt.Println(os.Rename("test", "test03"))
	fmt.Println(os.Rename("test02", "test04"))
	time.Sleep(3 * time.Second)

	fmt.Println(os.Remove("test04"))
	fmt.Println(os.Remove("test03"))
	fmt.Println(os.RemoveAll("test03"))

	// 写文件(文件不存在创建)
	file01, err := os.Create("test01.txt")
	if err == nil {
		defer file01.Close()

		//写文件
		file01.WriteString("hi, i am kk")
	}

	// 读文件(文件不存在创建)
	file02, err := os.Create("test01.txt")
	if err == nil {
		defer file02.Close()
		// 读文件
		cxt := make([]byte, 20)
		for {
			n, err := file02.Read(cxt)
			if err != io.EOF {
				fmt.Println(n, string(cxt[:n]))
			} else {
				break
			}
		}
	}

	// 读写文件(文件不存在创建)
	file03, err := os.Create("test03.txt")
	if err == nil {
		defer file03.Close()
		//写文件
		file03.WriteString("hi, i am kk")

		// 读文件
		cxt := make([]byte, 20)
		for {
			n, err := file03.Read(cxt)
			if err != io.EOF {
				fmt.Println(n, string(cxt[:n]))
			} else {
				break
			}
		}

		// 设置文件指针后读文件
		file03.Seek(0, os.SEEK_SET)
		for {
			n, err := file03.Read(cxt)
			if err != io.EOF {
				fmt.Println(n, string(cxt[:n]))
			} else {
				break
			}
		}
	}

	// 读写文件(文件不存在报错)
	file04, err := os.Open("test04.txt")
	if err == nil {
		defer file04.Close()
		file04.WriteString("hi2, i am kk")
	} else {
		fmt.Println(err)
	}

	// 追加文件
	file05, err := os.OpenFile("test05.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err == nil {
		defer file05.Close()
		file05.WriteString("hi3, i am kk")
	}

	// 只读文件（文件不存在报错）
	file06, err := os.OpenFile("test06.txt", os.O_RDONLY, os.ModePerm)
	if err == nil {
		defer file06.Close()
		cxt := make([]byte, 3)
		for {
			n, err := file06.Read(cxt)
			if err != io.EOF {
				fmt.Println(n, string(cxt[:n]))
			} else {
				break
			}
		}
	} else {
		fmt.Println(err)
	}

	// 读取目录所有文件名
	dir01, err := os.Open("e:/codes")
	if err == nil {
		defer dir01.Close()
		if info, err := dir01.Stat(); err == nil && info.IsDir() {
			fmt.Println(dir01.Readdirnames(-1))
		}
	} else {
		fmt.Println(err)
	}

	// 读取目录所有文件信息
	dir02, err := os.Open("e:/codes")
	if err == nil {
		defer dir02.Close()
		if info, err := dir02.Stat(); err == nil && info.IsDir() {
			if cinfos, err := dir02.Readdir(-1); err == nil {
				for _, cinfo := range cinfos {
					fmt.Println(cinfo.Name(), cinfo.IsDir(), cinfo.ModTime(), cinfo.Size(), cinfo.Mode())
				}
			}
		}
	} else {
		fmt.Println(err)
	}
}
