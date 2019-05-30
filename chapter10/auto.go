package main

type User struct {
	id   int
	name string
}

func (user User) SetId(id int) {
	user.id = id
}

func (user User) GetId() int {
	return user.id
}

/**
// 隐式
func (user *User) SetId(id int)
	// 获取user地址的值，并拷贝调用SetId, 只影响拷贝(*user)的值, 并不影响调用者的值
	// 与(user User)SetId(id int) 行为一致
	(*user).SetId(id)
}

func (user *User) GetId() int {
	return (*user).GetId()
}

//使用值和指针都不改变调用者, 行为一致
*/

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) GetName() string {
	return user.name
}

/**
func (user User) SetName(name string) {
	// user为值接收者, 拷贝, 在调用(&user).SetName只会影响接收者(user)的值，并不会影响调用者的值
	// 与(user *User)SetName(name string) 行为不一致
	(&user).SetName(name)
}

func (user User) GetName(name string) {
	return (&user).GetNae()
}

// 使用指针改变接收者，使用值不改变调用者, 行为不一致
*/
