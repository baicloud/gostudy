package main

import (
	"fmt"
	"strings"
)

func statWordCount(s string) map[string]int{
	fmt.Println(s)
	words := strings.Split(s, " ")
	fmt.Println(words)
	var result map[string]int = make(map[string]int, 128)
	for _, value := range words {
		count, ok := result[value]
		if !ok {
			result[value] = 1
		} else {
			result[value] = count + 1
		}
	}
	return result
}

func main() {
	s := "how do you do"
	result := statWordCount(s)
	fmt.Println(result)
}
