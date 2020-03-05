package main

import (
	"fmt"
	"strings"
)

//Person for full name
type Person struct {
	firstName string
	lastName  string
}

//  parameter of type *Person (so that the object itself can be changed!)
func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 1- struct as a value type:
	var pers1 Person //this is a value
	pers1.firstName = "Thanh"
	pers1.lastName = "Truong"
	upPerson(&pers1)
	fmt.Printf("The name of the person 1 is %s %s\n", pers1.firstName, pers1.lastName)

	// 2- struct as a pointer:
	pers2 := new(Person) // this is a pointer
	pers2.firstName = "Sam"
	pers2.lastName = "Hoang"
	(*pers2).lastName = "Hoang" // this is also valid
	upPerson(pers2)
	fmt.Printf("The name of the person 2 is %s %s\n", pers2.firstName, pers2.lastName)

	//3- struct as a literal
	pers3 := &Person{"Giang", "Truong"}
	upPerson(pers3)
	fmt.Printf("The name of the person 3 is %s %s\n", pers3.firstName, pers3.lastName)

}

/*
Structs in Go and the data they contain, even when a struct contains other structs, form a continuous block in memory:
this gives a huge performance benefit
*/
