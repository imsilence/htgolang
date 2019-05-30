package main

import "fmt"

func main() {
    var name string = "silence"
    var age int = 30
    fmt.Println("name=", name)

    fmt.Printf("name=%v, age=%v\n", name, age)
    fmt.Printf("name=%T, age=%T\n", name, age)
}