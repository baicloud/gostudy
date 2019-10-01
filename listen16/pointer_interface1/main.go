package main

import "fmt"

type Animal interface {
	Talk()
	Eat()
	Name() string
}

type Dog struct {
}

func (d Dog) Talk() {
	fmt.Println("汪汪汪")
}

func (d Dog) Eat() {
	fmt.Println("dog eat")
}

func (d Dog) Name() string {
	fmt.Println("dog name is wangcai")
	return "wangcai"
}

func main() {
	var a Animal
	var d Dog
	a = d
	a.Eat()
	fmt.Printf("%T %v\n", a, a)
	var d1 *Dog = &Dog{}
	a = d1
	a.Eat()
	fmt.Printf("%T %v\n", a, a)
}
