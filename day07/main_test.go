package main

import (
	"testing"
)

//destination range start, the source range start, and the range length.
var input =
`
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`


func TestDay01Part1(t *testing.T) {
	solution := SolvePart1(input);
	if solution != 6440 {
		t.Errorf("Solution incorrect: %d", solution)
	}
}

func TestDay01Part2(t *testing.T) {
 	solution := SolvePart2(input);
 	if solution != 5905 {
 		t.Errorf("Solution incorrect: %d", solution)
 	}
}



