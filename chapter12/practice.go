package main

import (
	. "chapter12/objs"
	"fmt"
	"reflect"
	"strconv"
)

func displayType(t reflect.Type, tab string) {
	if t == nil {
		fmt.Printf("<nil>")
		return
	}
	switch t.Kind() {
	case reflect.Int, reflect.Float32, reflect.Bool, reflect.String:
		fmt.Printf("%s%s", tab, t.Name())
	case reflect.Array:
		// fmt.Printf("%s[%d]", tab, t.Len())
		// displayType(t.Elem(), "")

		fmt.Printf("%s%s", tab, t.String())
	case reflect.Slice:
		// fmt.Printf("%s[]", tab)
		// displayType(t.Elem(), "")
		fmt.Printf("%s%s", tab, t.String())
	case reflect.Map:
		fmt.Printf("%smap{\n", tab)
		fmt.Printf("%s\tKey: ", tab)
		fmt.Printf("%s%s", tab, t.String())
		// displayType(t.Key(), "")
		fmt.Println()
		fmt.Printf("%s\tValue: ", tab)
		fmt.Printf("%s%s", tab, t.String())
		// displayType(t.Key(), "")
		fmt.Println()
		fmt.Printf("%s}", tab)
	case reflect.Ptr:
		// fmt.Printf("%s%s", tab, t.String())
		fmt.Printf("%s*", tab)
		displayType(t.Elem(), "")
	case reflect.Func:
		fmt.Printf("%sfunc %s(", tab, t.Name())
		for i := 0; i < t.NumIn(); i++ {
			// displayType(t.In(i), "")
			fmt.Printf("%s", t.In(i).String())
			if i != t.NumIn()-1 {
				fmt.Printf(", ")
			}
		}
		if t.IsVariadic() {
			fmt.Printf("...")
		}
		fmt.Printf(") ")
		if t.NumOut() > 0 {
			fmt.Printf("(")
			for i := 0; i < t.NumOut(); i++ {
				fmt.Printf("%s", t.Out(i).String())
				if i != t.NumOut()-1 {
					fmt.Printf(", ")
				}
			}
			fmt.Printf(") ")
		}
		fmt.Printf("{}")

	case reflect.Struct:
		fmt.Printf("%stype %s struct {\n", tab, t.Name())
		fmt.Printf("%s\tFields:\n", tab)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fmt.Printf("%s\t\t%s\t%s", tab, field.Name, field.Type.String())
			fmt.Printf(",\n")

		}
		fmt.Printf("%s\tMethods:\n", tab)
		for i := 0; i < t.NumMethod(); i++ {
			method := t.Method(i)
			displayMethod(method, tab+"\t\t")
			fmt.Printf(",\n")

		}
		fmt.Printf("%s}", tab)
	case reflect.Interface:
		fmt.Printf("%sinterface{", tab)
		fmt.Printf("}")
	}
}

func displayMethod(method reflect.Method, tab string) {
	t := method.Type
	fmt.Printf("%sfunc %s(", tab, method.Name)
	for i := 0; i < t.NumIn(); i++ {
		fmt.Printf("%s", t.In(i).String())
		if i != t.NumIn()-1 {
			fmt.Printf(", ")
		}
	}
	if t.IsVariadic() {
		fmt.Printf("...")
	}
	fmt.Printf(") ")
	if t.NumOut() > 0 {
		fmt.Printf("(")
		for i := 0; i < t.NumOut(); i++ {
			fmt.Printf("%s", t.Out(i).String())
			if i != t.NumOut()-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf(") ")
	}
	fmt.Printf("{}")
}

func displayValue(value reflect.Value, tab string) {
	return
	switch value.Kind() {
	case reflect.Int:
		fmt.Printf("%s[%s] %s", tab, value.Type().String(), strconv.FormatInt(value.Int(), 10))
	case reflect.Float32:
		fmt.Printf("%s[%s] %s", tab, value.Type().String(), strconv.FormatFloat(value.Float(), 'E', -1, 64))
	case reflect.Bool:
		fmt.Printf("%s[%s] %s", tab, value.Type().String(), strconv.FormatBool(value.Bool()))
	case reflect.String:
		fmt.Printf("%s[%s] %s", tab, value.Type().String(), value.String())
	case reflect.Array:
		fmt.Printf("%s[%s] {\n", tab, value.Type().String())
		for i := 0; i < value.Len(); i++ {
			displayValue(value.Index(i), tab+"\t")
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Slice:
		fmt.Printf("%s[%s] {\n", tab, value.Type().String())
		for i := 0; i < value.Len(); i++ {
			displayValue(value.Index(i), tab+"\t")
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Map:
		fmt.Printf("%s[%s] {\n", tab, value.Type().String())
		iter := value.MapRange()
		for iter.Next() {
			displayValue(iter.Key(), tab+"\t")
			fmt.Printf(" : ")
			displayValue(iter.Value(), "")
			fmt.Printf(",\n")

		}
		fmt.Printf("%s}", tab)
	case reflect.Ptr:
		fmt.Printf("%s[%s] (", tab, value.Type().String())
		displayValue(value.Elem(), "")
		fmt.Printf(")")
	case reflect.Func:
		fmt.Printf("%s%s", tab, value.Type().String())
	case reflect.Struct:
		fmt.Printf("%s[%s] {\n", tab, value.Type().String())
		fmt.Printf("%s\tFields(%d):\n", tab, value.NumField())
		structType := value.Type()
		for i := 0; i < value.NumField(); i++ {
			structField := structType.Field(i)
			field := value.Field(i)
			fmt.Printf("\t\t%s: ", structField.Name)
			displayValue(field, "")
			fmt.Printf("\t`%s`, \n", structField.Tag)
		}

		fmt.Printf("%s\tMethods(%d):\n", tab, value.NumMethod())
		for i := 0; i < structType.NumMethod(); i++ {
			structMethod := structType.Method(i)
			fmt.Print("\t\t")
			displayMethod(structMethod, "")
			fmt.Printf("%s\t\t\n", tab)
		}
		fmt.Printf("%s}", tab)
	case reflect.Interface:
		fmt.Printf("%sinterface{", tab)
		fmt.Printf("}")

	}
}

func main() {

	var intV int = 1

	displayType(reflect.TypeOf(intV), "")
	fmt.Println()
	displayValue(reflect.ValueOf(intV), "")
	fmt.Println()

	displayType(reflect.TypeOf(&intV), "")
	fmt.Println()
	displayValue(reflect.ValueOf(&intV), "")
	fmt.Println()

	var floatV float32 = 3.14
	displayType(reflect.TypeOf(floatV), "")
	fmt.Println()
	displayValue(reflect.ValueOf(floatV), "")
	fmt.Println()

	var boolV bool = true
	displayType(reflect.TypeOf(boolV), "")
	fmt.Println()
	displayValue(reflect.ValueOf(boolV), "")
	fmt.Println()

	var stringV string = "吾日三省吾身。为人谋而不忠乎？与朋友交而不信乎？传不习乎？"
	displayType(reflect.TypeOf(stringV), "")
	fmt.Println()
	displayValue(reflect.ValueOf(stringV), "")
	fmt.Println()

	var arrayV [5]int = [...]int{1, 2, 3, 4, 5}
	displayType(reflect.TypeOf(arrayV), "")
	fmt.Println()
	displayValue(reflect.ValueOf(arrayV), "")
	fmt.Println()

	var sliceV []int = make([]int, 3, 5)
	displayType(reflect.TypeOf(sliceV), "")
	fmt.Println()
	displayValue(reflect.ValueOf(sliceV), "")
	fmt.Println()

	var sliceV2 [][]int = make([][]int, 3, 5)
	displayType(reflect.TypeOf(sliceV2), "")
	fmt.Println()
	displayValue(reflect.ValueOf(sliceV2), "")
	fmt.Println()

	var mapV map[string]string = map[string]string{"name": "kk"}
	displayType(reflect.TypeOf(mapV), "")
	fmt.Println()
	displayValue(reflect.ValueOf(mapV), "")
	fmt.Println()

	var funcV func(x ...interface{}) error = func(x ...interface{}) error { fmt.Println(x...); return nil }

	displayType(reflect.TypeOf(funcV), "")
	fmt.Println()
	displayValue(reflect.ValueOf(funcV), "")
	fmt.Println()

	var funcV2 func(string, int) *Connection = NewConnection

	displayType(reflect.TypeOf(funcV2), "")
	fmt.Println()
	displayValue(reflect.ValueOf(funcV2), "")
	fmt.Println()

	var userV *User = NewUser(1, "kk")

	displayType(reflect.TypeOf(*userV), "")
	fmt.Println()
	displayValue(reflect.ValueOf(*userV), "")
	fmt.Println()

	displayType(reflect.TypeOf(userV), "")
	fmt.Println()
	displayValue(reflect.ValueOf(userV), "")
	fmt.Println()

	var userSlice []User = make([]User, 3, 5)

	displayType(reflect.TypeOf(userSlice), "")
	fmt.Println()
	displayValue(reflect.ValueOf(userSlice), "")
	fmt.Println()

	displayType(reflect.TypeOf(&userSlice), "")
	fmt.Println()
	displayValue(reflect.ValueOf(&userSlice), "")
	fmt.Println()

	var pUserSlice []*User = make([]*User, 3, 5)

	displayType(reflect.TypeOf(pUserSlice), "")
	fmt.Println()
	displayValue(reflect.ValueOf(pUserSlice), "")
	fmt.Println()

	displayType(reflect.TypeOf(&pUserSlice), "")
	fmt.Println()
	displayValue(reflect.ValueOf(&pUserSlice), "")
	fmt.Println()

	// var closer Closer = NewConnection("127.0.0.1", 8080)
	var closer Closer

	displayType(reflect.TypeOf(closer), "")
	fmt.Println()
	displayValue(reflect.ValueOf(closer), "")
	fmt.Println()
}
