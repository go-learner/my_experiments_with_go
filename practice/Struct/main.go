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

	//alex := person{"Alex", "Andreson"}

	//var alex person

	//alex.firstName = "alex"
	//alex.lastName = "andreson"
	//alex.contact.email = "alex.andreson@gmail.com"
	//alex.contact.zipCode = 12345

	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)

	lionel := person{
		firstName: "lionel",
		lastName:  "Messi",
		contact: contactInfo{
			email:   "lionel.Messi@gmail.com",
			zipCode: 12345,
		},
	}
	//leoPointer := &lionel
	lionel.updateFirstName("leo")
	lionel.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

}
