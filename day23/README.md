# Advent of Code 2023: Day 23: A Long Walk

https://adventofcode.com/2023/day/23

## Problem Statement

You've got a bit of down time so you figure it would be nice to go for a long walk.  This time we are given a maze / grid and we want to find the longest walk we can without hitting the same spot twice.

## Part 1

In part 1 we have a maze with slopes in it and the slopes are icy so we can only go in one direction when we hit a slope.

### Solution Overview

The longest path is interesting since most of the time we are looking for the shortest path.  I solved this using Breath First Search algorithm.  

The longest path is considered a NP-Hard so the solution will not be that efficient.  Use the BFS I search all the paths that follow the rules of part 1, and each time I find the end point I put the path in a completed slice.  At the end I just looked through each path and found the longest one.

To keep track path taken I push it onto the queue with the current position.  This actually was more tricky than I thought since in go `newPath = append(path, newP)` does not always return a new slice.  I was getting crazy paths as the different forks in the algorithm started to overwrite each others path.

I ended up solving it by checking if there was more than one option at a grid point then I would copy it to a new slice, otherwise, I could just append it.

## Part 2

Part 2 introduces an interesting twist: the slopes are not as icy, allowing movement in both directions. This change significantly increases the search space. The input slopes surround all points in the mazes to prevent loops.

Some solutions on Reddit suggested removing the rules to go down the slopes and letting it run. However, this approach led to memory exhaustion due to the vast number of paths.

The solution was to "compress" the graph. The input contains sections that are long paths without any branching. Instead of iterating along that each time, we can jump to the next branch and record that we walked n steps.

Here's the approach I took:
1. Searched the grid to find all spots with more than one path leading out. This step created a set of vertices in a new graph.
2. Performed a BFS to find the distance between each pair of connected vertices. This step created a weighted edge for the graph.
3. Performed a DFS between the start and end nodes in the graph to find the longest path, similar to the solution in part 1.

## Challenges and Learnings

This problem was quite engaging. It took me some time to realize that my slices were being reused and that I needed to copy them. In some ways, this was a hint for part two to compress the graph to just the vertices that you can branch from.
