package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name" db:"name"`
	Sex int
	Age int
	Score float32
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s *Student) Print() {
	fmt.Printf("通过反射进行调用:%#v\n", s)
}

func main()  {
	var s Student
	s.SetName("stu01")
	v:=reflect.ValueOf(&s)
	t:=v.Type()
	fmt.Println(t)
	field0:=t.Elem().Field(0)

	fmt.Println(v.Elem().Field(0))
	fmt.Println(field0)
	fmt.Printf("tag json=%s\n",field0.Tag.Get("json"))
	fmt.Printf("tag db=%s\n",field0.Tag.Get("db"))

}