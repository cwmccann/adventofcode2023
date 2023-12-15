package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"adventofcode2023/utils"
)
type TestCase = utils.TestCase

var input =
`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`


func TestHash(t *testing.T) {
    tests := []utils.TestCase{
        {
            Name:     "Test 1",
            Input:    "HASH",
            Expected: 52,
        },
        {
            Name:     "Test 2",
            Input:    "rn",
            Expected: 0,
        },
        // Add more test cases here
    }

    assert := assert.New(t)
    for _, tt := range tests {
        t.Run(tt.Name, func(t *testing.T) {
            assert.Equal(tt.Expected, Hash(tt.Input), tt.Name)
        })
    }
}

func TestPart1(t *testing.T) {
    tests := []utils.TestCase{
        {
            Name:     "Test 1",
            Input:    input,
            Expected: 1320,
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
            Expected: 145,
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




