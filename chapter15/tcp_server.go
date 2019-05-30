package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	addr := ":8888"

	// 启动监听服务
	server, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// 延迟关闭服务
	defer server.Close()
	fmt.Printf("Listen on: %s\n", server.Addr())

	for {
		// 获取客户端连接
		conn, err := server.Accept()
		if err == nil {

			// 使用例程处理与客户端连接
			go func(conn net.Conn) {
				defer conn.Close() //延迟关闭客户端连接

				fmt.Printf("client is connected: %s\n", conn.RemoteAddr())

				// 读取客户端发送的消息
				reader := bufio.NewReader(conn)
				cxt, _, _ := reader.ReadLine()
				fmt.Println(string(cxt))

				//向客户端发送消息
				fmt.Fprintf(conn, "Time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
			}(conn)
		}
	}
}
