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
	v:=reflect.ValueOf(s)
	fmt.Println("valueof: ",v)
	t:=v.Type()
	fmt.Println("type: ",t)
	k:=t.Kind()
	fmt.Println("kind: ",k)
	switch k {
	case reflect.Int64:
		fmt.Printf("s is int64\n")
	case reflect.Float32:
		fmt.Printf("s is Float32")
	case reflect.Struct:
		fmt.Printf("s is struct\n")
		fmt.Printf("field num of s is %d\n",v.NumField())
		for i:=0;i<v.NumField();i++{
			field:=v.Field(i)
			//fmt.Println("field",field)
			fmt.Printf("name:%s type:%v value:%v\n",
				t.Field(i).Name,field.Type().Kind(),field.Interface())
		}
	default:
		fmt.Printf("default\n")
	}
}
