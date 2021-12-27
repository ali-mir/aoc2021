package main

import "testing"

func TestLineToPoints(t *testing.T) {
	str := "217,490 -> 217,764"
	p1, p2 := lineToPoints(str)
	if p1.x != 217 && p1.y != 490 {
		t.Errorf("Expected p1: {x: 217, y: 490}. Received p1: {x: %v, y: %v}", p1.x, p1.y)
	}
	if p2.x != 217 && p2.y != 764 {
		t.Errorf("Expected p1: {x: 217, y: 764}. Received p1: {x: %v, y: %v}", p1.x, p1.y)
	}
}

func TestConvertLinesToPairs(t *testing.T) {
	str := []string{"217,490 -> 217,764", "44,270 -> 373,599"}
	expected := []Pair{Pair{Point{217, 490}, Point{217, 764}}, Pair{Point{44, 270}, Point{373, 599}}}

	pairs := convertLinesToPairs(str)
	for i, p := range pairs {
		exp := expected[i]
		if !pointsAreEqual(p.start, exp.start) || !pointsAreEqual(p.end, exp.end) {
			t.Errorf("Expected pair: %+v. Received pair: %+v", expected[i], p)
		}

	}
}

func TestFindMinMax(t *testing.T) {
	p := Pair{Point{1, 5}, Point{1, 2}}
	minX, maxX, minY, maxY := findMinMax(p)
	if minX != 1 && minX != maxX {
		t.Errorf("Expected minX and maxX to equal 1")
	}
	if minY != 2 && maxY != 5 {
		t.Errorf("Expected minY to equal 2 and maxY to equal 5. Got minY: %v and maxY: %v.", minY, maxY)
	}

}

func TestMarkMapHorizontalVerticalLinesOne(t *testing.T) {
	m := make(map[Point]int)
	expectedKeys := []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}}
	pair := Pair{Point{1, 1}, Point{1, 3}}

	markMapHorizontalVerticalLines(pair, m)
	for _, p := range expectedKeys {
		if m[p] != 1 {
			t.Errorf("Expected the map to value 1 for key %+v. Got value %v.", p, m[p])
		}
	}
}

func TestMarkMapHorizontalVerticalLinesTwo(t *testing.T) {
	m := make(map[Point]int)
	expectedKeys := []Point{Point{1, 1}, Point{2, 1}, Point{3, 1}}
	pair := Pair{Point{1, 1}, Point{3, 1}}

	markMapHorizontalVerticalLines(pair, m)
	for _, p := range expectedKeys {
		if m[p] != 1 {
			t.Errorf("Expected the map to value 1 for key %+v. Got value %v.", p, m[p])
		}
	}
}

func TestMarkMapDiagonalLines(t *testing.T) {
	m := make(map[Point]int)
	expectedKeys := []Point{Point{1, 1}, Point{2, 2}, Point{3, 3}}
	pair := Pair{Point{1, 1}, Point{3, 3}}

	markMapDiagonalLines(pair, m)
	for _, p := range expectedKeys {
		if m[p] != 1 {
			t.Errorf("Expected the map to value 1 for key %+v. Got value %v.", p, m[p])
		}
	}
}

func TestFindOverlaps(t *testing.T) {
	str := []string{"4,4 -> 4,6", "6,4 -> 4,4"}
	pairs := convertLinesToPairs(str)

	o := findOverlaps(pairs, false /* diagonals */)
	if o != 1 {
		t.Errorf("Expected overlap to be 1. Got %v.", o)
	}
}

func TestFindOverlapsWithDiagonals(t *testing.T) {
	str := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2"}
	pairs := convertLinesToPairs(str)

	o := findOverlaps(pairs, true /* diagonals */)
	if o != 12 {
		t.Errorf("Expected overlap to be 12. Got %v.", o)
	}
}
