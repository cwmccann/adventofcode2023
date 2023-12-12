package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strings"
)

type Point = utils.Point

func main() {
	text := utils.InputToString();

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text, 1000000)
	fmt.Printf("Solution Part 2: %d\n", solution)
}


func SolvePart1(input string) int {
	return SolvePart2(input, 2)
}

func SolvePart2(input string, expansionFactor int) int {
	lines := strings.Split(input, "\n")
	grid := utils.StringsToRune2D(lines)

	galaxies := findGalaxies(grid)
	expandingCols, expandingRows := findExpandingRowsAndCols(grid)
	sum := 0

	for i, galaxy1 := range galaxies {
		for j := i+1; j < len(galaxies); j++ {
			galaxy2 := galaxies[j]

			sum += DistanceBetweenPoints(galaxy1, galaxy2, expandingCols, expandingRows, expansionFactor)
		}
	}
	return sum
}

func findGalaxies(grid [][]rune) []Point {
	galaxies := make([]Point, 0)
	for i, row := range grid {
		for j, char := range row {
			if char == '#' {
				galaxies = append(galaxies, Point{X: j, Y: i})
			}
		}
	}
	return galaxies
}


func DistanceBetweenPoints(p1, p2 Point, expandingCols []int, expandingRows []int, expansionFactor int) int {
	distance := 0
	xRange := utils.RangeBetween(p1.X, p2.X)

	for i := 0; i < len(xRange); i++ {
		if ContainsPoint(expandingCols, xRange[i]) {
			xRange[i] = expansionFactor
		} else {
			xRange[i] = 1
		}
	}

	distance += utils.Sum(xRange)

	yRange := utils.RangeBetween(p1.Y, p2.Y)
	for i := 0; i < len(yRange); i++ {
		if ContainsPoint(expandingRows, yRange[i]) {
			yRange[i] = expansionFactor
		} else {
			yRange[i] = 1
		}
	}

	distance += utils.Sum(yRange)
	return distance
}


func ContainsPoint(slice []int, point int) bool {
	for _, p := range slice {
		if p == point {
			return true
		}
	}
	return false
}

func findExpandingRowsAndCols(grid [][]rune) ([]int, []int) {
	expandingRows := make([]int, 0)

	//Expand the grid if there are rows and columns of only '.'
	for i, row := range grid {
		if utils.All(row, func(r rune) bool { return r == '.' }) {
			expandingRows = append(expandingRows, i)
		}
	}

	expandingCols := make([]int, 0)
	for j := 0; j < len(grid[0]); j++ {
		allEmpty := true
		for i := 0; i < len(grid); i++ {
			if grid[i][j] != '.' {
				allEmpty = false
				break
			}
		}
		if allEmpty {
			expandingCols = append(expandingCols, j)
		}
	}

	return expandingCols, expandingRows
}


