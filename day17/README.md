# Advent of Code 2023: Day 17: Clumsy Crucible

Visit the problem statement at [Advent of Code 2023 Day 17](https://adventofcode.com/2023/day/17).

## Problem Statement

The problem for day 17 is another grid problem.  In this case we need to find a path from the top left corner to the lower right corner, that minimizes the cost of the path.

## Part 1

In Part 1 our cart that is taking the path has some special properties:
1. It can't go backwards
2. It can go in a straight path for more than 3 squares

### Solution Overview

This is standard path finding and chose to use Dijkstra's Algorithm for it.

To implement Dijkstra's algorithm, a priority queue is essential. I utilized the `container/heap` module for this purpose. The object in the priority queue has the following fields:

```
row, col   int // the position
dRow, dCol int // the direction going
dCount     int // the number of times we have gone in this direction
cost       int // The cost of the node
index      int // The index of the item in the heap
```

Go's unique approach to interfaces was used to implement the required methods for the PriorityQueue. The priority queue is sorted based on the least cost.

The following steps are executed once the queue is set up:

1. Retrieve the next item from the queue, which will be the option with the lowest cost.
2. Check if we have reached the end. If so, return the cost to get there.
3. Check if this state has already been encountered.
   - It's important to exclude cost from this check as it would make every path unique, defeating the purpose of the seen map.
4. Plan the next moves.
    1. Check if we can go straight (ensure we're within bounds and haven't exceeded the maximum distance in one direction)
    2. Check the 4 cardinal directions to determine valid moves
5. Push the new options onto the queue.

## Part 2

Part 2 introduces new movement dynamics:
1. We have to move between  4 and 10 squares in a direction before turning.
2. We still can't go backwards
3. We need at least 4 squares before it can stop at the end.

### Solution Overview

This part was just refactoring some of the rules around what are valid moves.  It wasn't so bad.

## Challenges and Learnings

This one was interesting as it was the first real time I needed to implement an interface in Go and deal with pointers.  It wasn't too bad.

I also liked refreshing my understanding of path searching.  My first attempt just used BFS and it was too slow.
