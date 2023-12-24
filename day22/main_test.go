package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"adventofcode2023/utils"
)
type TestCase = utils.TestCase

var input =
`
1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9
`


func TestPart1(t *testing.T) {
    tests := []utils.TestCase{
        {
            Name:     "Test 1",
            Input:    input,
            Expected: 5,
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
            Input:    input,
            Expected: 7,
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



