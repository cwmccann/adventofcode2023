# Advent of Code 2023: Day 2: Cube Conundrum

Visit the problem statement at [Advent of Code 2023 Day 2](https://adventofcode.com/2023/day/2).

## Problem Statement

In this challenge, we're playing a game with an elf. The game involves drawing cubes of three different colors - red, green, and blue - from a bag in each round.

## Part 1

In Part 1, we need to determine if the game is valid when we have only 12 red cubes, 13 green cubes, and 14 blue cubes.

### Solution Overview

The solution involves parsing each round of each game and checking their validity. The total number of valid games is then counted.

## Part 2

In Part 2, we need to find the fewest cubes of each color that we can play with, multiply these values together for each game, and then sum the results.

### Solution Overview

For each game, I determined the maximum number of red, blue, and green cubes required for all the rounds. I then multiplied these values together to get the result for each game.

## Challenges and Learnings

This being one of the early problems, I believe I may have over-engineered the solution a bit. However, it was a valuable exercise in learning more about how go works.
