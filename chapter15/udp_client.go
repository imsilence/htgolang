package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr := ":8889"

	// 创建UDP连接
	conn, err := net.Dial("udp", addr)

	defer conn.Close()

	// 向服务器发送数据
	conn.Write([]byte(fmt.Sprintf("UnixTime: %d", time.Now().Unix())))

	cxt := make([]byte, 1024)

	// 从服务器读取数据
	n, err := conn.Read(cxt)
	if err == nil {
		fmt.Println(string(cxt[:n]))
	}
}
