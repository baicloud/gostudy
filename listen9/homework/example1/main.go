package main

import "fmt"

func modify(val *int) {
	*val=100
}

func main() {
	var a  int = 3
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)
}
