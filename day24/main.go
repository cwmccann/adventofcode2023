package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strings"
	"gonum.org/v1/gonum/mat"
)

type Hail struct {
	x, y, z int
	dx, dy, dz int
}
func (h *Hail) String() string {
	return fmt.Sprintf("<pos=(%d,%d,%d), vel=(%d,%d,%d)>", h.x, h.y, h.z, h.dx, h.dy, h.dz)
}


func main() {
	text := utils.InputToString();

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}


func findIntersects2D(a, b Hail) (float64, float64, bool) {
	// Create matrix A
	A := mat.NewDense(2, 2, []float64{
		float64(a.dx), -float64(b.dx),
		float64(a.dy), -float64(b.dy),
	})

	// Create vector b
	B := mat.NewVecDense(2, []float64{
		float64(b.x - a.x),
		float64(b.y - a.y),
	})

	// Solve for t
	var t mat.VecDense
	if err := t.SolveVec(A, B); err != nil {
		return 0, 0, false // No solution (lines are parallel or collinear)
	}

	// Check if t1 and t2 are greater than 0
	if t.AtVec(0) > 0 && t.AtVec(1) > 0 {
		// Intersection point using t1
		intersectionX := float64(a.x) + t.AtVec(0)*float64(a.dx)
		intersectionY := float64(a.y) + t.AtVec(0)*float64(a.dy)
		return intersectionX, intersectionY, true
	}

	return 0, 0, false // No valid intersection

}

func countHailIntersects(hails []Hail, min , max int) int {
	total := 0
	for i :=0 ; i < len(hails); i++ {
		for j := i+1; j < len(hails); j++ {
			//fmt.Println(i, j)
			a := hails[i]
			b := hails[j]

			if x, y, intersects := findIntersects2D(a, b); intersects {

				if x >= float64(min) && x <= float64(max) && y >= float64(min) && y <= float64(max) {
					//fmt.Printf("%v intersects %v at %f %f", a, b, x, y)
					total++
				}
			}
		}
	}

	return total
}

func SolvePart1(input string) int {
	hails := parseInput(input)
	return countHailIntersects(hails, 200000000000000, 400000000000000)
}

func parseInput(input string) []Hail {
	lines := strings.Split(input, "\n")
	hails := make([]Hail, 0)

	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}
		hails = append(hails, parseLine(line))
	}
	return hails
}

func parseLine(line string) Hail {
	parts := strings.Split(line, "@");
	positions := utils.StringToIntSlice2(parts[0], ",")
	velocities := utils.StringToIntSlice2(parts[1], ",")

	return Hail{
		x: positions[0],
		y: positions[1],
		z: positions[2],
		dx: velocities[0],
		dy: velocities[1],
		dz: velocities[2],
	}
}

func SolvePart2(input string) int {
	return -1;
}
