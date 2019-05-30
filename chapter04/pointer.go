package main

import "fmt"

func main() {

    var pointer01 *int
    var pointer02 *float64
    var pointer03 *string

    fmt.Printf("%T, %T, %T\n", pointer01, pointer02, pointer03)
    fmt.Println(pointer01, pointer02, pointer03)
    fmt.Printf("%t, %t, %t\n", pointer01 == nil, pointer02 == nil, pointer03 == nil)

    var (
        age int = 30
        heigh float64 = 1.68
        motto string = "少年经不得顺境，中年经不得闲境，晚年经不得逆境"
    )

    //指针变量初始化
    pointer01, pointer02, pointer03 = &age, &heigh, &motto
    pointer04, pointer05, pointer06 := &age, &heigh, &motto
    pointer07, pointer08, pointer09 := new(int), new(float64), new(string)

    //打印变量地址
    fmt.Println(&age, &heigh, &motto)

    //打印指针变量
    fmt.Println(pointer01, pointer02, pointer03)
    fmt.Println(pointer04, pointer05, pointer06)
    fmt.Println(pointer07, pointer08, pointer09)

    fmt.Println(age, heigh, motto) //打印变量值
    fmt.Println(*pointer01, *pointer02, *pointer03) //通过指针变量访问位置存储的值
    fmt.Printf("%v, %v, %q\n", *pointer07, *pointer08, *pointer09)

    //通过指针变量访问修改存储的值
    *pointer01 += 1
    *pointer02 = 1.70
    *pointer03 += "--曾国藩"

    fmt.Println(*pointer01, *pointer02, *pointer03)  //通过指针变量访问位置存储的值
    fmt.Println(age, heigh, motto)  //打印变量值

    fmt.Println(&age, &heigh, &motto) //打印变量地址
    fmt.Println(pointer01, pointer02, pointer03) //打印指针变量

    //与赋值新变量对比
    age2, heigh2, motto2 := age, heigh, motto

    //修改新变量值
    age2 += 1
    heigh2 = 1.72
    motto2 += "家书"

    //打印变量
    fmt.Println(age, heigh, motto)
    fmt.Println(age2, heigh2, motto2)

    //打印变量地址
    fmt.Println(&age, &heigh, &motto)
    fmt.Println(&age2, &heigh2, &motto2)


    //定义声明指针的指针
    var ppointer01 **int
    var ppointer02 **float64 = &pointer02
    ppointer03 := &pointer03

    fmt.Println(ppointer01, ppointer02, ppointer03)

    ppointer01 = &pointer01
    fmt.Println(ppointer01, ppointer02, ppointer03)

    //通过指针的指针访问变量地址和变量值
    fmt.Println(*ppointer01, *ppointer02, *ppointer03)
    fmt.Println(**ppointer01, **ppointer02, **ppointer03)

    //通过指针的指针修改和变量值
    **ppointer01 += 1
    **ppointer02 = 1.72
    **ppointer03 += "家属"

    fmt.Println(**ppointer01, **ppointer02, **ppointer03)
    fmt.Println(*pointer01, *pointer02, *pointer03)
    fmt.Println(age, heigh, motto)


}
