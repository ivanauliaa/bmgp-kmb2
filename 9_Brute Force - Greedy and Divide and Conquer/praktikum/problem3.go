package main

import "fmt"

func DragonOfLoowater(dragonHead, knightHeight []int) {
	if len(knightHeight) < len(dragonHead) {
		fmt.Println("knight fall")
		return
	}

	sort(&dragonHead)
	sort(&knightHeight)

	count := 0
	minValue := 0

	for i := 0; i < len(dragonHead); i++ {
		for j := 0; j < len(knightHeight); j++ {
			if knightHeight[j] >= dragonHead[i] {
				count++
				minValue += knightHeight[j]

				break
			}
		}
	}

	if count != len(dragonHead) {
		fmt.Println("knight fall")
	} else {
		fmt.Println(minValue)
	}
}

func sort(array *[]int) {
	for i := 0; i < len(*array)-1; i++ {
		for j := i + 1; j < len(*array); j++ {
			if (*array)[j] < (*array)[i] {
				temp := (*array)[i]
				(*array)[i] = (*array)[j]
				(*array)[j] = temp
			}
		}
	}
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})
	DragonOfLoowater([]int{5, 10}, []int{5})
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})
}
