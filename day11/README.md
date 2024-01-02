# Advent of Code 2023: Day 11: Cosmic Expansion

Visit the problem statement at [Advent of Code 2023 Day 11](https://adventofcode.com/2023/day/11).

## Problem Statement

Today we are given a map if the universe.  It has empty space `.` and galaxies `#`
```
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
```

We need to find the distance between each galaxy.  The trick is that if the row or column is empty it is expanding.

## Part 1

In Part 1, the universe expands by a factor of 2.

So the above example is actually:
```
....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......
```

I needed to find the distance between each galaxy and sum them up.

### Solution Overview

Finding the distance between two points in a grid, when you can only move in the 4 cardinal direction is given by the manhattan distance.  $d = |x_2 - x_1| + |y_2 - y_1|$

In part 1 I just expanded the map by the factor of 2 and calculated the manhattan distance for each galaxy.

## Part 2

Part 2, changes the expansion factor to `1000000`!
### Solution Overview

Obviously I couldn't just expand the map by a million so I need to find a different way.

If you look at the manhattan distance it's just the difference between the $x$ values plus the difference of the $y$ values.

So what I did was find all the rows and columns that needed to be expanded.
The I created an array of all the rows between the galaxy, if it wasn't an expanding row I put a `1` if it was I put the expansion factor.  I then summed up the array.

I did the same for the columns.

## Challenges and Learnings

This was a fun one.  It took a little leap to recall exactly what the manhattan distance was, but once that "clicked" it was pretty easy.  I went back and refactored my part one solution to use the same logic as part 2.
