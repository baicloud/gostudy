package main

import "fmt"

func describe(a interface{}) {
	fmt.Printf("%T %v\n",a,a)
}

type Student struct {
	Name string
	Sex int
}

func main() {
	a:=65
	describe(a)
	str:="hello"
	describe(str)
	var stu Student = Student{
		Name: "user01",
		Sex: 1,
	}
	describe(stu)
}
