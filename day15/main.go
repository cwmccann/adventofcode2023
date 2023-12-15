package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	text := utils.InputToString();

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}

func Hash(input string) int {
	hash := 0

	for _, char := range input {
		hash += int(char)
		hash *= 17
		hash %= 256
	}
	return hash
}

func SolvePart1(input string) int {
	total := 0
	input = strings.Replace(input, "\n", "", -1)
	for _, line := range strings.Split(input, ",") {
		total += Hash(line)
	}
	return total
}

type Lens struct {
	label string
	focalLength int
}


func SolvePart2(input string) int {
    boxes := make(map[int][]Lens)
	input = strings.Replace(input, "\n", "", -1)
	for _, line := range strings.Split(input, ",") {
		lens, op := parseLens(line)
		hash := Hash(lens.label)
		if op == '=' {
			addToBox(boxes, hash, lens)
		} else {
			removeFromBox(boxes, hash, lens)
		}
	}

	total := 0
	for hash, box := range boxes {
		boxTotal := 0
		for i, lens := range box {
			boxTotal += (hash + 1) * (i + 1) * lens.focalLength
		}
		total += boxTotal
	}

	return total
}

func printBoxes(boxes map[int][]Lens) {
	for hash, box := range boxes {
		if (len(box) == 0) {
			continue
		}
		fmt.Printf("Box %d: ", hash)
		for _, lens := range box {
			fmt.Printf("[%s %d] ", lens.label, lens.focalLength)
		}
		fmt.Printf("\n")
	}
}

func addToBox(boxes map[int][]Lens, hash int, lens Lens) {
	if _, ok := boxes[hash]; !ok {
		boxes[hash] = make([]Lens, 0)
	}
	box := boxes[hash]

	for i, l := range box {
		if l.label == lens.label {
			box[i] = lens
			return
		}
	}
	boxes[hash] = append(boxes[hash], lens)
}

func removeFromBox(boxes map[int][]Lens, hash int, lens Lens) {
	if _, ok := boxes[hash]; !ok {
		return
	}
	box := boxes[hash]

	for i, l := range box {
		if l.label == lens.label {
			boxes[hash] = append(box[:i], box[i+1:]...)
			return
		}
	}
}

func parseLens(input string) (Lens, rune) {
	label := ""
	focalLengthStr := ""
	operation := 'a'

	for _, char := range input {
		if unicode.IsLetter(char) {
            label += string(char)
        }
        if unicode.IsDigit(char) {
            focalLengthStr += string(char)
        }
		if char == '=' || char == '-' {
			operation = char
		}
	}

	focalLength, _ := strconv.Atoi(focalLengthStr)
	return Lens{label, focalLength}, operation

}
