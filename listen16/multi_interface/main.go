package main

import "fmt"

type Animal interface {
	Talk()
	Eat()
	Name() string
}

type PuruDongWu interface {
	Taisheng()
}

type Dog struct {
}

func (d Dog) Talk() {
	fmt.Println("dog voice")
}

func (d Dog) Eat() {
	fmt.Println("我在吃骨头")
}

func (d Dog) Name() string {
	fmt.Println("dog is name")
	return "name"
}

func (d Dog) Taisheng() {
	fmt.Println("狗是taisheng")
}

func main() {
	var d Dog
	var a Animal
	fmt.Println("%v %T %p", a, a, a)
	if a == nil {
		fmt.Println("a is nil")
	}
	a = d
	a.Eat()
	var b PuruDongWu
	b = d
	b.Taisheng()
	d.Taisheng()
}
