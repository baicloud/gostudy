package main

import "fmt"

func funcA() (y int) {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func funcB() (x int) {
	defer func() {
		x += 1
	}()
	return 5
}

func funcC() (y int) {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func funcD() (x int) {
	// 传值相当于复制一份进去
	defer func(x int) {
		x+=1
	}(x)
	return 5
}

func main() {
	fmt.Println(funcA())
	fmt.Println(funcB())
	fmt.Println(funcC())
	fmt.Println(funcD())
}
