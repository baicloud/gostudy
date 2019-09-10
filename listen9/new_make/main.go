package main

import "fmt"

func test_new() {
	//new⽤用来分配除引⽤用类型的所有其他类型的内存，⽐比如 int、数组等
	var a *int = new(int)
	fmt.Println(*a)
	*a=100
	fmt.Println(*a)

	var b *[4]int=new([4]int)
	fmt.Println(b)
	(*b)[1]=1

	var c *[]int = new([]int)
	fmt.Printf("*c = %v\n", *c)

	(*c) = make([]int, 5, 100)
	fmt.Println(*c)
	(*c)[0] = 100
	(*c)[1] = 200
	fmt.Printf("*c = %v\n", *c)
}

func main() {
	test_new()
}
