package main

import (
	r "aoc2021/reader"
	"fmt"
	"strings"
)

type Pair struct {
	start Point
	end   Point
}

type Point struct {
	x int
	y int
}

func main() {
	input := r.Read()

	// part 1
	pairs := convertLinesToPairs(input)
	o := findOverlaps(pairs, false /* diagonals */)
	fmt.Println("Part 1:", o)

	// part 2
	o = findOverlaps(pairs, true /* diagonals */)
	fmt.Println("Part 2:", o)

}

func convertLinesToPairs(input []string) []Pair {
	var pairs []Pair
	for _, line := range input {
		p1, p2 := lineToPoints(line)
		pairs = append(pairs, Pair{p1, p2})
	}
	return pairs
}

func lineToPoints(line string) (Point, Point) {
	var points []Point
	split := strings.Split(line, " -> ")
	for _, p := range split {
		positions := strings.Split(p, ",")
		points = append(points, Point{r.StringToInt(positions[0]), r.StringToInt(positions[1])})
	}
	return points[0], points[1]

}

func pointsAreEqual(p1, p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func filterForVerticalHorizontalLines(pairs []Pair) []Pair {
	var filtered []Pair
	for _, p := range pairs {
		start := p.start
		end := p.end
		if start.x == end.x || start.y == end.y {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func filterForDiagonalLines(pairs []Pair) []Pair {
	var filtered []Pair
	for _, p := range pairs {
		start := p.start
		end := p.end
		if start.x != end.x && start.y != end.y {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func findOverlaps(pairs []Pair, diagonal bool) int {
	var overlaps int
	m := make(map[Point]int)

	filtered := filterForVerticalHorizontalLines(pairs)
	for _, p := range filtered {
		markMapHorizontalVerticalLines(p, m)
	}

	if diagonal {
		filtered = filterForDiagonalLines(pairs)
		for _, p := range filtered {
			markMapDiagonalLines(p, m)
		}
	}

	for _, e := range m {
		if e > 1 {
			overlaps++
		}
	}
	return overlaps
}

// this function is not clean at all
func findMinMax(p Pair) (int, int, int, int) {
	var minX int
	var maxX int
	var minY int
	var maxY int

	if p.start.x <= p.end.x {
		minX = p.start.x
		maxX = p.end.x
	} else if p.start.x > p.end.x {
		minX = p.end.x
		maxX = p.start.x
	}

	if p.start.y <= p.end.y {
		minY = p.start.y
		maxY = p.end.y
	} else if p.start.y > p.end.y {
		minY = p.end.y
		maxY = p.start.y
	}

	return minX, maxX, minY, maxY
}

func markMapHorizontalVerticalLines(p Pair, m map[Point]int) {
	minX, maxX, minY, maxY := findMinMax(p)
	// mark x lines
	if minY == maxY {
		for i := minX; i <= maxX; i++ {
			p := Point{i, minY}
			m[p]++
		}
	}
	// mark y lines
	if minX == maxX {
		for i := minY; i <= maxY; i++ {
			p := Point{minX, i}
			m[p]++
		}
	}
}

// also not clean
func markMapDiagonalLines(p Pair, m map[Point]int) {
	startX, endX := p.start.x, p.end.x
	startY, endY := p.start.y, p.end.y

	if r.Abs(endX-startX) != r.Abs(endY-startY) {
		panic("Diagonal line is not at 45 degree angle")
	}

	for {
		p := Point{startX, startY}
		m[p]++

		if startX == endX && startY == endY {
			break
		}

		if startX < endX {
			startX++
		} else {
			startX--
		}

		if startY < endY {
			startY++
		} else {
			startY--
		}

	}
}
