package main

import "fmt"


type People struct {
	Name string
	Country string
}

func (p People) Print() {
	fmt.Printf("name=%s country=%s\n",p.Name,p.Country)
}

func (p People) Set(name string,country string) {
	p.Name = name
	p.Country=country
}

func (p *People) SetV2(name string,country string) {
	p.Country = country
	p.Name = name
}

func main() {
	var p1 People=People{
		Name: "plp01",
		Country:"ch",
	}
	p1.Print()
	p1.Set("plp02","en")
	p1.Print()
	p1.SetV2("p2","c2")
	p1.Print()
}