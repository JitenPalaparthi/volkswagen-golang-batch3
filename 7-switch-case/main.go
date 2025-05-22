package main

func main() {

	var day uint8 = 6

	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("no day")
	}

	switch char := 'a'; char {
	case 'a', 'e', 'i', 'o', 'u':
		println(string(char), "is lower case vovel")
	case 'A', 'E', 'I', 'O', 'U':
		println(string(char), "is upper case vovel")
	default:
		println(string(char), "is consonent or a special unicode char")
	}

	// empty switch
	num := -150
	min, max := 0, 50
	switch {
	case num >= min && num < max:
		{
			println(num, "is bwt 0-50")
		}
	case num >= max && num < 100:
		println(num, "is bwt 50-100")
	case num >= 100:
		println(num, "is greater than or eq 100")
	default:
		println(num, "is a negative numer")
	}

	println("fallthrough in positive cases")
	num = 48 // mutated the existing variable
	// where ever you avoid break in other programming(c family), to get the similar result
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough // fallthrough does not check the case but executes the body
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

	println("fallthrough in false negative cases")
	num = 2
	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough // fallthrough does not check the case but executes the body
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	}

}
