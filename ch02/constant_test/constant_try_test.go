package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuestday
	Wensday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuestday, Wensday)
}

func TestConstantTry1(t *testing.T) {
	a := 7
	t.Log(a)
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
