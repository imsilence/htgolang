package main

import "fmt"

type TypeName Formatter

// 自定义int类型
type Counter int

// 自定义map[string]string类型
type User map[string]string

// 自定义函数类型
type Callback func(...string)

/**
* 定义接收自定义类型(函数类型)为参数的函数
 */
func printResult(pf Callback, list ...string) {
	pf(list...)
}

func column(list ...string) {
	for _, e := range list {
		fmt.Println(e)
	}
	fmt.Println()
}

func main() {

	// 使用自定义结构Counter
	var counter Counter

	fmt.Printf("%T:%v", counter, counter)
	counter++
	fmt.Println(counter)

	//使用自定义结构体User
	var me User = make(User)
	me["name"] = "kk"
	me["addr"] = "西安"

	fmt.Printf("%T: %#v\n", me, me)

	names := []string{"kk", "woniu", "ada"}

	//将函数column传递给printResult的pf(Callback类型)参数
	printResult(column, names...)
}
