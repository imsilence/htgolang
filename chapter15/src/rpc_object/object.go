package rpc_object

import "errors"

// 定义计算结构体，用于定义远程过程方法
type Calculate struct {
}

// 定义请求参数结构体
type Request struct {
	Left  int
	Right int
}

// 定义响应结果结构体
type Response struct {
	Result int
}

// 定义远程方法用于计算请求参数中的和
// 第一个参数为请求参数指针
// 第二个参数为响应结果指针
// 返回值为error类型
func (c *Calculate) Sum(request *Request, response *Response) error {
	response.Result = request.Left + request.Right
	return nil
}

// 定义远程方法用于计算请求参数中的差
func (c *Calculate) Difference(request *Request, response *Response) error {
	response.Result = request.Left - request.Right
	return nil
}

// 定义远程方法用于计算请求参数中的积
func (c *Calculate) Product(request *Request, response *Response) error {
	response.Result = request.Left * request.Right
	return nil
}

// 定义远程方法用于计算请求参数中的商
func (c *Calculate) Quotient(request *Request, response *Response) error {
	if request.Right == 0 {
		return errors.New("除数为0")
	}
	response.Result = request.Left / request.Right
	return nil
}
