package main

func main() {
	cd1 := ColourCode{int: 101, string: "Red"}
	cd2 := ColourCode{101, "Red", "Thick Red"}
	println("Code:", cd1.int, "Colour:", cd1.string, "Colour2:", cd1.mystring)
	println("Code:", cd2.int, "Colour:", cd2.string, "Colour2:", cd2.mystring)
	cd3 := New(101, "Red", "Thick Red") // can manage to store in heap
	println("Code:", cd3.int, "Colour:", cd3.string, "Colour2:", cd3.mystring)
	//println(&cd1, &cd2, &cd3)
}

type mystring = string

type ColourCode struct { // anonymous fields
	int      // no field name,field type as the filed name
	string   // no filed name,ield type as the filed name
	mystring // no filed name,ield type as the filed name
}

type Person struct {
	No            int
	Name          string
	Email, Mobile string // can also declare files with comma
}

func NewPerson(no int, name, email, mobile string) *Person {
	return &Person{no, name, email, mobile}
}

func NewColourCode(code int, colour, colour2 string) *ColourCode {
	return &ColourCode{code, colour, colour2}
}
