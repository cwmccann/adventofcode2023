package main

import (
	"adventofcode2023/utils"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gammazero/deque"
	"strings"
	"time"
)

type Point = utils.Point

func main() {
	start := time.Now()
	text := utils.InputToString()

	solution := SolvePart1(text)
	elapsed := time.Since(start)

	fmt.Printf("Solution Part 1: %d took %s\n", solution, elapsed)

	start = time.Now()
	solution = SolvePart2(text)
	elapsed = time.Since(start)
	fmt.Printf("Solution Part 2: %d took %sa\n", solution, elapsed)
}

func SolvePart1(input string) int {
	lines := utils.TrimAndRemoveEmptyLines(strings.Split(input, "\n"))
	grid := utils.StringsToRune2D(lines)

	start := findStart(grid)

	return CountSquares(grid, 64, start)
}

type WorkItem struct {
	point Point
	step  int
}

func CountSquares(grid [][]rune, steps int, start Point) int {
	R := len(grid)
	C := len(grid[0])

	q := deque.New[WorkItem]()
	seen := mapset.NewSet[Point]()
	q.PushBack(WorkItem{point: start, step: steps})

	for q.Len() > 0 {
		w := q.PopFront()
		p := w.point
		n := w.step

		if seen.Contains(p) {
			continue
		}

		if n%2 == 0 {
			seen.Add(p)
		}

		if n == 0 {
			continue
		}

		for _, neighbor := range p.GetCardinalNeighbors() {
			c := GetCharFromPoint(grid, neighbor)

			if neighbor.IsInGrid(R, C) && !seen.Contains(neighbor) && utils.RuneInString(c, ".S") {
				q.PushBack(WorkItem{point: neighbor, step: n - 1})
			}
		}
	}

	return seen.Cardinality()
}

func CountSquareInfinity(grid [][]rune, steps int, start Point) int {
	q := deque.New[WorkItem]()
	seen := mapset.NewSet[Point]()
	q.PushBack(WorkItem{point: start, step: steps})

	//lastSeen := 0
	for q.Len() > 0 {
		w := q.PopFront()
		p := w.point
		n := w.step

		if seen.Contains(p) {
			continue
		}

		if n%2 == 0 {
			seen.Add(p)

			// if GetCharFromPoint(grid, p) == 'S' {
			// 	stepsTaken := steps - n
			// 	d := stepsTaken - lastSeen
			// 	lastSeen = stepsTaken
			// 	if d != 0 {
			// 		fmt.Printf("%d,%d\n", stepsTaken, seen.Cardinality())
			// 	}
			// }

		}

		if n == 0 {
			continue
		}

		for _, neighbor := range p.GetCardinalNeighbors() {
			c := GetCharFromPoint(grid, neighbor)

			if !seen.Contains(neighbor) && utils.RuneInString(c, ".S") {
				q.PushBack(WorkItem{point: neighbor, step: n - 1})
			}
		}
	}
	return seen.Cardinality()
}

func GetCharFromPoint(grid [][]rune, p Point) rune {
	m := MapPoint(p, len(grid), len(grid[0]))
	return grid[m.Y][m.X]
}

func MapPoint(p Point, R int, C int) Point {
	return Point{
		X: (p.X%C + C) % C,
		Y: (p.Y%R + R) % R,
	}
}

func findStart(grid [][]rune) Point {
	for y, row := range grid {
		for x, char := range row {
			if char == 'S' {
				return Point{X: x, Y: y}
			}
		}
	}
	return Point{X: -1, Y: -1}
}

func SolvePart2(input string) int {
	lines := utils.TrimAndRemoveEmptyLines(strings.Split(input, "\n"))
	grid := utils.StringsToRune2D(lines)
	R := len(grid)
	C := len(grid[0])

	start := findStart(grid)

	//Check somethings that should be true
	if R != C {
		panic("grid must be square")
	}
	if start.X != (C-1)/2 || start.Y != (R-1)/2 {
		panic("start must be in the middle of the grid")
	}
	if R%2 == 0 || C%2 == 0 {
		panic("grid must be odd")
	}

	//Use Newtons' Interpolation method to find the formula for the number of squares

	//Get the first 3 y values.  Use the number of tiles as the x value
	y0 := CountSquareInfinity(grid, start.X+C*0, start)
	y1 := CountSquareInfinity(grid, start.X+C*1, start)
	y2 := CountSquareInfinity(grid, start.X+C*2, start)

	points := []Point{
		{X: 0, Y: y0},
		{X: 1, Y: y1},
		{X: 2, Y: y2},
	}
	coefficients := NewtonInterpolation(points)
	//fmt.Println("coefficients: ", coefficients)

	//Figure out which x value we need to get to
	//which is the number of tiles in the end grid
	steps := 26501365
	x := (steps - start.X) / C
	//fmt.Printf("x: %d \n", x)

	ans := EvaluatePolynomial(coefficients, points, x)
	return ans
}

// Calculate the coefficients of the Newton polynomial
// Given a set of points, calculate the coefficients of the Newton polynomial
func NewtonInterpolation(points []Point) []int {
	n := len(points)
	coefficients := make([]int, n)
	for i := 0; i < n; i++ {
		coefficients[i] = points[i].Y
	}

	for i := 1; i < n; i++ {
		for j := n - 1; j >= i; j-- {
			coefficients[j] = (coefficients[j] - coefficients[j-1]) / (points[j].X - points[j-i].X)
		}
	}
	return coefficients
}

// Evaluate the Newton polynomial
func EvaluatePolynomial(coefficients []int, points []Point, x int) int {
	n := len(coefficients)
	result := coefficients[n-1]

	for i := n - 2; i >= 0; i-- {
		result = result*(x-points[i].X) + coefficients[i]
	}

	return result
}
