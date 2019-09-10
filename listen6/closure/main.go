package main

import (
	"fmt"
	"time"
)

func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func testClosure1() {
	f1 := Adder()
	fmt.Printf("type of f1 %T\n", f1)
	fmt.Printf("f1(1)=%d\n", f1(1))
	fmt.Printf("f1(1)=%d\n", f1(1))
	fmt.Printf("f1(1)=%d\n", f1(1))
}

func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func testClosure2() {
	f1 := add(20)
	fmt.Printf("f1(3)=%d\n", f1(3))
	fmt.Printf("f1(10)=%d\n", f1(10))
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func testClosure3() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))
	fmt.Println(f1(7), f2(8))
}
func testClosure4() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}

func testClosure5() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second)
}

func testClosure6() {
	for i := 0; i < 5; i++ {
		go func(index int) {
			fmt.Println(index)
		}(i)
	}
	time.Sleep(time.Second)
}
func main() {
	//testClosure1()
	//testClosure2()
	//testClosure3()
	//testClosure4()   //输出为 5 5 5 5 5应为for循环已经执行完毕，go func才开始执行
	//testClosure5()	// 在for循环中加延迟，则输出为0，1，2，3，4
	testClosure6()		//	随机输出0，1，2，3，4；应为go func执行顺序不同
}
