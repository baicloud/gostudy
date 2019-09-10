package main

import (
	"fmt"
)

func testSlice0() {
	var a []int // 定义一个切片
	if a == nil {
		fmt.Printf("a is nil\n")
	} else {
		fmt.Printf("a = %v\n", a)
	}
	//a[0] = 100
}

func testSlice1() {
	a := [5]int{76, 77, 78, 79, 80}

	var b []int = a[1:4] //基于数组a创建⼀一个切⽚片，包括元素a[1] a[2] a[3]

	fmt.Println(b)
	fmt.Printf("type of a %T\n", a)
}

func testSlice2() {
	c := []int{6, 7, 8} // 创建⼀一个数组并返回⼀一个切⽚片
	fmt.Println(c)
	fmt.Printf("type of c %T", c)
}

func testSlice3() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Printf("slice b %v\n", b)
	c := a[2:]
	fmt.Printf("slice c %v\n", c)
	d := a[:4]
	fmt.Printf("slice d %v\n", d)
	e := a[:]
	fmt.Printf("slice e %v\n", e)
}

func testSlice4() {
	// 创建⼀一个数组，其中[...]是编译器器确定数组的⻓长度,darr的⻓长度是 9
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}

	//基于darr 创建⼀一个切⽚片 dslice,包括darr[2],darr[3],darr[4] 三个元素
	dslice := darr[2:5]
	darr[0] = 30
	fmt.Println("array before", darr)

	for i := range dslice {
		//对于 dslice中每个元素进⾏行行 +1，其实修改是darr[2],darr[3],darr[4]
		dslice[i]++
	}

	fmt.Println("array after", darr)

}

func testSlice5() {
	numa := [3]int{78, 79, 80}
	//创建⼀一个切⽚片，包含整个数组的所有元素

	nums1 := numa[:]
	nums2 := numa[:]

	fmt.Println("array before change 1", numa)

	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101

	fmt.Println("array after modification to slice nums2", numa)
}

func testSlice6() {
	// 使用make创建切片
	// var 切片名 []type = make([], len, [cap])；参数说明：type是数据类型、len是大小、cap是切片容量（容量必须>=长度）
	i := make([]int, 5, 5)
	var j []int = make([]int, 6, 6)
	fmt.Println(i)
	fmt.Println(j)
}

func testSlice7() {
	// 切片的容量和长度
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	sliceb := a[1:4]
	fmt.Println(sliceb)
	fmt.Printf("len=%v,cap=%v\n", len(sliceb), cap(sliceb))
	c := sliceb[:len(sliceb)]
	fmt.Printf("len=%v,cap=%v\n", len(c), cap(c))
}

func testSlice8() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	//长度和容量量都等于3

	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	//容量量等于6
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))
}

func testSlice9() {
	//定义names是⼀一个空切⽚片，⻓长度和容量量都等于0
	//不能对空切⽚片进⾏行行访问，否则panic
	var names []string

	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")

		fmt.Println("names contents:", names)
	}
}

func testSlice10() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}
	b = append(b, a...)
	fmt.Printf("b=%v", b)
}

func subtactOne(numbers []int) {

	for i := range numbers {
		numbers[i] -= 2
	}
}

func testSlice11() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("change before %v\n", a)
	// 切片传参
	subtactOne(a)
	fmt.Printf("change after %v", a)
}

func testSlice12() {
	var a []int = []int{1}
	var b []int = []int{4, 5, 6}
	// 切片拷贝，对应位置拷贝
	copy(a, b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

func sumArray(a []int) int {
	var sum int = 0
	for _, v := range a {
		sum = sum + v
	}
	return sum
}

func testSumArray() {
	var a [10]int = [10]int{1, 3, 3, 4, 5, 5, 8}
	sum := sumArray(a[:])
	fmt.Println("sum:", sum)
}
func main() {
	//testSlice0()
	//testSlice1()
	//testSlice2()
	//testSlice3()
	//testSlice4()
	//testSlice5()
	//testSlice6()
	//testSlice7()
	//testSlice8()
	//testSlice9()
	//testSlice10()
	//testSlice11()
	//testSlice12()
	testSumArray()
}
