package main

import "fmt"

func playWithAsterisk(n int) {
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			fmt.Print(" ")
		}

		for j := 0; j < i+1; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}

func main() {
	playWithAsterisk(5)
}
