# Advent of Code 2023: Day 13: Point of Incidence


Visit the problem statement at [Advent of Code 2023 Day 13](https://adventofcode.com/2023/day/13).

## Problem Statement

The problem involves finding a line of reflection in a grid. The line can be either vertical or horizontal and does not necessarily need to be in the center of the grid. The goal is to find a line such that the parts of the grid on either side of the line are mirror images of each other.

## Part 1

In Part 1, we need to find the line of reflection for each grid in our puzzle input. If the line is horizontal, we multiply the reflection point by 100; if it's vertical, we simply report the result.  We need to do this for each grid in the input, and sum the results.

### Solution Overview

I solved this problem by first parsing the puzzle input into individual grids. For each grid, I checked for a horizontal line of reflection by iterating through each row and performing the following steps:

1. Split the grid into two parts, above and below the current row.
2. Reverse the order of the rows in the upper part to match the reflection with the lower part row by row.
3. Truncate the larger part to be the same size as the smaller part.
4. Compare the two parts. If they are equal, we have found a line of reflection.

To check for a vertical line of reflection, I rotated the grid by 90 degrees and repeated the same process.

## Part 2

In Part 2, we need to find a line of reflection even if one character is off. This adds an extra layer of complexity to the problem.

### Solution Overview

To handle this, I refactored the find HRReflect method to taker a comparison function.   Instead of checking for strict equality between the two parts of the grid, I wrote a function that considers two strings equal if exactly one character is different. This function was then used in the same process as in Part 1 to find the line of reflection.

## Challenges and Learnings

This one was tricky to get right.  I tried detecting the reflection point in a bunch of ways and I'm happy with the final result but it took a while to get there.  Part two was pretty easy in comparison. 
