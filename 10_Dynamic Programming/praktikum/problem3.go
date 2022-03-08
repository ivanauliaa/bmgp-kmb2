package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	minJumps := make([]int, n)

	minJumps[0] = 0
	minJumps[1] = int(math.Abs(float64(jumps[0] - jumps[1])))

	for i := 2; i < n; i++ {
		singleJump := int(math.Abs(float64(jumps[i] - jumps[i-1])))
		doubleJump := int(math.Abs(float64(jumps[i] - jumps[i-2])))

		singleJumpCost := minJumps[i-1] + singleJump
		doubleJumpCost := minJumps[i-2] + doubleJump

		minJumps[i] = int(math.Min(float64(singleJumpCost), float64(doubleJumpCost)))
	}

	return minJumps[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}
