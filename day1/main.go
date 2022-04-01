package main

import (
	r "aoc2021/reader"
	"fmt"
)

func main() {
	input := r.StringsToInts(r.Read())

	// Part 1
	count := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			count++
		}
	}
	fmt.Println("Part 1:", count)

	// Part 2
	count = 0
	for i := 3; i < len(input); i++ {
		if input[i] > input[i-3] {
			count++
		}
	}
	fmt.Println("Part 2:", count)
}
