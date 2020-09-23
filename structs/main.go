package main

import "fmt"

type contactInfo struct {
	zipcode int
	email   string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Pearson",
		contactInfo: contactInfo{
			zipcode: 12345,
			email:   "jimbo@pearson.com",
		},
	}

	jim.setFirstName("Laila")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) setFirstName(name string) {
	(*p).firstName = name
}
