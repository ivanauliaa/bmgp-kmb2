package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(angka string) []int {
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
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
