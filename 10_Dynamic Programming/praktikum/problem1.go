package main

import "fmt"

func fibo(n int, memo *map[int]int) int {
	// TODO define base condition
	if n == 0 || n == 1 {
		return n
	}

	// TODO return nilai memoization kalo != 0
	if (*memo)[n] != 0 {
		return (*memo)[n]
	}

	// TODO assign nilai ke memoization[n]
	(*memo)[n] = fibo(n-1, memo) + fibo(n-2, memo)

	// TODO return memoization[n]
	return (*memo)[n]
}

func main() {
	fmt.Println(fibo(0, &map[int]int{}))
	fmt.Println(fibo(1, &map[int]int{}))
	fmt.Println(fibo(2, &map[int]int{}))
	fmt.Println(fibo(3, &map[int]int{}))
	fmt.Println(fibo(5, &map[int]int{}))
	fmt.Println(fibo(6, &map[int]int{}))
	fmt.Println(fibo(7, &map[int]int{}))
	fmt.Println(fibo(9, &map[int]int{}))
	fmt.Println(fibo(10, &map[int]int{}))
}
