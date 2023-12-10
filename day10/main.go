package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strings"
	"github.com/gammazero/deque"
)

func main() {
	text := utils.InputToString()

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}

type Point struct {
	X, Y int
}

func SolvePart1(input string) int {
	lines := strings.Split(input, "\n")
	grid := utils.StringsToRune2D(lines)
	loop, _ := findLoop(grid)
	return (len(loop) + 1) / 2 //Go rounds down so so add one
}

//Find the loop and the starting character
func findLoop(grid [][]rune) ([]Point, rune) {
	start := findStart(grid)

	sPossibilities := "|-LJF7"

	q := deque.New[Point]()
	q.PushBack(start)

	loop := make([]Point, 0)
	loop = append(loop, start)


	for q.Len() > 0 {
		p := q.PopFront()
		char := grid[p.Y][p.X]

		//Check up
		if p.Y > 0 &&
			strings.ContainsRune("S|LJ", char) &&
			strings.ContainsRune("|F7", grid[p.Y-1][p.X]) &&
			!inList(loop, p.Y - 1, p.X) {

			q.PushBack(Point{X: p.X, Y: p.Y - 1})
			loop = append(loop, Point{X: p.X, Y: p.Y - 1})

			if char == 'S' {
				sPossibilities = utils.RemoveAll(sPossibilities, "-F7")
			}
		}
		//Check down
		if p.Y < len(grid) - 1 &&
			strings.ContainsRune("S|7F", char) &&
			strings.ContainsRune("|LJ", grid[p.Y+1][p.X]) &&
			!inList(loop, p.Y + 1, p.X) {

			q.PushBack(Point{X: p.X, Y: p.Y + 1})
			loop = append(loop, Point{X: p.X, Y: p.Y + 1})

			if char == 'S' {
				sPossibilities = utils.RemoveAll(sPossibilities, "-JL")
			}
		}

		//Check left
		if p.X > 0 &&
			strings.ContainsRune("S-7J", char) &&
			strings.ContainsRune("-LF", grid[p.Y][p.X-1]) &&
			!inList(loop, p.Y, p.X - 1) {

				q.PushBack(Point{X: p.X - 1, Y: p.Y})
				loop = append(loop, Point{X: p.X - 1, Y: p.Y})
				if char == 'S' {
					sPossibilities = utils.RemoveAll(sPossibilities, "|FL")
				}
		}

		//Check right
		if p.X < len(grid[p.Y]) - 1 &&
			strings.ContainsRune("S-LF", char) &&
			strings.ContainsRune("-7J", grid[p.Y][p.X+1]) &&
			!inList(loop, p.Y, p.X + 1) {
				q.PushBack(Point{X: p.X + 1, Y: p.Y})
				loop = append(loop, Point{X: p.X + 1, Y: p.Y})

				if char == 'S' {
					sPossibilities = utils.RemoveAll(sPossibilities, "|7J")
				}
		}
	}

	if len(sPossibilities) != 1 {
		panic("Could not find start char: " + sPossibilities)
	}

	// fmt.Printf("Points: %v\n", len(loop))
	// fmt.Printf("Loop: %v\n", loop)
	return loop, rune(sPossibilities[0])
}

func inList(list []Point, y, x int) bool {
	for _, p := range list {
		if p.X == x && p.Y == y {
			return true
		}
	}
	return false
}

func SolvePart2(input string) int {
	lines := strings.Split(input, "\n")
	grid := utils.StringsToRune2D(lines)
	loop, startChar := findLoop(grid)

	//Get rid of the extra pipe characters that are not in the loop
	for y := range grid {
		for x := range grid[y] {
			if !inList(loop, y, x) {
				grid[y][x] = '.'
			}
		}
	}
	//printGrid(grid)

	start := findStart(grid)
	grid[start.Y][start.X] = startChar

	//Expand the grid to 2x2 blocks
	grid = replaceCharsWithBlocks(grid)

	changedAny := true

	//Flood Fill
	for changedAny {
		changedAny = false

		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] == '.' {
					//Outer edges
					if y == 0 || y == len(grid)-1 || x == 0 || x == len(grid[y])-1 {
						changedAny = true
						grid[y][x] = 'O'
						continue
					}

					//check is any adjacent is open
					if anyOpen(grid, y, x) {
						changedAny = true
						grid[y][x] = 'O'
					}
				}
			}
		}
	}

	//Count the remaining 2x2 blocks of .
	count := 0
	for i := 0; i < len(grid)-1; i += 2 {
		for j := 0; j < len(grid[i])-1; j += 2 {
			//Only the top left char is important
			if grid[i][j] == '.' {
				count++
			}
		}
	}
	//printGrid(grid)

	return count
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func replaceCharsWithBlocks(grid [][]rune) [][]rune {
	blocks := map[rune]string{
		'S': "SS\nSS",
		'|': "|.\n|.",
		'-': "--\n..",
		'L': "L-\n..",
		'J': "J.\n..",
		'7': "7.\n|.",
		'F': "F-\n|.",
		'.': "..\n..",
	}

	newGrid := make([][]rune, len(grid)*2)
	for i := range newGrid {
		newGrid[i] = make([]rune, len(grid[0])*2)
	}

	for i, row := range grid {
		for j, r := range row {
			block := strings.Split(strings.Trim(blocks[r], "\n"), "\n")
			for bi, brow := range block {
				for bj, br := range brow {
					newGrid[i*2+bi][j*2+bj] = br
				}
			}
		}
	}

	return newGrid
}

func anyOpen(grid [][]rune, y, x int) bool {
	if y > 0 && grid[y-1][x] == 'O' {
		return true
	}
	if x < len(grid[y]) && grid[y][x+1] == 'O' {
		return true
	}
	if y < len(grid) && grid[y+1][x] == 'O' {
		return true
	}
	if x > 0 && grid[y][x-1] == 'O' {
		return true
	}
	return false
}

func findStart(rune2D [][]rune) (Point) {
	for i, row := range rune2D {
		for j, r := range row {
			if r == 'S' {
				return Point{X: j, Y: i}
			}
		}
	}
	return Point{X: -1, Y: -1}
}

