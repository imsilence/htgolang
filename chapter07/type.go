package main

import "fmt"


func main() {
    name, age, heigh, isBoy := "Silence", 30, 1.68, false //定义字符串、数值、布尔类型
    pointer := new(int) //定义指针类型
    scores := [...]int{1, 2, 3} // 定义数组类型
    names := make([]string, 1, 3) //定义切片类型
    user := make(map[int]string) //定义映射类型

    name2, age2, heigh2, isBoy2, pointer2, scores2, names2, user2 := name, age, heigh, isBoy, pointer, scores, names, user

    name2 = "kk"
    age2 = 31
    heigh2 = 1.70
    isBoy2 = true
    scores2[0] = 0
    pointer2 = &age
    names2[0] = "kk"
    user2[1] = "kk"

    fmt.Println(name, age, heigh, isBoy, scores, pointer, names, user)
    fmt.Println(name2, age2, heigh2, isBoy2, scores2, pointer2, names2, user2)



    fmt.Printf("%p, %p, %p, %p, %p, %p, %p, %p\n", &name, &age, &heigh, &isBoy, &scores, &pointer, &names, &user)
    fmt.Printf("%p, %p, %p, %p, %p, %p, %p, %p\n", &name2, &age2, &heigh2, &isBoy2, &scores2, &pointer2, &names2, &user2)
}