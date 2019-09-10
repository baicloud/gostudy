package main

import "fmt"

func testSwap1(a int, b int) {
	fmt.Printf("a=%d,b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("a=%d,b=%d\n", a, b)
}

func testSwap2(a *int,b *int) {
	fmt.Printf("a=%d,b=%d\n", *a, *b)
	*a, *b = *b, *a
	fmt.Printf("a=%d,b=%d\n", *a, *b)
}

func main() {
	var a int = 11
	var b int = 22
	// int类型为值类型，所以函数内部a,b值交换了，但是在main函数里面a,b的值不变
	testSwap1(a, b)
	fmt.Printf("a=%d,b=%d\n", a, b)
	fmt.Println()
	// 指针为引用类型，指针变量里的值修改了，main函数里面的值也就被修改了
	testSwap2(&a, &b)
	fmt.Printf("a=%d,b=%d\n", a, b)
}
