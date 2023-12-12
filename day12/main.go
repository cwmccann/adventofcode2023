package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strings"
)

var DEBUG = false

func main() {
	text := utils.InputToString();

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}

func SolvePart1(input string) int {
	solution := Solve(input, false)
	fmt.Printf("Part2: %d\n", solution)
	return solution;
}

func SolvePart2(input string) int {
	solution := Solve(input, true)
	fmt.Printf("Part2: %d\n", solution)
	return solution;
}

func parseLine(line string) (string, []int) {
	if (strings.Trim(line, " ") == "") {
		panic("Empty line")
	}
	parts := strings.Split(line, " ")
	s := parts[0]
	groups := utils.StringToIntSlice2(parts[1], ",")
	return s, groups
}

func Solve(input string, part2 bool) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}
		s, groups := parseLine(line)

		if (part2) {
			newS := s
			newGroups := groups

			//Make it 5 times bigger
			for i := 1; i < 5; i++ {
				newS += "?" + s
				newGroups = append(newGroups, groups...)
			}

			// fmt.Printf("Springs: %s %v\n", newSprings, newGroups)
			s = newS
			groups = newGroups
		}

		cleanCache()
		sum += countArrangements(s, groups, NewState(s, 0, 0, 0))
	}
	return sum
}

type State struct {
	s string
	i, gi, current int
}
func NewState(s string, stringPosition, groupPosition, currentBlockLength int) State {
    return State{s, stringPosition, groupPosition, currentBlockLength}
}

var cache map[State]int = make(map[State]int)
func cleanCache() {
	cache = make(map[State]int)
}

func countArrangements(s string, groups []int, state State) int {
	i := state.i

	if val, ok := cache[state]; ok {
		return val
	}

	//Are we at the end
	if i == len(s) {
		//Processed all the groups
		if state.gi == len(groups) && state.current == 0 {
			if DEBUG {
				fmt.Printf("Found a solution1: %s\n", s)
			}
			return 1
		}
		//We matched the last group
		if state.gi == (len(groups) - 1) && groups[state.gi] == state.current {
			if DEBUG {
				 fmt.Printf("Found a solution2: %s\n", s)
			}
			return 1
		}

		if DEBUG {
			fmt.Printf("Not a solution2: %s\n", s)
		}
		return 0
	}

	sum := 0

	char := rune(s[i])

	if char == '.' {
		//if the char is a ., we are either ending a group or just proceeding
		if state.current == 0 {
			sum += countArrangements(s, groups, NewState(s[i+1:], i+1, state.gi, 0))
		} else if state.current > 0 && state.gi < len(groups) && state.current == groups[state.gi] {
			sum += countArrangements(s, groups, NewState(s[i+1:], i+1, state.gi+1, 0))
		}
	}

	if char == '#' {
		//if the char is a #, we are either starting a group or we add another item to the current count
		//it's the same state either way
		sum += countArrangements(s, groups, NewState(s[i+1:], i+1, state.gi, state.current+1))
	}

	if char == '?' {
		//Try replacing it with a .
		newS := utils.ReplaceAtIndex(s, '.', i)
		sum += countArrangements(newS, groups, NewState(newS[i:], i, state.gi, state.current))

		//Try replacing it with a #
		newS = utils.ReplaceAtIndex(s, '#', i)
		sum += countArrangements(newS, groups, NewState(newS[i:], i, state.gi, state.current))
	}

	cache[state] = sum
	//fmt.Printf("Cache: %d\n", len(cache))
	return sum
}



