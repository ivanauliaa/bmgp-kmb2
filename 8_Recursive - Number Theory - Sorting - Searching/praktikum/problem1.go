package main

import "fmt"

func primeX(number int) int {
	count := 0
	i := 1

	// TODO perulangan sampe ketemu prima ke x
	for count < number {
		i++

		if isPrime(i) {
			count++
		}
	}

	return i
}

// TODO fungsi cek prima
func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
