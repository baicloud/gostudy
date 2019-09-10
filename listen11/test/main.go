package main

import "fmt"

func main() {
	s := "hello world"
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}
}
