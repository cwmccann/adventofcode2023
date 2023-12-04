package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		// handle the error here
		fmt.Println(err)
		return
	}

	// Convert []byte to string
	text := string(input)
	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}

func SolvePart1(input string) int {
	schematic := StringTo2DArray(input)
	parts := findParts(schematic)
	return SumSlice(parts)
}

func SolvePart2(input string) int {
	schematic := StringTo2DArray(input)
	gears := findGears(schematic)
	sum := 0
	for _, gear := range gears {
		sum += gear.GetGearRatio()
	}
	return sum
}

func StringTo2DArray(input string) [][]rune {
	lines := strings.Split(input, "\n")
	result := make([][]rune, 0) // Initialize with length 0

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) == 0 {
			continue // Skip blank lines
		}
		result = append(result, []rune(trimmedLine)) // Only append non-blank lines
	}

	return result
}

// findParts returns a slice of Part structs, each representing a part in the schematic
// a part is any number that has an adjacent symbol to it, including diagonals.
// a '.' is not a symbol
func findParts(schematic [][]rune) []int {
	parts := make([]int, 0)

	currentPart := ""
	adjacentSymbol := false

	for y, row := range schematic {
		for x, char := range row {
			//Do we have a digit?
			if unicode.IsDigit(char) {
				currentPart += string(char)
				//if we don't already have an adjacent symbol, check for one
				if !adjacentSymbol {
					adjacentSymbol = HasAdjacentSymbol(schematic, y, x, isSymbol).IsValid()
				}
			}
			//Are we at the end of the row or the end of a number?
			if x == len(schematic[y])-1 || !unicode.IsDigit(char) {
				//we're at the end of the number, so we need to check if we have a part
				if len(currentPart) > 0 && adjacentSymbol {
					partNumber, _ := strconv.Atoi(currentPart)
					//fmt.Printf("Found part: %d\n", partNumber)
					parts = append(parts, partNumber)
				} else if currentPart != "" {
					//fmt.Printf("Skipping part: %s\n", currentPart)
				}
				currentPart = ""
				adjacentSymbol = false
			}
		}
	}

	return parts
}


type Position struct {
	x, y int
}
func InvalidPosition() Position {
	return Position {
		x: -1,
		y: -1,
	}
}
func (p Position) IsValid() bool {
	return p.x >= 0 && p.y >= 0
}

type Gear struct {
	position Position
	connectedParts []int
}
func (g *Gear) AddConnectedPart(part int) {
    g.connectedParts = append(g.connectedParts, part)
}
func (g Gear) GetGearRatio() int {
    if len(g.connectedParts) == 2 {
		return g.connectedParts[0] * g.connectedParts[1]
	} else {
		return 0
	}
}

// findGears returns a slice of Gear structs, each representing a gear in the schematic
// a gear is a *
func findGears(schematic [][]rune) []Gear {

	gearMap := make(map[Position]Gear)

	currentPart := ""
	connectedToGear := false
	gearPosition := InvalidPosition()

	for y, row := range schematic {
		for x, char := range row {
			//Do we have a digit?
			if unicode.IsDigit(char) {
				currentPart += string(char)
				//if we don't already have an adjacent symbol, check for one
				if !connectedToGear {
					tmpGearPosition := HasAdjacentSymbol(schematic, y, x, isGear)
					if tmpGearPosition.IsValid() {
						gearPosition = tmpGearPosition
						connectedToGear = true
						//Add it to the map
						if _, ok := gearMap[gearPosition]; !ok {
							gearMap[gearPosition] = Gear{position: gearPosition, connectedParts: []int{}}
						}
					}
				}
			}
			//Are we at the end of the row or the end of a number
			if x == len(schematic[y])-1 || !unicode.IsDigit(char) {
				//we're at the end of the number, so we need to check if we are connected to a gear
				if len(currentPart) > 0 && connectedToGear {
					partNumber, _ := strconv.Atoi(currentPart)
					//fmt.Printf("Found part connected to a gear: %d\n", partNumber)
					gear := gearMap[gearPosition]
					gear.AddConnectedPart(partNumber)
					gearMap[gearPosition] = gear
				} else if currentPart != "" {
					//fmt.Printf("Skipping part: %s\n", currentPart)
				}
				currentPart = ""
				connectedToGear = false
				gearPosition = Position{0,0}
			}
		}
	}

	//Convert the map to a slice
	var gearList []Gear
	for _, gear := range gearMap {
    	gearList = append(gearList, gear)
	}
	return gearList
}

func HasAdjacentSymbol(grid [][]rune, y int, x int, symbolCheck func(char rune) bool) Position  {
	// Directions to move in the grid (up, down, left, right, and diagonals)
	dirs := []struct{ dx, dy int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{ 0, -1},          { 0, 1},
		{ 1, -1}, { 1, 0}, { 1, 1},
	}

	// Iterate through each direction
	for _, dir := range dirs {
		newX, newY := x+dir.dx, y+dir.dy

		// Check bounds
		if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) {
			//fmt.Printf("Checking %d,%d, char = %s\n", newX, newY, string(grid[newY][newX]));
			if symbolCheck(grid[newY][newX]) {
				return Position{newX, newY}
			}
		}
	}
	return InvalidPosition()
}

func isSymbol(char rune) bool {
	return !unicode.IsDigit(char) && char != '.'
}

func isGear(char rune) bool {
	return char == '*'
}

func SumSlice(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}
