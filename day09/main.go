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
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}
		nums := utils.StringToIntSlice(line)
		sum += FindNext2(nums)
	}

	return sum
}

func FindNext(nums []int) int {
	triangle := make([][]int, 1)
	triangle[0] = nums

	curRow := 1
	done := false

	for !done {
		//Each row is one shorter than the previous
		triangle = append(triangle, make([]int, len(triangle[curRow-1])-1))

		for i := 0; i < len(triangle[curRow]); i++ {
			triangle[curRow][i] = triangle[curRow-1][i+1] - triangle[curRow-1][i]
		}
		if allZeros(triangle[curRow]) {
			done = true
		}
		curRow++
	}

	for i := len(triangle) - 1; i >= 0; i-- {
		if i == len(triangle)-1 {
			triangle[i] = append(triangle[i], 0)
			continue
		}

		bottom := triangle[i+1][len(triangle[i+1])-1]
		next := triangle[i][len(triangle[i])-1]
		triangle[i] = append(triangle[i], bottom+next)
	}

	// for i, row := range triangle {
	//  	fmt.Printf("Row %d: %v\n", i, row)
	// }

	return triangle[0][len(triangle[0])-1]
}

func allZeros(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func SolvePart2(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		// fmt.Println("****************")
		if strings.Trim(line, " ") == "" {
			continue
		}
		nums := utils.StringToIntSlice(line)
		sum += FindPrev2(nums)
	}

	return sum
}

func FindPrev(nums []int) int {
	triangle := make([][]int, 1)
	triangle[0] = nums

	curRow := 1
	done := false

	for !done {
		//Each row is one shorter than the previous
		triangle = append(triangle, make([]int, len(triangle[curRow-1])-1))

		for i := 0; i < len(triangle[curRow]); i++ {
			triangle[curRow][i] = triangle[curRow-1][i+1] - triangle[curRow-1][i]
		}
		if allZeros(triangle[curRow]) {
			done = true
		}
		curRow++
	}

	for i := len(triangle) - 1; i >= 0; i-- {
		if i == len(triangle)-1 {
			triangle[i] = utils.Prepend(triangle[i], 0)
			continue
		}

		bottom := triangle[i+1][0]
		next := triangle[i][0]
		triangle[i] = utils.Prepend(triangle[i], next-bottom)
	}

	// for i, row := range triangle {
	//  	fmt.Printf("Row %d: %v\n", i, row)
	// }

	return triangle[0][0]
}

func FindNext2(nums []int) int {
	if allZeros(nums) {
		return 0
	}
	deltas := make([]int, len(nums)-1)
	zipped := utils.Zip(nums, nums[1:])
	i := 0
	for pair := range zipped {
		deltas[i] = pair[1] - pair[0]
		i++
	}
	//fmt.Printf("Deltas: %v\n", deltas)

	diff := FindNext2(deltas)
	return nums[len(nums)-1] + diff
}

func FindPrev2(nums []int) int {
	if allZeros(nums) {
		return 0
	}
	deltas := make([]int, len(nums)-1)
	zipped := utils.Zip(nums, nums[1:])
	i := 0
	for pair := range zipped {
		deltas[i] = pair[1] - pair[0]
		i++
	}
	//fmt.Printf("Deltas: %v\n", deltas)

	diff := FindPrev2(deltas)
	return nums[0] - diff
}

