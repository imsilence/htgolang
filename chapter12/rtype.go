package main

import (
	. "chapter12/objs"
	"fmt"
	"reflect"
)

// 打印reflect.Type类型变量t的信息
func displayType(t reflect.Type, tab string) {
	// 处理类型为nil
	if t == nil {
		fmt.Printf("<nil>")
		return
	}
	//获取类型对应的枚举值使用选择语句分别处理每种类型
	switch t.Kind() {
	case reflect.Int, reflect.Float32, reflect.Bool, reflect.String:
		// 针对基本数据类型显示类型名
		fmt.Printf("%s%s", tab, t.Name())
	case reflect.Array, reflect.Slice:
		// 针对数组和切片,直接打印Type对象
		fmt.Printf("%s%s", tab, t)
	case reflect.Map:
		// 对于映射类型打印键和值的Type对象
		fmt.Printf("%smap{\n", tab)
		fmt.Printf("%s\tKey: ", tab)
		fmt.Printf("%s%s", tab, t.Key()) //获取键的Type对象
		fmt.Println()
		fmt.Printf("%s\tValue: ", tab)
		fmt.Printf("%s%s", tab, t.Elem()) //获取值的Type对象
		fmt.Println()
		fmt.Printf("%s}", tab)
	case reflect.Func:
		//针对函数打印参数和返回值,对于可变参数在最后一个参数之后添加...
		fmt.Printf("%sfunc (", tab)

		//打印参数信息
		//获取参数数量并遍历
		for i := 0; i < t.NumIn(); i++ {
			fmt.Printf("%s", t.In(i)) //根据索引获取第i个参数的Type对象
			if i != t.NumIn()-1 {
				fmt.Printf(", ")
			}
		}
		if t.IsVariadic() {
			fmt.Printf("...")
		}
		fmt.Printf(") ")

		///打印返回值信息
		if t.NumOut() > 0 {
			fmt.Printf("(")
			//获取返回值数量并遍历
			for i := 0; i < t.NumOut(); i++ {
				fmt.Printf("%s", t.Out(i)) //根据索引获取第i个返回值的Type对象
				if i != t.NumOut()-1 {
					fmt.Printf(", ")
				}
			}
			fmt.Printf(") ")
		}
		fmt.Printf("{}")
	case reflect.Struct:
		// 针对结构体显示结构体属性和方法
		fmt.Printf("%stype %s struct {\n", tab, t.Name())

		//获取属性数量并遍历
		fmt.Printf("%s\tFields(%d):\n", tab, t.NumField())
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)                                                      //根据索引获取第i个属性的StructField对象
			fmt.Printf("%s\t\t%s\t%s\t`%s`", tab, field.Name, field.Type, field.Tag) //显示属性名,属性的Type对象和标签
			fmt.Printf(",\n")
		}

		//获取方法数量并遍历
		fmt.Printf("\n%s\tMethods(%d):\n", tab, t.NumMethod())
		for i := 0; i < t.NumMethod(); i++ {
			//打印方法信息
			displayMethod(t.Method(i), tab+"\t\t")
			fmt.Printf(",\n")

		}
		fmt.Printf("%s}", tab)
	case reflect.Ptr:
		//对于指针类型,递归分析其引用值
		fmt.Printf("%s*{\n", tab)
		displayType(t.Elem(), tab+"\t") //获取指针的引用值,并递归调用displayType函数进行分析显示

		// 打印指针变量的方法
		if t.NumMethod() > 0 { //获取方法数量
			fmt.Printf("\n%s\tMethods(%d):\n", tab, t.NumMethod())
			for i := 0; i < t.NumMethod(); i++ {
				//打印方法信息
				displayMethod(t.Method(i), tab+"\t\t")
				if i != t.NumMethod()-1 {
					fmt.Printf(",\n")
				}
			}
		}
		fmt.Printf("\n%s}", tab)
	default:
		fmt.Printf("%sUnkonw[%s]", tab, t)
	}
}

// 打印reflect.Method类型变量method的信息
func displayMethod(method reflect.Method, tab string) {
	// 获取方法接收者类型
	t := method.Type
	fmt.Printf("%sfunc %s(", tab, method.Name) //显示方法名

	//打印参数信息
	//获取参数数量并遍历
	for i := 0; i < t.NumIn(); i++ {
		fmt.Printf("%s", t.In(i)) //根据索引获取第i个参数的Type对象
		if i != t.NumIn()-1 {
			fmt.Printf(", ")
		}
	}

	//打印可变参数信息
	if t.IsVariadic() {
		fmt.Printf("...")
	}

	///打印返回值信息
	fmt.Printf(") ")

	if t.NumOut() > 0 {
		fmt.Printf("(")
		//获取返回值数量并遍历
		for i := 0; i < t.NumOut(); i++ {
			fmt.Printf("%s", t.Out(i)) //根据索引获取第i个返回值的Type对象
			if i != t.NumOut()-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf(") ")
	}
	fmt.Printf("{}")
}

func main() {
	vars := make([]interface{}, 0, 20)

	var intV int = 1
	var floatV float32 = 3.14
	var boolV bool = true
	var stringV string = "吾日三省吾身。为人谋而不忠乎？与朋友交而不信乎？传不习乎？"
	var arrayV [5]int = [...]int{1, 2, 3, 4, 5}
	var sliceV []int = make([]int, 3, 5)
	var mapV map[string]string = map[string]string{"name": "kk"}
	var funcV1 func(...interface{}) error = func(x ...interface{}) error { fmt.Println(x...); return nil }
	var funcV2 func(string, int) *Connection = NewConnection
	var userV *User = NewUser(1, "kk", "15200000000", 1.68, "少年经不得顺境，中年经不得闲境，晚年经不得逆境", 72)
	var closerV Closer
	vars = append(vars, intV, &intV, floatV, boolV, stringV, arrayV, sliceV, mapV, funcV1, funcV2, *userV, userV, closerV)
	for _, v := range vars {
		displayType(reflect.TypeOf(v), "")
		fmt.Println()
	}
}
