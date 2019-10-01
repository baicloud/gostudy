package main

import (
	"fmt"

	"github.com/baicloud/gostudy/listen6/calc"
)

func main() {
	fmt.Println(calc.A)
	result := calc.Add(2, 8)
	fmt.Printf("result:%d\n", result)
}
