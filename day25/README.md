# Advent of Code 2023: Day 25: Snowverload

https://adventofcode.com/2023/day/25

## Problem Statement

Today is a graph problem.  We have a much of components that are connected and we need to split them into two groups by only disconnecting 3 wires.

This is a standard finding the min-cut graph problem, although I didn't know that when I first tackled the problem.  I researched a bunch of ways to do it.  Most of the references referred to the [Max-Flow Min-Cut theorem](https://en.wikipedia.org/wiki/Max-flow_min-cut_theorem).

There was only one part to this day.  The second star you can get by just pushing the button.

## Solution Overview
I did this one a few ways since my graph theory knowledge isn't super strong.

1. I tried to implement a max flow algorithm, but couldn't quite get it right so I switched paths.
2. Then I switch languages to python and used the [igraph](https://python.igraph.org/en/stable/) library to solve it more me.  That was pretty easy but quite dissatisfying.
3. I tried a different approach back using Go and implemented:
   - pick two nodes in the graph at random
   - find the shortest paths between the two nodes
   - for each edge in the path update a counter to say it was used
   - repeat it a 1000 times and pick the 3 edges that were visted the most.

   This worked on the real input but the sample input was too small for it to work.

## Challenges and Learnings
I'm going to come back to this a try to implement the [Stoer Wagner Algorithm](https://en.wikipedia.org/wiki/Stoer%E2%80%93Wagner_algorithm) and maybe one of the variations of [Ford-Fulkerson](https://en.wikipedia.org/wiki/Ford%E2%80%93Fulkerson_algorithm)
