package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	blue  int
	red   int
	green int
}

func MakeRound(str string) Round {
	balls := strings.Split(str, ",")
	blue, red, green := 0, 0, 0

	for _, ball := range balls {
		parts := strings.Split(strings.TrimSpace(ball), " ")
		if len(parts) == 2 {
			num, _ := strconv.Atoi(parts[0])
			color := parts[1]

			switch color {
			case "green":
				green = num
			case "red":
				red = num
			case "blue":
				blue = num
			}
		}
	}
	return Round{blue, red, green}
}
func IsRoundValid(round Round) bool {
	return round.red <= 12 && round.green <= 13 && round.blue <= 14
}

type Game struct {
	id     int
	rounds []Round
}

func MakeGame(line string) Game {
	//Remove the game at the front of the line
	str := strings.Replace(line, "Game", "", -1)

	parts := strings.Split(str, ":")
	id, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	rounds := []Round{}

	for _, round := range strings.Split(parts[1], ";") {
		rounds = append(rounds, MakeRound(round))
	}
	return Game{id, rounds}
}
func IsGameValid(game Game) bool {
	for _, round := range game.rounds {
		if !IsRoundValid(round) {
			return false
		}
	}
	return true
}
func (game Game) CalculatePower() int {
	//Find the min number of balls for each color
	blue, red, green := 0, 0, 0

	for _, round := range game.rounds {
		if round.blue > blue {
			blue = round.blue
		}
		if round.red > red {
			red = round.red
		}
		if round.green > green {
			green = round.green
		}
	}
	//The power is the product of the min number of balls for each color
	return blue * red * green
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		// handle the error here
		fmt.Println(err)
		return
	}

	// Convert []byte to string
	text := string(input)
	solution := SolvePart1(text)
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

		game := MakeGame(line)
		if IsGameValid(game) {
			sum += game.id
		}
	}
	return sum
}

func SolvePart2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		//Skip blank lines
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		game := MakeGame(line)
		sum += game.CalculatePower()
	}
	return sum
}
