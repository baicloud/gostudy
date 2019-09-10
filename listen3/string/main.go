package main

import (
	"fmt"
)

func testString() {
	var str string = "hello"
	fmt.Println(str)
	fmt.Printf("str[0]=%c len(str)=%d\n", str[0], len(str))
	for index, val := range str {
		fmt.Printf("str[%d]=%c\n", index, val)
	}

	var byteSlice = []byte(str)
	//fmt.Println(byteSlice)
	byteSlice[0] = 'p'
	str = string(byteSlice)
	fmt.Println(str)

	str = "hello,世界"
	fmt.Printf("len(str)=%d\n", len(str))

	var b rune = '国'
	fmt.Printf("b=%c\n", b)

	var runeSlice []rune
	runeSlice = []rune(str)
	fmt.Printf("str 长度:%d, len(str)=%d\n", len(runeSlice), len(str))

}

func testReverseStringV1() {
	var str = "hellogo"
	fmt.Println(str)
	var byteSlice []byte
	byteSlice = []byte(str)
	fmt.Println(byteSlice)
	for i := 0; i < len(byteSlice)/2; i++ {
		var tmp = byteSlice[i]
		byteSlice[i] = byteSlice[len(byteSlice)-i-1]
		byteSlice[len(byteSlice)-i-1] = tmp
	}
	str = string(byteSlice)
	fmt.Println(str)
}

func testReverseStringV2() {
	var str = "你好中国abcd"
	var strRune = []rune(str)
	fmt.Println(len(strRune))
	for i := 0; i < len(strRune)/2; i++ {
		var tmp = strRune[i]
		strRune[i] = strRune[len(strRune)-i-1]
		strRune[len(strRune)-i-1] = tmp
	}
	str = string(strRune)
	fmt.Println(str)

}

func testHuiWen() {
	var str = "上海自来水来自海上"
	var r []rune = []rune(str)
	for i := 0; i < len(r)/2; i++ {
		var tmp = r[i]
		r[i] = r[len(r)-i-1]
		r[len(r)-i-1] = tmp
	}
	var str1 = string(r)
	if str == str1 {
		fmt.Println("it is huiwen")
	}

}

func main() {
	// testString()
	// testReverseStringV2()
	testHuiWen()

}
