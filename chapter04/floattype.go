package main

import "fmt"

func main() {

    var (
        f1 float32 = 3.1415926
        f2 float32 = 3E-3
        f3 float64 = 3.1E10
    )

    // 算术运算
    fmt.Println(f1 + f2)
    fmt.Println(f1 - f2)
    fmt.Println(f1 * f2)
    fmt.Println(f1 / f2)
    f1++
    f2--
    fmt.Println(f1, f2)

    // 关系运算
    fmt.Println(f1 > f2)
    fmt.Println(f1 >= f2)
    fmt.Println(f1 < f2)
    fmt.Println(f1 <= f2)

    // 赋值运算
    f1 += f1
    fmt.Println(f1)

    // 类型转换
    fmt.Printf("%T %T\n", f1, float64(f1))

    fmt.Printf("%f %f %F %F\n", f1, f3, f1, f3)
    fmt.Printf("%e %e %E %E\n", f1, f3, f1, f3)
    fmt.Printf("%g %g %G %G\n", f1, f3, f1, f3)
    fmt.Printf("%5.2f %5.2f\n", f1, f3)
}