package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("hello go")
}

func calc(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

func testFunc1(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return

}

func calc_v1(b ...int) int {
	sum := 0
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
	}
	return sum
}

func calc_v2(a int, b ...int) int {
	sum := a
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
	}
	return sum
}

func main() {
	//sayHello()
	//sum,sub:=calc(100,23)
	//sum, sub := testFunc1(300, 123)
	//fmt.Printf("sum=%d,sub=%d", sum, sub)
	sum := calc_v2(12,39,99)
	fmt.Printf("sum = %d", sum)

}
