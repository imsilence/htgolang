package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(fmt.Errorf("index out of range"))

	fmt.Print("func fmt.Print\n")
	fmt.Printf("func %s \n", "fmt.Printf")
	fmt.Println("fmt.Println")

	fmt.Fprint(os.Stdout, "func fmt.Fprint\n")
	fmt.Fprintf(os.Stdout, "fun %s\n", "fmt.Fprintf")
	fmt.Fprintln(os.Stdout, "func fmt.Fprintln")

	fmt.Printf("%q\n", fmt.Sprint("func fmt.Sprint"))
	fmt.Printf("%q\n", fmt.Sprintf("func %s", "fmt.Sprintf"))
	fmt.Printf("%q\n", fmt.Sprintln("func fmt.Sprintln"))

	var (
		name   string
		age    int
		weight float32
	)

	fmt.Print("请输入信息:")
	fmt.Scan(&name, &age, &weight)
	fmt.Println(name, age, weight)

	fmt.Print("请输入信息:")
	fmt.Scanf("%s %d %f", &name, &age, &weight)
	fmt.Println(name, age, weight)

	fmt.Print("请输入信息:")
	fmt.Scanln(&name, &age, &weight)
	fmt.Println(name, age, weight)
}
