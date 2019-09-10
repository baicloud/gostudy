package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func testBufio() {
	file, err := os.Open("./bufio.go")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file failed,err:", err)
			return
		}
		fmt.Println(line)
	}
}

func main() {
	testBufio()
}
