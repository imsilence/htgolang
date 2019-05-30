package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(strconv.Atoi("30"))
	fmt.Println(strconv.Atoi("30x"))

	fmt.Println(strconv.ParseInt("10", 10, 64))  // 10进制
	fmt.Println(strconv.ParseUint("10", 16, 64)) // 16进制
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseFloat("3.14", 64))

	fmt.Println(strconv.Itoa(30))
	fmt.Println(strconv.FormatInt(30, 10))  // 10进制
	fmt.Println(strconv.FormatUint(30, 16)) // 16进制
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatFloat(3.14, 'g', -1, 64))

}
