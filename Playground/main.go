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
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "email@test.com",
			zipCode: 94555,
		},
	}
	jim.print()
	jim.updateName("Biju")
	//jim.updateName("Biju")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(n string) {
	(*p).firstName = n
	//p.print()
}
