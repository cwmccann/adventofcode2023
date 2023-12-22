package main

import (
	"adventofcode2023/utils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)
type TestCase = utils.TestCase

var input =
`
...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
`


func TestPart1(t *testing.T) {
    tests := []struct {
        name string
        input string
        steps int
        expected int
    }{
        {
            name: "Test 1",
            input: input,
            steps: 1,
            expected: 2,
        },
        {
            name: "Test 2",
            input: input,
            steps: 2,
            expected: 4,
        },
        {
            name: "Test 6",
            input: input,
            steps: 6,
            expected: 16,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            lines := utils.TrimAndRemoveEmptyLines(strings.Split(tt.input, "\n"))
	        grid := utils.StringsToRune2D(lines)
            start := findStart(grid)

            result := CountSquares(grid, tt.steps, start);
            assert.Equal(t, tt.expected, result)
        })
    }
}

func TestPart2(t *testing.T) {
    tests := []struct {
        name string
        steps int
        expected int
    }{
        {
            name: "Test 6 steps",
            steps: 6,
            expected: 16,
        },
        {
            name: "Test 10 steps",
            steps: 10,
            expected: 50,
        },
        {
            name: "Test 50 steps",
            steps: 50,
            expected: 1594,
        },
        {
            name: "Test 10 steps",
            steps: 100,
            expected: 6536,
        },
        {
            name: "Test 1000 steps",
            steps: 1000,
            expected: 668697,
        },


    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            lines := utils.TrimAndRemoveEmptyLines(strings.Split(input, "\n"))
	        grid := utils.StringsToRune2D(lines)
            start := findStart(grid)

            result := CountSquareInfinity(grid, tt.steps, start);
            assert.Equal(t, tt.expected, result)
        })
    }
}

