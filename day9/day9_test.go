package main

import "testing"

func TestExample(t *testing.T) {
	input := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}

	parsed := parse(input, 5, 10)

	lows := findLows(parsed)
	s := calculateSumOfRisk(lows)
	if s != 15 {
		t.Errorf("Expected sum of risk levels to be 15 but received %d", s)
	}
}

func TestExampleSmall(t *testing.T) {
	input := []string{
		"333",
		"313",
		"333",
	}

	parsed := parse(input, 3, 3)

	lows := findLows(parsed)
	s := calculateSumOfRisk(lows)
	if s != 2 {
		t.Errorf("Expected sum of risk levels to be 2 but received %d", s)
	}
}
