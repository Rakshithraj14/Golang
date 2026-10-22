package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("square(6):", square(6))

	q, r := divmod(17, 5)
	fmt.Println("17 / 5 =", q, "remainder", r)

	if _, err := safeDivide(10, 0); err != nil {
		fmt.Println("error:", err)
	}

	sorted := []int{2, 5, 8, 12, 16}
	fmt.Println("index of 12:", binarySearch(sorted, 12))
	fmt.Println("index of 7:", binarySearch(sorted, 7))
}

func square(n int) int {
	return n * n
}

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func binarySearch(sorted []int, target int) int {
	lo, hi := 0, len(sorted)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		switch {
		case sorted[mid] == target:
			return mid
		case sorted[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}
