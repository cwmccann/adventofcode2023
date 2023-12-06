package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"regexp"
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	text := string(input)
	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}


func SolvePart1(input string) int {
	times := getField(input, "Time")
	distances := getField(input, "Distance")

	fmt.Printf("Times: %v\n", times)
	fmt.Printf("Distances: %v\n", distances)

	total := 1

	for i := 0; i < len(times); i++ {
		total *= solve(times[i], distances[i])
		fmt.Println()
	}

	return total
}

func solve(time int, targetDist int) int {
	minTime := 1
	maxTime := time - 1

	for holdTime := minTime; holdTime <= maxTime; holdTime++ {
		dist := holdTime *(time - holdTime)
		if dist > targetDist {
			minTime = holdTime
			break
		}
	}

	for holdTime := maxTime; holdTime >= minTime; holdTime-- {
		dist := holdTime *(time - holdTime)
		if dist > targetDist {
			maxTime = holdTime
			break
		}
	}

	return maxTime - minTime + 1
}

func getField(input string, field string) []int {
	regexp := regexp.MustCompile(field + `:\s+([0-9\s]+)`)
	matches := regexp.FindStringSubmatch(input)

	fields := strings.Fields(matches[1])
    times := make([]int, len(fields))


    for i, f := range fields {
        time, _ := strconv.Atoi(f)
        times[i] = time
    }

    return times
}

func SolvePart2(input string) int {
	lines := strings.Split(input, "\n")
	time := 0
	distance := 0
	for _, line := range lines {
		if strings.HasPrefix(line, "Time:") {
			line = strings.TrimPrefix(line, "Time:")
			line = strings.ReplaceAll(line, " ", "")
			time, _ = strconv.Atoi(line)
		}
		if strings.HasPrefix(line, "Distance:") {
			line = strings.TrimPrefix(line, "Distance:")
			line = strings.ReplaceAll(line, " ", "")
			distance, _ = strconv.Atoi(line)
		}
	}

	return solve(time, distance)
}
