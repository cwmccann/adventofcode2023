package main

import (
	"testing"
)

var input1a =
`
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`

var input1b =
`
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

var input2 =
`
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`

func TestPart1(t *testing.T) {
	solution := SolvePart1(input1a);
	if solution != 2 {
		t.Errorf("Part1 Solution incorrect: %d", solution)
	}
}

func TestPart1b(t *testing.T) {
	solution := SolvePart1(input1b);
	if solution != 6 {
		t.Errorf("Part1 Solution incorrect: %d", solution)
	}
}

func TestPart2(t *testing.T) {
 	solution := SolvePart2(input2);
 	if solution != 6 {
 		t.Errorf("Part2 Solution incorrect: %d", solution)
 	}
}



