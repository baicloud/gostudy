package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//测试bufio write
	f, err := os.OpenFile("./1.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file faile,err:", err)
		return
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello i")
	}
	writer.Flush()
}
