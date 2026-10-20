package main

import "fmt"

func main() {
	a, b := 17, 5

	fmt.Println(a+b, a-b, a*b, a/b, a%b)

	fmt.Println(a > b, a == b, a > 10 && b > 10)

	age := 24
	if age >= 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("minor")
	}

	if n := a % 2; n == 0 {
		fmt.Println(a, "is even")
	} else {
		fmt.Println(a, "is odd")
	}

	switch day := 3; day {
	case 1:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("some other day")
	}
}
