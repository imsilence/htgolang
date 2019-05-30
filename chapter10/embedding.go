package main

import "fmt"

// 定义User结构体和方法
type User struct {
	name string
	addr string
}

func NewUser(name, tel string) *User {
	return &User{name, tel}
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetAddr(addr string) {
	user.addr = addr
}

func (user *User) GetAddr() string {
	return user.addr
}

// 定义Employee结构体和方法
type Employee struct {
	*User
	title  string
	salary float64
}

func NewEmployee(name, addr, title string, salary float64) *Employee {
	return &Employee{
		NewUser(name, addr),
		title,
		salary,
	}
}

func (employee *Employee) GetName() string {
	return fmt.Sprintf("[%s]%s", employee.title, employee.name)
}

func (employee *Employee) SetTitle(title string) {
	employee.title = title
}

func (employee *Employee) GetTitle() string {
	return employee.title
}

func (employee *Employee) SetSalary(salary float64) {
	employee.salary = salary
}

func (employee *Employee) GetSalary() float64 {
	return employee.salary
}

func main() {

	me := NewEmployee("kk", "北京市朝阳区", "开发工程师", 15000)
	fmt.Printf("%#v\n", me)

	fmt.Println(me.GetName())      //调用Employee.GetName
	fmt.Println(me.User.GetName()) //调用User.GetName

	me.SetAddr("西安市高新区") //调用User.SetAddr

	fmt.Println(me.GetAddr()) //调用User.GetAddr

	me.SetSalary(13000)         //调用Employee.SetSalary
	fmt.Println(me.GetSalary()) //调用Employee.GetSalary

	fmt.Printf("%#v\n", me)
}
