package main

import (
	"fmt"
	"strings"
	"unicode"
	"strconv"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
        // handle the error here
        fmt.Println(err)
        return
    }

    // Convert []byte to string
    text := string(input)
	solution := SolvePart1(text);
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text);
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

        //remove the non-numeric characters
		str := RemoveNonNumeric(line)
		//convert the first and last number to an int
		num, err := strconv.Atoi(string(str[0]) + string(str[len(str)-1]))
		if err != nil {
			panic(err)
		}
		//fmt.Printf("line %s, num %d\n", line, num)
		sum += num
	}
    return sum
}

func RemoveNonNumeric(str string) string {
	newStr := ""
	for _, ch := range str {
		if unicode.IsDigit(ch) {
			newStr += string(ch)
		}
	}
	return newStr
}

func SolvePart2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		//Skip blank lines
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		//convert the string into the numbers
		str := convertToNumbers(line)
		//convert the first and last number to an int
		num, err := strconv.Atoi(string(str[0]) + string(str[len(str)-1]))
		if err != nil {
			panic(err)
		}
		//fmt.Printf("line %s, num %d\n", line, num)
		sum += num
	}
	return sum
}

func convertToNumbers(line string) string {
	numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	newStr := ""
	//Go through the string one character at a time
	for i, ch := range line {
		//Check if the char is a digit
		if unicode.IsDigit(ch) {
			newStr += string(ch)
			continue
		}

		//See if the string starts with one of the numbers
		for j, num := range numbers {
			if strings.HasPrefix(line[i:], num) {
				newStr += strconv.Itoa(j)
			}
		}
	}
	return newStr
}



