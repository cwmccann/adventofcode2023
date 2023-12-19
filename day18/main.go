package main

import (
	"adventofcode2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point utils.Point

func main() {
	text := utils.InputToString()

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}

func SolvePart1(input string) int64 {
	return Solve(input, parseLine1)
}

func SolvePart2(input string) int64 {
	return Solve(input, parseLine2)
}

func Solve(input string, parseLine func(string) (int, int)) int64 {
	lines := strings.Split(input, "\n")

	vertices := make([]Point, 0)

	cur := Point{0, 0}
	vertices = append(vertices, cur)
	perimeter := 0

	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}
		direction, distance := parseLine(line)
		perimeter += distance


		switch direction {
		case 0:
			//Right
			cur = Point{cur.X + distance, cur.Y}
		case 1:
			//Down
			cur = Point{cur.X, cur.Y - distance}
		case 2:
			//Left
			cur = Point{cur.X - distance, cur.Y}
		case 3:
			//Up
			cur = Point{cur.X, cur.Y + distance}
		default:
			panic(fmt.Sprintf("Unknown direction: %d in line %s", direction, line))
		}
		vertices = append(vertices, cur)
	}

	//fmt.Println("perimeter:", perimeter)
	shoelaceArea := ShoelaceArea(vertices)
	fmt.Println("shoelaceArea:", shoelaceArea)

	//Pick's theorem  A = i + b/2 - 1
	//Rearrange to get i = A - b/2 + 1
	//A = shoelaceArea
	interiorArea := shoelaceArea - float64(perimeter)/2.0 + 1.0
	fmt.Println("interiorArea:", interiorArea)

	//We have the interior area, but we need to add the perimeter back on
	return int64(interiorArea) + int64(perimeter)
}

func ShoelaceArea(points []Point) float64 {
    area := 0.0
    j := len(points) - 1 // The last vertex is the 'previous' one to the first

    for i := 0; i < len(points); i++ {
        area += (float64(points[j].X) + float64(points[i].X)) * (float64(points[j].Y) - float64(points[i].Y))
        j = i // j is previous vertex to i
    }

    return math.Abs(area / 2)
}

// parses a line of the form "R 2 (#7807d2)"
func parseLine2(line string) (int, int) {
	parts := strings.Split(line, " ")
	distance := 0
	dir := 0
	fmt.Sscanf(parts[2], "(#%5x%1x)", &distance, &dir)
	return dir, distance
}

// parses a line of the form "R 2 (#7807d2)"
func parseLine1(line string) (int, int) {
	parts := strings.Split(line, " ")
	direction := parts[0]
	dir := 0
	switch direction {
	case "R":
		dir = 0
	case "L":
		dir = 2
	case "U":
		dir = 3
	case "D":
		dir = 1
	default:
		panic("Unknown direction")
	}
	distance, _ := strconv.Atoi(parts[1])
	return dir, distance
}
