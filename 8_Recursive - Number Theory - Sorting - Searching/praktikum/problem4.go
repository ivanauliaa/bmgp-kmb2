package main

import "fmt"

func MaxSequence(arr []int) int {
	max := 0
	// TODO iterate satu item terhadap seluruh sisa item
	// tiap iterasi, cari sum dan selalu bandingin dengan nilai max
	for i := 0; i < len(arr)-1; i++ {
		sum := arr[i]

		if sum > max {
			max = sum
		}
		for j := i + 1; j < len(arr); j++ {
			sum += arr[j]

			if sum > max {
				max = sum
			}
		}
	}

	return max
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
