package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)


type ScratchCard struct {
	instanceCount int
	winningNumbers []int
	cardNumbers []int
	matchCount int
}

func NewScratchCard(line string) *ScratchCard {
	instanceCount := 1
	winningNumbers, cardNumbers := parseLine(line)
	matchCount := len(intersect(winningNumbers, cardNumbers))

	return &ScratchCard{instanceCount, winningNumbers, cardNumbers, matchCount}
}

func parseLine(line string) ([]int, []int) {
	//Discard the Card x: part
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		log.Fatal("Invalid input")
	}
	line = strings.TrimSpace(parts[1])

	//Split on the | to get the winning numbers and card numbers
	parts = strings.Split(line, "|")
	if len(parts) != 2 {
		log.Fatal("Invalid input")
	}
	winningNumbers := parseNumbers(parts[0])
	cardNumbers := parseNumbers(parts[1])

	return winningNumbers, cardNumbers
}

func parseNumbers(line string) []int {
	parts := strings.Fields(line)
	numbers := make([]int, len(parts))
	for i, part := range parts {
		numbers[i], _ = strconv.Atoi(part)
	}
	return numbers
}


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
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		//Skip blank lines
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		scratchCard := NewScratchCard(line)
		cardScore := int(math.Pow(2, float64(scratchCard.matchCount - 1)))
		sum += cardScore
	}
	return sum
}



func SolvePart2(input string) int {
	lines := strings.Split(input, "\n")
	scratchcards := make([]*ScratchCard, 0)

	for _, line := range lines {
		//Skip blank lines
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		scratchcards = append(scratchcards, NewScratchCard(line))
	}

	totalCardCount := 0
	for i, scratchcard := range scratchcards {
		//Add the instance count to the total card count
		totalCardCount += scratchcard.instanceCount

		j := scratchcard.matchCount
		for j > 0  && i + j < len(scratchcards) {
			scratchcards[i + j].instanceCount += scratchcard.instanceCount
			j = j - 1
		}
	}
	return totalCardCount
}

func intersect(slice1, slice2 []int) []int {
    m := make(map[int]bool)
    for _, item := range slice1 {
        m[item] = true
    }

    var result []int
    for _, item := range slice2 {
        if _, ok := m[item]; ok {
            result = append(result, item)
        }
    }
    return result
}
