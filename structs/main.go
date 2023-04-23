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
		lastName:  "Morrison",
		contactInfo: contactInfo{

			email:   "jim@morrison.hotel",
			zipCode: 66600,
		},
	}

	jim.updateName("jimbo")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {

	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
