# Advent of Code 2023: Day 22: Sand Slabs

Visit the problem statement at [Advent of Code 2023 Day 22](https://adventofcode.com/2023/day/22).

## Problem Statement

Here we are given a number of blocks (rectangular prisms) in 3D space.  They are in the process of falling and have some special properties to make it easier, such at if a block is supported in any way it won't fall.  The also don't twist.

## Part 1

In Part 1, we need to figure out which ones are safe to destroy.  A block is safe to destroy if it doesn't cause other black to drop.

### Solution Overview

The first step after parsing is to sink the blocks to the ground to see how they end up.

To sink the blocks:
1. Sorted them by the Z value so they lowest blocks got process first.
2. I then chose 1 as the min, since it was the ground and checked if there were any block under it that would stop it from falling. 
   - If there was I would set the min to be +1 the lower blocks z value.

Once the blocks were on the ground I create a map of all the block that supported other blocks, and a reverse map of each block that is supportedBy another block.

If the block doesn't support any other it can be disintegrated.  Also, if there is another block that can take the load this block can be destroyed.  Odd logic but it solves the puzzle.

## Part 2

In part two, we need to figure out what will happen if we cause a chain reaction and destroy blocks that other are support by.

### Solution Overview

Part 2 starts the same way as part 1, read the blocks in, sink them, and create the supports and supported by maps.

The we need to figure out the chain reaction.  I used a FIFO for this and tracked the list of falling blocks.  The FIFO was use to track the potential falling blocks that need to be checked.

## Challenges and Learnings

I really enjoyed this days.  It was tricky but could be tackled with just logic and didn't really require any new algorithms or formulas.

