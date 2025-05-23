package main

import "fmt"

func main() {
	var p1 struct{ Name, Email string } // variable of type struct
	p1.Name = "Jiten"
	p1.Email = "Jitenp@outlook.com"

	var p2 struct{ Name, Email string } = struct {
		Name  string
		Email string
	}{Name: "Jiten", Email: "JitenP@Outlook.Com"}
	fmt.Println(p1, p2)
}
