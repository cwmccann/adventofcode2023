package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointAdd(t *testing.T) {
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 3, Y: 4}
	expected := Point{X: 4, Y: 6}

	result := p1.Add(p2)

	assert.Equal(t, expected, result, "Addition failed")
}

func TestPointSub(t *testing.T) {
	p1 := Point{X: 5, Y: 8}
	p2 := Point{X: 2, Y: 3}
	expected := Point{X: 3, Y: 5}

	result := p1.Sub(p2)

	assert.Equal(t, expected, result, "Subtraction failed")
}

func TestPointAbs(t *testing.T) {
	p := Point{X: -3, Y: -4}
	expected := Point{X: 3, Y: 4}

	result := p.Abs()

	assert.Equal(t, expected, result, "Abs failed")
}

func TestPointManhattan(t *testing.T) {
	p := Point{X: 3, Y: 4}
	expected := 7

	result := p.Manhattan()

	assert.Equal(t, expected, result, "Manhattan failed")
}

func TestPointManhattanTo(t *testing.T) {
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 4, Y: 6}
	expected := 7

	result := p1.ManhattanTo(p2)

	assert.Equal(t, expected, result, "ManhattanTo failed")
}

func TestPointGetCardinalNeighbors(t *testing.T) {
	p := Point{X: 0, Y: 0}
	expected := []Point{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
	}

	result := p.GetCardinalNeighbors()

	assert.ElementsMatch(t, expected, result, "GetCardinalNeighbors failed")
}

// Add tests for other methods here...
