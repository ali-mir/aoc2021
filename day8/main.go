package main

import (
	r "aoc2021/reader"
	"fmt"
	"strings"
)

type Signal struct {
	patterns []string // size 10
	output   []string // size 4
}

func main() {
	input := r.Read()
	signals := parse(input)

	fmt.Println("Part 1:", findUniqueSegmentOcurrences(signals))
}

func parse(input []string) []Signal {
	var signals []Signal
	for _, i := range input {
		split := strings.Split(i, "|")
		patterns := strings.Split(split[0], " ")
		output := strings.Split(split[1], " ")
		signals = append(signals, Signal{patterns, output})
	}
	return signals
}

func findUniqueSegmentOcurrences(signals []Signal) int {
	sum := 0
	for _, s := range signals {
		for _, o := range s.output {
			if hasUniqueNumberOfSegments(o) {
				sum++
			}
		}
	}
	return sum
}

func hasUniqueNumberOfSegments(s string) bool {
	segments := len(s)
	return segments == 2 || segments == 4 || segments == 3 || segments == 7
}
