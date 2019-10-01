package main

import (
	"fmt"
	"time"
)

func process(i int,ch chan bool){
	fmt.Println("started Goroutine ",i)
	time.Sleep(2*time.Second)
	fmt.Printf("Goroutine %d ended\n",i)
	ch<-true
}

func main() {
	no:=3
	exitChan:=make(chan bool,no)
	for i:=0;i<no;i++{
		go process(i,exitChan)
	}
	/*
	如何等待一组goroutine结束
	在goroutine中插入channnel,在主线程中输出，
	channel长度等于goroutine数量
	*/
	for i:=0;i<no;i++{
		<-exitChan
	}
	fmt.Println("All go routines finished executing")
}
