package main

import (
	"fmt"
	"net/rpc/jsonrpc"
	"rpc_object"
)

func main() {
	// 创建rpc客户端
	client, err := jsonrpc.Dial("tcp", "localhost:8888")
	defer client.Close()
	fmt.Println(err)

	req := rpc_object.Request{7, 3}
	var resp rpc_object.Response

	var methods []string = []string{"Calculate.Sum", "Calculate.Difference", "Calculate.Product", "Calculate.Quotient"}

	for _, method := range methods {
		fmt.Println(method, ":")
		// 同步调用rpc服务
		err := client.Call(method, &req, &resp)
		if err == nil {
			fmt.Println(resp.Result)
		} else {
			fmt.Println(err)
		}
	}
}
