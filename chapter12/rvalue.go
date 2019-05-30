package main

import (
	. "chapter12/objs"
	"fmt"
	"reflect"
	"strconv"
)

// 修改reflect.Value类型变量v的信息
func displayValue(value reflect.Value, tab string) {

	//获取值对应的枚举类型使用选择语句分别处理每种类型
	switch value.Kind() {
	case reflect.Int:
		// 针对整数类型,获取值对应的Type对象并使用Int函数转化为对应的基本类型,并使用strconv将转化为string类型
		fmt.Printf("%s[%s] %s", tab, value.Type(), strconv.FormatInt(value.Int(), 10))
	case reflect.Float32:
		// 针对浮点类型,获取值对应的Type对象并使用Float函数转化为对应的基本类型,并使用strconv将转化为string类型
		fmt.Printf("%s[%s] %s", tab, value.Type(), strconv.FormatFloat(value.Float(), 'E', -1, 64))
	case reflect.Bool:
		// 针对布尔类型,获取值对应的Type对象并使用Bool函数转化为对应的基本类型,并使用strconv将转化为string类型
		fmt.Printf("%s[%s] %s", tab, value.Type(), strconv.FormatBool(value.Bool()))
	case reflect.String:
		// 针对字符串类型,获取值对应的Type对象并打印值
		fmt.Printf("%s[%s] %s", tab, value.Type(), value)
	case reflect.Array:
		// 针对数组类型,获取值对应的Type对象
		fmt.Printf("%s[%s] {\n", tab, value.Type())

		//获取数组的长度
		for i := 0; i < value.Len(); i++ {
			displayValue(value.Index(i), tab+"\t") //根据索引获取数组的每个元素,并调用displayValue递归显示
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Slice:
		// 针对切片类型,获取值对应的Type对象,长度和容量
		fmt.Printf("%s[%s](%d:%d) {\n", tab, value.Type(), value.Len(), value.Cap())

		//获取切片的长度
		for i := 0; i < value.Len(); i++ {
			displayValue(value.Index(i), tab+"\t") //根据索引获取数组的每个元素,并调用displayValue递归显示
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Map:
		//针对映射类型,获取值对应的Type对象
		fmt.Printf("%s[%s] {\n", tab, value.Type())
		//获取映射迭代对象遍历键值对
		iter := value.MapRange()
		for iter.Next() { //判断迭代对象是否到末尾
			displayValue(iter.Key(), tab+"\t") //根据从迭代对象获取当前键,并调用displayValue递归显示
			fmt.Printf(" : ")
			displayValue(iter.Value(), "") //根据从迭代对象获取当前值,并调用displayValue递归显示
			fmt.Printf(",\n")

		}
		fmt.Printf("%s}", tab)
	case reflect.Struct:
		//针对结构体类型,获取值对应的Type对象
		structType := value.Type()
		fmt.Printf("%s[%s] {\n", tab, structType)

		//获取属性数量并遍历
		for i := 0; i < value.NumField(); i++ {
			structField := structType.Field(i)           //根据索引获取属性的类型
			field := value.Field(i)                      //根据索引获取属性的值
			fmt.Printf("%s\t%s:", tab, structField.Name) //打印类型名
			displayValue(field, tab+"\t")                //调用displayValue递归显示
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Ptr:
		//针对指针类型,获取值对应的Type对象
		fmt.Printf("%s[%s] (\n", tab, value.Type())

		//获取指针的引用值,并递归调用displayValue函数递归分析显示
		displayValue(value.Elem(), tab+"\t")
		fmt.Printf("\n%s)", tab)
	default:
		fmt.Printf("%sUnkonw[%#v]", tab, value)
	}
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
	var userV *User = NewUser(1, "kk", "15200000000", 1.68, "少年经不得顺境，中年经不得闲境，晚年经不得逆境", 72)
	var closerV Closer
	vars = append(vars, intV, &intV, floatV, boolV, stringV, arrayV, sliceV, mapV, *userV, userV, closerV)
	for _, v := range vars {
		displayValue(reflect.ValueOf(v), "")
		fmt.Println()
	}
}
