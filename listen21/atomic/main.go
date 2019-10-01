package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var x int32
var mutex sync.Mutex

func addAtomic() {
	for i := 0; i < 5000000; i++ {
		atomic.AddInt32(&x, 1)
	}
	wg.Done()
}

func addMutex() {
	for i := 0; i < 5000000; i++ {
		mutex.Lock()
		x = x + 1
		mutex.Unlock()
	}
	wg.Done()
}

func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go addAtomic()
		// go addMutex()
	}
	wg.Wait()
	end := time.Now().UnixNano()
	cost := (end - start) / 1000 / 1000
	fmt.Println(x)
	fmt.Println("cost:", cost, "ms")
}
