# Advent of Code 2023: Day 9: XXXXMirage Maintenance

Visit the problem statement at [Advent of Code 2023 Day 9](https://adventofcode.com/2023/day/9).

## Problem Statement

The problem involves a sequence of numbers and a unique way of generating the next number in the sequence. The sequence is represented as a triangle, and the next number is calculated by adding a column from the bottom up. Here's an example:

For the sequence  `0 3 6 9 12 15`

We expand it to get:
```
0   3   6   9  12  15
  3   3   3   3   3
    0   0   0   0
```
Then to get the next number in the sequence we start at the bottom and add a column going up:L
```
0   3   6   9  12  15   B
  3   3   3   3   3   A
    0   0   0   0   0
```
`A: 3 - A = 0` ->   `A = 3`

`B: 15 - B = A` ->  `A = 3` so `B = 18`

I've seen this, or something like it, somewhere else but I can't find the reference.  It looks like the method of differences but it's not.

## Part 1

In Part1 we need to find the next number in each sequence and add up the numbers.

### Solution Overview

I initially solved this problem iteratively, building the triangle from the top down, then adding a column from the bottom up. This approach worked. Later, I revisited the problem and implemented a recursive solution, which made the problem more interesting.

## Part 2

In Part 2, we need to find the previous number in the sequence. This requires a slight modification to the logic used in Part 1, where we now look at the first numbers instead of the last.


## Challenges and Learnings

Writing recursive code is always a bit of a brain twister for me, as I don't often get to write such code. However, it's good practice and always a fun challenge.
