package main

import (
	"fmt"
	"math"
)

func main() {
	var tinggi, jariJari int

	fmt.Println("Masukkan nilai tinggi dan jari-jari tabung (pisahkan dengan spasi)")
	fmt.Scanf("%d %d", &tinggi, &jariJari)

	hasil := 2 * math.Pi * float64(jariJari) * float64(jariJari+tinggi)
	fmt.Printf("Luas permukaan tabung: %f\n", hasil)
}
