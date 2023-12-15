package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strings"
)

func main() {
	text := utils.InputToString();

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}


func SolvePart1(input string) int {
	lines := strings.Split(input, "\n")
	grid := utils.StringsToRune2D(lines)
	grid = tiltGridNorth(grid)
	weight := calcWeight(grid)
	return weight
}


func SolvePart2(input string) int {
	lines := strings.Split(input, "\n")
	grid := utils.StringsToRune2D(lines)

	grid = ApplyNTimes(grid, 1000000000)
	weight := calcWeight(grid)
	return weight
}


func calcWeight(grid [][]rune) int {
	total := 0

	maxWeight := len(grid)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			char := grid[row][col]
			if char == 'O' {
				total += maxWeight - row
			}
		}
	}
	return total
}

func spinGrid(grid [][]rune) [][]rune {
	grid = tiltGrid(grid, -1, 0)
	grid = tiltGrid(grid, 0, -1)
	grid = tiltGrid(grid, 1, 0)
	grid = tiltGrid(grid, 0, 1)
	return grid
}

func tiltGridNorth(grid [][]rune) [][]rune {
	tiltGrid(grid, -1, 0)
	return grid
}

func tiltGrid(grid [][]rune, yDirection, xDirection int) [][]rune {
	anyChanges := true
	for anyChanges {
		anyChanges = false

		for row := 0; row < len(grid); row++ {
			for col := 0; col < len(grid[row]); col++ {
				char := grid[row][col]
				if char == '.' || char == '#' {
					continue
				}

				if row + yDirection < 0 || row + yDirection >= len(grid) {
					continue
				}

				if col + xDirection < 0 || col + xDirection >= len(grid[row]) {
					continue
				}

				if char == 'O' && grid[row + yDirection][col + xDirection] == '.' {
					anyChanges = true
					grid[row + yDirection][col + xDirection] = 'O'
					grid[row][col] = '.'
				}
			}
		}
	}

	return grid
}


func gridAsString(grid [][]rune) string {
    var gridStr strings.Builder
    for _, row := range grid {
        gridStr.WriteString(string(row))
        gridStr.WriteString("\n")
    }
    return gridStr.String()
}

func printGrid(grid [][]rune) {
	fmt.Println(gridAsString(grid))
}

func ApplyNTimes(grid [][]rune, n int) [][]rune {
	seen := make(map[string]int)
	cycleDetected := false
	i := 0

	for ; i < n; i++ {
		hash := gridAsString(grid)
		if _, ok := seen[hash]; ok {
			cycleDetected = true
			break
		}
		seen[hash] = i
		grid = spinGrid(grid)
	}

	if !cycleDetected {
		return grid
	}

	cycleStart := seen[gridAsString(grid)]
	cycleLength := i - cycleStart
	remaining := (n - i) % cycleLength

	return ApplyNTimes(grid, remaining)
}

