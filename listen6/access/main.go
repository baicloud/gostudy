package main

import (
	"fmt"
	"github.com/baicloud/gostudy/listen6/calc"
)


func main() {
	// 不能使用calc包小写字母开头的变量或者函数
	sum := calc.Add(3, 5)
	fmt.Printf("sum=%d\n", sum)
	fmt.Println(calc.A)
	fmt.Println(calc.Test())
}
