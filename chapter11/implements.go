package main

import "fmt"

// 定义Sender接口
type Sender interface {
	Send(msg string) error
}

// 定义Sender接口
type Receiver interface {
	Receive() (string, error)
}

// 定义Connection接口,由Sender和Receiver组合
type Connection interface {
	Sender
	Receiver
	Open() error
	Close() error
}

//定义TCPConnection结构体，并实现Connection接口中Open、Send、Receive、Close方法
type TCPConnction struct {
	addr string
	port int
}

func (conn TCPConnction) Open() error {
	fmt.Println("打开")
	return nil
}

func (conn TCPConnction) Send(msg string) error {
	fmt.Println("发送消息")
	return nil
}

func (conn TCPConnction) Receive() (string, error) {
	fmt.Println("接收消息")
	return "", nil
}

func (conn TCPConnction) Close() error {
	fmt.Println("关闭")
	return nil
}

func main() {
	var conn Connection = TCPConnction{"127.0.0.1", 8888}
	conn.Open()
	conn.Send("Hi")
	msg, _ := conn.Receive()
	fmt.Printf("%q\n", msg)
	conn.Close()

	// 通过匿名接口声明接口变量
	var closer interface {
		Close() error
	}
	fmt.Printf("%T, %#v\n", closer, closer)
	closer = conn
	closer.Close()

}
