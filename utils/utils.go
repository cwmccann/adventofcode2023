package utils

import (
	"log"
	"os"
)

func InputToString() string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}

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
