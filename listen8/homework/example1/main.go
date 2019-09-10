package main

import (
	"fmt"
	"sort"
)

func testSort() {
	var a [5]int = [5]int{5, 4, 3, 2, 1}
	sort.Ints(a[:])
	fmt.Printf("a:%v\n", a)
	var b [5]string = [5]string{"ew", "ef", "aa", "la", "ek"}
	sort.Strings(b[:])
	fmt.Printf("b:%v\n", b)
	var c [5]float64 = [5]float64{29.38, 22.32, 0.8, 99191.2}
	sort.Float64s(c[:])

	fmt.Println("c:", c)
}

func main() {
	testSort()
}
