package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	alex := person {
		firstName : "travis",
		lastName : "head",
		contact:  contactInfo {
			  email: "travis@gmail.com",
			  zipCode: 1234,
		},

	}

	
	alex.updateFirstName()
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v",p)
}

func (pointerToPerson *person) updateFirstName() {
	(*pointerToPerson).firstName = "Pat"
}
	