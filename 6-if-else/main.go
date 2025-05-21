package main

import "strconv"

func main() {

	a, b := 10, 20
	c := a + b*(b-a) + a/10 + b*5 + 100 // > 100 && true // Expression
	println(c)
	// operator precedence -->(), * ,/ + . - etc.
	// Lowest Precedence = or :=
	// r2 = (b-a) == 10
	// a + b * r2 + a/10 +b *5 +100
	// mul r2 b --> a + 200 + b/10 +b * 5 +100
	// a+ 200 + 1 + b  * 5 +100
	// a+ 200 + 1+ 100 + 100
	// 20 + 200 + 1 + 100 +100
	// 411
	// 1. Is it an atomic operation ?
	// 2. how does this expression is evaluated

	if a+b*(b-a)+a/10+b*5+100 > 100 {
		println("Yes more than 100")
	} else {
		println("less than 100")
	}

	if age := 18; age >= 18 {
		println("Eligible for vote. Becase age is", age)
	} else {
		println("Not Eligible for vote. Becase age is", age)
	}
	//println(age) // it is not allowed

	{
		a1, b1 := 10, 20
		println(a1 + b1)
	}
	// println(a1, b1) // Why I am unable to use a1 and b1

	if age, gender := 41, 'M'; age >= 18 && gender == 'F' || gender == 'f' {
		println("She is eligible for marriage, bcz of age is ", age)
	} else if age >= 21 && gender == 77 || gender == 'm' {
		println("He is eligible for marriage, bcz of age is ", age)
	} else {
		println("not elibible")
	}

	str1 := "1231231"

	if v, err := strconv.Atoi(str1); err != nil {
		println(err.Error())
	} else {
		println(v)
	}

	var any1 interface{} = "12312d"

	if v, ok := any1.(int); !ok {
		if v1, ok1 := any1.(string); ok1 {
			if v2, err := strconv.Atoi(v1); err != nil {
				println(err.Error())
			} else {
				println(v2)
			}
		}
	} else {
		println(v)
	}

	// v or err cannot be used here
}

// arithmetic --> *, / ,+,-
// comparision --> == , >=,<= ,!=, >,<
// logical --> &&, ||
// bitwise --> &, |, <<,>>
