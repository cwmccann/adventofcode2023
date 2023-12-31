# Advent of Code 2023: Day 3: Gear Ratios

Visit the problem statement at [Advent of Code 2023 Day N](https://adventofcode.com/2023/day/3).

## Problem Statement

In this challenge, we are given a grid containing numbers and symbols. The goal is to identify all the "part numbers", which are numbers adjacent to a symbol.

```
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
```

## Part 1

In Part 1, we need to find all the part numbers and sum them up.

### Solution Overview

To solve this, I looped through each character in the grid. If I encountered a digit, I stored it in the current part and checked if there was an adjacent symbol. If there was, I set a flag. When I reached the end of a number (either a dot or end of line), if the adjacent symbol flag was true, I kept the part number. Finally, I summed up all the part numbers.

## Part 2

In Part 2, we need to find gears. Gears are represented by `*` with exactly two numbers adjacent to it.

### Solution Overview

I took a similar approach as in Part 1, except I looked for `*` instead of any symbol. When I found a `*`, I stored the position as a key in a map and the value being the list of numbers.

At the end, I looked for all the gears that had 2 numbers, multiplied the 2 numbers attached to them, and added up the results.

## Challenges and Learnings

This problem helped me learn more about maps and structs in Go. It also got me thinking about grids, which is beneficial since there are more grid-based problems to come.
