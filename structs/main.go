package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
