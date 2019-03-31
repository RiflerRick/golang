package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct { // struct like c
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"} // Alex is going to be the firstName and Anderson is going to
	// be lastName. NOT RECOMMENDED
	romanof := person{firstName: "Alex", lastName: "Romanof"} // recommended way
	// fmt.Println(alex)
	fmt.Println(romanof)

	var bob person // this is another way to instantiate a variable of struct type in go
	// when we do not actually instantiate a struct with a value in go, go populates that with what is
	// called a zero value

	//fmt.Printf("%+v", alex) // %+v is printout all values within a struct

	//updating structs
	bob.firstName = "Bob"
	bob.lastName = "Hawke"

	fmt.Printf("%+v", bob)

	// structs are basically like python dictionaries
	// embedded struct
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000, // this comma is required
		}, //even if it is the last value in the struct, the comma is required
	}

	jim.updateNamePassByValue("jimmy") // when we call this function, from the struct actually its a pass by value
	jimPointer := &jim
	jim.print()
	jimPointer.updateNamePassByRef("jimmy") // as a matter of fact it is also possible to call this
	// in the following way
	jim.updateNamePassByRef("jimmy")
	// this is a go shortcut, although jim is not a pointer, go figures out from the function prototype
	// that the function actually expects a pointer and not a value and what you get inside the function
	// is still a pointer.
	jim.print()

}

//receiver functions for structs
func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateNamePassByValue(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) updateNamePassByRef(newFirstName string) {
	(*p).firstName = newFirstName
}

// ints, strings, floats, bools, structs are all passed by value however
