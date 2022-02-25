package main

import "fmt"

func pow(base, power int) int {
	result := 1

	for power > 0 {
		if power%2 == 1 {
			power -= 1
			result *= base
		}

		power /= 2
		base *= base
	}

	return result
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
