package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strings"
)

func main() {
	text := utils.InputToString()

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}

func SolvePart1(input string) int {
	blocks := strings.Split(input, "\n\n")
	total := 0
	for _, block := range blocks {
		// fmt.Printf("block:\n %s\n", block)
		lines := strings.Split(block, "\n")
		lines = utils.TrimAndRemoveEmptyLines(lines)

		hReflect := findHReflect(lines, equals)
		// fmt.Printf("hReflect: %d\n", hReflect)

		rotated := rotate(lines)
		vReflect := findHReflect(rotated, equals)
		// fmt.Printf("vReflect: %d\n", vReflect)

		score := hReflect*100 + vReflect
		// fmt.Printf("%d\n", score)

		total += score
	}
	fmt.Println()
	return total
}

func findHReflect(lines []string, compare func(string, string) bool) int {
	for i := 1; i < len(lines); i++ {
		up := lines[:i]
		down := lines[i:]

		rows := utils.Min(len(up), len(down))

		up = utils.Reverse(up)
		up = up[:rows]
		down = down[:rows]

		upStr := strings.Join(up, ",")
		downStr := strings.Join(down, ",")

		if compare(upStr, downStr) {
			return i
		}
	}
	return 0
}

func rotate(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}

	result := make([]string, len(lines[0]))
	for i := range lines[0] {
		for j := range lines {
			result[i] += string(lines[j][i])
		}
	}

	return result
}

func offByOne(s1, s2 string) bool {
	for i := range s1 {
		if s1[i] != s2[i] {
			return equals(s1[i+1:], s2[i+1:])
		}
	}
	return false
}

func equals(s1, s2 string) bool {
	return s1 == s2
}

func SolvePart2(input string) int {
	blocks := strings.Split(input, "\n\n")
	total := 0
	for _, block := range blocks {
		// fmt.Printf("block:\n %s\n", block)
		lines := strings.Split(block, "\n")
		lines = utils.TrimAndRemoveEmptyLines(lines)

		hReflect := findHReflect(lines, offByOne)
		// fmt.Printf("hReflect: %d\n", hReflect)

		rotated := rotate(lines)
		vReflect := findHReflect(rotated, offByOne)
		// fmt.Printf("vReflect: %d\n", vReflect)

		score := hReflect*100 + vReflect

		// fmt.Printf("%d\n", score)

		total += score
	}
	return total
}
