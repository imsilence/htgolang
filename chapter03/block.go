package main

import "fmt"


func main() {
    outer := 1
    {
        inner := 2
        fmt.Println(outer, inner)
        outer := 3
        fmt.Println(outer, inner)
    }

    //fmt.Println(outer, inner)
    fmt.Println(outer)
}