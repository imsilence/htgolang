package objs

import "fmt"

// 定义User结构体,并为每个属性定义标签
type User struct {
	id     int     `json:"id"`
	name   string  `json:"name"`
	Tel    string  `json:"addr"`
	Height float32 `json:"height"`
	Desc   *string `json:"desc"`
	Weight *int    `json:"weight"`
}

func NewUser(id int, name string, tel string, height float32, desc string, weight int) *User {
	return &User{id, name, tel, height, &desc, &weight}
}

//定义String方法
func (user User) String() string {
	return fmt.Sprintf("User{id: %d, name: %s, tel: %s, height: %e, desc: %s, weight: %d}",
		user.id, user.name, user.Tel, user.Height, *user.Desc, *user.Weight,
	)
}

func (user User) GetId() int {
	return user.id
}

func (user User) SetId(id int) {
	user.id = id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

// 定义接口Closer
type Closer interface {
	Close() error
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
	fmt.Printf("%s:%d\n", ip, port)
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
