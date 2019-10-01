package main

import "fmt"

func main() {
	ch:=make(chan string,2)
	ch<-"hello"
	ch<-"world"
	//ch<-"!个数据
	//channel类似队列，先进先出
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
