package main

import (
	"fmt"
	"math"
)

func SimpleEquations(a, b, c int) {
	// TODO cari nilai tertinggi antara a, b, c buat jadi max iter
	max := 0

	if a > b {
		if a > c {
			max = a
		} else {
			max = c
		}
	} else {
		if b > c {
			max = b
		} else {
			max = c
		}
	}

	// TODO perulangan 3 level
	// Trus cek hasil 3 persamaan, kalo memenuhi semua berarti print + return
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			for k := 0; k < max; k++ {
				if i+j+k == a && i*j*k == b && math.Pow(float64(i), 2)+math.Pow(float64(j), 2)+math.Pow(float64(k), 2) == float64(c) {
					fmt.Printf("%d %d %d", i, j, k)
					return
				}
			}
		}
	}

	fmt.Println("no solution")
}

func main() {
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
}
