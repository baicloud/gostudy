package main

import (
	"fmt"
)

func main() {
	var a int
	var b bool
	var (
		c string
		d float32
	)
	const f  = 100
	fmt.Printf("a=%d b=%t c=%s d=%f f=%d\n", a, b, c, d, f)

	a = 100
	b = true
	c = "hello world"
	d = 5.6
	fmt.Printf("a=%d b=%t c=%s d=%f\n", a, b, c, d)
}
