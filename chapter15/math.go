package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxUint32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.SmallestNonzeroFloat32)

	fmt.Printf("%T\n", math.Abs(-1))

	fmt.Println(math.Abs(-1))
	fmt.Println(math.Ceil(1.2))
	fmt.Println(math.Floor(1.2))
	fmt.Println(math.Trunc(1.2))
	fmt.Println(math.Dim(1.3, 1.4))
	fmt.Println(math.Dim(1.4, 1.3))
	fmt.Println(math.Max(1.3, 1.4))
	fmt.Println(math.Min(1.3, 1.4))
	fmt.Println(math.Round(1.4))
	fmt.Println(math.Round(1.5))
	fmt.Println(math.Round(-1.4))
	fmt.Println(math.Round(-1.5))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Sqrt(9))
	fmt.Println(math.Mod(2, 3))
	fmt.Println(math.Hypot(3, 4))
}
