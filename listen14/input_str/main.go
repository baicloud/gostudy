package main

import "fmt"

func testSscanf() {
	var a int
	var b string
	var c float32

	fmt.Sscanf("1 ajl 1.1", "%d%s%f", &a, &b, &c)

	//fmt.Scanf("%d\n", &a)
	//fmt.Scanf("%s\n", &b)
	//fmt.Scanf("%f\n", &c)
	//fmt.Scanf("%d %s %f",&a,&b,&c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func testSscan() {
	var a int
	var b string
	var c float32
	var str string = "1 abc\n 1.4"
	fmt.Sscan(str,&a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func testSscanln() {
	var a int
	var b string
	var c float32
	var str string = "1 \nabc 1.5"
	fmt.Sscanln(str,&a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)

}

func main() {
	//testSscanf()
	//testSscan()
	testSscanln()
}
