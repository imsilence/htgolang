package main

import "fmt"

func main() {


    var tels map[string]string
    var points map[[2]int]float64
    fmt.Printf("%T, %t, %v\n", tels, tels == nil, tels)
    fmt.Printf("%T, %t, %v\n", points, points == nil, points)

    tels = map[string]string{"kk": "15200000000", "woniu" : "18500000000"}
    fmt.Printf("%q\n", tels)

    points = map[[2]int]float64{{1, 2} : 3, {4, 5} : 6}
    fmt.Println(points)

    scores := map[string]int{"kk" : 80, "woniu" : 90}
    fmt.Println(scores)

    heighs := map[string]float64{}
    fmt.Println(heighs)

    weights := make(map[string]float64)
    fmt.Println(weights)


    fmt.Println(len(tels), len(points), len(scores), len(heighs), len(weights))

    students := map[int]string{1: "kk", 2: "woniu"}
    students01 := map[int]map[string]string{1 : map[string]string{"name" : "kk", "tel" : "152xxxxxxxx"}}

    fmt.Printf("%v, %q, %q\n", students, students[1], students[3])
    fmt.Printf("%v, %q, %q, %t\n", students01, students01[1], students01[3], students01[3] == nil)

    student, ok := students[1]
    fmt.Printf("%t, %v\n", ok, student)

    student, ok = students[2]
    fmt.Printf("%t, %v\n", ok, student)


    student01, ok := students01[1]
    fmt.Printf("%t, %v\n", ok, student01)

    student01, ok = students01[2]
    fmt.Printf("%t, %v\n", ok, student01)

    students[1] = "KK"  //key存在, 修改
    students[3] = "WD"  //key不存在, 增加

    fmt.Println(students)

    students01[1]["tel"] = "15200000000"  //key存在, 修改
    students01[1]["addr"] = "西安市" //key不存在, 增加
    students01[2] = map[string]string{"name" : "woniu", "tel" : "15800000000", "addr" : "北京市"} //key不存在, 增加
    fmt.Println(students01)

    students01[2] = map[string]string{"name" : "woniu", "tel" : "15811111111", "addr" : "北京市"} //key存在, 修改
    fmt.Println(students01)


    delete(students, 1)
    delete(students, 4)
    fmt.Println(students)
    delete(students01[1], "addr")
    delete(students01, 2)
    fmt.Println(students01)

    for k, v := range students {
        fmt.Printf("%v:%v\n", k, v)
    }
}