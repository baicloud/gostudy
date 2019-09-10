package main

import (
	"fmt"
)

func testSwitch1() {
	finger := 2
	switch finger {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("输入值非法")
	}
}

func testSwitch2() {
	switch letter := "p"; letter {
	case "a", "b", "c":
		fmt.Println("abc")
	case "h", "j", "i":
		fmt.Println("hji")
	default:
		fmt.Println("default")
	}
}

func testSwitch3() {
	// 条件判断
	num := 75
	switch {
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50 ")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	default:
		fmt.Println("输入值非法")
	}
}

func testMulti() {
	for i:=1;i<=9;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%d * %d = %d\t",i,j,i*j)
		}
		fmt.Println()
	}
}
func main() {
	// testSwitch1()
	//testSwitch3()
	testMulti()
}
