package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	addr := ":8889"

	// 启动服务监听
	server, err := net.ListenPacket("udp", addr)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer server.Close()

	cxt := make([]byte, 1024)
	for {
		// 从客户端读取数据
		n, addr, err := server.ReadFrom(cxt)
		if err == nil {
			fmt.Println(addr, string(cxt[:n]))

			// 向客户端发送数据
			server.WriteTo([]byte(fmt.Sprintf("Time: %s", time.Now().Format("2006-01-02 15:04:05"))), addr)
		}
	}
}
