package main

import "fmt"

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	freq := map[string]int{}
	result := []pair{}

	// TODO hitung frekuensi pake map
	for _, value := range items {
		freq[value] += 1
	}

	// TODO konversi dari map ke []pair
	for key, value := range freq {
		result = append(result, pair{name: key, count: value})
	}

	// TODO sort berdasarkan count
	for i := 0; i < len(result)-1; i++ {
		if result[i+1].count < result[i].count {
			temp := result[i]
			result[i] = result[i+1]
			result[i+1] = temp
		}
	}

	return result
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
