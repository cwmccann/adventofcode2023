package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var input =
`
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`
// `
// ...1...... 0
// .......2.. 1
// 3......... 2
// .......... 3
// ......4... 4
// .5........ 5
// .........6 6
// .......... 7
// .......7.. 8
// 8...9..... 9
// 0123456789
// `

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(374, SolvePart1(input), "input")
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(374, SolvePart2(input, 2), "input")
	assert.Equal(1030, SolvePart2(input, 10), "input")
	assert.Equal(8410, SolvePart2(input, 100), "input")
}




