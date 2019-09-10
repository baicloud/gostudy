package main

import "fmt"

func add(a, b int) int {
	sum := a + b
	return sum
}

func sub(a, b int) int {
	return a- b
}

func testFunc1() {
	//函数也是⼀一种类型，因此可以定义⼀一个函数类型的变量
	f1 := add
	fmt.Printf("type of f1=%T", f1)
}

func testFunc2() {
	// 匿名函数，即没有名字的函数
	f2 := func(a, b int) int {
		return a + b
	}
	sum := f2(12,342)
	fmt.Println(sum)
	fmt.Printf("type of f2=%T", f2)
}

func testFunc3() {
	var i int = 0
	// defer 已经注入了i的值，所以i=0
	defer fmt.Printf("defer i=%d\n", i)
	i = 100
	fmt.Printf("i=%d\n", i)
	return
}

func testFunc4() {
	var i int = 0
	// defer 此时i的值是100,defer的匿名函数是在整个函数执行结束前的一瞬间执行的，即延时执行
	defer func() {
		fmt.Printf("defer i=%d\n", i)
	}()

	i = 100
	fmt.Printf("i=%d\n", i)
	return
}

func testFunc5() {
	var i int = 0

	defer func(i int) {
		fmt.Printf("defer i = %d\n", i)
	}(i)    // -----这里发生了传参

	i=101

	fmt.Printf("i = %d\n", i)
}

func testFunc6(a,b int,op func(a,b int)int) int{
	// 函数作为参数
	return op(a, b)

}

func main() {
	//testFunc1()
	//testFunc2()
	//testFunc3()
	//testFunc5()
	re := testFunc6(1,2,sub)
	fmt.Printf("result=%d",re)
}
