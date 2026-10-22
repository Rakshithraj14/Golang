package main

import "fmt"

func main() {
	arr := [3]int{10, 20, 30}
	fmt.Println(arr, len(arr))

	number := []int{10, 20, 30}
	number = append(number, 40)

	fmt.Println(number)
	fmt.Println("The number added is:", number[3])
	fmt.Println("length:", len(number))

	scores := []int{5, 10, 15}
	scores = append(scores, 20)
	fmt.Println(scores[2])

	fmt.Println(number[1:3])

	for i, v := range number {
		fmt.Println(i, v)
	}
}
