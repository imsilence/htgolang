package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.MatchString(`^\d+$`, "1"))
	fmt.Println(regexp.QuoteMeta(`\.+*?()|[]^$`))

	var reg *regexp.Regexp = regexp.MustCompile(`(?U)\d+`) //标记为非贪婪模式
	var regl *regexp.Regexp = regexp.MustCompile(`\d+`)

	fmt.Println(reg.FindString("010-12456789-123"))
	fmt.Println(reg.FindAllString("010-12456789-123", 2))
	fmt.Println(reg.FindAllString("010-12456789-123", -1))
	fmt.Println(regl.FindString("010-12456789-123"))
	fmt.Println(regl.FindAllString("010-12456789-123", 2))
	fmt.Println(regl.FindAllString("010-12456789-123", -1))

	reg.Longest() //切换为贪婪模式
	fmt.Println(reg.FindString("010-12456789-123"))
	fmt.Println(reg.FindAllString("010-12456789-123", 2))
	fmt.Println(reg.FindAllString("010-12456789-123", -1))

	fmt.Println(reg.ReplaceAllString("010-12456789-123", "xxx"))

	reg01 := regexp.MustCompile(`^010-`)

	fmt.Println(reg01.MatchString("010-1245678-123"))
	fmt.Println(reg01.MatchString("020-1245678-123"))

	reg02 := regexp.MustCompile(regexp.QuoteMeta(`-`))
	fmt.Println(reg02.Split("010-11111-1111-1111", -1))

	fmt.Println(reg02.String())

}
