# Advent of Code 2023: Day 16: The Floor Will Be Lava

Visit the problem statement at [Advent of Code 2023 Day 16](https://adventofcode.com/2023/day/16).

## Problem Statement

We have another grid problem.  Today we have a beam of light entering a grid containing empty space `.`, mirrors `/` and `\`, and splitters `|` and `-`.

- If the beam hits a mirror it gets reflected by 90 degrees depending on the angle of the mirror.
- If it hits a splitter on the pointer side nothing happens, it just passes through.
- If it hits a splitter on the flat side, it splits into two beams

## Part 1

In part 1, the beam enters on the top left then travels through the grid.  We need to figure out how many squares get energized by it travel.

### Solution Overview

I created a Beam object that has it's current positions (x,y), direction, and the path it has taken so far.

I added a FIFO queue to process the beams.  We start with one beam coming from the top left moving to the right.  It's important to set the initial beam to have a starting position off the grid since in the actual input the first `grid[0][0] = \`

The queue works like this:
1. Update the beams position based on the direction
2. Is the beam off the grid?  If so we are done with it
3. Have we seen a beam in the same position and direction?  If so we have detected a loop and we are done.
4. Add the current square to the energized list
5. The figure out what to do next:
   1. If the char = `.` keep going in the same direction.  Push the beam back onto the queue.
   2. If it's a mirror, change the direction amd push it back on the queue
   3. If it's a splitter
      1. Flat side? add two beams to the queue going in opposite directions.
      2. Pointy side? push the beam back on the queue

When the queue is empty we are done and just need to return the length of the energized list.

## Part 2

Part 2 changes it in that the beam can come in from any direction.  We need to find the maximum number of tiles that can be energized.

### Solution Overview

We just need to loop around the outside of the grid and start the part 1 from all the directions and find the max.


