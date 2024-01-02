# Advent of Code 2023: Day 12: Hot Springs

Visit the problem statement at [Advent of Code 2023 Day 12](https://adventofcode.com/2023/day/12).

## Problem Statement

The problem involves processing a string record representing a sequence of springs. The springs can be intact (`#`), broken (`.`), or uncertain (`?`). The goal is to find a configuration of the uncertain springs that satisfies a given set of group conditions.

For example an undamaged record could look like this:
```
#.#.### 1,1,3
.#...#....###. 1,1,3
.#.###.#.###### 1,3,1,6
####.#...#... 4,1,1
#....######..#####. 1,6,5
.###.##....# 3,2,1
```

However some of the records are damaged and we have `?` instead of `.#`.

For example:
```
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
```

## Part 1

In Part 1, we need to figure out how many different configurations are possible for each row when replacing the `?` with a `.` or a `#`.

We then need to sum up the results from each row.

### Solution Overview

I implemented a recursive algorithm to solve this problem. The state of the recursion is tracked using a struct with the following fields:

```go
    s string //The string record
	i int //The index into the string we are looking at
    gi int // the group index we are processing,
    current int // the current number of broken springs we have found in the group
```

I also pass in the a slice of all the groups.

Recursive Logic:
1. Check if we are done.
    - We are done when we have processed the entire string.
    - We have found a solution of we have processed all of the groups and we have `current = 0`
    - We are also found a solution if we are processing the last group and the `current = the last group`
2. If the current character is a `.` we are either ending a group or just proceeding.  Increment the index and optionally end the group.
3. If the character is a `#` add to the current counter.
4. If the character is a `?` try replacing it with both a `.` and a `#` and see if we get a solution.

## Part 2

Part two makes the search space way larger!  The record is actually folded up and it's 5 times the size.

For example:
`???.### 1,1,3` becomes `???.###????.###????.###????.###????.### 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3`

### Solution Overview

Trying my part 1 solution on part 2 didn't work.  The search space is too large.  The solution is to use memorization to cache the results.  That speeds up the results enough to get a result in a few seconds.

## Challenges and Learnings

This problem was interesting and challenging. Figuring out the state tracking for the recursive function was tricky, and implementing the logic was finicky. I enjoyed the added complexity of Part 2, where memoization was needed to handle the larger search space.
