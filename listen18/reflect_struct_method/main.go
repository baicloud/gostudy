package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Sex   int
	Age   int
	Score float32
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s *Student) Print() {
	fmt.Printf("通过反射进行调用：%#v\n", s)
}

func main() {
	var s Student
	s.SetName("stu01")
	v := reflect.ValueOf(&s)
	fmt.Println(v)
	t:=v.Type()
	//t := reflect.TypeOf(s)
	fmt.Println(t)
	fmt.Printf("struct student have %d methods\n", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("struct %d method,name:%s type:%v\n", i, method.Name, method.Type)
	}
	//通过reflect.Value获取对应方法并调用
	m1:=v.MethodByName("Print")
	var args []reflect.Value
	m1.Call(args)

	m2:=v.MethodByName("SetName")
	var args2 []reflect.Value
	name:="stu02"
	nameVal:=reflect.ValueOf(name)
	args2=append(args2,nameVal)
	m2.Call(args2)
	m1.Call(args)
}
