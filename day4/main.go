package main

import (
	r "aoc2021/reader"
	"fmt"
	"strings"
)

// Each bingo board is 5x5.
var length int = 5

type Board struct {
	entries [][]int
}

func main() {
	// parsing
	input := r.Read()
	numbers := r.StringsToInts(strings.Split(input[0], ","))
	// Remove bingo numbers line
	input = input[1:]

	// Part 1
	boards := createBoards(input)
	var winner Board
	var winningNum int
	outer:
		for i := range numbers {
			for _, b := range boards {
				won, num := hasWon(b, numbers[0:i])
				if won {
					winner = b
					winningNum = num
					break outer
				}
			}
		}
	fmt.Println("Part 1:", score(winner)*winningNum)

	// Part 2
	boards = createBoards(input)
	var lastWinningNum int
	var lastWinner Board
	for len(boards) > 0 {
	outer2:
		for i := range numbers {
			for j, b := range boards {
				won, num := hasWon(b, numbers[0:i])
				if won {
					// If a board has won, remove it from the list of boards.
					boards = append(boards[:j], boards[j+1:]...)
					lastWinningNum = num
					lastWinner = b
					break outer2
				}
			}
		}
	}
	fmt.Println("Part 2:", score(lastWinner)*lastWinningNum)

}

// Parses input into all the boards
func createBoards(input []string) []Board {
	var boards []Board
	for i := 1; i < len(input); i += 6 {
		board := createSingleBoard(input[i : i+length])
		boards = append(boards, board)
	}
	return boards
}

// Parses input into a single board
func createSingleBoard(input []string) Board {
	var entries [][]int
	for _, row := range input {
		entries = append(entries, r.StringsToInts(strings.Fields(row)))
	}
	return Board{entries}
}

// Given a board and a slice of numbers, does the board win?
func hasWon(b Board, numbers []int) (bool, int) {
	// This isn't clean, but it works.
	var won bool
	for _, num := range numbers {
		for row := 0; row < length; row++ {
			for col := 0; col < length; col++ {
				// If the entry matches the number,
				// we mark it by setting it to -1.
				if b.entries[row][col] == num {
					b.entries[row][col] = -1
				}
			}
		}
		won = searchBoardForWin(b)
		if won {
			return true, num
		}
	}
	return false, -1
}

// Searches the rows and columns for a winning slice.
func searchBoardForWin(b Board) bool {
	win := []int{-1, -1, -1, -1, -1}
	// Check the rows
	for row := 0; row < length; row++ {
		if r.EqualSlices(b.entries[row], win) {
			return true
		}
	}
	// Check the columns
	for col := 0; col < length; col++ {
		var currCol []int
		// Collect the numbers from each column in each row
		for row := 0; row < length; row++ {
			currCol = append(currCol, b.entries[row][col])
		}
		if r.EqualSlices(currCol, win) {
			return true
		}
	}
	return false
}

// Calculates the score given a marked up board.
func score(b Board) int {
	var product int
	for row := 0; row < length; row++ {
		for col := 0; col < length; col++ {
			if b.entries[row][col] != -1 {
				product += b.entries[row][col]
			}
		}
	}
	return product
}
