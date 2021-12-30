package main

import (
	r "aoc2021/reader"
	"fmt"
)

const size int = 100

type Point struct {
	row int
	col int
	val int
}

func main() {
	input := r.Read()
	heightmap := parse(input, size, size)
	lows := findLows(heightmap)
	fmt.Println("Part 1:", calculateSumOfRisk(lows))
}

func parse(input []string, rows, cols int) [][]int {
	// initialize slice of slices
	heightmap := make([][]int, rows)
	for i := 0; i < rows; i++ {
		heightmap[i] = make([]int, cols)
	}
	// parse each char into an int
	for i, str := range input {
		for j, char := range str {
			parsed := r.StringToInt(string(char))
			heightmap[i][j] = parsed
		}
	}
	return heightmap
}

func findLows(h [][]int) []int {
	var result []int
	for i := 0; i < len(h); i++ {
		for j := 0; j < len(h[i]); j++ {
			val := h[i][j]
			if isLowestAdjacent(h, val, i, j) {
				result = append(result, val)
			}
		}
	}
	return result
}

func isLowestAdjacent(h [][]int, val, row, col int) bool {
	isLowest := true
	rowSize := len(h)
	colSize := len(h[0])

	// up
	if isWithinBounds(row+1, col, rowSize, colSize) && h[row+1][col] <= val {
		isLowest = false
	}
	// down
	if isWithinBounds(row-1, col, rowSize, colSize) && h[row-1][col] <= val {
		isLowest = false
	}
	// right
	if isWithinBounds(row, col+1, rowSize, colSize) && h[row][col+1] <= val {
		isLowest = false
	}
	// left
	if isWithinBounds(row, col-1, rowSize, colSize) && h[row][col-1] <= val {
		isLowest = false
	}
	return isLowest
}

func isWithinBounds(row, col, rowSize, colSize int) bool {
	rowValid := row >= 0 && row < rowSize
	colValid := col >= 0 && col < colSize
	return rowValid && colValid
}

func calculateSumOfRisk(lows []int) int {
	sum := 0
	for _, l := range lows {
		sum += (l + 1)
	}
	return sum
}

