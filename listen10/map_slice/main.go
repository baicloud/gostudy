package main

import (
	"fmt"
	"math/rand"
	"time"
)

func testMapSlice1() {
	//var mapa map[string]int
	//mapa=make(map[string]int,5)
	//fmt.Println(mapa)
	var mapSlice []map[string]int

	mapSlice=make([]map[string]int,5)
	fmt.Println("map before init")
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)

	}
	//mapSlice[0]=make(map[string]int,5)
	//mapSlice[0]["stu01"]=1000
	//fmt.Println(mapSlice[0])
	mapSlice[0] = make(map[string]int, 10)
	mapSlice[0]["a"] = 1000
	mapSlice[0]["b"] = 2000
	mapSlice[0]["c"] = 3000
	mapSlice[0]["d"] = 4000
	mapSlice[0]["e"] = 5000
	fmt.Println("after map init")
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

func mapSlice() {
	rand.Seed(time.Now().UnixNano())
	var s map[string][]int
	s = make(map[string][]int, 16)
	key := "stu01"
	value, ok := s[key]
	if !ok {
		s[key] = make([]int, 0, 16)
		value = s[key]
	}

	value = append(value, 100)
	value = append(value, 200)
	value = append(value, 300)
	s[key] = value
	fmt.Printf("map:%v\n", s)
}

func main() {
	//testMapSlice1()
	mapSlice()
}
