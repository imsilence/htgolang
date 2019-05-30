package main

import "fmt"
import "errors"

/**
* 定义无参无返回值的函数
* 在控制台答应Hello World
 */
func sayHello() {
	fmt.Println("Hello World")
}

/**
* 定义一个参数且无返回值的函数
* 根据传递参数在控制台打印Hi, name
 */
func sayHi(name string) {
	fmt.Printf("Hi, %s\n", name)
}

/**
* 定义两个参数且有一个返回值的函数
* 计算传递的两个参数的和并返回
 */
func add(n1 int, n2 int) int {
	return n1 + n2
}

/**
* 合并相同类型参数类型名
 */
func mergeFuncArgsType(n1, n2 int, s1, s2, s3 string, b1 bool) {
	fmt.Printf("%T, %T, %T, %T, %T, %T\n", n1, n2, s1, s2, s3, b1)
	fmt.Println(n1, n2, s1, s2, s3, b1)
}

/**
* 定义可变参数列表函数, 至少有2个参数
* 打印所有参数到控制台
 */
func printArgs(n1, n2 int, args ...string) {
	fmt.Printf("%T, %T, %T\n", n1, n2, args)
	fmt.Println(n1, n2, args)
}

/**
* 定义有多个返回值的函数
* 计算两个参数的四则运算结果并返回
 */
func calc(n1, n2 int) (int, int, int, int) {
	return n1 + n2, n1 - n2, n1 * n2, n1 / n2
}

/**
* 定义命名返回值函数
 */
func calcReturnNamecalc(n1, n2 int) (sum, difference, product, quotient int) {
	sum, difference, product, quotient = n1+n2, n1-n2, n1*n2, n1/n2
	return
}

/**
* 计算n的阶乘
* n < 0:错误
* n = 0:1
* n > 0: n! = n * n-1!
 */
func factorial(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}

}

/**
* 汉罗塔游戏
* 将所有a柱上的的圆盘借助b柱移动到c柱, 在移动过程中保证每个柱子的上面圆盘比下面圆盘小
* n: a -> c(b)
* n=1: a->c
* n>1: n-1 (a -> b(c)); a -> c; n-1(b -> c(a))
 */
func tower(a, b, c string, layer int) {
	if layer <= 0 {
		return
	}
	if layer == 1 {
		fmt.Printf("%s -> %s\n", a, c)
		return
	}
	tower(a, c, b, layer-1)
	fmt.Printf("%s -> %s\n", a, c)
	tower(b, a, c, layer-1)
}

/**
* 定义接收函数类型作为参数的函数
 */
func printResult(pf func(...string), list ...string) {
	pf(list...)
}

func line(list ...string) {
	fmt.Print("|")
	for _, e := range list {
		fmt.Print(e)
		fmt.Print("\t|")
	}
	fmt.Println()
}

func column(list ...string) {
	for _, e := range list {
		fmt.Println(e)
	}
	fmt.Println()
}

/**
* 定义闭包函数, 返回一个匿名函数用于计算与base元素的和
 */
func addBase(base int) func(int) int {
	return func(num int) int {
		return base + num
	}
}

/*
* 定义除法函数, 若除数为0则使用error返回错误信息
 */
func division(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, errors.New("除数为0")
	}
	return n1 / n2, nil
}

func main() {
	//调用无参无返回值函数
	sayHello()

	//调用有参无返回值函数
	sayHi("KK")
	sayHi("woniu")

	//调用有参有返回值函数
	n1, n2 := 1, 2
	fmt.Printf("%d + %d = %d\n", n1, n2, add(n1, n2))

	n3 := add(4, 5)
	fmt.Println(n3)

	//忽略函数返回值
	add(3, 4)

	mergeFuncArgsType(1, 2, "3", "4", "5", true)

	//调用可变参数函数
	printArgs(1, 2)
	printArgs(1, 2, "3", "4", "5")

	//通过切片解包并调用可变参数函数
	printArgs(1, 2, []string{"3", "4", "5", "6"}...)
	args := []string{"3", "4", "5", "6", "8"}
	printArgs(1, 2, args...)
	printArgs(1, 2, args[:3]...)

	fmt.Println(calc(5, 2))

	fmt.Println(calcReturnNamecalc(5, 2))

	fmt.Println(factorial(5))

	fmt.Println("layer 1:")
	tower("A", "B", "C", 1)

	fmt.Println("layer 2:")
	tower("A", "B", "C", 2)

	fmt.Println("layer 3:")
	tower("A", "B", "C", 3)

	// 定义函数类型变量, 并使用零值nil进行初始化
	var callback func(int, int) (int, int, int, int)
	fmt.Printf("%T, %v", callback, callback)

	callback = calc             // 赋值为函数calc
	fmt.Println(callback(5, 2)) // 调用calc函数

	callback = calcReturnNamecalc // 赋值为函数calcReturnNamecalc
	fmt.Println(callback(8, 2))   // 赋值为函数calcReturnNamecalc

	// 调用参数类型为函数的函数
	names := []string{"kk", "woniu", "ada"}
	printResult(line, names...)
	printResult(column, names...)

	// 定义匿名函数并赋值给hi
	hi := func(name string) {
		fmt.Printf("Hi, %s\n", name)
	}

	//调用匿名函数hi
	hi("kk")
	hi("woniu")

	// 定义匿名函数并进行调用
	func() {
		fmt.Println("我是匿名函数, 我在工作")
	}()

	// 使用匿名函数作为printResult的参数
	printResult(func(list ...string) {
		for i, v := range list {
			fmt.Printf("%d: %s\n", i, v)
		}
	}, names...)

	//使用闭包函数
	base2 := addBase(2)
	base10 := addBase(10)

	fmt.Println(base2(1), base10(1))
	fmt.Println(base2(5), base10(5))
	fmt.Println(base2(10), base10(10))

	e1, e2 := "kk", []string{"kk", "silence"}

	// 值传递
	// 在函数内修改值类型
	fmt.Printf("e1: %p %v\n", &e1, e1)
	func(e string) {
		fmt.Printf("e: %p %v\n", &e, e)
		e = "silence"
	}(e1)

	// 在函数内修改引用类型
	fmt.Printf("e2: %p %v\n", &e2, e2)
	func(e []string) {
		fmt.Printf("e: %p %v\n", &e, e)
		e[1] = "woniu"
	}(e2)

	fmt.Println(e1, e2)

	// 引用传递
	// 在函数内修改值类型
	fmt.Printf("e1: %p %v\n", &e1, e1)
	func(e *string) {
		fmt.Printf("e: %p %v\n", e, *e)
		*e = "silence"
	}(&e1)

	// 在函数内修改引用类型
	fmt.Printf("e2: %p %v\n", &e2, e2)
	func(e *[]string) {
		fmt.Printf("e: %p %v\n", e, *e)
		(*e)[1] = "woniu"
	}(&e2)

	fmt.Println(e1, e2)

	// 处理函数返回的错误
	for _, v := range [...]int{0, 3} {
		if r, err := division(6, v); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(r)
		}
	}

	err1, err2 := errors.New("error: 1"), fmt.Errorf("error: %d", 2)
	fmt.Printf("%T, %T, %v, %v", err1, err2, err1, err2)
}
