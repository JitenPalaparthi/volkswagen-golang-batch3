package greet

import "fmt"

// There is not public or private
// restricted/unexported --> private --> The symbols starts with lowerCase
// unrestricted/exported --> public
func Greet() {
	greet()
}

func greet() { // restricted
	fmt.Println("Hello VolksWagen")
}

type Public struct { // unrestricted
	PublicField  string // unrestricted
	privateField string
}

func (p *Public) Greet() { // unrestricted
	greet()
}

func (p *Public) sayHi() { // starts with lower case
	greet()
}
