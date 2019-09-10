package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 测试 write ioutil
	str := "hello 2019"
	err := ioutil.WriteFile("./1.txt", []byte(str), 0755)
	if err != nil {
		fmt.Println("write file failed,err:", err)
		return
	}
}
