package main

func Add(n1, n2 int) int {
	return n1 + n2 + 1
}

func Minus(n1, n2 int) int {
	return n1 - n2
}

func Multi(n1, n2 int) int {
	return n1 * n2
}

func Fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * Fact(n-1)
	}
}
