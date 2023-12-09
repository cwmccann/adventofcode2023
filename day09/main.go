package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	text := utils.InputToString();

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}


func SolvePart1(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}
		nums := StringToIntSlice(line)
		sum += FindNext(nums)
	}

	return sum
}

func FindNext(nums []int) int {
	triangle := make([][]int, 1)
	triangle[0] = nums;

	curRow := 1
	done := false

	for !done {
		//Each row is one shorter than the previous
		triangle = append(triangle, make([]int, len(triangle[curRow-1]) - 1))

		for i := 0; i < len(triangle[curRow]); i++ {
			triangle[curRow][i] = triangle[curRow-1][i+1] - triangle[curRow-1][i]
		}
		if allZeros(triangle[curRow]) {
			done = true
		}
		curRow++
	}

	for i := len(triangle) - 1; i >= 0; i-- {
		if (i == len(triangle) - 1) {
			triangle[i] = append(triangle[i], 0)
			continue
		}

		bottom := triangle[i+1][len(triangle[i+1]) - 1]
		next := triangle[i][len(triangle[i]) - 1]
		triangle[i] = append(triangle[i], bottom + next)
	}

	// for i, row := range triangle {
	//  	fmt.Printf("Row %d: %v\n", i, row)
	// }

	return triangle[0][len(triangle[0]) - 1];
}

func allZeros(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func StringToIntSlice(input string) []int {
	tokens := strings.Split(input, " ")
	nums := make([]int, len(tokens))
	for i, token := range tokens {
		nums[i], _ = strconv.Atoi(token)
	}
	return nums
}

func SolvePart2(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		// fmt.Println("****************")
		if strings.Trim(line, " ") == "" {
			continue
		}
		nums := StringToIntSlice(line)
		sum += FindPrev(nums)
	}

	return sum
}


func FindPrev(nums []int) int {
	triangle := make([][]int, 1)
	triangle[0] = nums;

	curRow := 1
	done := false

	for !done {
		//Each row is one shorter than the previous
		triangle = append(triangle, make([]int, len(triangle[curRow-1]) - 1))

		for i := 0; i < len(triangle[curRow]); i++ {
			triangle[curRow][i] = triangle[curRow-1][i+1] - triangle[curRow-1][i]
		}
		if allZeros(triangle[curRow]) {
			done = true
		}
		curRow++
	}

	for i := len(triangle) - 1; i >= 0; i-- {
		if (i == len(triangle) - 1) {
			triangle[i] = prepend(triangle[i], 0)
			continue
		}

		bottom := triangle[i+1][0]
		next := triangle[i][0]
		triangle[i] = prepend(triangle[i], next - bottom)
	}

	// for i, row := range triangle {
	//  	fmt.Printf("Row %d: %v\n", i, row)
	// }

	return triangle[0][0];
}

func prepend(nums []int, num int) []int {
	return append([]int{num}, nums...)
}

