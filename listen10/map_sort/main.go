package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func testMapSort1() {
	var a map[string]int = make(map[string]int, 10)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		a[key] = i
	}
	for key, value := range a {
		fmt.Printf("key:%s = %d\n", key, value)
	}
	var keys []string
	for key, _ := range a {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Println(keys)
	for key, value := range a {
		fmt.Printf("key:%s = %d\n", key, value)
	}
	fmt.Println("----分割线------")
	for _, key := range keys {
		fmt.Printf("key:%s = %d\n", key, a[key])
	}

}

func testMapSort2() {
	rand.Seed(time.Now().UnixNano())
	var a map[string]int = make(map[string]int, 1024)

	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("stu%d", i)
		value := rand.Intn(1000)
		a[key] = value
	}

	var keys []string = make([]string, 0, 128)
	for key, value := range a {
		fmt.Printf("map[%s]=%d\n", key, value)
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, value := range keys {
		fmt.Printf("key:%s val:%d\n", value, a[value])
	}
}
func main() {
	//testMapSort1()
	testMapSort2()
}
