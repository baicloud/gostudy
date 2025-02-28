package main

import "fmt"

type Animal struct {
	Name string
	Sex string
}

func (a *Animal) Talk() {
	fmt.Printf("i'talk, i'm %s\n", a.Name)
}

type PuruAnimal struct {

}

func (p *PuruAnimal) Talk() {
	fmt.Println("buru dongwu talk")
}

type Dog struct {
	Feet string
	*Animal
	*PuruAnimal
}

func (d *Dog) Eat() {
	fmt.Println("dog is eat")
}

func main() {
	var d *Dog=&Dog{
		Feet: "four feet",
		Animal:&Animal{
			Name: "dog",
			Sex: "xigong",
	 	},
	}
	d.Eat()
	d.Animal.Talk()
	d.PuruAnimal.Talk()
}