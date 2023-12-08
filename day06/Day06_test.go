package main

import (
	"testing"
)

var input =
`
Time:      7  15   30
Distance:  9  40  200
`


func TestPart1(t *testing.T) {
	solution := SolvePart1(input);
	if solution != 288 {
		t.Errorf("Solution incorrect: %d", solution)
	}
}

func TestPart2(t *testing.T) {
 	solution := SolvePart2(input);
 	if solution != 71503 {
 		t.Errorf("Solution incorrect: %d", solution)
 	}
}



