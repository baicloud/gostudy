package main

import (
	"fmt"
)

func testDefer1() {
	defer fmt.Println("hello defer1")
	defer fmt.Println("hello defer2")
}

func testDefer2() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("i=%d\n", i)
	}

	fmt.Printf("running\n")
	fmt.Printf("return\n")
}

func testDefer3() {
	var i int = 0
	defer fmt.Printf("defer i=%d\n", i)
	i= 1000
	fmt.Printf("i=%d\n", i)
}

func main() {
	// defer遵循栈原理，先进后出
	//testDefer1()
	//testDefer2()
	testDefer3()
}
