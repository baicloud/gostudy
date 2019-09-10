package main

import "fmt"

func test1() {
	var sa []string = make([]string,5,10)
	fmt.Println("a:", sa)
	for i:=0;i<10;i++{
		sa=append(sa,fmt.Sprintf("%d", i))
	}
	fmt.Println("sa:", sa)
}


func main() {
	test1()
}
