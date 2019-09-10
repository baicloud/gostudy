package main

import "fmt"

func testPoint1() {
	var a int32
	a = 1000
	fmt.Printf("the addr of a :%p, a:%d\n", &a, a)

	var b *int32
	fmt.Printf("the addr of b: %p, b:%v\n", &b, b)
	if b == nil {
		fmt.Println("b is nil addr")
	}
	//*b = 100
	b = &a
	fmt.Printf("the addr of b: %p, b:%v\n", &b, b)

	c := 255

	var d *int = &c

	fmt.Printf("Type of c is %T\n", c)

	fmt.Println("address of c is", d)
}

func testPoint2() {
	var a int = 200
	var b *int = &a

	fmt.Printf("b指向的地址存储的值为:%d\n", *b)
	*b = 1000
	fmt.Printf("b指向的地址存储的值为:%d\n", *b)
	fmt.Printf("a = %d\n", a)
}

func testPoint3() {
	var a int = 21
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b is", b)
	}
}

func testPoint4() {
	// 地址指针变量指向地址的值
	var a int = 43
	var b *int = &a
	fmt.Printf("a的地址 %p\n", b)
	fmt.Println("b指向地址存储的值", *b)
}

func motify(val *int) {
	*val = 100
}

func testPoint5() {
	var a int = 12
	fmt.Println("a is", a)
	motify(&a)
	fmt.Printf("被改变后a的值 %d", a)
}

func motify_arr(val *[3]int) {
	(*val)[1]=44
}
func testPoint6() {
	a := [3]int{11,22,33}
	fmt.Println(a)
	b:=&a
	motify_arr(b)
	fmt.Println(a)

}

func modify_slice(val []int) {
	val[0]=10

}

func testPoint7(){
	a := [3]int{11,22,33}
	fmt.Println(a)
	modify_slice(a[:])
	fmt.Println(a)
}

func main() {
	//testPoint1()
	//testPoint2()
	//testPoint3()
	//testPoint4()
	//testPoint5()
	//testPoint6()
	testPoint7()
}
