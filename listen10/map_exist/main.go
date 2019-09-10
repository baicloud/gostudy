package main

import "fmt"

func judgeKeyExist() {
	a := map[string]int{
		"stu01": 1000,
		"stu02": 2000,
		"stu03": 3000,
	}
	b := "stu01"
	value, ok := a[b]
	fmt.Printf("the value of a is %d\n", value)
	fmt.Println("key是否存在：", ok)
	if ok == true {
		fmt.Println("student num of", b, "is", value)
	} else {
		fmt.Println(b, "not found")
	}
}

func main() {
	judgeKeyExist()
}
