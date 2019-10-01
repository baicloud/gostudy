package main

import (
	"fmt"
	"time"
)

func hello(c chan bool) {
	time.Sleep(5*time.Second)
	fmt.Println("hello goroutine")
	c<-true
}
func main() {
	var exitChan chan bool
	//有出队的chan，便不用初始化容量make(chan bool,1)
	exitChan=make(chan bool)
	go hello(exitChan)
	fmt.Println("main thread terminate")
	<-exitChan
}
