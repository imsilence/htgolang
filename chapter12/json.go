package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// 定义User结构体, 并使用标签声明json属性
type User struct {
	ID           int       `json: `                            //json键使用属性名ID
	Name         string    `json:"name"`                       //json键重命名为name
	Password     string    `json:"-"`                          //json序列化忽略该字段
	gender       bool      `json:"gender"`                     //非公开字段不进行序列化
	Online       bool      `json:"is_online,string,omitempty"` //json键重命名为is_online, 并设置值类型为string
	RegisterTime time.Time `json:"register_time"`              //json键重命名为register_time
	Status       int       `json:",omitempty"`                 //json键使用属性名Status, 当为0值时不进行序列化
}

func main() {

	//定义用户类型容器(切片和映射)
	users_slice := make([]User, 10, 20)
	users_map := make(map[int]User)

	//初始化切片和映射
	for i := 0; i < 10; i++ {
		u := User{
			ID:           i,
			Name:         fmt.Sprintf("user.%d", i),
			Password:     fmt.Sprintf("password.%d", i),
			gender:       i%2 == 0,
			Online:       i%3 == 0,
			RegisterTime: time.Now(),
			Status:       i % 5,
		}
		users_slice[i] = u
		users_map[i] = u
	}

	// 对切片进行json序列化
	b, err := json.Marshal(users_slice)
	if err == nil {
		fmt.Println(string(b))

		// 将json字符串格式化并输出到bytes.Buffer变量中
		var buffer bytes.Buffer
		json.Indent(&buffer, b, "", "\t")
		buffer.WriteTo(os.Stdout)
	}

	// 对映射进行json序列化(并进行格式化)
	b, err = json.MarshalIndent(users_map, "", "\t")
	if err == nil {
		fmt.Println(string(b))
	}

	//定义json格式字符串
	jsonStr := `
	[
		{
				"ID": 0,
				"Name": "silence",
				"Password" : "test@123",
				"is_online": "true",
				"register_time": "2019-04-19T14:34:41.4164701+08:00"
		},
		{
				"ID": 1,
				"Name": "kk",
				"is_online": "false",
				"register_time": "2019-04-19T14:34:41.4164701+08:00",
				"status": 1
		}
	]
	`

	// 将json字符串进行反序列化到切片对象中
	var users []User
	err = json.Unmarshal([]byte(jsonStr), &users)
	if err == nil {
		fmt.Println(users)
	}

	// 验证字符串是否满足json格式
	fmt.Println(json.Valid([]byte("1")))
	fmt.Println(json.Valid([]byte("true")))
	fmt.Println(json.Valid([]byte("1.1")))
	fmt.Println(json.Valid([]byte(`{}`)))
	fmt.Println(json.Valid([]byte(`[]`)))
	fmt.Println(json.Valid([]byte(`[{}, {}]`)))
	fmt.Println(json.Valid([]byte(`{"name" : "kk"}`)))
	fmt.Println(json.Valid([]byte(`{"name" : kk}`)))    //json格式字符串需要使用"包含
	fmt.Println(json.Valid([]byte(`{"name" : "kk",}`))) //json格式最后一个字段后不允许有,

	// 使用Encoder将json结果输出到标准输出
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(users_slice)
	if err != nil {
		fmt.Println(err)
	}

	// 使用Encoder将序列化结果输出到bytes.Buffer/strings.Builder
	var out bytes.Buffer
	// var out strings.Builder

	encoder = json.NewEncoder(&out)
	err = encoder.Encode(users_map)
	if err == nil {
		fmt.Println(out.String())
	} else {
		fmt.Println(err)
	}

	// 使用Decoder从bytes.Buffer或strings.Reader中读取结果并反序列化

	// 将json字符串写入到bytes.Buffer中供Decoder读取
	in := bytes.NewBufferString(jsonStr)
	// 将json字符串写入到strings.Reader中供Decoder读取
	// in := strings.NewReader(jsonStr)

	var users2 []User
	decoder := json.NewDecoder(in)
	err = decoder.Decode(&users2)
	if err == nil {
		fmt.Println(users2)
	} else {
		fmt.Println(err)
	}

}
