package main

import "testing"

func setup() []Board {
	var boards []Board
	var b Board
	entries := [][]int{
		{72, 99, 88, 8, 59},
		{61, 96, 92, 2, 70},
		{1, 32, 18, 10, 95},
		{33, 20, 31, 66, 43},
		{26, 24, 91, 44, 11},
	}
	b = Board{entries}
	boards = append(boards, b)
	return boards
}

func winHelper(t *testing.T, numbers []int) {
	boards := setup()
	won, winningNum := hasWon(boards[0], numbers)
	if won != true {
		t.Errorf("Board %q expected to win but did not", boards[0])
	}
	if winningNum != numbers[4] {
		t.Errorf("Winning num expected to be %v but was %v", numbers[4], winningNum)
	}
}

func TestHasWonRow(t *testing.T) {
	numbers := []int{72, 99, 88, 8, 59}
	winHelper(t, numbers)
}

func TestHasWonCol(t *testing.T) {
	numbers := []int{72, 61, 1, 33, 26}
	winHelper(t, numbers)
}
