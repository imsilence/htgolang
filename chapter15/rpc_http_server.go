package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"rpc_object"
)

func main() {
	// 注册RPC服务
	rpc.Register(&rpc_object.Calculate{})
	rpc.HandleHTTP() // 使用HTTP服务

	server, err := net.Listen("tcp", ":8888") // 监听服务
	if err == nil {
		defer server.Close()
		http.Serve(server, nil) //启动http服务
	} else {
		fmt.Println(err)
	}
}
