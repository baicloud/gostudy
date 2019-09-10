package main

import "fmt"

type Employer interface {
	CalcSalary() float32
}

type Programer struct {
	name  string
	base  float32
	extra float32
}

func NewProgramer(name string, base float32, extra float32) Programer {
	return Programer{
		name:  name,
		base:  base,
		extra: extra,
	}
}

func (p Programer) ClacSlary() float32 {
	return p.base
}

func main() {
}
