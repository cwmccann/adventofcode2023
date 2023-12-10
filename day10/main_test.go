package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const part1Solution = 4


var input =
`
-L|F7
7S-7|
L|7||
-L-J|
L|-JF
`

var input2 =
`
7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
`

var input3 =
`
..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........
`

var input4 =
`
FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
`

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(4, SolvePart1(input), "input")
	assert.Equal(8, SolvePart1(input2), "input2")
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, SolvePart2(input), "input")
	assert.Equal(1, SolvePart2(input2), "input2")
	assert.Equal(4, SolvePart2(input3), "input3")
	assert.Equal(10, SolvePart2(input4), "input4")
}



