package main

import "fmt"

func prismaSegiEmpat(high, wide, start int) {
	// TODO bikin deret
	nums := []int{}
	length := high * wide
	count := 0
	sum := 0

	for count < length {
		start++

		if isPrime(start) {
			nums = append(nums, start)
			count++
			sum += start
		}
	}

	// TODO print
	cursor := 0
	for i := 0; i < wide; i++ {
		for j := 0; j < high; j++ {
			fmt.Printf("%d ", nums[cursor])
			cursor++
		}
		fmt.Println()
	}

	fmt.Printf("%d\n", sum)
}

func isPrime(number int) bool {
	if number == 0 || number == 1 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	prismaSegiEmpat(2, 3, 13)
	prismaSegiEmpat(5, 2, 1)
}
