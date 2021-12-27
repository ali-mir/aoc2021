package main

import "testing"

func TestFindMedian(t *testing.T) {
	nums := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	median := findMedian(nums)
	if median != 2 {
		t.Errorf("Expected median to be 2 but received %d", median)
	}
}

func TestFindFuelConstant(t *testing.T) {
	nums := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	fuel := findFuelConstant(nums)
	if fuel != 37 {
		t.Errorf("Expected cheapest fuel to be 37, received %v.", fuel)
	}
}

func TestFindFuelExponential(t *testing.T) {
	nums := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	fuel := findFuelExponential(nums)
	if fuel != 168 {
		t.Errorf("Expected cheapest fuel to be 168, received %v.", fuel)
	}
}
