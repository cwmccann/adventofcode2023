# Advent of Code 2023: Day 10: Pipe Maze

Visit the problem statement at [Advent of Code 2023 Day 10](https://adventofcode.com/2023/day/10).

## Problem Statement

We have a pipe maze defined by the characters `|-LJ7F` and we also have `S` for the starting position and `.` as the ground without a pipe.  The whole grid is densely packed with pipes so some of them are not part of the maze.

## Part 1

In part 1, we need to find the farthest point in the maze away from the starting paoint.

### Solution Overview

I solved part 1 by:
1. Parsing the whole grid into a 2D array of characters (runes in Go).
2. Find the starting position
3. Find the pipe loop
   1. I did a BFS search starting from the `S`
   2. At each position I checked the 4 cardinal directions and checked if it had a valid pipe connections or not.
   3. I also check to see if it was already in the pipe loop so it wasn't added twice.
   4. Keep going until we find the `S` again
   5. Return the path found.
4. Return the `(length of the pipe loop + 1) / 2`.  I added the `+1` since Go rounds down for integer division.

## Part 2

Part 2 was tough.  We need to find all the positions that where contained within the pipe maze.  Here's an example:
```
...........
.S-------7.
.|F-----7|.
.||OOOOO||.
.||OOOOO||.
.|L-7OF-J|.
.|II|O|II|.
.L--JOL--J.
.....O.....
```
In this case there are 4 points, the `I`s, contained in the pipe maze.  The rest of the points are either the pipes in the loop or outside.

The second wrinkle, and the part that made it hard was:
..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........
Even if the pipes are together you can slide between them.  So this case even though the there isn't a clear path to the middle, we still only have 4 points inside the maze.

### Solution Overview

In my first attempt, which I knew wouldn't work, was to flood fill the grid.  
  1. I found the pipe loop, 
  2. changed any pipe that was not in the loop to a `.`, 
  3. added an extra `.` around the whole outside of the grid
  4. Starting from one corner flood filled the whole maze
     1. For each character in the grid
        1. If the character was on edge, change it to a `O`
        2. If we are on a `.` and if the position has a neighbor that is an `O` change it to an `O`
     2. Keep doing this until we go through the maze without changing any characters
   
This got me pretty far but didn't handle the second wrinkle of the problem.

What I ended up doing was expanding the grid so that each grid location was a 2x2 block.  It's important to note that I used the top left character in each mapping and added any spaces from there.  That mapping needed to be consistent for it to work properly.

```
....................
....................
..F-------------7...
..|.............|...
..|.F---------7.|...
..|.|.........|.|...
..|.|.........|.|...
..|.|.........|.|...
..|.|.........|.|...
..|.|.........|.|...
..|.L---7.F---J.|...
..|.....|.|.....|...
..|.....|.|.....|...
..|.....|.|.....|...
..L-----J.L-----J...
....................
....................
....................
```
Also, note I replaced the `S` with the correct pipe piece.

I was then able to apply my flood fill algorithm from above.  The remaining step was to compress the grid back to normal size to get the correct number of 'contained' spaces.

## Challenges and Learnings

It was pretty tough for me.  I racked my brain to find a solution and ended up needing to look at the mega thread for a way to solve it.  I like this approach over the others since it was clear to me what was happening and I could use my flood fill code.  Another solution was to count the inside and out turns for each pipe but I found it confusing.

I also created a pretty print grid that used extended ascii characters to print the grid which made it nicer to look at and reminded me of the mazes in [Roguelike](https://en.wikipedia.org/wiki/Roguelike) games.

```
..........
.┌──────┐.
.│┌────┐│.
.││....││.
.││....││.
.│└─┐┌─┘│.
.│..││..│.
.└──┘└──┘.
..........
```
