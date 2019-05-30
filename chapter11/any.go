package main

import "fmt"

func printType(vs ...interface{}) {
	for _, v := range vs {
		switch v.(type) {
		case nil:
			fmt.Println("nil")
		case int:
			fmt.Println("int")
		case bool:
			fmt.Println("bool")
		case string:
			fmt.Println("string")
		case [5]int:
			fmt.Println("[5]int")
		case []int:
			fmt.Println("[]int")
		case map[string]string:
			fmt.Println("map[string]string")
		default:
			fmt.Println("unknow")
		}
	}
}

func main() {
	var i01 interface{} = 1
	var i02 interface{} = true
	var i03 interface{} = "我是kk"
	var i04 interface{} = &i03
	var i05 interface{} = [...]int{1, 2, 3, 4, 5}
	var i06 interface{} = []int{1, 2, 3, 4, 5}
	var i07 interface{} = map[string]string{"name": "kk"}
	var i08 interface{} = struct{ X, Y int }{1, 2}

	fmt.Printf("%#v\n", i01)
	fmt.Printf("%#v\n", i02)
	fmt.Printf("%#v\n", i03)
	fmt.Printf("%#v\n", i04)
	fmt.Printf("%#v\n", i05)
	fmt.Printf("%#v\n", i06)
	fmt.Printf("%#v\n", i07)
	fmt.Printf("%#v\n", i08)

	printType(i01)
	printType(i02)
	printType(i03)
	printType(i04)
	printType(i05, i06, i07, i08, nil)
}
