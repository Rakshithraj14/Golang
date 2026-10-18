package main

import "fmt"

const Pi = 3.14159

func main() {
	var name string = "Rakshith"
	age := 24

	var score float64
	var done bool
	var city string
	fmt.Println(name, age, score, done, city)

	fmt.Println("name")
	fmt.Println(name)

	fmt.Printf("%s is %d years old\n", name, age)
	fmt.Printf("Pi = %.2f, type of age = %T\n", Pi, age)
}
