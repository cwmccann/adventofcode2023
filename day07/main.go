package main

import (
	"adventofcode2023/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"sort"
)

const (
	part1CardOrdering = "23456789TJQKA"
	cardOrderingPart2 = "J23456789TQKA"
)

func main() {
	text := utils.InputToString();

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}


func SolvePart1(input string) int {
	hands := ParseInput(input, CalculateHandValuePart1)
	sort.Slice(hands, func(i, j int) bool {
		return CompareHands(hands[i], hands[j], part1CardOrdering) < 0
	})

	//fmt.Println(hands)
	total:= 0
	for i, hand := range hands {
		total += hand.Bid * (i+1)
	}
	//251106089
	return total
}

func CalculateHandValuePart1(cards string) int {
	freq := make(map[rune]int)
    for _, card := range cards {
        freq[card]++
    }
	numItems := len(freq)
	_, maxFreq := getMaxFreq(freq)

	switch numItems {
	case 1:
		return FiveOfAKind
	case 2:
		if maxFreq == 4 {
			return FourOfAKind
		} else {
			return FullHouse
		}
	case 3:
		if maxFreq == 3 {
			return ThreeOfAKind
		} else {
			return TwoPair
		}
	case 4:
		return Pair
	case 5:
		return HighCard
	}
	return -1
}


func SolvePart2(input string) int {
	hands := ParseInput(input, CalculateHandValuePart2)
	sort.Slice(hands, func(i, j int) bool {
		return CompareHands(hands[i], hands[j], cardOrderingPart2) < 0
	})

	//fmt.Println(hands)
	total:= 0
	for i, hand := range hands {
		total += hand.Bid * (i+1)
	}
	//249620106
	return total
}
func CalculateHandValuePart2(cards string) int {
	freq := make(map[rune]int)
    for _, card := range cards {
        freq[card]++
    }

	//Get the count of the jokers and remove them from the map
	jokerCount := freq['J']
	if jokerCount > 0 {
		delete(freq, 'J')
	}
	//Find the card with the highest frequency and add the jokers to it
	maxFreqCard, _ := getMaxFreq(freq)
	freq[maxFreqCard] += jokerCount

	numItems := len(freq)
	_, maxFreq := getMaxFreq(freq)

	switch numItems {
	case 1:
		return FiveOfAKind
	case 2:
		if maxFreq == 4 {
			return FourOfAKind
		} else {
			return FullHouse
		}
	case 3:
		if maxFreq == 3 {
			return ThreeOfAKind
		} else {
			return TwoPair
		}
	case 4:
		return Pair
	case 5:
		return HighCard
	}
	return -1
}



func ParseInput(input string, valueCalculator func(cards string) int) []Hand {
	hands := []Hand{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		hands = append(hands, ParseHand(line, valueCalculator))
	}
	return hands
}

type Hand struct {
	Cards string
	Bid int
	HandValue int
}

func ParseHand(line string, valueCalculator func(cards string) int) Hand {
	parts := strings.Split(strings.TrimSpace(line), " ")
	cards := parts[0]
	bid, _ := strconv.Atoi(parts[1])
	hand := Hand{Cards: cards, Bid: bid, HandValue: valueCalculator(cards)}
	return hand
}

const (
	HighCard = 0
	Pair = 1
	TwoPair = 2
	ThreeOfAKind = 3
	FullHouse = 4
	FourOfAKind = 5
	FiveOfAKind = 6
)

func getMaxFreq(freq map[rune]int) (rune, int) {
	var maxFreq int
	var maxFreqCard rune

    for card, freq := range freq {
        if freq > maxFreq {
            maxFreq = freq
			maxFreqCard = card
        }
    }
    return maxFreqCard, maxFreq
}

func CompareHands(h1, h2 Hand, cardOrdering string) int {
    if h1.HandValue > h2.HandValue {
        return 1
    } else if h1.HandValue < h2.HandValue {
        return -1
    } else {
        //They have the same hand value, so compare the cards
		for i := 0; i < len(h1.Cards); i++ {
			c1 := string(h1.Cards[i])
			c2 := string(h2.Cards[i])
			compare := CompareCards(c1, c2, cardOrdering)
			if compare != 0 {
				return compare
			}
		}
		//They are the same!
		return 0
    }
}

func CompareCards(c1, c2 string, cardOrdering string) int {
	c1Index := strings.Index(cardOrdering, c1)
	c2Index := strings.Index(cardOrdering, c2)

	if c1Index > c2Index {
		return 1
	} else if c1Index < c2Index {
		return -1
	} else {
		return 0
	}
}
