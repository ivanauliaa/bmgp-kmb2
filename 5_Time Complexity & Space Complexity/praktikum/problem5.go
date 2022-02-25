package main

import "fmt"

func PairSum(arr []int, target int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
}
