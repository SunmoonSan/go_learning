package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int  {
	ret := 0
	for _, op := range ops{
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3))
	t.Log(Sum(1, 3, 5, 7, 9))
}

func Clear()  {
	fmt.Println("Clear resources")
}

func TestDefer(t *testing.T)  {
	//defer Clear()
	defer func() {
		fmt.Println("Clear resources")
	}()
	fmt.Println("Start")
	panic("error")
	fmt.Println("End")
}