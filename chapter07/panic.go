package main

import "fmt"


func main() {
    defer func() {
        fmt.Println("defer 01")
    }()

    panic("error 00")
}