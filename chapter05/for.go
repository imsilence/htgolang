package main

import "fmt"

func main() {
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    //continue
    for i := 1; i <= 5; i++ {
        if i == 3 {
            continue
        }
        fmt.Println(i)
    }

    //break
    for i := 1; i <= 5; i++ {
        if i == 3 {
            break
        }
        fmt.Println(i)
    }

    sum = 0
    i := 1
    for i <= 100 {
        sum += i
        i++
    }
    fmt.Println(sum)

    i = 1
    sum = 0
    for {
        if i > 100 {
            break
        }
        sum += i
        i++
    }
    fmt.Println(sum)

    // i = 0
    // for {
    //     i++
    //     fmt.Println(i)
    // }

    text := "我爱中国"
    for i, e := range text {
        fmt.Printf("%d: %c\n", i, e)
    }

    for _, e := range text {
        fmt.Printf("%c\n", e)
    }

}