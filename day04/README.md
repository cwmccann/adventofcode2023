# Advent of Code 2023: Day 4: Scratchcards

Visit the problem statement at [Advent of Code 2023 Day 4](https://adventofcode.com/2023/day/4).

## Problem Statement

An elf has a stack of scratch cards and we need to find out what he won!

## Part 1

In Part 1, we have a list of cards.  On each card are two sets of numbers; the winning numbers and the card numbers.  To win you have to match the cards numbers to the winning numbers.  The amount you win starts with with one point and doubles for each number past it.

### Solution Overview

I parsed the cards into a `ScratchCard` struct that has the following
```go
type ScratchCard struct {
	winningNumbers []int
	cardNumbers []int
	matchCount int
}
```
When creating the ScratchCard, I performed an intersection between the two lists of numbers and counted the elements.

To get the score of the card, I calculated $2^{matchCount - 1}$. I then summed up the scores for all the cards.

## Part 2

Part 2 adds a twist: instead of winning points, you win extra cards. Specifically, if a card has a certain number of matching numbers, you win one copy of each of the cards below it, up to the number of matches.


### Solution Overview

I added a field to the ScratchCard struct, instanceCount, and set its default value to 1. After parsing all the cards, I looped through them to calculate how many additional cards were won. The instanceCount field represents the number of instances I have for the current card, and I incremented the instanceCount of future cards by this amount.

## Challenges and Learnings
This problem provided more insights into Go, particularly in managing instance counts and incrementing them. It was a fun and enlightening challenge.
