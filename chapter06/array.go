package main

import "fmt"

func main() {
	var names [10]string
	var scores [10]int

	fmt.Printf("%T, %T\n", names, scores)
	fmt.Printf("%q\n", names)
	fmt.Println(scores)

	var users [3]string = [3]string{"KK", "蜗牛", "ada"}
	bounds := [...]int{10, 20, 30, 40, 50}
	teachers := [5]string{"KK"}
	nums := [10]int{1: 10, 3: 30, 5: 50, 8: 80}

	fmt.Printf("%q\n", users)
	fmt.Println(bounds)
	fmt.Printf("%q\n", teachers)
	fmt.Println(nums)

	fmt.Println(bounds == [...]int{10, 20, 30, 40, 50})
	fmt.Println(bounds != [...]int{20, 10, 30, 40, 50})

	langs := [...]string{"GO", "Python", "C#", "JAVA", "F#", "C", "C++", "Lua", "Lisp", "PHP", "Rust"}

	fmt.Println(len(langs))

	fmt.Println(langs[0], langs[5], langs[len(langs)-1])

	fmt.Printf("%q\n", langs)

	langs[0] = "Go"

	fmt.Printf("%q\n", langs)

	fmt.Printf("%T, %q\n", langs[1:3], langs[1:3])
	fmt.Printf("%T, %q\n", langs[1:3:3], langs[1:3:3])

	for i := 0; i < len(langs); i++ {
		fmt.Printf("%d: %q\n", i, langs[i])
	}

	for i, v := range langs {
		fmt.Printf("%d: %q\n", i, v)
	}

	//声明&初始化
	var students00 [6][8]string
	students01 := [...][2]string{{"kk", "woniu"}, {"ada", "wd"}, {"雪糕", "小林"}} //多维数组只有第一维长度可使用变量数量推测
	students02 := [3][3]string{{"kk", "woniu", "ada"}, {"wd"}, {"雪糕", "小林"}}
	students03 := [3][3]string{0: {"kk", "woniu", "ada"}, 2: {"wd", "雪糕", "小林"}}
	students04 := [3][3]string{0: {0: "kk", 2: "woniu"}, 2: {1: "wd", 2: "雪糕"}}

	fmt.Printf("%T, %T, %T, %T, %T\n", students00, students01, students02, students03, students04)
	fmt.Printf("%q\n", students00)
	fmt.Printf("%q\n", students01)
	fmt.Printf("%q\n", students02)
	fmt.Printf("%q\n", students03)
	fmt.Printf("%q\n", students04)

	//访问元素
	fmt.Printf("%q\n", students01[1])
	fmt.Println(students01[0][0])

	//修改元素
	students01[1] = [2]string{"无敌", "卧底"}
	students01[0][0] = "KK"
	fmt.Printf("%q\n", students01)

	//遍历元素
	for i := 0; i < len(students01); i++ {
		for j := 0; j < len(students01[i]); j++ {
			fmt.Printf("[%d, %d]: %q\n", i, j, students01[i][j])
		}
	}

	for i, line := range students01 {
		for j, v := range line {
			fmt.Printf("[%d, %d]: %q\n", i, j, v)
		}
	}

}
