package main

import "fmt"

func testAssert(i interface{}) {
	s:=i.(int)
	fmt.Println(s)
}

func test(a interface{}) {
	s,ok:=a.(int)
	if ok{
		fmt.Println(s)
		return
	}
	str,ok:=a.(string)
	if ok{
		fmt.Println(str)
		return
	}
	f,ok:=a.(float32)
	if ok{
		fmt.Println(f)
		return
	}
	fmt.Println("can not define the type of a")
}

func testInterface1() {
	var a int=100
	test(a)
	var b string="hello"
	test(b)
}

func testSwitch1(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Printf("a is string,value:%v\n",a.(string))
	case int:
		fmt.Printf("a is int, value:%v\n",a.(int))
	case int32:
		fmt.Printf("a is int, value:%v\n", a.(int))
	default:
		fmt.Println("not support type\n")
	}
}

func testInterface2() {
	var a int=100
	testSwitch1(a)
}

func testSwitch2(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Printf("a is string, value:%v\n", v)
	case int:
		fmt.Printf("a is int, value:%v\n", v)
	case int32:
		fmt.Printf("a is int, value:%v\n", v)
	default:
		fmt.Println("not support type\n")
	}
}
func testInterface3() {
	var a int = 100
	testSwitch2(a)
	var b string = "hello"
	testSwitch2(b)
}

func testAssert2(i interface{}) {
	v,ok:=i.(int)
	fmt.Println(v,ok)
}
func main() {
	//var s interface{}=56
	//testAssert(s)
	//testInterface1()
	//testInterface2()
	//testInterface3()

	var s interface{}="hello"
	testAssert2(s)
	var i interface{}=88
	testAssert2(i)
}
