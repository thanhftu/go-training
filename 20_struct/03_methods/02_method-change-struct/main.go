package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) changeName() {
	p.first = "Thanh"
	p.last = "Truong"
	p.age = 28
}

func (p *person) changeName2() {
	p.first = "Thanh"
	p.last = "Truong"
	p.age = 28
}

func main() {
	p1 := person{"James", "Bond", 20}
	fmt.Println(p1) // {James Bond 20}
	// cant change p1
	p1.changeName()
	fmt.Println(p1) // {James Bond 20}

	p2 := person{"James", "Bond", 20}
	p2.changeName2()
	fmt.Println(p2)
}
