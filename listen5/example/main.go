package main

import "fmt"

func justify(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeNumber() {
	for i := 2; i <= 100; i++ {
		test := justify(i)
		if test == true {
			fmt.Printf("[%d] is prime number\n", i)
		}
	}
}

func shuiXianHua(n int) bool {
	first := n % 10
	second := (n / 10) % 10
	thrid := (n / 100) % 100
	//fmt.Printf("%d,%d,%d",first,second,thrid)
	if first*first*first+second*second*second+thrid*thrid*thrid == n {
		return true
	}
	return false
}

func isShuiXianHua() {
	for i := 100; i <= 999; i++ {
		bool1 := shuiXianHua(i)
		if bool1 == true {
			fmt.Printf("%d is 水仙花数\n", i)
		}
	}
}

func calc(s string) (charCount int, numCount int, spaceCount int, otherCount int) {
	sRune := []rune(s)
	for i := 0; i < len(sRune); i++ {
		if (sRune[i] >= 'a' && sRune[i] <= 'z' || sRune[i] >= 'A' && sRune[i] <= 'Z') {
			charCount++
			continue
		}

		if (sRune[i] >= '0' && sRune[i] <= '9') {
			numCount++
			continue
		}

		if (sRune[i] == ' ') {
			spaceCount++
			continue
		}

		otherCount++
	}
	return
}

func testCalc() {
	str := "jalfja  中国 啊2231SAjdl"
	charCount, numCount, SpaceCount, otherCount := calc(str)
	fmt.Printf("charCount=%d,numCount=%d,SpaceCount=%d,otherCount=%d", charCount, numCount, SpaceCount, otherCount)
}
func main() {
	//primeNumber()
	//shuiXianHua(192)
	//isShuiXianHua()
	testCalc()
}
