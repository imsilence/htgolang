package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println(runtime.GOROOT())       //获取Go安装路径
	fmt.Println(runtime.NumCPU())       // 获取可使用CPU逻辑核数
	fmt.Println(runtime.GOMAXPROCS(1))  // 设置可使用CPU逻辑核数
	fmt.Println(runtime.NumGoroutine()) //获取例程数量
	group := &sync.WaitGroup{}
	group.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(10)
			group.Done()
		}()
	}
	fmt.Println(runtime.NumGoroutine()) //获取例程数量
	group.Wait()

	fmt.Println(runtime.NumGoroutine()) //获取例程数量
}
