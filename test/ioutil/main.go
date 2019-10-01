package main

import (
	"fmt"
	"io/ioutil"
)

func test_writefile() {
	content := "hello world"
	if err := ioutil.WriteFile("./text.txt", []byte(content), 0644); err != nil {
		fmt.Println(err)
	}
	fmt.Println("write file success")
}

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
