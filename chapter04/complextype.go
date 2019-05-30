package main

import "fmt"

func main() {
    var (
        c1 complex64 = 1 + 2i
        c2 complex64 = complex(3, 4)
    )

    fmt.Println(c1 + c2)
    fmt.Println(real(c1), imag(c1))
}
