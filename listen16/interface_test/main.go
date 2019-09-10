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
	fmt.Println("dog voice")
}

func (d Dog) Eat() {
	fmt.Println("dog eat")
}

func (d Dog) Name() string {
	fmt.Println("dog is name")
	return "dog01"
}
func testInterface1() {
	var d Dog
	var a Animal
	a = &d
	a.Eat()
	a.Talk()
	a.Name()
	fmt.Printf("%T %v\n", a, a)
}

func main() {
	testInterface1()
}
