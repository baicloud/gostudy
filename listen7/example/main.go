package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumArray(a [10]int) int {
	//求数组所有元素之和
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}
	return sum
}

func testSumArray() {
	// 初始化随机种子
	rand.Seed(time.Now().Unix())
	var b [10]int
	for i := 0; i < len(b); i++ {
		//产生一个0到999的随机数
		b[i] = rand.Intn(1000)
		//产生一个0到Int最大值的随机数
		//b[i] = rand.Int()
		fmt.Printf("b[%d]=%d\n",i,b[i])
	}

	sum := sumArray(b)
	fmt.Printf("sum=%d\n", sum)
}

func TwoSum(a [5]int, target int) {
	for i:=0;i<len(a);i++{
		for j:=i+1;j<len(a);j++{
			if a[i]+a[j] == target{
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}

func testTwoSum() {
	a :=[5]int{1,3,5,8,7}
	TwoSum(a,8)
}

func main() {
	//testSumArray()
	testTwoSum()
}
