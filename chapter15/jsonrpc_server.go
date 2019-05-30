package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc_object"
)

func main() {
	// 注册RPC服务
	rpc.Register(&rpc_object.Calculate{})
	server, err := net.Listen("tcp", ":8888")
	defer server.Close()
	if err == nil {
		for {
			client, err := server.Accept() // 监听客户端连接
			if err == nil {
				go jsonrpc.ServeConn(client) //处理客户端请求
			}
		}
	} else {
		fmt.Println(err)
	}
}
