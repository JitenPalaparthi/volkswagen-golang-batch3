package main

import "fmt"

func main() {
	var p1 Person // declaration
	p1.Name = "Jiten"
	p1.Email = "JitenP@outlook.com"
	p1.Mobile = "9618558500"

	fmt.Println(p1) /// dont use it in production

	// if p1 == nil {

	// }

	p2 := Person{Name: "Jiten", Email: "JitenP@Outlook.Com", Mobile: "9618558500"}

	p3 := Person{101, "Jiten", "JitenP@Outlook.Com", "9618558500"}
	println(&p2, &p3)

	p4 := New(101, "Jiten", "JitenP@Outlook.Com", "9618558500")
	println(p4)
}

type Person struct {
	No            int
	Name          string
	Email, Mobile string // can also declare files with comma
}

func New(no int, name, email, mobile string) *Person {
	return &Person{no, name, email, mobile}
}

func Greet() {
	println("Hello VolksWagen")
}
