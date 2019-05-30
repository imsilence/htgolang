package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// 定义counter由 incr和decr两个函数分别进行递增和递减
var counter *int32

// 原子操做,对counter递增1000次
func incr_atomic(group *sync.WaitGroup) {

	for i := 0; i < 1000; i++ {
		atomic.AddInt32(counter, 1)
		runtime.Gosched()
	}
	group.Done()
}

// 原子操作,对counter激减1000次
func decr_atomic(group *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(counter, -1)
		runtime.Gosched()
	}
	group.Done()
}
func main() {

	var value int32 = 0
	counter = &value

	group := &sync.WaitGroup{}
	//例程递增和递减
	for i := 0; i < 5; i++ {
		group.Add(2)
		go incr_atomic(group)
		go decr_atomic(group)
	}
	group.Wait()
	fmt.Println(value)
}
