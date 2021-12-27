package main

import "testing"

func TestExample(t *testing.T) {
	fish := []int{3, 4, 3, 1, 2}
	population := modelFish(fish, 18 /* days */)

	if population != 26 {
		t.Errorf("Expected fish population to be 26 but received %d", population)
	}

	population = modelFish(fish, 80 /* days */)
	if population != 5934 {
		t.Errorf("Expected fish population to be 5934 but received %d", population)
	}

	population = modelFish(fish, 256 /* days */)
	if population != 26984457539 {
		t.Errorf("Expected fish population to be 26984457539 but received %d", population)
	}
}
