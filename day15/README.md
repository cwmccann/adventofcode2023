# Advent of Code 2023: Day 15: Lens Library

Visit the problem statement at [Advent of Code 2023 Day 15](https://adventofcode.com/2023/day/15).

## Problem Statement

We are implementing a hashmap.

## Part 1

Part 1, we need to implement the hashing algorithm.  It's just a serier of operations on a each character of a string:
1. Determine the ASCII code for the current character of the string.
2. Increase the current value by the ASCII code you just determined.
3. Set the current value to itself multiplied by 17.
4. Set the current value to the remainder of dividing itself by 256.

### Solution Overview

It was pretty easy to just implement the hashing algorithm as described and hash the input.

## Part 2

In Part 2, we need to implement a hashmap to store lenses in boxes. We have two operations: add (`=`) and remove (`-`). Collision resolution is handled using a list.


### Solution Overview

The solution involves implementing a hash table with collision resolution using a list. Here's an example input: `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`.

S
Steps:
1. Hash the key to determine the box it goes into.
2. If the operation is `=`, add the item to the box.
3. If the operation is `-`, remove the item from the box.

When adding an item to a box that already contains an item, we need to:
1. Check if an item with the same label is already in the box. If so, update the value.
2. If it's not in the box, add it to the end.

When removing an item:
1. If the item is in a box, remove it while preserving the order. I used slices to handle this.
2. If it's not in a box, don't do anything.

## Challenges and Learnings

This problem was fun and involved a straightforward implementation of a hash table.
