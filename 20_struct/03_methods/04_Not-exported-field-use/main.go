package main

import (
	"fmt"

	person "github.com/thanhftu/udemytraining/20_struct/03_methods/03_Not-exported-field"
)

func main() {
	p := new(person.Person)
	// p.firstName undefined
	// (cannot refer to unexported field or method firstName)
	// p.firstName = "Thanh"

	p.FirstNameSettler("Thanh")
	fmt.Println(p.FirstNameCaller())
}
