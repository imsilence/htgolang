package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 定义counter由 incr和decr两个函数分别进行递增和递减
var counter int

// 对counter递增1000次
func incr(group *sync.WaitGroup) {
	defer group.Done()
	for i := 0; i < 1000; i++ {
		counter++
		runtime.Gosched()
	}
}

// 对counter激减1000次
func decr(group *sync.WaitGroup) {
	defer group.Done()
	for i := 0; i < 1000; i++ {
		counter--
		runtime.Gosched()
	}
}

func main() {
	group := &sync.WaitGroup{}

	//例程递增和递减
	for i := 0; i < 5; i++ {
		group.Add(2)
		go incr(group)
		go decr(group)
	}
	group.Wait()
	fmt.Println(counter)
}
