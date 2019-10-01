package main

import (
	"fmt"
	"reflect"
)

func reflect_example1() {
	var x float32 = 32.4
	fmt.Println("type", reflect.TypeOf(x))

}

func reflect_example2(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("type of a is: %v\n", t)

	k := t.Kind()
	fmt.Printf("type of a kind is: %v\n", k)

	switch k {
	case reflect.Int64:
		fmt.Printf("a is int64\n")
	case reflect.String:
		fmt.Printf("a is string\n")
	}
}

type Stu struct {
}

func test_type_kind() {
	/*type类型与kind类型的区别
	kind获取的是基础类型，type获取的是类型，如type Stu struct
	type输出的类型是Stu,kind输出的是struce
	*/
	var v Stu
	typeV := reflect.TypeOf(v)
	fmt.Printf("type=%v,kind=%v", typeV.Name(), typeV.Kind())
}

func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()
	fmt.Printf("a of type: %v\n", t)
	k := v.Kind()
	fmt.Printf("a of kind: %v\n", k)
	switch k {
	case reflect.Int64:
		fmt.Printf("a is int64, store value is: %d\n", v.Int())
	case reflect.Float64:
		fmt.Printf("a is float64, store value is: %f\n", v.Float())
	}
}

func test_set_reflect_value() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	fmt.Println(v)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is float64:", v.Kind())
	v.Elem().SetFloat(6.8)
	fmt.Println("value:", x)
}

func reflect_set_value(a interface{}) {
	v:=reflect.ValueOf(a)
	fmt.Println(v)
	k:=v.Kind()
	switch k {
	case reflect.Int64:
		v.SetInt(100)
		fmt.Printf("a is int64, store value is:%d\n", v.Int())
	case reflect.Float64:
		v.SetFloat(6.9)
		fmt.Printf("a is float64, store value is:%f\n", v.Float())
	case reflect.Ptr:
		fmt.Printf("set a to 6.8\n")
		v.Elem().SetFloat(6.8)
	default:
		fmt.Printf("default switch\n")
	}
}

func main() {
	//reflect_example1()
	var x float64=4.4
	//reflect_example2(x)
	//test_type_kind()
	//reflect_value(x)
	test_set_reflect_value()
	reflect_set_value(&x)
}
