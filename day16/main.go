package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strings"
	"github.com/gammazero/deque"

)

type Point = utils.Point

func main() {
	text := utils.InputToString();

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}

type Beam struct {
	pos Point
	dir string
	path []Beam
}
func NewBeam(row, col int, dir string) Beam {
    return Beam{pos: Point{X: col, Y: row}, dir: dir, path: make([]Beam, 0)}
}
func (b *Beam) Copy() Beam {
	return Beam{pos: b.pos, dir: b.dir, path: b.path}
}
func (b *Beam) ChangeDir(dir string) Beam {
	b.dir = dir
	return *b
}
func (b *Beam) Move() {
    b.path = append(b.path, *b)

	switch b.dir {
    case "R":
        b.pos.X++
    case "L":
        b.pos.X--
    case "U":
        b.pos.Y--
    case "D":
        b.pos.Y++
    }
}
func (b *Beam) IsValidPosition(R, C int) bool {
    return b.pos.X >= 0 && b.pos.X < C && b.pos.Y >= 0 && b.pos.Y < R
}
func (b *Beam) String() string {
	return fmt.Sprintf("Pos: %v, Dir: %s", b.pos, b.dir)
}

func SolvePart1(input string) int {
	lines := strings.Split(input, "\n")
	grid := utils.StringsToRune2D(lines)
	return Solve(grid, NewBeam(0, -1, "R"))
}

func Solve(grid [][]rune, startingBeam Beam) int {
	R := len(grid)
	C := len(grid[0])

	beams := deque.New[Beam]()
	beams.PushBack(startingBeam)

	energized := make([]Point, 0)
	seen := make(map[Point][]string)
	
	for beams.Len() > 0 {
		beam := beams.PopFront()
		beam.Move()

		//Drop beams if it's off teh grid
		if !beam.IsValidPosition(R, C) {
			continue
		}

		//Drop beams if it's already been here
		if _, ok := seen[beam.pos]; !ok {
			seen[beam.pos] = make([]string, 0)
		}
		if utils.Contains(seen[beam.pos], beam.dir) {
			continue
		}
		seen[beam.pos] = append(seen[beam.pos], beam.dir)

		//Add the beam position to the energized list
		if (!inList(energized, beam.pos.Y, beam.pos.X)) {
			energized = append(energized, beam.pos)
		}

		char := grid[beam.pos.Y][beam.pos.X]

		//Keep going if it's clear
		if char == '.' {
			beams.PushBack(beam)
			continue
		}

		//Change direction if it's a \
		if char == '\\' {
			if beam.dir == "R" {
				beams.PushBack(beam.ChangeDir("D"))
			} else if beam.dir == "L" {
				beams.PushBack(beam.ChangeDir("U"))
			} else if beam.dir == "U" {
				beams.PushBack(beam.ChangeDir("L"))
			} else if beam.dir == "D" {
				beams.PushBack(beam.ChangeDir("R"))
			}
			continue
		}
		//Change direction if it's a /
		if char == '/' {
			if beam.dir == "R" {
				beams.PushBack(beam.ChangeDir("U"))
			} else if beam.dir == "L" {
				beams.PushBack(beam.ChangeDir("D"))

			} else if beam.dir == "U" {
				beams.PushBack(beam.ChangeDir("R"))
			} else if beam.dir == "D" {
				beams.PushBack(beam.ChangeDir("L"))
			}
			continue
		}

		//Split if we hit a - and traveling vertically
		if char == '-' {
			if beam.dir == "R" || beam.dir == "L" {
				beams.PushBack(beam)
			} else {
				beams.PushBack(beam.ChangeDir("R"))
				newBeam := beam.Copy()
				beams.PushBack(newBeam.ChangeDir("L"))
			}
			continue
		}

		//Split if we hit a | and traveling horizontally
		if char == '|' {
			if beam.dir == "U" || beam.dir == "D" {
				beams.PushBack(beam)
			} else {
				beams.PushBack(beam.ChangeDir("U"))
				newBeam := beam.Copy()
				beams.PushBack(newBeam.ChangeDir("D"))
			}
			continue
		}
	}
	// printEnergized(energized, grid)
	return len(energized)
}

func printEnergized(energized []Point, grid [][]rune) {
	R := len(grid)
	C := len(grid[0])

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if (grid[r][c] == '.') {
				if inList(energized, r, c) {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(string(grid[r][c]))
			}
		}
		fmt.Println()
	}
}

func inList(points []Point, r, c int) bool {
	for _, p := range points {
		if p.X == c && p.Y == r {
			return true
		}
	}
	return false
}

func SolvePart2(input string) int {
	max := 0
	lines := strings.Split(input, "\n")
	grid := utils.StringsToRune2D(lines)
	R := len(grid)
	C := len(grid[0])

	for r := 0; r < R; r++ {
		max = utils.Max(max, Solve(grid, NewBeam(r, -1, "R")))
		max = utils.Max(max, Solve(grid, NewBeam(r, C, "L")))
	}

	for c := 0; c < C; c++ {
		max = utils.Max(max, Solve(grid, NewBeam(-1, c, "D")))
		max = utils.Max(max, Solve(grid, NewBeam(R, c, "U")))
	}

	return max;
}
