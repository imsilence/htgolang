package main

import "fmt"

const (
    c1 int = iota
    c2
    c3 int = iota
    c4
)

func main() {
    const (
        c5 = iota
        c6
        c7 = 7
        c8 = iota
        c9
    )

    fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9)
}
