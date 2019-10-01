package main

import "fmt"

func test_channel1() {
	var a chan int
	fmt.Println(a)
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
		fmt.Println(a)
	}
}

func test_channel2() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a=make(chan int)
		fmt.Printf("a=%v\n",a)
		fmt.Printf("a=%T\n",a)
		a<-10
		fmt.Printf("Type of a is %T",a)
	}

}

func test_channel3() {
	var c chan int
	fmt.Printf("c=%v\n", c)
	c = make(chan int, 1)
	fmt.Printf("c=%v\n", c)
	//入队
	c <- 100
	//出队
	data:=<-c
	fmt.Println(data)
}

func main() {
	//test_channel1()
	test_channel3()
}
