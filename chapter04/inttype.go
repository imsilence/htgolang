package main

import "fmt"

func main() {

    // 分别打印十进制，八进制，十六进制标识发的字面量
    fmt.Println(10, 010, 0X10)

    var n1, n2, n3 int = 1, 8, -2
    var n4 uint = 2

    // 算术运算
    fmt.Println(n1 + n2) // 9
    fmt.Println(n1 - n2) // -7
    fmt.Println(n1 * n2) // 8
    fmt.Println(n1 / n2) // 0
    fmt.Println(n1 % n2) // 1

    n1++
    n2--
    fmt.Println(n1, n2) // 2 7

    // 关系运算
    fmt.Println(n1 > n2)
    fmt.Println(n1 >= n2)
    fmt.Println(n1 < n2)
    fmt.Println(n1 <= n2)
    fmt.Println(n1 == n2)
    fmt.Println(n1 != n2)

    // 位运算
    // 2 => 0010
    // 7 => 0111
    // -2 => 1110
    fmt.Println(n1 & n2)  // 0010
    fmt.Println(n1 | n2)  // 0111
    fmt.Println(n1 ^ n2)  // 0101
    fmt.Println(n1 &^ n2) // 0000
    fmt.Println(n2 << n4) // 0001 1100
    fmt.Println(n2 >> n4) // 0001
    fmt.Println(n3 << n4) // 1111 1000
    fmt.Println(n3 >> n4) // 1111 1111

    // 赋值运算
    n1 += n2
    n2 -= n1
    n4 <<= n4
    fmt.Println(n1, n2, n4)

    //类型转换
    fmt.Printf("%T %T %T %T %T %T\n", n1, uint(n1), n4, int(n4), uint8(n4), int64(n4))

    //打印
    var (
        n5 int  = 10
        n6 int  = -10
        c0 byte = 'a'
        u0 rune = '牛'
    )

    fmt.Printf("%d %b %o %x %X %#o %#x %#X\n", n5, n5, n5, n5, n5, n5, n5, n5)
    fmt.Printf("%d|%+d|%10d|%-10d|%010d|%+-10d|%+010d|\n", n5, n5, n5, n5, n5, n5, n5)
    fmt.Printf("%d|%+d|%10d|%-10d|%010d|%+-10d|%+010d|\n", n6, n6, n6, n6, n6, n6, n6)
    fmt.Printf("%c %c %q %q\n", c0, u0, c0, u0)
    fmt.Printf("%U %#U\n", u0, u0)
}
