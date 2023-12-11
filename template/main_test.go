package main

import (
	"testing"
	"github.com/stretchr/testify/assert"

)

var input =
`
TEST INPUT
`


func TestPart1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(-1, SolvePart1(input), "input")
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(-1, SolvePart2(input), "input")
}



