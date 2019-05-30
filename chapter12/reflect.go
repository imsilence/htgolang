package main

import (
	"fmt"
	"reflect"
)

const MAXDEPTH = 10

func reflectType(t reflect.Type, tab string, depth int) {
	if depth == MAXDEPTH {
		return
	}
	if t.String() == "*reflect.rtype" {
		return
	}

	fmt.Printf("%s------------------------------\n", tab)

	// fmt.Printf("%s# %T, %#v\n", tab, t, t)
	// fmt.Printf("%s* Align: %#v\n", tab, t.Align())
	// fmt.Printf("%s* FieldAlign: %#v\n", tab, t.FieldAlign())
	fmt.Printf("%s* Name: %#v\n", tab, t.Name())
	fmt.Printf("%s* PkgPath: %#v\n", tab, t.PkgPath())
	fmt.Printf("%s* Size: %#v\n", tab, t.Size())
	fmt.Printf("%s* String: %#v\n", tab, t.String())
	fmt.Printf("%s* Kind: %#v\n", tab, t.Kind())

	fmt.Printf("%s* Comparable: %#v\n", tab, t.Comparable())

	// fmt.Printf("%s* Implements Sender: %#v\n", tab, t.Implements(reflect.TypeOf((*Sender)(nil)).Elem()))
	// fmt.Printf("%s* AssignableTo: %#v\n", tab, t.AssignableTo(reflect.TypeOf((*Sender)(nil)).Elem()))
	// fmt.Printf("%s* ConvertibleTo Connection: %#v\n", tab, t.ConvertibleTo(reflect.TypeOf(Integer(1))))

	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		fmt.Printf("%s* Bits: %#v\n", tab, t.Bits())
	case reflect.Array:
		fmt.Printf("%s* Len: %#v\n", tab, t.Len())
		// fmt.Printf("%s* Elem: %#v\n", tab, t.Elem())
		fmt.Printf("%s*Element\n", tab+"\t")
		reflectType(t.Elem(), tab+"\t", depth+1)
	case reflect.Slice:
		// fmt.Printf("%s* Elem: %#v\n", tab, t.Elem())
		fmt.Printf("%s*Element\n", tab+"\t")
		reflectType(t.Elem(), tab+"\t", depth+1)
	case reflect.Map:
		// fmt.Printf("%s* Key: %#v\n", tab, t.Key())
		fmt.Printf("%s*Key\n", tab+"\t")
		reflectType(t.Key(), tab+"\t", depth+1)
		// fmt.Printf("%s* Elem: %#v\n", tab, t.Elem())
		fmt.Printf("%s*Value\n", tab+"\t")
		reflectType(t.Elem(), tab+"\t", depth+1)
	case reflect.Ptr:
		// fmt.Printf("%s* Elem: %#v\n", tab, t.Elem())
		fmt.Printf("%s*Value\n", tab+"\t")
		reflectType(t.Elem(), tab+"\t", depth+1)
	case reflect.Func:
		fmt.Printf("%s* IsVariadic: %#v\n", tab, t.IsVariadic())
		fmt.Printf("%s* NumIn: %#v\n", tab, t.NumIn())
		fmt.Printf("%s* Parameters:\n", tab)
		for i := 0; i < t.NumIn(); i++ {
			// fmt.Printf("%s\t*** Parameter(%d): %#v\n", tab, i, t.In(i))
			fmt.Printf("%s\t*** Parameter(%d):\n", tab, i)
			reflectType(t.In(i), tab+"\t\t", depth+1)
		}
		fmt.Printf("%s* Returns:\n", tab)
		for i := 0; i < t.NumOut(); i++ {
			// fmt.Printf("%s\t*** Return(%d): %#v\n", tab, i, t.Out(i))
			fmt.Printf("%s\t*** Return(%d):\n", tab, i)
			reflectType(t.Out(i), tab+"\t\t", depth+1)
		}
	case reflect.Chan:
		fmt.Printf("%s* ChanDir: %#v\n", tab, t.ChanDir())
	case reflect.Struct:
		fmt.Printf("%s* NumField: %#v\n", tab, t.NumField())
		fmt.Printf("%s* Fields:\n", tab)
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Printf("%s\t** Field(%d): %#v\n", tab, i, f.Name)
			fmt.Printf("%s\t\t*** PkgPath: %#v\n", tab, f.PkgPath)
			// fmt.Printf("%s\t\t*** Type: %#v\n", tab, f.Type)

			fmt.Printf("%s\t\t***Type\n", tab)
			reflectType(f.Type, tab+"\t\t\t", depth+1)
			fmt.Printf("%s\t\t*** Tag: %#v\n", tab, f.Tag)
			fmt.Printf("%s\t\t*** Tag: %#v\n", tab, f.Tag.Get("ip"))
			fmt.Printf("%s\t\t*** Offset: %#v\n", tab, f.Offset)
			fmt.Printf("%s\t\t*** Index: %#v\n", tab, f.Index)
			fmt.Printf("%s\t\t*** Anonymous: %#v\n", tab, f.Anonymous)
		}
		// f, ok := t.FieldByName("ip")
		// fmt.Printf("%s* FieldByName: %#v, %v\n", tab, f, ok)
		// fmt.Printf("%s* FieldByIndex: %#v\n", tab, t.FieldByIndex([]int{1}))

	}

	fmt.Printf("%s* NumMethod: %#v\n", tab, t.NumMethod())
	fmt.Printf("%s* Methods:\n", tab)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		fmt.Printf("%s\t** Method(%d): %#v\n", tab, i, m.Name)
		fmt.Printf("%s\t\t*** PkgPath: %#v\n", tab, m.PkgPath)
		// fmt.Printf("%s*** Type: %#v\n", tab, m.Type)
		fmt.Printf("%s\t\t***Type\n", tab)
		reflectType(m.Type, tab+"\t\t\t", depth+1)
		fmt.Printf("%s\t\t*** Func: %T, %#v\n", tab, m.Func, m.Func)
		fmt.Printf("%s\t\t*** Index: %#v\n", tab, m.Index)
	}

	// m, ok := t.MethodByName("Compare")
	// fmt.Printf("%s* MethodByName: %#v, %v\n", tab, m, ok)

}

func reflectValue(v reflect.Value, tab string, depth int) {
	fmt.Printf("%s%s------------------------------\n", tab, v.Type().String())
	fmt.Printf("%sCanAddr: %t\n", tab, v.CanAddr())
	if v.CanAddr() {
		fmt.Printf("%sAddr:%#v\n", tab, v.Addr())
	}

	fmt.Printf("%sCanInterface: %t\n", tab, v.CanInterface())
	if v.CanInterface() {

	}
	fmt.Printf("%sCanSet: %t\n", tab, v.CanSet())

	fmt.Printf("%sInterface: %#v\n", tab, v.Interface())
	fmt.Printf("%sIsValid: %t\n", tab, v.IsValid())

	switch v.Kind() {
	case reflect.Int:
		fmt.Printf("%sInt: %d\n", tab, v.Int())
		if v.CanSet() {
			v.SetInt(30)
		}
	case reflect.Float32:
		fmt.Printf("%sFloat: %f\n", tab, v.Float())
	case reflect.Bool:
		fmt.Printf("%sBool: %t\n", tab, v.Bool())
	case reflect.String:
		fmt.Printf("%sIndex(0): %#v\n", tab, v.Index(0))
		fmt.Printf("%sLen: %d\n", tab, v.Len())

	case reflect.Array:
		fmt.Printf("%sCap: %d\n", tab, v.Cap())
		fmt.Printf("%sIndex(0): %#v\n", tab, v.Index(0))
		fmt.Printf("%sLen %d\n", tab, v.Len())
	case reflect.Slice:
		fmt.Printf("%sCap: %d\n", tab, v.Cap())
		fmt.Printf("%sIndex(0): %#v\n", tab, v.Index(0))
		fmt.Printf("%sIsNil: %t\n", tab, v.IsNil())
		fmt.Printf("%sLen %d\n", tab, v.Len())
		fmt.Printf("%sPointer: %d\n", tab, v.Pointer())
	case reflect.Map:
		fmt.Printf("%sIsNil: %t\n", tab, v.IsNil())
		fmt.Printf("%sLen %d\n", tab, v.Len())
		fmt.Printf("%sMapKeys: %#v\n", tab, v.MapKeys())
		fmt.Printf("%sPointer: %d\n", tab, v.Pointer())

	case reflect.Ptr:
		fmt.Printf("%sIsNil: %t\n", tab, v.IsNil())
		fmt.Printf("%sPointer: %d\n", tab, v.Pointer())
		reflectValue(v.Elem(), tab+"\t", depth+1)
	case reflect.Chan:
		fmt.Printf("%sCap: %d\n", tab, v.Cap())
		fmt.Printf("%sIsNil: %t\n", tab, v.IsNil())
		fmt.Printf("%sPointer: %d\n", tab, v.Pointer())
	case reflect.Func:
		fmt.Printf("%sIsNil: %t\n", tab, v.IsNil())
	case reflect.Struct:
		fmt.Printf("%s* NumField: %#v\n", tab, v.NumField())
		fmt.Printf("%s* Fields:\n", tab)
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)

			fmt.Printf("%s\t** Field(%d): %#v\n", tab, i, f)
		}
	case reflect.Interface:
		fmt.Printf("%sInterfaceData: %#v\n", tab, v.InterfaceData())
		fmt.Printf("%sIsNil: %t\n", tab, v.IsNil())

	}

	fmt.Printf("%s* NumMethod: %#v\n", tab, v.NumMethod())
	fmt.Printf("%s* Methods:\n", tab)
	for i := 0; i < v.NumMethod(); i++ {
		m := v.Method(i)

		fmt.Printf("%s\t** Method(%d): %#v\n", tab, i, m)

	}

	vv := v.MethodByName("Compare")
	if vv.Kind() != reflect.Invalid {
		fmt.Println(vv.Call([]reflect.Value{reflect.ValueOf((Integer)(1))})[0].Int())
		fmt.Println(vv.Call([]reflect.Value{reflect.ValueOf((Integer)(10))})[0].Int())
		fmt.Println(vv.Call([]reflect.Value{reflect.ValueOf((Integer)(-10))})[0].Int())
	}

	// Bytes
	// Call
	//CallSlice
	//Convert
	//Elem
	//Field
	//FieldByIndex
	// FieldByName
}

type Integer int

func (integer Integer) Compare(other Integer) int {
	switch {
	case integer == other:
		return 0
	case integer > other:
		return 1
	case integer < other:
		return -1
	}
	return 0
}

type User struct {
	id   int    "json: id"
	name string "json:name"
}

func NewUser(id int, name string) *User {
	return &User{id, name}
}

func (user *User) GetId() int {
	return user.id
}

func (user *User) SetId(id int) {
	user.id = id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

type Sender interface {
	Send(msg string) error
}
type Receiver interface {
	Receive() error
}

type Address struct {
	ip   string "json:ip"
	port int    "json:port"
}

func (address Address) GetIp() string {
	return address.ip
}

func (address Address) GetPort() int {
	return address.port
}

type Connection struct {
	Address
	status int
}

func NewConnection(ip string, port int) *Connection {
	return &Connection{Address: Address{ip, port}}
}

func (conn *Connection) Send(msg string) error {
	fmt.Printf("发送消息给[%s:%d]: %s", conn.ip, conn.port, msg)
	return nil
}
func (conn *Connection) Close() error {
	fmt.Printf("关闭连接[%s:%d]", conn.ip, conn.port)
	return nil
}

func main() {

	var intV int = 1
	// reflectType(reflect.TypeOf(intV), "", 0)
	reflectValue(reflect.ValueOf(intV), "", 0)

	// reflectType(reflect.TypeOf(&intV), "", 0)
	reflectValue(reflect.ValueOf(&intV), "", 0)
	fmt.Println("intV:", intV)

	var floatV float32 = 2
	// reflectType(reflect.TypeOf(floatV), "", 0)
	reflectValue(reflect.ValueOf(floatV), "", 0)

	var boolV bool = true
	// reflectType(reflect.TypeOf(boolV), "", 0)
	reflectValue(reflect.ValueOf(boolV), "", 0)

	var stringV string = "吾日三省吾身。为人谋而不忠乎？与朋友交而不信乎？传不习乎？"
	// reflectType(reflect.TypeOf(stringV), "", 0)
	reflectValue(reflect.ValueOf(stringV), "", 0)

	var arrayV [5]int = [...]int{1, 2, 3, 4, 5}
	// reflectType(reflect.TypeOf(arrayV), "", 0)
	reflectValue(reflect.ValueOf(arrayV), "", 0)

	var sliceV []int = make([]int, 3, 5)
	// reflectType(reflect.TypeOf(sliceV), "", 0)
	reflectValue(reflect.ValueOf(sliceV), "", 0)

	var mapV map[string]string = map[string]string{"name": "kk"}
	// reflectType(reflect.TypeOf(mapV), "", 0)
	reflectValue(reflect.ValueOf(mapV), "", 0)

	var integerV Integer = 1
	// reflectType(reflect.TypeOf(integerV), "", 0)
	reflectValue(reflect.ValueOf(integerV), "", 0)

	// var funcV func(func(interface{}), ...interface{}) error
	// reflectType(reflect.TypeOf(funcV), "", 0)

	var funcV2 func() = func() { fmt.Println("Hi") }
	reflectValue(reflect.ValueOf(funcV2), "", 0)

	var userV *User = NewUser(1, "kk")
	// reflectType(reflect.TypeOf(userV), "", 0)
	reflectValue(reflect.ValueOf(userV), "", 0)

	var addressV Address = Address{"127.0.0.1", 8080}
	// reflectType(reflect.TypeOf(addressV), "", 0)
	reflectValue(reflect.ValueOf(addressV), "", 0)

	var connectionV *Connection = NewConnection("127.0.0.1", 8080)

	// reflectType(reflect.TypeOf(connectionV), "", 0)
	reflectValue(reflect.ValueOf(connectionV), "", 0)

	var senderV Sender = connectionV
	// reflectType(reflect.TypeOf(senderV), "", 0)
	reflectValue(reflect.ValueOf(senderV), "", 0)
}
