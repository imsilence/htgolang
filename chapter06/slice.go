package main

import "fmt"

func main() {

    var names []string
    fmt.Printf("%T, %t, %v\n", names, names == nil, names)

    scores := [...]int{60, 68, 70, 80, 95, 85}

    var scores00 []int = []int{80, 90, 95}
    scores01 := []int{}
    scores02 := []int{0:80, 2:95}
    scores03 := make([]int, 3)
    scores04 := make([]int, 3, 5)
    scores05 := scores[1:3]
    scores06 := scores[1:3:3]

    fmt.Printf("%T, %T, %T, %T, %T, %T\n", scores00, scores01, scores02, scores03, scores05)

    fmt.Println(scores00)
    fmt.Println(scores01)
    fmt.Println(scores02)
    fmt.Println(scores03)
    fmt.Println(scores04)
    fmt.Println(scores05)
    fmt.Println(scores06)

    students := make([]string, 3, 5)

    fmt.Println(len(students), cap(students))
    fmt.Printf("%q\n", students)


    fmt.Printf("%q, %q\n", students[0], students[1])

    students[0] = "KK"
    students[1] = "woniu"
    students[2] = "ada"
    fmt.Printf("%q, %q, %q\n", students[0], students[1], students[2])
    fmt.Printf("%q\n", students)

    teachers := [...]string{"kk", "woniu", "ada", "雪糕", "wd"}
    teachers00 := teachers[:]

    teachers01 := teachers[0:3]
    teachers02 := teachers[1:4]
    teachers03 := teachers00[1:3]

    fmt.Printf("%q\n", teachers)
    fmt.Printf("%d, %d, %q\n", len(teachers), cap(teachers01), teachers01)
    fmt.Printf("%d, %d, %q\n", len(teachers01), cap(teachers01), teachers01)
    fmt.Printf("%d, %d, %q\n", len(teachers02), cap(teachers02), teachers02)
    fmt.Printf("%d, %d, %q\n", len(teachers03), cap(teachers03), teachers03)

    teachers01[2] = "小林"
    fmt.Printf("%q\n", teachers)
    fmt.Printf("%d, %d, %q\n", len(teachers), cap(teachers01), teachers01)
    fmt.Printf("%d, %d, %q\n", len(teachers01), cap(teachers01), teachers01)
    fmt.Printf("%d, %d, %q\n", len(teachers02), cap(teachers02), teachers02)
    fmt.Printf("%d, %d, %q\n", len(teachers03), cap(teachers03), teachers03)


    teachers04 := teachers[1:4:4]
    teachers05 := teachers00[1:3:3]

    fmt.Printf("%d, %d, %q\n", len(teachers02), cap(teachers02), teachers02)
    fmt.Printf("%d, %d, %q\n", len(teachers03), cap(teachers03), teachers03)
    fmt.Printf("%d, %d, %q\n", len(teachers04), cap(teachers04), teachers04)
    fmt.Printf("%d, %d, %q\n", len(teachers05), cap(teachers05), teachers05)

    langs := []string{"GO", "Python", "C#", "JAVA", "F#", "C", "C++", "Lua", "Lisp", "PHP", "Rust"}

    for i:= 0; i < len(langs); i++ {
        fmt.Printf("%d, %q\n", i, langs[i])
    }

    for i, v := range langs {
        fmt.Printf("%d, %q\n", i, v)
    }

    nums := []int{1, 2, 3, 4, 5}

    nums2 := nums[:]
    fmt.Printf("%d, %d, %v\n", len(nums), cap(nums), nums)
    fmt.Printf("%d, %d, %v\n", len(nums2), cap(nums2), nums2)

    nums2 = append(nums2, 6)
    fmt.Printf("%d, %d, %v\n", len(nums), cap(nums), nums)
    fmt.Printf("%d, %d, %v\n", len(nums2), cap(nums2), nums2)

    nums2[0] = 0
    fmt.Printf("%d, %d, %v\n", len(nums), cap(nums), nums)
    fmt.Printf("%d, %d, %v\n", len(nums2), cap(nums2), nums2)

    nums3 := []int{80:5}
    fmt.Printf("%d, %d\n", len(nums3), cap(nums3))
    nums3 = append(nums3, 0, 1)
    fmt.Printf("%d, %d\n", len(nums3), cap(nums3))



    users01 := []string{"00", "01"}
    users02 := []string{"10", "11", "12"}
    users03 := []string{"20", "21", "22", "23"}

    fmt.Printf("%q, %q, %q\n", users01, users02, users03)
    copy(users01, users02)
    fmt.Printf("%q, %q, %q\n", users01, users02, users03)
    copy(users03, users02)
    fmt.Printf("%q, %q, %q\n", users01, users02, users03)


    elements := []int{0, 1, 2, 3, 4, 5}
    copy(elements[3:], elements[4:])
    fmt.Println(elements[:len(elements) - 1])

    queue := []int{}
    queue = append(queue, 1)
    queue = append(queue, 3)
    queue = append(queue, 2)

    fmt.Println(queue[0])
    queue = queue[1:]
    fmt.Println(queue[0])
    queue = queue[1:]
    fmt.Println(queue[0])

    stack := []int{}

    stack = append(stack, 1)
    stack = append(stack, 3)
    stack = append(stack, 2)

    fmt.Println(stack[len(stack) - 1])
    stack = stack[:len(stack) - 1]
    fmt.Println(stack[len(stack) - 1])
    stack = stack[:len(stack) - 1]
    fmt.Println(stack[len(stack) - 1])

    //声明&初始化
    points := [][]int{{1, 1}, {1, 2, 3}}
    fmt.Printf("%T, %v, %v, %d\n", points, points, points[0], points[0][0])

    //修改
    points[0] = []int{2, 2}
    points[1][1] = 3

    fmt.Println(points)

    //切片
    fmt.Println(points[0:2])

    //遍历
    for i := 0; i < len(points); i++ {
        for j := 0; j < len(points[i]); j++ {
            fmt.Printf("[%d, %d]: %v\n", i, j, points[i][j])
        }
    }

    for i, line := range points {
        for j, v := range line{
            fmt.Printf("[%d, %d]: %v\n", i, j, v)
        }
    }

    //append
    points = append(points, []int{2, 3, 1})
    points[0] = append(points[0], 1)
    fmt.Println(points)

    //copy
    points2 := [][]int{{}, {}}
    copy(points2, points)
    fmt.Println(points2)

}