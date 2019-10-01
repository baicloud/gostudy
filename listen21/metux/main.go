package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup

var mutex sync.Mutex

func test_add1() {
	for i := 0; i < 5000; i++ {
		// 该步骤非原子性，多个goroutine同时执行是得不到想要的效果
		x = x + 1
	}
	wg.Done()
}

func test_add2() {
	for i := 0; i < 5000; i++ {
		mutex.Lock()
		x = x + 1
		mutex.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go test_add2()
	go test_add2()
	wg.Wait()
	fmt.Println("x:", x)
}
