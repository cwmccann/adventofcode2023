# Advent of Code 2023: Day 14: Parabolic Reflector Dish

Visit the problem statement at [Advent of Code 2023 Day 14](https://adventofcode.com/2023/day/14).

## Problem Statement

The problem involves a platform that can tilt in the four cardinal directions. When the platform tilts, rounded rocks (`O`) roll to the edge being tilted until they hit a stationary rock (`#`) or an edge. The rocks will not roll off the platform. An empty square is represented by (`.`).

## Part 1

In Part 1, we need to calculate the load on the north supporter. The load is calculated by counting the number of rounded rocks on the board and their positions. The closer a rock is to the north edge, the more it contributes to the load.


```
OOOO.#.O.. 10
OO..#....#  9
OO..O##..O  8
O..#.OO...  7
........#.  6
..#....#.#  5
..O..#.O.O  4
..O.......  3
#....###..  2
#....#....  1
```
So for a part 1, we need to tilt to the north and calculate the load on the northern support.

### Solution Overview

After reading in the grid, I tilt the grid to the north. To do this, I iterate through each character starting from the top left. If I find a rounded rock (`O`) and there is a free space above it, the rock rolls. I repeat this process until there are no changes.

Then, I iterate through the grid and find all the rocks, assigning them a load value based on their final position.

## Part 2

Part 2 introduces a new challenge. We need to rotate the grid in all four cardinal directions (North, West, South, and East) and repeat this process a billion times!

### Solution Overview

To solve this, I first updated the logic to allow the grid to rotate in all four directions. This was achieved by passing in a delta (x, y) to guide the movement of the rocks.

The next challenge was to handle a billion rotations. Brute force was not an option due to the sheer number of iterations. I hypothesized that there must be a repeating pattern after a certain number of rotations. After analyzing the data, I indeed found a cycle.

To detect this cycle, I hashed the state of the grid after each rotation and checked for collisions. Once a collision was detected, I could determine the length of the cycle and when it started. This allowed me to skip ahead by the cycle length and finish off the remaining steps.

## Challenges and Learnings

This problem was quite challenging. Detecting cycles was a difficult task, and it took me some time to come up with the idea of hashing the grid. However, once I implemented this approach, it worked excellently and significantly reduced the computation time.
