package main

import "fmt"

func testInitMap1() {
	//创建时初始化
	var a map[string]int = map[string]int{
		"stu01": 0001,
		"stu02": 0002,
		"stu03": 0003,
	}
	fmt.Println(a)
	fmt.Printf("%#v\n", a)
	fmt.Println("-----end-----")

	//map插入操作
	a["stu04"] = 0004
	fmt.Printf("%#v\n", a)
}

func testInitMap2() {
	//使用make初始化
	var a map[string]int
	a = make(map[string]int)
	a["tec01"] = 0001
	a["tec02"] = 0002
	a["tec03"] = 0003
	fmt.Println(a)

}

func testAccessMap() {
	// 通过key值访问map
	var a map[string]int = map[string]int{
		"stu01": 1000,
		"stu02": 2000,
		"stu03": 3000,
	}
	fmt.Printf("student number of student stu01 is %d\n", a["stu01"])
	b := "123"

	fmt.Println("Salary of", b, "is", a[b])
}

func main() {
	//testInitMap1()
	//testInitMap2()
	testAccessMap()
}
