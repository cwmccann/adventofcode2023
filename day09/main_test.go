package main

import (
	"testing"
)

const part1Solution = 114
const part2Solution = 2

var input =
`
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func TestFindNext(t *testing.T) {
	nums := []int{0, 3, 6, 9, 12, 15}
	solution := FindNext(nums)
	if solution != 18 {
		t.Errorf("FindNext incorrect: %d", solution)
	}

	nums = []int{1, 3, 6, 10, 15, 21}
	solution = FindNext(nums)
	if solution != 28 {
		t.Errorf("FindNext incorrect: %d", solution)
	}

	nums = []int{10, 13, 16, 21, 30, 45}
	solution = FindNext(nums)
	if solution != 68 {
		t.Errorf("FindNext incorrect: %d", solution)
	}
}





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



