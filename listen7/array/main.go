package main

import (
	"fmt"
)

func testArray1() {
	//数组初始化
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Printf("数组a: %d\n", a)
	// 定义时初始化
	var b [4]int = [4]int{4, 5, 6, 7}
	fmt.Printf("数组b: %d\n", b)
	c := [5]int{8, 9, 10, 11, 12}
	fmt.Printf("数组c: %d\n", c)
	d := [...]int{13, 14, 15, 16, 17, 18}
	fmt.Printf("数组d: %d\n", d)
	//定义式初始化指定位置
	e := [7]int{1: 10, 4: 20}
	fmt.Printf("数组e: %d\n", e)
}

func testArray2() {
	a := [5]int{3: 100, 4: 300}
	fmt.Println(a, len(a))

	var b [5]int
	// 数组赋值，只有同类型，同长度可以赋值
	b = a
	fmt.Println(b)
}

func testArray3() {
	// 数组遍历
	a := [6]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]=%d\n", i, a[i])
	}
	fmt.Println()
	// 数组遍历，range
	for index, val := range a {
		fmt.Printf("a[%d]=%d\n", index, val)
	}
}

func testArray4() {
	var a [3][2]int
	a[0][0] = 10
	a[0][1] = 20
	a[1][0] = 30
	a[1][1] = 30
	a[2][0] = 30
	a[2][1] = 30
	fmt.Println(a)
	//遍历
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}

	// 遍历，range
	fmt.Println("other method")
	for i, val := range a {
		fmt.Printf("row[%d]=%v\n", i, val)
		for j, val2 := range val {
			fmt.Printf("(%d,%d)=%d ",i, j, val2)
		}
		fmt.Println()
	}
}

func testArray5() {
	//变量是值类型，赋值相当于拷贝一份存放在内存中
	var a int = 1000
	b := a
	b = 3000
	fmt.Printf("a=%d b=%d\n", a, b)
}



func testArray6() {
	// 数组是值类型，赋值相当于拷贝一份
	a := [3]int{10, 20, 30}
	b := a
	b[0] = 1000
	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v\n", b)
}

func modify(b [3]int) {
	b[0] = 1000
}

func testArray7() {
	// 数组是值类型，函数参数也会拷贝一份,原来数组a不变
	var a [3]int = [3]int{10, 20, 30}
	modify(a)
	fmt.Println(a)

}
func main() {
	//testArray1()
	//testArray2()
	//testArray3()
	//testArray4()
	//testArray5()
	//testArray6()
	testArray7()
}
