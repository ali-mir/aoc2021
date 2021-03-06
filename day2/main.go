package main

import (
	r "aoc2021/reader"
	"fmt"
	"strconv"
	"strings"
)

type command struct {
	dir  Direction
	dist int
}

type Direction = string

const (
	Up      Direction   = "up"
	Down                = "down"
	Forward             = "forward"
)

func main() {
	var commands []command
	for _, c := range r.Read() {
		parsed := strings.Split(c, " ")
		dir, distStr := parsed[0], parsed[1]
		dist, err := strconv.Atoi(distStr)
		if err != nil {
			panic("Converting string to int failed")
		}
		commands = append(commands, command{dir, dist})
	}

	// Part 1
	pos := 0
	depth := 0
	for _, command := range commands {
		switch command.dir {
		case Down:
			depth += command.dist
		case Up:
			depth -= command.dist
		case Forward:
			pos += command.dist
		}
	}
	fmt.Println("Part 2:", pos*depth)

	// Part 2
	pos = 0
	depth = 0
	aim := 0
	for _, command := range commands {
		switch command.dir {
		case "down":
			aim += command.dist
		case "up":
			aim -= command.dist
		case "forward":
			pos += command.dist
			depth += aim * command.dist
		}
	}
	fmt.Println("Part 2:", pos*depth)
}
