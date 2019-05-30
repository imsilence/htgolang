package main

import "fmt"

// 当未发生panic则recover函数得到的结果为nil
func success() {
    defer func() {
        fmt.Println(recover())
    }()
    fmt.Println("success")
}

// 当未发生panic则recover函数得到的结果为panic传递的参数
func failure() {
    defer func() {
        fmt.Println(recover())
    }()
    fmt.Println("failure")
    panic("error")
}

// recover只能获取最后依次的panic信息
func failure2() {
    defer func() {
        fmt.Println(recover())
    }()
    defer func() {
        fmt.Println("failure 02")
        panic("error 02")
    }()

    fmt.Println("failure")
    panic("error")
}

func tpl() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("发生异常: %v\n", r)
        }
    }()

    fmt.Println("failure")
    panic("error")
    return
}

func main() {
    failure()
    success()
    failure2()
    fmt.Println(tpl())
}