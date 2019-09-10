package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

const (
	NumStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

func parseArgs() {
	flag.IntVar(&length, "l", 16, "密码长度")
	flag.StringVar(&charset, "t", "num",
		`-t 制定密码生成的字符集,
		num:只使用数字[0-9], 
		char:只使用英文字母[a-zA-Z], 
		mix: 使用数字和字母， 
		advance:使用数字、字母以及特殊字符`)
	flag.Parse()
	//fmt.Printf("长度%d，字符集%s",length,charset)
}

func generatePasswd() string {
	var password []byte = make([]byte, length, length)
	var sourceStr string
	if charset == "num" {
		sourceStr = NumStr
	} else if charset == "char" {
		sourceStr = CharStr
	} else if charset == "mix" {
		sourceStr = NumStr + CharStr
	} else if charset == "advance" {
		sourceStr = NumStr + CharStr + SpecStr
	} else {
		sourceStr = NumStr
	}
	fmt.Println(sourceStr)
	for i := 0; i < length; i++ {
		index :=rand.Intn(len(sourceStr))
		password[i]=sourceStr[index]
	}

	return string(password)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano())
	parseArgs()
	fmt.Printf("长度%d，字符集%s\n", length, charset)
	passwd := generatePasswd()
	fmt.Printf("随机生成的密码是: %s\n",passwd)
}
