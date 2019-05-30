package main

import (
	"fmt"

	gpkg01 "gpkgname/pkg01"
	gpkg02 "gpkgname/pkg02"
	vpkg01 "vpkgname/pkg01"
	vpkg02 "vpkgname/pkg02"
)

func main() {
	fmt.Println("gvmain")
	fmt.Println(gpkg01.Name)
	fmt.Println(gpkg02.Name)
	fmt.Println(vpkg01.Name)
	fmt.Println(vpkg02.Name)
}
