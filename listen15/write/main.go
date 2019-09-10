package main

import "fmt"
import "os"

func main() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	str := "hello 1"
	file.Write([]byte(str))
	file.WriteString(str)
}
