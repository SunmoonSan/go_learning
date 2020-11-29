package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter 是所有goroutine都要增加其值的变量
	counter int

	// wg用来等待程序结束
	wg sync.WaitGroup

	// mutex用来定义一段代码临界区
	mutex sync.Mutex
)

// main 是所有Go程序的入口
func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)
	go incCounter(3)
	go incCounter(4)

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)

}

// incCounter使用互斥锁来同步保证访问安全
func incCounter(id int) {
	defer wg.Done()

	fmt.Println(id)

	for count := 0; count < 20; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}
