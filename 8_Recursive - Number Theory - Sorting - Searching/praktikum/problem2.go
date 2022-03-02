package main

import "fmt"

func fibonacci(number int) int {
	// TODO base condition
	if number <= 1 {
		return number
	}

	return fibonacci(number-1) + fibonacci(number-2)
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci(12))
}
