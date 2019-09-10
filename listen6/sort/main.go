package main

import (
	"fmt"
)

func insert_sort(a [8]int) [8]int {
	// 插入排序
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
	return a

}
func select_sort(a [8]int) [8]int {
	// 选择排序
	for i:=0;i<len(a);i++{
		for j:=i+1;j<len(a);j++{
			if a[i] > a[j]{
				a[i],a[j] = a[j],a[i]
			}
		}
	}
	return a
}

func bubble_sort(a [8]int) [8]int {
	// 冒泡排序
	for i:=0;i<len(a);i++{
		for j:=0;j<len(a)-i-1;j++  {
			if a[j]>a[j+1]{
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
	return a
}

func main() {
	var i [8]int = [8]int{29, 13, 0, 4, 1, 42, 12, 32}
	j := bubble_sort(i)
	fmt.Println(i)
	fmt.Println(j)
}
