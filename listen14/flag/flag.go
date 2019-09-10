package main

import (
	"flag"
	"fmt"
)

var a int
var b string

func init() {
	flag.IntVar(&a, "l", 18, "输入年龄")
	flag.StringVar(&b, "s", "xing", "姓名")
	flag.Parse()

}

func main() {
	fmt.Printf("name=%s,age=%d\n", b, a)
}
