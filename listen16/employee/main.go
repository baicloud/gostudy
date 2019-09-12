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

func (p Programer) CalcSalary() float32 {
	return p.base
}

type Sale struct {
	name  string
	base  float32
	extra float32
}

func NewSale(name string, base, extra float32) Sale {
	return Sale{
		name:  name,
		base:  base,
		extra: extra,
	}
}

func (p Sale) CalcSalary() float32 {
	return p.base + p.extra*p.base*0.5
}

func caclAll(e []Employer) float32 {
	var cost float32
	for _, v := range e {
		cost = cost + v.CalcSalary()
	}
	return cost
}

func main() {
	p1 := NewProgramer("it1", 1500, 0)
	p2 := NewProgramer("it2", 1500, 0)
	p3 := NewProgramer("it3", 1500, 0)

	s1 := NewProgramer("销售1", 800, 2.5)
	s2 := NewProgramer("销售2", 800, 2.5)
	s3 := NewProgramer("销售3", 800, 2.5)
	var employList []Employer
	employList = append(employList, p1)
	employList = append(employList, p2)
	employList = append(employList, p3)
	employList = append(employList, s1)
	employList = append(employList, s2)
	employList = append(employList, s3)
	cost := caclAll(employList)
	fmt.Printf("这个月人力成本：%f\n", cost)
}
