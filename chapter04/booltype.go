package main

import "fmt"

func main() {
    var (
        isBody bool = true
        isGirl bool = false
    )

    // 逻辑运算
    // 与
    fmt.Println(isBody && isBody)
    fmt.Println(isBody && isGirl)
    fmt.Println(isGirl && isBody)
    fmt.Println(isGirl && isGirl)

    // 或
    fmt.Println(isBody || isBody)
    fmt.Println(isBody || isGirl)
    fmt.Println(isGirl || isBody)
    fmt.Println(isGirl || isGirl)

    // 非
    fmt.Println(!isBody)
    fmt.Println(!isGirl)

    // 关系运算
    fmt.Println(isBody == isGirl)
    fmt.Println(isBody != isGirl)

    // 打印
    fmt.Println(isBody, isGirl)
    fmt.Printf("isBody=%t, isGirl=%t\n", isBody, isGirl)

}
