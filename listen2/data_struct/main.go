package main

import (
	"fmt"
	"strings"
)

func test_bool() {
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)

	if a == true {
		fmt.Println("a is ture")
	} else {
		fmt.Println("a is false")
	}
	b := !a
	fmt.Printf("b is %t", b)

}

func testInt() {
	var a int8
	a = 18
	fmt.Println("a=", a)
	a = -12
	fmt.Println("a=", a)

	var b int
	b = 138338338

	b = int(a)
	fmt.Println("b=", b)

	var c = 5.6
	fmt.Println(c)

	fmt.Printf("a=%d a=%x c=%f\n", a, a, c)
}

func test_str() {
	var a string
	fmt.Println(a)
	a = "hello a"
	fmt.Println(a)
	str := "http://baidu.com"
	str_len := len(str)
	fmt.Println(str_len)
	str_c := "/"
	var d string
	d = str + str_c
	fmt.Println(d)
	g := fmt.Sprint(str, "/", "index.html")
	fmt.Println(g)
	ips := "192.168.56.71, 192.168.56.72, 192.168.56.73"
	fmt.Println(strings.Split(ips, ","))
	fmt.Println(strings.Split(ips, ",")[1])

	if strings.HasPrefix(str, "http"){
		fmt.Println("it is url")
	}else {
		fmt.Println("it is not url")
	}

	if strings.HasSuffix(str, "com"){
		fmt.Println("it is TLD com")
	}else {
		fmt.Println("it is not TLD com")
	}

	fmt.Println(strings.Index(str, "baidu"))

	var strArr []string  = []string{"10.237.8.2", "10.237.8.3", "10.237.8.3"}
	resultStr := strings.Join(strArr, ";")
	fmt.Printf("result=%s\n", resultStr)
}
func main() {
	//test_bool()
	//testInt()
	test_str()

}
