package main

import "fmt"

func main() {
    var (
        s1 string = "abc\nbcd"
        s2 string = `abc\nbcd`
        s3 string = `abc
        bcd
        `
        s4 string = "abcdefhijklmnopqrstuvwxyz"
        s5 string = "我爱中国"
    )

    fmt.Println(s1)
    fmt.Println(s2)
    fmt.Println(s3)
    fmt.Println(s4)
    fmt.Println(s5)

    fmt.Println(s4 + "ABCDEFGHIJKLMNOPQRSTUVWXYZ")

    s5 += "--来自Silence"
    fmt.Println(s5)

    fmt.Println(s1 == s2)
    fmt.Println(s1 != s3)
    fmt.Println("abc" > "abe")
    fmt.Println("abc" < "abe")
    fmt.Println("abc" >= "abe")
    fmt.Println("abc" <= "abe")

    fmt.Println(s4[0])
    fmt.Println(s4[:], s4[3:], s4[:8], s4[3:8])

    fmt.Printf("%s len %d\n", s1, len(s1))
    fmt.Printf("%s len %d\n", s5, len(s5))
}
