package main

import "testing"

func TestAdd(t *testing.T) {
	if 3 != Add(1, 2) {
		t.Errorf("1 + 2 != 3")
	}
}

func TestMinus(t *testing.T) {
	if -1 != Minus(1, 2) {
		t.Errorf("1 - 2 != -1")
	}
}

func BenchmarkFact(b *testing.B) {
	for i := 0; i < 10; i++ {
		Fact(10)
	}
}
