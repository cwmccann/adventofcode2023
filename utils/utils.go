package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// InputToString reads the contents of input.txt and returns it as a string
func InputToString() string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}

// StringToIntSlice converts a string of space separated ints to a slice of ints
func StringToIntSlice(input string) []int {
	tokens := strings.Split(input, " ")
	nums := make([]int, len(tokens))
	for i, token := range tokens {
		nums[i], _ = strconv.Atoi(token)
	}
	return nums
}

// MinInSlice returns the minimum value in a slice of ints
func MinInSlice(slice []int) int {
    if len(slice) == 0 {
        log.Fatal("minInSlice called with empty slice")
    }
    min := slice[0]
    for _, v := range slice {
        if v < min {
            min = v
        }
    }
    return min
}

// Prepend returns a new slice with num prepended to nums
func Prepend(nums []int, num int) []int {
	return append([]int{num}, nums...)
}

// Zip returns a channel that will yield the elements of a and b interleaved
func Zip(a, b []int) (<-chan [2]int) {
    minLen := len(a)
    if len(b) < minLen {
        minLen = len(b)
    }

    ch := make(chan [2]int, minLen)

    go func() {
        for i := 0; i < minLen; i++ {
            ch <- [2]int{a[i], b[i]}
        }
        close(ch)
    }()

    return ch
}


// StringsToRune2D converts a slice of strings to a 2D slice of runes
func StringsToRune2D(lines []string) [][]rune {
	rune2D := make([][]rune, 0)
	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}
		rune2D = append(rune2D, []rune(line))
	}
	return rune2D
}

// RemoveAll removes all instances of charsToRemove from s
func RemoveAll(s string, charsToRemove string) string {
	for _, r := range charsToRemove {
        s = strings.ReplaceAll(s, string(r), "")
    }
    return s
}
