package main

import "fmt"

func main() {

    // 老婆
    fmt.Println("老婆:")

    fmt.Println("有卖西瓜的:")
    has_watermelon := true
    fmt.Println("买10个包子")
    if has_watermelon {
        fmt.Println("买1个西瓜")
    }

    fmt.Println("没有卖西瓜的:")
    has_watermelon = false
    fmt.Println("买10个包子")
    if has_watermelon {
        fmt.Println("有卖西瓜的, 买1个西瓜")
    }


    // 老公
    fmt.Println("老公:")

    fmt.Println("有卖西瓜的:")
    has_watermelon = true
    if has_watermelon {
        fmt.Println("买1个包子")
    } else {
        fmt.Println("买10个包子")
    }

    fmt.Println("没有卖西瓜的:")
    has_watermelon = false
    if has_watermelon {
        fmt.Println("买1个包子")
    } else {
        fmt.Println("买10个包子")
    }


    score := 80

    if score >= 90 {
        fmt.Println("优秀(A)")
    } else {
        if score >= 80 {
            fmt.Println("良好(B)")
        } else {
            if score >= 60 {
                fmt.Println("及格(C)")
            } else {
                fmt.Println("不及格(D)")
            }
        }
    }


    if score >= 90 {
        fmt.Println("优秀(A)")
    } else if score >= 80 {
        fmt.Println("良好(B)")
    } else if score >= 60 {
        fmt.Println("及格(C)")
    } else {
        fmt.Println("不及格(D)")
    }

    if flag := true; flag {
        fmt.Println("flag:", flag)
    }

    if flag := false; flag {
        fmt.Println("flag if:", flag)
    } else {
        fmt.Println("flag else:", flag)
    }
}