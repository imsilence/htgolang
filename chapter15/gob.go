package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Address struct {
	City   string
	Street string
	No     int
}

type User struct {
	ID       int
	Name     string
	Gender   bool
	Birthday time.Time
	Tel      string
	Addr     interface{}
}

func init() {
	// 向gob注册需要编码和解码的记录值类型(在使用interface{}定义类型并接受的变量)
	gob.Register(Address{})
	gob.RegisterName("User", User{})
}

func main() {
	var users []interface{} = make([]interface{}, 0, 10)
	users = append(users, User{
		ID:       1,
		Name:     "Silence",
		Gender:   true,
		Birthday: time.Date(1988, 10, 11, 0, 0, 0, 0, time.UTC),
		Tel:      "152********",
		Addr:     Address{"西安市", "锦业路", 1},
	})

	// gob编码User切片对象为二进制字符串并写入到文件
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(users); err == nil {
		ioutil.WriteFile("users.gob", buffer.Bytes(), os.ModePerm)
	} else {
		fmt.Println(err)
	}

	// 读取文件二进制内容并解码为User切片
	if file, err := os.Open("users.gob"); err == nil {
		reader := bufio.NewReader(file)
		decoder := gob.NewDecoder(reader)
		var us []interface{}
		if err := decoder.Decode(&us); err == nil {
			fmt.Println(us)
		} else {
			fmt.Println(err)
		}
	}
}
