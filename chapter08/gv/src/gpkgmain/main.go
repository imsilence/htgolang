package main

import (
	"fmt"

	"gpkgname/pkg01" //导入包pkg01, 路径gpkgname/pkg01
	"gpkgname/pkg02" //导入包pkg02, 路径gpkgname/pkg02

	"github.com/astaxie/beego"
)

func init() {
	fmt.Println("main")
}

func main() {
	fmt.Println("gpkgmain")
	fmt.Println(pkg01.Name) //调用pkg01包中的成员Name
	fmt.Println(pkg02.Name) //调用pkg01包中的成员Name

	beego.Run()
}
