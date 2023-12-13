package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`

func TestHReflect(t *testing.T) {
	assert := assert.New(t)
	input :=
		`#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

	lines := strings.Split(input, "\n")
	assert.Equal(4, findHReflect(lines, equals), "input")
}

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(405, SolvePart1(input), "input")
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(400, SolvePart2(input), "input")
}
