package main

import (
	"adventofcode2023/utils"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"strings"
	"math"
	"time"
)

type Vector struct {
	x, y, z    int
	dx, dy, dz int
}

func (h *Vector) String() string {
	return fmt.Sprintf("<pos=(%d,%d,%d), vel=(%d,%d,%d)>", h.x, h.y, h.z, h.dx, h.dy, h.dz)
}

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

func findIntersects2D(a, b Vector) (float64, float64, bool) {
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

func countHailIntersects(hails []Vector, min, max int) int {
	total := 0
	for i := 0; i < len(hails); i++ {
		for j := i + 1; j < len(hails); j++ {
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

func parseInput(input string) []Vector {
	lines := strings.Split(input, "\n")
	hails := make([]Vector, 0)

	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}
		hails = append(hails, parseLine(line))
	}
	return hails
}

func parseLine(line string) Vector {
	parts := strings.Split(line, "@")
	positions := utils.StringToIntSlice2(parts[0], ",")
	velocities := utils.StringToIntSlice2(parts[1], ",")

	return Vector{
		x:  positions[0],
		y:  positions[1],
		z:  positions[2],
		dx: velocities[0],
		dy: velocities[1],
		dz: velocities[2],
	}
}

func findRock(hails []Vector) Vector {
	h0 := hails[0]
	h1 := hails[1]
	h2 := hails[2]

	aData := []int{
		h1.dy - h0.dy, h0.dx - h1.dx, 0, h0.y - h1.y, h1.x - h0.x, 0,
		h2.dy - h0.dy, h0.dx - h2.dx, 0, h0.y - h2.y, h2.x - h0.x, 0,
		h1.dz - h0.dz, 0, h0.dx - h1.dx, h0.z - h1.z, 0, h1.x - h0.x,
		h2.dz - h0.dz, 0, h0.dx - h2.dx, h0.z - h2.z, 0, h2.x - h0.x,
		0, h1.dz - h0.dz, h0.dy - h1.dy, 0, h0.z - h1.z, h1.y - h0.y,
		0, h2.dz - h0.dz, h0.dy - h2.dy, 0, h0.z - h2.z, h2.y - h0.y,
	}
	a := mat.NewDense(6, 6, convertToFloat(aData))

	bData := make([]int, 6)
	bData[0] = (h0.y*h0.dx - h1.y*h1.dx) - (h0.x*h0.dy - h1.x*h1.dy)
	bData[1] = (h0.y*h0.dx - h2.y*h2.dx) - (h0.x*h0.dy - h2.x*h2.dy)
	bData[2] = (h0.z*h0.dx - h1.z*h1.dx) - (h0.x*h0.dz - h1.x*h1.dz)
	bData[3] = (h0.z*h0.dx - h2.z*h2.dx) - (h0.x*h0.dz - h2.x*h2.dz)
	bData[4] = (h0.z*h0.dy - h1.z*h1.dy) - (h0.y*h0.dz - h1.y*h1.dz)
	bData[5] = (h0.z*h0.dy - h2.z*h2.dy) - (h0.y*h0.dz - h2.y*h2.dz)
	b := mat.NewDense(6, 1, convertToFloat(bData))

	// fmt.Printf("a = \n%.2f\n\n", mat.Formatted(a, mat.Prefix(""), mat.Squeeze()))
	// fmt.Printf("b = \n%.2f\n\n", mat.Formatted(b, mat.Prefix(""), mat.Squeeze()))

	var x mat.Dense
	if err := x.Solve(a, b); err != nil {
		panic(err)
	}
	// fmt.Printf("x = \n%.2f\n\n", mat.Formatted(&x, mat.Prefix(""), mat.Squeeze()))
	rock := Vector{
		x:  int(math.Round(x.At(0, 0))),
		y:  int(math.Round(x.At(1, 0))),
		z:  int(math.Round(x.At(2, 0))),
		dx: int(math.Round(x.At(3, 0))),
		dy: int(math.Round(x.At(4, 0))),
		dz: int(math.Round(x.At(5, 0))),
	}
	return rock
}

func convertToFloat(data []int) []float64 {
	result := make([]float64, len(data))
	for i, v := range data {
		result[i] = float64(v)
	}
	return result
}

func SolvePart2(input string) int {
	hails := parseInput(input)
	rock := findRock(hails)
	return rock.x + rock.y + rock.z
}
