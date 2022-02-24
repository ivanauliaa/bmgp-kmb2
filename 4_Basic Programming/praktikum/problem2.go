package main

import (
	"fmt"
)

func main() {
	var nama string
	var nilai int

	fmt.Println("Masukkan nama dan nilai siswa (pisahkan dengan spasi)")
	fmt.Scanf("%s %d", &nama, &nilai)

	var hasil string
	if nilai >= 80 && nilai <= 100 {
		hasil = "Nilai A"
	} else if nilai >= 65 && nilai <= 79 {
		hasil = "Nilai B"
	} else if nilai >= 50 && nilai <= 64 {
		hasil = "Nilai C"
	} else if nilai >= 35 && nilai <= 49 {
		hasil = "Nilai D"
	} else if nilai >= 0 && nilai <= 34 {
		hasil = "Nilai E"
	} else {
		hasil = "Nilai Invalid"
	}

	fmt.Printf("%s: %s\n", nama, hasil)
}
