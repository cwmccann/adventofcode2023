package main

import (
	"adventofcode2023/utils"
	"bufio"
	"fmt"
	"strings"
	"math/big"
)

func main() {
	text := utils.InputToString();

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	part2Solution := SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", part2Solution)
}

type Path struct {
	L, R string
}

func parseInput(input string) (string, map[string]Path) {
	path := ""
	graph := make(map[string]Path)

	scanner := bufio.NewScanner(strings.NewReader(input))
    for scanner.Scan() {
        line := scanner.Text()
		if strings.Trim(line, " ") == "" {
			continue
		}
		if path == "" {
			path = line
			continue
		}
		parts := strings.Split(line, " = ")
		key := parts[0]
		paths := strings.Split(strings.Trim(parts[1], "()"), ", ")
		l := paths[0]
		r := paths[1]
		graph[key] = Path{L: l, R: r}
    }
	return path, graph
}

func SolvePart1(input string) int {
	path, graph := parseInput(input)

	current := "AAA"
	last := "ZZZ"

	steps := 0;
	pathIndex := 0;
	for current != last {
		steps++

		c := path[pathIndex]
		if c == 'L' {
			current = graph[current].L
		} else if c == 'R' {
			current = graph[current].R
		} else {
			panic("Invalid path character")
		}

		pathIndex++
		if pathIndex >= len(path) {
			pathIndex = 0
		}
	}
	return steps
}

func SolvePart2(input string) int64 {
	path, graph := parseInput(input)
	starts := make([]string, 0)

	for k := range graph {
		if strings.HasSuffix(k, "A") {
			starts = append(starts, k)
		}
	}

	stepCounts := make([]int64, len(starts))

	for i := range starts {
		current := starts[i]
		steps := 0
		pathIndex := 0

		for !strings.HasSuffix(current, "Z") {
			steps++

			c := path[pathIndex]
			if c == 'L' {
				current = graph[current].L
			} else if c == 'R' {
				current = graph[current].R
			} else {
				panic("Invalid path character")
			}

			pathIndex++
			if pathIndex >= len(path) {
				pathIndex = 0
			}
		}
		//fmt.Printf("Found path for %s in %d steps\n", starts[i], steps)
		stepCounts[i] = int64(steps)
	}
	lcm := LCMOfSlice(stepCounts)
	return lcm

}

func GCD(a, b int64) int64 {
    return new(big.Int).GCD(nil, nil, big.NewInt(a), big.NewInt(b)).Int64()
}

func LCM(a, b int64) int64 {
    return a / GCD(a, b) * b
}

func LCMOfSlice(numbers []int64) int64 {
    result := numbers[0]
    for _, number := range numbers[1:] {
        result = LCM(result, number)
    }
    return result
}
