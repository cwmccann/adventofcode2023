# Advent of Code 2023: Day 1: Trebuchet?!

Visit the problem statement at [Advent of Code 2023 Day 1](https://adventofcode.com/2023/day/1).

## Problem Statement

The task is straightforward: find the first and last number in each line, combine them to form a two-digit number, and then sum all these numbers.

## Part 1
### Solution Overview

This is my first Go program. I had to learn how to perform simple tasks like reading a file, splitting it into lines, and filtering out empty lines. Although I'm used to a functional approach, it's not idiomatic in Go.

Here's my solution:
1. Read the file.
2. Split it into lines, filtering out the empty ones.
3. Remove all non-numeric characters.
4. Get the first and last digit, combine them into a string, and convert it to an integer.
5. Sum up the numbers.

## Part 2

Part 2 introduced a twist: some numbers are written out. So, we now have digits and written-out numbers.

### Solution Overview

The solution for Part 2 is similar to Part 1. I iterated through the line and created a new string. If the character was a digit, I added it. If the string started with one of the written-out digits, I added the index of that digit to the new string.

## Challenges and Learnings

The main challenge was learning how to do things in Go. The puzzle input had a gotcha: some of the written-out numbers shared letters, like `oneight`. Since I looked at each part of the string and didn't replace the values in the original string, I avoided this issue.
