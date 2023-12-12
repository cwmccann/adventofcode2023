package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input =
`
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`

func TestCountArrangements(t *testing.T) {
	assert := assert.New(t)
	lines := strings.Split(input, "\n")

	s, groups := parseLine(lines[1])
	assert.Equal(1, countArrangements(s, groups, NewState(s, 0, 0, 0)), s)

	cleanCache()
	s, groups = parseLine(lines[6])
	assert.Equal(10, countArrangements(s, groups, NewState(s, 0, 0, 0)), s)
}

func TestPart1(t *testing.T) {
 	assert := assert.New(t)
 	assert.Equal(21, SolvePart1(input), "input")
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(525152, SolvePart2(input), "input")
}

func TestPart2_2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, SolvePart2("???.### 1,1,3"), "input")
	assert.Equal(16384, SolvePart2(".??..??...?##. 1,1,3"), "input")
}


