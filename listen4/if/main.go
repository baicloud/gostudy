package main

import (
	"fmt"
)

func testIf1() {
	num := 10
	if num%2 == 0 {
		fmt.Printf("num%d is even", num)
	} else {
		fmt.Printf("num%d is odd", num)
	}
}

func testIf2() {
	num := 99
	if num < 50 {
		fmt.Printf("number is less than 50")
	} else if num > 50 && num < 100 {
		fmt.Printf("number is between 50 and 100")
	} else {
		fmt.Printf("number is greater than 100")
	}
}

func getNum() int {
	return 10
}

func testIf3() {
	if num := getNum(); num%2 == 0 {
		fmt.Printf("num:%d is even\n", num)
	} else {
		fmt.Printf("num:%d is odd\n", num)
	}
}
func main() {
	//testIf1()
	testIf3()
}
