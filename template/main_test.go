package main

import (
	"testing"
)

const part1Solution = -1
const part2Solution = -1

var input =
`
TEST INPUT
`


func TestPart1(t *testing.T) {
	solution := SolvePart1(input);
	if solution != part1Solution {
		t.Errorf("Part1 Solution incorrect: %d", solution)
	}
}

func TestPart2(t *testing.T) {
 	solution := SolvePart2(input);
 	if solution != part2Solution {
 		t.Errorf("Part2 Solution incorrect: %d", solution)
 	}
}



