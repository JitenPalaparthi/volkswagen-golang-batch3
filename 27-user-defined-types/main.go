package main

import "fmt"

func main() {

	p1 := Person{No: 101, Name: "Jiten", Email: "JitenP@outlook.com", Mobile: "9618558500", Address: &Address{City: "Bangalore", PinCode: "560096"}}

	p2 := new(Person) // creating a pointer using new builtin function

	p2.No = 101
	p2.Name = "Jiten"
	p2.Email = "JitenP@Outlook.com"
	p2.Address = new(Address)
	p2.Address.City = "Bangalore"
	p2.Address.PinCode = "560086"
	a1 := new(Address)
	a1.City = "Trivandrum"
	a1.PinCode = "690051"
	//p2.Address = a1
	fmt.Println(p1, *p2)
}

type Person struct { // composition of another type
	No            int
	Name          string
	Email, Mobile string   // can also declare files with comma
	Address       *Address // it okay to give the type name as the fieldname
}

// Go does not suppor direct(oops based) inheritance but it supports composition

type Address struct {
	City, PinCode string
}

func NewPerson(no int, name, email, mobile string) *Person {
	return &Person{No: no, Name: name, Email: email, Mobile: mobile}
}
