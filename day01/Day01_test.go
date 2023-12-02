package main

import ( "testing" )

var input =
`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`
var input2 =
`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

func TestDay01Part1(t *testing.T) {
	solution := SolvePart1(input);
	if solution != 142 {
		t.Errorf("Solution incorrect: %d", solution)
	}
}

func TestDay01Part2(t *testing.T) {
	solution := SolvePart2(input2);
	if solution != 281 {
		t.Errorf("Solution incorrect: %d", solution)
	}
}
