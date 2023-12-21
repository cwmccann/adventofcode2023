package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"adventofcode2023/utils"
)
type TestCase = utils.TestCase

var input =
`
broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`

var input2 =
`
broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
`

func TestPart1(t *testing.T) {
    tests := []utils.TestCase{
        {
            Name:     "Test 1",
            Input:    input,
            Expected: 32000000,
        },
        {
            Name:     "Test 2",
            Input:    input2,
            Expected: 11687500,
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


