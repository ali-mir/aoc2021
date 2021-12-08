package main

import (
	r "aoc2021/reader"
	"fmt"
)

func main() {
	input := r.StringsToDigits(r.Read())

	// Part 1
	gamma, epsilon := commonBits(input)

	fmt.Println("Part 1:", r.BinaryToDecimal(r.DigitArrayToString(gamma))*
		r.BinaryToDecimal(r.DigitArrayToString(epsilon)))

	// Part 2
	oxygenRating := filter(input, gamma, true)
	co2Rating := filter(input, epsilon, false)

	lifeSupportRating := r.BinaryToDecimal(r.DigitArrayToString(oxygenRating)) *
		r.BinaryToDecimal(r.DigitArrayToString(co2Rating))
	fmt.Println("Part 2:", lifeSupportRating)
}

// Filters the slice of binary numbers. This is not as clean as I wish it was.
func filter(input [][]int, common []int, mostCommon bool) []int {
	for i := 0; i < 12; i++ {
		var filtered [][]int
		for _, number := range input {
			if number[i] == common[i] {
				filtered = append(filtered, number)
			}
		}
		input = filtered
		if mostCommon {
			common, _ = commonBits(input)
		} else {
			_, common = commonBits(input)

		}
		if len(input) == 1 {
			return input[0]
		}
	}
	panic("Failed to find correct reading from bit criteria")
}

// Takes in a slice of binary numbers, and outputs the most
// common and least common digits for each position.
func commonBits(input [][]int) ([]int, []int) {
	var mostCommon []int
	var leastCommon []int
	for i := 0; i < 12; i++ {
		var zeroes int
		var ones int
		for _, number := range input {
			if number[i] == 0 {
				zeroes++
			} else {
				ones++
			}
		}
		if zeroes > ones {
			mostCommon = append(mostCommon, 0)
			leastCommon = append(leastCommon, 1)
		} else {
			mostCommon = append(mostCommon, 1)
			leastCommon = append(leastCommon, 0)
		}
	}
	return mostCommon, leastCommon
}
