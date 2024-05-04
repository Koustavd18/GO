package main

import "fmt"

type person struct{
	firstName string
	lastName string
	contactInfo
}

type contactInfo struct{
	email string
	zipCode int
}

func main() {

	jim := person{
		firstName: "Dwight",
		lastName: "Schrute",
		contactInfo: contactInfo{
			email: "email@gmail.com",
			zipCode: 560068,
		},
	}
	
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}