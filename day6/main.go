package main

import (
	r "aoc2021/reader"
	"fmt"
	"strings"
)

func main() {
	input := r.Read()
	nums := r.StringsToInts(strings.Split(input[0], ","))

	// part 1
	population := modelFish(nums, 80 /* days */)
	fmt.Println("Part 1:", population)

	// part 2
	population = modelFish(nums, 256 /* days */)
	fmt.Println("Part 2:", population)

}

// We can efficiently solve this problem by maintaining a slice with each index i
// mapping to the number of fish that have exactly i days left until reproduction.
// At the end of each day, we can shift each element down one index. Each new fish
// is added to the end of the slice (with 8 days before reproduction), and each fish
// that just reproduced is added to the 6th index (signalling 6 days before reproduction)
func modelFish(nums []int, days int) int {
	fish := make([]int, 9)

	for _, num := range nums {
		fish[num]++
	}

	for i := 0; i < days; i++ {
		// 'ready' is the number of fish ready to spawn
		ready := fish[0]
		fish = fish[1:]

		fish = append(fish, ready)
		fish[6] += ready
	}
	sum := 0
	for _, e := range fish {
		sum += e
	}
	return sum
}
