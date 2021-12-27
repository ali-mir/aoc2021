package main

import (
	r "aoc2021/reader"
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := r.Read()
	nums := r.StringsToInts(strings.Split(input[0], ","))

	fmt.Println("Part 1:", findFuelConstant(nums))
	fmt.Println("Part 2:", findFuelExponential(nums))

}

func findFuelConstant(nums []int) int {
	median := findMedian(nums)
	fuel := 0
	for _, n := range nums {
		fuel += r.Abs(n - median)
	}
	return fuel
}

func findFuelExponential(nums []int) int {
	avg := findAverage(nums)
	fuel := 0
	for _, n := range nums {
		delta := r.Abs(n - avg)
		fuel += (delta * (delta + 1)) / 2
	}
	return fuel
}

func findMedian(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func findAverage(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return r.Round(float64(sum) / float64(len(nums)))
}
