package main

import "fmt"

// 定义结构体Dog
type Dog struct {
	name string
}

// 为结构体Dog定义方法Call
func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

// 为结构体Dog定义方法SetName
func (dog Dog) SetName(name string) {
	dog.name = name
}

// 为结构体Dog定义方法PSetName(接收者为Dog的指针对象)
func (dog *Dog) PSetName(name string) {
	dog.name = name
}

// 为结构体Dog定义方法PCall(接收者为Dog的指针对象)
func (dog *Dog) PCall() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

func main() {
	// 初始化结构体对象
	dog := Dog{"豆豆"}

	// 调用结构体对象Call方法
	dog.Call()

	//调用结构体对象SetName方法
	dog.SetName("大黄")

	dog.Call() // 这里结果是什么???为什么???改怎么做

	//调用结构体指针对象的PSetName方法
	(&dog).PSetName("大黄")

	dog.Call()

	dog2 := &Dog{"大毛"}

	// 调用结构体对象的Call方法
	(*dog2).Call()

	// 使用结构体对象调用指针接收者的PSetName方法
	dog.PSetName("小黄")

	dog.Call()

	dog2.PSetName("二毛")

	// 使用结构体指针对象调用值接收者的Call方法
	dog2.Call()

	dog2.SetName("三毛")
	dog2.Call() //???

	var dog3 *Dog
	dog3.Call()

	// 方法值
	methodv01 := dog.PSetName
	methodv02 := dog.Call
	methodv03 := dog2.PSetName
	methodv04 := dog2.Call

	fmt.Printf("%T, %T, %T, %T\n", methodv01, methodv02, methodv03, methodv04)

	//通过方法值调用方法
	fmt.Printf("%p %p\n", &dog, dog2)
	methodv01("乐乐1")
	methodv02() //???
	methodv03("乐乐2")
	methodv04() //???

	fmt.Println(dog.name, dog2.name)

	methodv05 := dog.PCall
	methodv06 := dog2.PCall

	methodv05()
	methodv06()

	// 方法表达式
	methode01 := (*Dog).PSetName
	methode02 := Dog.Call

	fmt.Printf("%T, %T\n", methode01, methode02)

	//通过方法表达式指定对象调用方法
	methode01(&dog, "贝贝1") // (*Dog).PSetName(&dog, "贝贝1")
	methode01(dog2, "贝贝2") // (*Dog).PSetName(dog2, "贝贝2")

	methode02(dog)   //Dog.Call(dog)
	methode02(*dog2) //Dog.Call(*dog2)

	// 取引用
	// methode03 := Dog.PSetName
	methode04 := (*Dog).Call

	fmt.Printf("%T\n", methode04)

	methode04(&dog)
	methode04(dog2)

}
