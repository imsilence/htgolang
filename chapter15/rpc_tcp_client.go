package main

import (
	"fmt"
	"net/rpc"
	"rpc_object"
)

func main() {

	// 创建rpc客户端
	client, err := rpc.Dial("tcp", "localhost:8888")
	defer client.Close()
	fmt.Println(err)

	req := rpc_object.Request{8, 3}
	var resp rpc_object.Response

	var methods []string = []string{"calc.Sum", "calc.Difference", "calc.Product", "calc.Quotient"}

	for _, method := range methods {
		fmt.Println(method, ":")
		// 异步调用rpc服务, 通过注册服务器调用方法
		call := client.Go(method, &req, &resp, nil)
		<-call.Done
		if err == nil {
			fmt.Println(resp.Result)
		} else {
			fmt.Println(err)
		}
	}
}
