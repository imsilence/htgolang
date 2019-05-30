package main

import (
	"fmt"
)

// 定义结构体
type Author struct {
	id   int
	name string
	addr string
	tel  string
	desc string
}

// 定义工厂函数用户创建Author对象
func NewAuthor(id int, name, addr, tel, desc string) *Author {
	return &Author{id, name, addr, tel, desc}
}

// 定义地址结构体(区域, 街道, 门牌号)
type Address struct {
	region string
	street string
	no     string
}

// 定义收货人结构体(名字, 电话, 地址)
type User struct {
	name string
	tel  string
	addr Address
}

// 定义员工, 匿名嵌入User结构体
type Employee struct {
	User
	salary float64
	title  string
}

// 定义员工，匿名嵌入User结构体，并于User结构体具有相同的属性名
type Employee2 struct {
	User
	salary float64
	title  string
	addr   string
}

// 命名嵌入结构体指针
type PUser struct {
	name string
	tel  string
	addr *Address
}

// 匿名嵌入结构体指针
type PEmployee struct {
	*PUser
	salary float64
	title  string
}

func main() {

	var kk Author = Author{}

	fmt.Printf("%T: %#v\n", kk, kk)

	var pkk *Author

	fmt.Printf("%T: %#v\n", pkk, pkk)

	// 按照结构体属性定义顺序初始化变量
	var kk01 Author = Author{
		1001,
		"kk",
		"西安市",
		"15200000000",
		"少年经不得顺境，中年经不得闲境，晚年经不得逆境", //结尾逗号不能省略
	}

	// 只初始化一部分属性或不按照定义顺序初始化使用命名方式
	kk02 := Author{name: "kk2", id: 1002}

	fmt.Printf("%T, %#v\n", kk01, kk01)
	fmt.Printf("%T, %#v\n", kk02, kk02)

	//  使用new创建结构体结构体指针
	var kk03 *Author = new(Author)
	fmt.Printf("%T: %#v, %#v\n", kk03, kk03, *kk03)

	// 使用字面量初始化结构体指针
	kk04 := &Author{id: 1003, name: "kk4"}
	fmt.Printf("%T: %#v, %#v\n", kk04, kk04, *kk04)

	kk05 := NewAuthor(1004, "kk5", "西安市", "15300000000", "看山是山,看山不是山,看山还是山")
	fmt.Printf("%T: %#v\n", kk05, kk05)

	// 使用结构体对象访问和修改属性
	fmt.Printf("%d: %q\n", kk.id, kk.name)
	fmt.Printf("%d: %q\n", kk01.id, kk01.name)
	fmt.Printf("%d: %q\n", kk02.id, kk02.name)

	kk02.addr = "西安市"
	fmt.Println(kk02)

	// 使用结构体指针变量访问属性
	fmt.Println((*kk03).name)
	fmt.Println(kk04.name) //使用指针直接访问

	// 使用结构体指针变量修改属性
	(*kk03).name = "kk3"
	kk04.addr = "西安市" //使用指针直接修改
	fmt.Println(kk03, kk04)

	// 匿名结构体
	var coordinate struct {
		longitude float64
		latitude  float64
	}

	me := struct {
		name string
		age  int
		addr string
	}{"silence", 30, "西安"}

	// 匿名结构体数组
	studs := []struct {
		name string
		addr string
	}{{"silence", "西安"}, {"woniu", "北京"}}

	fmt.Println(coordinate, coordinate.longitude, coordinate.latitude)

	fmt.Println(me, me.name, me.age, me.addr)

	fmt.Println(studs)

	for i, stud := range studs {
		fmt.Println(i, stud, stud.name, stud.addr)
	}

	// 声明&初始化组合对象
	var u1 User

	fmt.Printf("%T, %#v\n", u1, u1)

	var a1 Address = Address{
		"陕西省西安市",
		"锦业路",
		"001",
	}
	var u2 User = User{"kk", "152000000", a1}
	u3 := User{
		name: "woniu",
		tel:  "18500000000",
		addr: Address{
			"北京市",
			"海淀路",
			"001",
		},
	}

	fmt.Println(u1, u2, u3)

	// 属性的访问
	fmt.Println(u2.addr)
	fmt.Println(u2.addr.no)

	// 属性的修改
	u2.addr.no = "002"
	fmt.Println(u2)

	// 声明&初始化嵌入对象
	var ekk Employee
	fmt.Printf("%T: %#v\n", ekk, ekk)

	// 使用已有变量初始化
	ekk01 := Employee{u2, 20000, "开发工程师"}
	ekk02 := Employee{
		User{
			"kk",
			"15200002222",
			Address{
				"陕西省西安市",
				"锦业路",
				"002",
			},
		},
		22000,
		"安全开发工程师",
	}

	// 使用指定名称初始化
	ekk03 := Employee{
		User: User{
			"kk",
			"15200002222",
			Address{
				"陕西省西安市",
				"锦业路",
				"003",
			},
		},
		salary: 25000,
		title:  "安全架构师",
	}

	fmt.Println(ekk01)
	fmt.Println(ekk02)
	fmt.Println(ekk03)

	// 针对匿名嵌入结构体属性，可以通过对象.结构体名称.属性名访问和修改属性值
	fmt.Println(ekk03.User.name, ekk03.User.tel, ekk03.User.addr, ekk03.salary, ekk03.title)

	ekk03.User.tel = "15200003333"

	fmt.Println(ekk03)

	// 针对匿名嵌入结构体属性值访问和修改时可省略结构体名称
	fmt.Println(ekk03.name, ekk03.tel, ekk03.addr, ekk03.salary, ekk03.title)

	ekk03.tel = "15200004444"

	fmt.Println(ekk03)

	// 使用指定名称初始化
	ekk04 := Employee2{
		User: User{
			"kk",
			"15200002222",
			Address{
				"陕西省西安市",
				"锦业路",
				"003",
			},
		},
		salary: 25000,
		title:  "安全架构师",
		addr:   "陕西省西安市高新路001",
	}

	fmt.Printf("%#v\n", ekk04)

	// 分别访问被嵌入和嵌入结构体的addr属性
	fmt.Println(ekk04.addr)
	fmt.Println(ekk04.User.addr)

	// 分别修改被嵌入和嵌入结构体的addr属性
	ekk04.addr = "陕西省西安市高新路002"
	fmt.Println(ekk04.addr)
	fmt.Println(ekk04.User.addr)

	ekk04.User.addr.no = "004"
	fmt.Println(ekk04.addr)
	fmt.Println(ekk04.User.addr)

	// 声明和初始化嵌入指针结构体对象
	paddr := Address{"陕西省西安市", "锦业路", "001"}

	puser := PUser{"kk", "152000000", &paddr}

	pemployee := PEmployee{
		PUser:  &puser,
		salary: 25000,
		title:  "安全架构师",
	}

	fmt.Printf("%#v\n", pemployee)

	//属性访问和修改
	fmt.Println(paddr)      // no:001
	fmt.Println(puser.addr) // no:001

	puser.addr.no = "002"
	fmt.Println(paddr)      // no:002
	fmt.Println(puser.addr) // no:002

	fmt.Println(puser.name)           //name: kk
	fmt.Println(pemployee.PUser.name) //name: kk
	fmt.Println(pemployee.name)       //name: kk

	pemployee.name = "silence"

	fmt.Println(puser.name)           //name: silence
	fmt.Println(pemployee.PUser.name) //name: silence
	fmt.Println(pemployee.name)       //name: silence

	// 声明和初始化嵌入值结构体对象
	vaddr := Address{"陕西省西安市", "锦业路", "001"}

	vuser := User{"kk", "152000000", vaddr}

	vemployee := Employee{
		User:   vuser,
		salary: 25000,
		title:  "安全架构师",
	}

	fmt.Printf("%#v\n", vemployee)

	//属性访问和修改
	fmt.Println(vaddr)      // no:001
	fmt.Println(vuser.addr) // no:001

	vuser.addr.no = "002"
	fmt.Println(vaddr)      // no:001
	fmt.Println(vuser.addr) // no:002

	fmt.Println(vuser.name)          //name: kk
	fmt.Println(vemployee.User.name) // name: kk
	fmt.Println(vemployee.name)      // name: kk

	vemployee.name = "silence"

	fmt.Println(vuser.name)          // name: kk
	fmt.Println(vemployee.User.name) // name: silence
	fmt.Println(vemployee.name)      // name: silence
}
