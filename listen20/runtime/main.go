package main

import (
	"fmt"
	"runtime"
	"time"
)

var i int

func calc() {
	for {
		i++
	}
}

func main() {
	cpu := runtime.NumCPU()
	fmt.Println("cpu:", cpu)
	// runtime.GOMAXPROCS(2)
	for i := 0; i < 10; i++ {
		go calc()
	}
	time.Sleep(time.Hour)
}
