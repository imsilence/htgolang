package main

import (
	. "chapter12/objs"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// 修改reflect.Value类型变量value的信息
func changeValue(value reflect.Value, path string) {
	//获取值对应的枚举类型使用选择语句分别处理每种类型
	switch value.Kind() {
	case reflect.Int:
		// 对于可设置的Int类型使用SetInt更新内存数据
		if value.CanSet() {
			fmt.Printf("Int CanSet: %s.%s\n", path, value.Type())
			value.SetInt(value.Int() + rand.Int63n(100))
		}
	case reflect.Float32:
		// 对于可设置的Float类型使用SetFloat更新内存数据
		if value.CanSet() {
			fmt.Printf("Float CanSet: %s.%s\n", path, value.Type())
			value.SetFloat(value.Float() + rand.Float64())
		}
	case reflect.Bool:
		// 对于可设置的Float类型使用SetInt更新内存数据
		if value.CanSet() {
			fmt.Printf("Bool CanSet: %s.%s\n", path, value.Type())
			value.SetBool(!value.Bool())
		}
	case reflect.String:
		// 对于可设置的String类型使用SetString更新内存数据
		if value.CanSet() {
			fmt.Printf("String CanSet: %s.%s\n", path, value.Type())
			value.SetString("chanage:" + value.String())
		}
	case reflect.Array:
		// 对于可设置的数组类型递归调用ChangeValue对每个元素更新内存数据
		if value.CanSet() {
			fmt.Printf("Array CanSet: %s.%s\n", path, value.Type())
		}
		//获取数组的长度
		for i := 0; i < value.Len(); i++ {
			changeValue(value.Index(i), path+".array") //根据索引获取数组的每个元素,并调用changeValue递归修改
		}
	case reflect.Slice:
		// 对于切片类型递归调用ChangeValue对每个元素更新内存数据
		for i := 0; i < value.Len(); i++ {
			changeValue(value.Index(i), path+".slice") //根据索引获取切片的每个元素,并调用changeValue递归修改
		}
	case reflect.Map:
		// 对于映射类型调用SetMapIndex设置Key对应的值
		keys := value.MapKeys() //获取映射所有key组成的Value切片
		for _, key := range keys {
			value.SetMapIndex(key, reflect.ValueOf("change:"+value.MapIndex(key).String()))
		}
	case reflect.Struct:
		if value.CanSet() {
			fmt.Printf("Struct CanSet: %s.%s\n", path, value.Type())
		}
		// 对于结构体类型递归调用ChangeValue对每个元素更新内存数据
		for i := 0; i < value.NumField(); i++ {
			changeValue(value.Field(i), path+".struct") //根据索引获取结构体的每个属性,并调用changeValue递归修改
		}
	case reflect.Ptr:
		if value.CanSet() {
			fmt.Printf("CanSet: %s.%s\n", path, value.Type())
		}
		// 对于指针类型, 递归调用对其引用值进行修改
		changeValue(value.Elem(), path+".pointer")
	default:
		fmt.Printf("Unkonw[%#v]\n", value)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
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
	vars = append(vars, intV, &intV, floatV, &floatV, boolV, &boolV, stringV, &stringV, arrayV, &arrayV)
	// intV, floatV, boolV, stringV, stringV, arrayV 不可修改
	// &intV, &floatV, &boolV, &stringV, &stringV, &arrayV 可修改, 指针类型通过Elem获取的引用值可获取地址

	vars = append(vars, sliceV, mapV, *userV, userV, closerV)
	// sliceV, mapV 可修改
	// userV 通过Elem获取的引用值可获取地址且公开的属性修改
	// *userV  公开属性可修改(结构体通过Elem获取的引用值可获取地址)

	for _, v := range vars {
		fmt.Println("------------------------")
		fmt.Printf("%v\n", v)
		vv := reflect.ValueOf(v)
		changeValue(vv, "")
		fmt.Printf("%v\n", reflect.Indirect(vv))
	}
}
