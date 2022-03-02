package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	count := map[int]int{}

	// TODO iterate cards
	// kalo ada angka yang sama dengan yang ada di deck, cari sum dan simpan di map dengan key index
	for index, card := range cards {
		found := false

		for _, cardValue := range card {

			for _, deckValue := range deck {
				if cardValue == deckValue {
					found = true
					break
				}
			}

			if found {
				break
			}
		}

		if found {
			sum := 0
			for _, cardValue := range card {
				sum += cardValue
			}

			count[index] = sum
		}
	}

	// TODO kalo len map 0 return "tutup kartu"
	// kalo ga, perulangan cari nilai map terbesar dan return it

	if len(count) == 0 {
		return "tutup kartu"
	}

	maxIndex := -1
	for key, value := range count {
		if maxIndex == -1 {
			maxIndex = key
			continue
		}

		if value > count[maxIndex] {
			maxIndex = key
		}
	}

	return cards[maxIndex]
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}
