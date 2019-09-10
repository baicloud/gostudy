package main

import (
	"fmt"
)

func testAddr1() {
	var a int
	a=1000
	fmt.Printf("变量地址%v\n值%d\n",&a,a)
	var b int8
	b=10
	fmt.Printf("变量地址%p\n值%d\n",&b,b)
}

func main() {
	testAddr1()
}
