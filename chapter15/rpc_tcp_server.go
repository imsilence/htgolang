package main

import (
	"net"
	"net/rpc"
	"rpc_object"
)

func main() {
	// 注册RPC服务 并设置注册名
	rpc.RegisterName("calc", &rpc_object.Calculate{})
	server, err := net.Listen("tcp", ":8888") // 监听服务
	if err == nil {
		defer server.Close()
		rpc.Accept(server) //监听客户端请求并处理
	}
}
