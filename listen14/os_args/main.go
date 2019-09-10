package main

import (
	"fmt"
	"os"
)

func main() {
	a:=os.Args
	fmt.Println(os.Args[0])
	if len(a)>1{
		for index,value:=range a{
			if index == 0 {
				continue
			}
			fmt.Printf("index=%d,value=%s\n",index,value)
		}
	}
}
