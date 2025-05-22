package main

import "fmt"

func main() {

	r := SumOf()
	fmt.Println("Sum:", r)

	r = SumOf(12, 30)
	fmt.Println("Sum:", r)

	r = SumOf(12, 12.3, 4, int8(30), int32(12321), 12312.123, true, "hello World")
	fmt.Println("Sum:", r)

	arr := [8]any{12, 12.3, 4, int8(30), int32(12321), 12312.123, true, "hello World"}
	r = SumOf(arr[:]...)
	fmt.Println("Sum:", r)

	slice := []any{12, 12.3, 4, int8(30), int32(12321), 12312.123, true, "hello World"}
	r = SumOf(slice...)
	fmt.Println("Sum:", r)

	slice = append(slice, 10, 11, 12, 13, 14)                                        // append func is an example of variadic                                        // variadic
	fmt.Println(12, 12.3, 4, int8(30), int32(12321), 12312.123, true, "hello World") // variadic example fmt.Println
}

// variadic parameter must be the last parameter to a func or a method
// variadic argumen(s) can also be no arguments
// variadic type can only be used in funcs/methods, cannot create a field or variable
// a slice can be passed as variadic argument using ... symbol
// can use range loop to iterate all variadic arguments

func SumOf(nums ...any) float64 {
	sum := 0.0
	for _, val := range nums {
		if IsNumber(val) {
			switch v := val.(type) {
			case int:
				sum += float64(val.(int))
			case uint:
				sum += float64(v)
			case int8:
				sum += float64(v)
			case int16:
				sum += float64(v)
			case int32:
				sum += float64(v)
			case int64:
				sum += float64(v)
			case uint8:
				sum += float64(v)
			case uint16:
				sum += float64(v)
			case uint32:
				sum += float64(v)
			case uint64:
				sum += float64(v)
			case float32:
				sum += float64(v)
			case float64:
				sum += v
			}
		}
	}
	return sum
}

// func MulAndSumOf(nums ...int, mul int) int {// it is error bcz variadic has given as first param.

// }

func MulAndSumOf(mul int, nums ...int) int { // variadic must be the last parameter
	sum := 0
	for _, v := range nums {
		sum += v * mul
	}
	return sum
}

func IsNumber(a any) bool {
	switch a.(type) { // type switch
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64:
		return true
	default:
		return false
	}
}
