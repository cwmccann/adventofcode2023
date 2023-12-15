package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"adventofcode2023/utils"
)
type TestCase = utils.TestCase

var input =
`
TEST INPUT
`


func TestPart1(t *testing.T) {
    tests := []utils.TestCase{
        {
            Name:     "Test 1",
            Input:    "TEST INPUT",
            Expected: -1,
        },
        // Add more test cases here
    }

    assert := assert.New(t)
    for _, tt := range tests {
        t.Run(tt.Name, func(t *testing.T) {
            assert.Equal(tt.Expected, SolvePart1(tt.Input), tt.Name)
        })
    }
}

func TestPart2(t *testing.T) {
    tests := []utils.TestCase{
        {
            Name:     "Test 1",
            Input:    "TEST INPUT",
            Expected: -1,
        },
        // Add more test cases here
    }

    assert := assert.New(t)
    for _, tt := range tests {
        t.Run(tt.Name, func(t *testing.T) {
            assert.Equal(tt.Expected, SolvePart2(tt.Input), tt.Name)
        })
    }
}



