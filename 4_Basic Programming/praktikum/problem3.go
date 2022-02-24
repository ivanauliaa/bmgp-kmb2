package main

import "fmt"

func main() {
	var bilangan int

	fmt.Println("Masukkan suatu bilangan")
	fmt.Scanf("%d", &bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}
