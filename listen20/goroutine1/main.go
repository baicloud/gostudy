package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello goroutine")
}

func main() {
	go hello()
	time.Sleep(time.Second)
	fmt.Print("main thread terminate")
}
