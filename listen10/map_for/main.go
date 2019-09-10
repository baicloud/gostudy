package main

import "fmt"

func testMapFor() {
	a := map[string]int{
		"stu01":1000,
		"stu02":2000,
		"stu03":3000,
	}
	a["stu04"]=4000
	for key,value := range a{
		fmt.Printf("key=%s,value=%d\n",key,value)
	}
}

func main() {
	testMapFor()
}
