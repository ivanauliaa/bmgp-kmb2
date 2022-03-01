package main

import "fmt"

func caesar(offset int, input string) string {
	offset %= 26
	result := ""
	strBytes := []byte(input)

	for _, value := range strBytes {
		converted := (value + byte(offset)) % 123

		if converted < 97 {
			converted += 97
		}

		result += string(converted)
	}

	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
