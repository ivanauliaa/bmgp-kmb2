package main

import "fmt"

func BinarySearch(array []int, x int) {
	result := -1
	left := 0
	right := len(array) - 1

	for left <= right && result == -1 {
		mid := (left + right) / 2

		if x < array[mid] {
			right = mid - 1
		} else if x > array[mid] {
			left = mid + 1
		} else {
			result = mid
		}
	}

	fmt.Println(result)
}

func main() {
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100)
}
