package main

import (
	"fmt"
	"runtime"
	"sync"
	"unicode"
)

func LetterFrequency(sentence string) {
	letterFreq := map[string]int{}
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}

	for _, letter := range sentence {
		wg.Add(1)

		go func(char rune) {
			mtx.Lock()

			if unicode.IsLetter(char) {
				letterFreq[string(char)] += 1
			}

			mtx.Unlock()
			wg.Done()
		}(letter)
	}

	wg.Wait()
	fmt.Println(letterFreq)
}

func main() {
	runtime.GOMAXPROCS(2)
	LetterFrequency("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
}
