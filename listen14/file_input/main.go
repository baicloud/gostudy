package main

import (
	"fmt"
	"os"
)

func main() {
	var a int
	var b string
	fmt.Fscanf(os.Stdin, "%d%s", &a, &b)
	fmt.Println(a,b)
	fmt.Fprintf(os.Stdout,"a=%d,b=%s",a,b)
}
