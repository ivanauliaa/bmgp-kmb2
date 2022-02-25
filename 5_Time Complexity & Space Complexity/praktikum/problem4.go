package main

import (
	"fmt"
	"strconv"
	"strings"
)

func angkaMunculSekali(angka string) []int {
	angkaArr := strings.Split(angka, "")
	freq := make(map[string]int)

	for _, val := range angkaArr {
		freq[val]++
	}

	result := make([]int, 0)

	for key, val := range freq {
		if val == 1 {
			toAppend, _ := strconv.Atoi(key)
			result = append(result, toAppend)
		}
	}

	return result
}

func main() {
	fmt.Println(angkaMunculSekali("1234123"))
}
