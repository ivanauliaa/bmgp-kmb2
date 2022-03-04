package main

import "fmt"

func moneyCoins(money int) []int {
	result := []int{}

	for money > 0 {
		if money > 10000 {
			substractAndAppend(10000, &money, &result)
		} else if money > 5000 {
			substractAndAppend(5000, &money, &result)
		} else if money > 2000 {
			substractAndAppend(2000, &money, &result)
		} else if money > 1000 {
			substractAndAppend(1000, &money, &result)
		} else if money > 500 {
			substractAndAppend(500, &money, &result)
		} else if money > 200 {
			substractAndAppend(200, &money, &result)
		} else if money > 100 {
			substractAndAppend(100, &money, &result)
		} else if money > 50 {
			substractAndAppend(50, &money, &result)
		} else if money > 20 {
			substractAndAppend(20, &money, &result)
		} else if money > 10 {
			substractAndAppend(10, &money, &result)
		} else {
			substractAndAppend(1, &money, &result)
		}
	}

	return result
}

func substractAndAppend(value int, money *int, result *[]int) {
	*money -= value
	*result = append(*result, value)
}

func main() {
	fmt.Println(moneyCoins(123))
	fmt.Println(moneyCoins(432))
	fmt.Println(moneyCoins(543))
	fmt.Println(moneyCoins(7752))
	fmt.Println(moneyCoins(15321))
}
