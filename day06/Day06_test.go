package main

import (
	"testing"
)

//destination range start, the source range start, and the range length.
var input =
`
Time:      7  15   30
Distance:  9  40  200
`


func TestDay01Part1(t *testing.T) {
	solution := SolvePart1(input);
	if solution != 288 {
		t.Errorf("Solution incorrect: %d", solution)
	}
}

func TestDay01Part2(t *testing.T) {
 	solution := SolvePart2(input);
 	if solution != 71503 {
 		t.Errorf("Solution incorrect: %d", solution)
 	}
}



