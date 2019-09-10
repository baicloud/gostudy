package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile() (written int64, err error) {
	var srcFile string = "./main.go"
	var desFile string = "./1.txt"
	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("open file faile,err:", err)
		return
	}
	defer src.Close()
	des, err := os.OpenFile(desFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("openfile failed,err:", err)
		return
	}
	defer des.Close()
	return io.Copy(des, src)
}

func main() {
	copyFile()
}
