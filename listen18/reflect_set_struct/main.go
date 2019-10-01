package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Sex int
	Age int
	Secore float32
}

func main() {
	var s Student
	v:=reflect.ValueOf(&s)
	v.Elem().Field(0).SetString("xing01")
	v.Elem().FieldByName("Sex").SetInt(1)
	v.Elem().FieldByName("Age").SetInt(18)
	v.Elem().FieldByName("Secore").SetFloat(4.5)
	fmt.Printf("s:%#v\n",s)
}