package main

import "fmt"

func main() {

    //单值
    operate := "add"
    switch operate {
    case "add":
        fmt.Println("添加")
    case "delete":
        fmt.Println("删除")
    case "modify":
        fmt.Println("修改")
    case "query":
        fmt.Println("查询")
    default:
        fmt.Println("操作错误")
	}

	input := "y"
    switch input {
	case "yes", "y":
		fmt.Println("确认")
	case "no", "n":
		fmt.Println("取消")
    default:
        fmt.Println("操作错误")
    }

    //表达式
    score := 80
    switch {
    case score >= 90:
        fmt.Println("优秀(A)")
    case score >= 80:
        fmt.Println("良好(B)")
    case score >= 60:
        fmt.Println("及格(C)")
    default:
        fmt.Println("不及格(D)")
    }


    //初始化单值
    switch operate := "add"; operate {
    case "add":
        fmt.Println("添加")
    case "delete":
        fmt.Println("删除")
    case "modify":
        fmt.Println("修改")
    case "query":
        fmt.Println("查询")
    default:
        fmt.Println("操作错误")
    }

    //初始化表达式
    switch score := 80; {
    case score >= 90:
        fmt.Println("优秀(A)")
    case score >= 80:
        fmt.Println("良好(B)")
    case score >= 60:
        fmt.Println("及格(C)")
    default:
        fmt.Println("不及格(D)")
    }

    //fallthrough
    switch score := 80; {
    case score >= 90:
        fmt.Println("优秀(A)")
    case score >= 80:
        fmt.Println("良好(B)")
        fallthrough
    case score >= 60:
        fmt.Println("及格(C)")
    default:
        fmt.Println("不及格(D)")
    }

}