package main

import (
	"fmt"
	"adventofcode2023/utils"
)

func main() {
	text := utils.InputToString();

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}


func SolvePart1(input string) int {
	fmt.Println(Hello())
	return -1
}

func SolvePart2(input string) int {
	return -1
}
