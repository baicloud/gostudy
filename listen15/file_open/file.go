package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func testOpenFile1() {
	f, err := os.Open("./1.txt")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	fmt.Println(string(b))
}
func testOpenFile2() {
	file, err := os.Open("./1.txt")
	if err != nil {
		fmt.Println("open file failed,err:", err)
	}
	defer file.Close()
	var buf [128]byte
	n, _ := file.Read(buf[:])
	fmt.Println(string(buf[:]), n)
}
func testOpenFile3() {
	file, err := os.Open("./1.txt")
	if err != nil {
		fmt.Println("open file failed,err:", err)
	}
	defer file.Close()
	var content []byte
	var buf [128]byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file:", err)
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}
func main() {
	// testOpenFile()
	testOpenFile2()
}
