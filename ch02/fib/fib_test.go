package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 10; i++ {
		t.Log(" ", b)
		a, b = a+b, a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
