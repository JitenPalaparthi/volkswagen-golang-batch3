package main

import "fmt"

func main() {

	p1 := Person{Name: "Jiten", Email: "JitenP@Outlook.Com", Address: struct {
		City        string
		PinCode     string
		LineDetails struct {
			Line1 string
			Line2 string
			Line3 string
		}
	}{City: "Bangalore", PinCode: "560086", LineDetails: struct {
		Line1 string
		Line2 string
		Line3 string
	}{Line1: "Srikanteswara Nagar", Line2: "Some Line2", Line3: "Some line3"}}}
	fmt.Println("Line1:", p1.Address.LineDetails.Line1)
	fmt.Println("Line2:", p1.Address.LineDetails.Line2)

}

type Person struct {
	Name, Email string
	Address     struct { // unnamed structure
		City, PinCode string
		LineDetails   struct {
			Line1 string
			Line2 string
			Line3 string
		}
	}
}

// type Employee struct {
// 	No          int
// 	Name, Email string
// 	Address     struct {
// 		City, PinCode string
// 	}
// }
