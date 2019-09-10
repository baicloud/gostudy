package main

import "fmt"

func testFor1() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func testFor2() {
	var i int
	for i = 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("final:i=%d", i)
}

func testFor3() {
	var i int
	for i = 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("i=%d\n", i)
	}
	fmt.Printf("final:i=%d\n", i)
}
func testFor7() {
	for {
		fmt.Printf("hello\n")
	}
}
func main() {
	// testFor1()
	// testFor2()
	// testFor3()
	testFor7()
}
