package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("count", i)
	}

	for i := 0; i < 4; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	i := 1
	for i <= 3 {
		fmt.Println("while-style", i)
		i++
	}

	nums := []int{4, 7, 2, 9}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("sum:", sum)

	target := 9
	for idx, n := range nums {
		if n == target {
			fmt.Println("found", target, "at index", idx)
			break
		}
	}
}
