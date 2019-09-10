package main

import "fmt"

func testMapNil() {
	//map类型的变量默认初始化为nil
	var a map[string]int
	fmt.Println(a)
	//a["stu"]=1
	if a==nil {
		fmt.Println("map is nil,go to make one")
		a=make(map[string]int,16)
		a["stu01"]=001
		a["stu02"]=002
		fmt.Printf("%#v",a)
	}

}

func main() {
	testMapNil()
}
