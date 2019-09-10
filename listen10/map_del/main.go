package main

import "fmt"

func testMapDel1() {
	a := map[string]int{
		"stu01":1000,
		"stu02":2000,
		"stu03":3000,
	}
	fmt.Println("map before deletion:",a)
	delete(a,"stu02")
	fmt.Println("map after deletion:",a)
}

func testMapDel2() {
	var a map[string]int
	fmt.Printf("a:%v\n", a)
	//a["stu01"] = 100
	if a == nil {
		a = make(map[string]int, 16)
		fmt.Printf("a=%v\n", a)
		a["stu01"] = 1000
		a["stu02"] = 1000
		a["stu03"] = 1000
		fmt.Printf("a=%#v\n", a)
		delete(a, "stu02")
		fmt.Printf("a=%#v\n", a)

		for key, _ := range a {
			delete(a, key)
		}
		fmt.Printf("after delete a=%#v\n", a)
	}
}

func testMapLength() {
	a := map[string]int{
		"stu01":1000,
		"stu02":2000,
		"stu03":3000,
	}
	fmt.Printf("length of map is %d",len(a))
}

func testMap() {
	// map是引用类型
	a := map[string]int{
		"stu01":1000,
		"stu02":2000,
		"stu03":3000,
	}
	fmt.Printf("a=%#v\n", a)
	b:=a
	b["stu01"]=4000
	fmt.Printf("b=%#v\n", b)
	fmt.Printf("a=%#v\n", a)

}

func main() {
	//testMapDel2()
	//testMapLength()
	testMap()
}
