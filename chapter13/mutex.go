package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 定义counter由 incr和decr两个函数分别进行递增和递减
var counter int = 0
// 加锁,对counter递增1000次
func incr_mutex(lock *sync.Mutex, group *sync.WaitGroup) {
	lock.Lock() // 加锁
	defer func() {
		lock.Unlock() //释放锁
		group.Done()
	}()
	for i := 0; i < 1000; i++ {
		counter++
		runtime.Gosched()
	}
}
// 加锁,对counter激减1000次
func decr_mutex(lock *sync.Mutex, group *sync.WaitGroup) {
	lock.Lock() // 加锁
	defer func() {
		lock.Unlock() //释放锁
		group.Done()
	}()
	for i := 0; i < 1000; i++ {
		counter--
		runtime.Gosched()
	}
}
func main() {
	group := &sync.WaitGroup{}
	// 定义锁
	lock := &sync.Mutex{}
	//例程递增和递减, 加锁
	for i := 0; i < 5; i++ {
		group.Add(2)
		go incr_mutex(lock, group)
		go decr_mutex(lock, group)
	}
	group.Wait()
	fmt.Println(counter)
}
