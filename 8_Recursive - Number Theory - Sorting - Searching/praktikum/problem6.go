package main

import "fmt"

func MaximumBuyProduct(money int, productPrice []int) {
	count := 0

	// TODO sort
	for i := 0; i < len(productPrice)-1; i++ {
		if productPrice[i+1] < productPrice[i] {
			temp := productPrice[i]
			productPrice[i] = productPrice[i+1]
			productPrice[i+1] = temp
		}
	}

	// TODO perulangan sampe money ga cukup
	for _, value := range productPrice {
		if money-value > 0 {
			count++
			money -= value
		} else {
			break
		}
	}

	// TODO print total
	fmt.Println(count)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})
	MaximumBuyProduct(50000, []int{15000, 10000, 12000, 5000, 3000})
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})
	MaximumBuyProduct(0, []int{10000, 30000})
}
