# Advent of Code 2023: Day 6: Camel Cards

Visit the problem statement at [Advent of Code 2023 Day 6](https://adventofcode.com/2023/day/6).

## Problem Statement

Today, we're playing Camel Cards, a simplified version of poker that's easier to play on a camel. We're given a hand and a bid. The winnings are calculated as the rank of the hand multiplied by the bid.

## Part 1

In Part 1, we need to calculate the total winnings based on the rank * bid for all the hands we're given.

### Solution Overview

I parsed the cards for each hand into a map, counting the frequency of each card. While doing so, I also tracked the maximum frequency of the cards in the hand. Then, I determined the rank of the hand based on the number of distinct cards and the maximum frequency.

## Part 2

In Part 2, the Jack is changed into the Joker card (J) as a wild card.

### Solution Overview

I counted the number of Jokers in the hand and increased the maximum frequency by that amount. I also removed the Jokers from the frequency map to avoid double counting. Additionally, I had to adjust the order of the high card mapping. I did this by creating arrays of the cards and using their index in the array to determine their value.
