package main

import "fmt"

func main() {

	p1 := Person{No: 101, Name: "Jiten", Email: "JitenP@outlook.com", Mobile: "9618558500", Address: Address{City: "Bangalore", PinCode: "560096"}}
	p1.Linkedin = "linkedin.com/jpalapathi"
	p1.Instagram = "jiten_1231"
	p1.Facebook = "facebook.com/jitenp"
	p1.Status = "active"
	p1.Address.Status = "active"
	p1.Gmail = "jiten.palaparthi@gmail.com"
	p1.Twitter = "jiten_1981"
	println("City:", p1.City) // directly can access City thru p1 bcz Address is a promoted field
	fmt.Println(p1)
}

type Person struct {
	No            int
	Name          string
	Email, Mobile string // can also declare files with comma
	Status        string
	Address       // promoted field, somewhere features of inheritance can there but it is purely composition
	SocialMedia   // promoted field
}

// Go does not suppor direct(oops based) inheritance but it supports composition

type Address struct {
	City, PinCode, Status string
}

type SocialMedia struct {
	Linkedin                  string
	Twitter                   string
	Instagram                 string
	Gmail, Facebook, Whatsapp string
}

func NewPerson(no int, name, email, mobile string) *Person {
	return &Person{No: no, Name: name, Email: email, Mobile: mobile}
}
