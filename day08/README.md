# Advent of Code 2023: Day 8: Haunted Wasteland

Visit the problem statement at [Advent of Code 2023 Day 8](https://adventofcode.com/2023/day/8).

## Problem Statement

In this problem, we navigate a map with instructions. We start at position AAA and aim to reach position ZZZ. Each position provides an instruction to go left or right, leading us to the next node. Here's an example:

```
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
```

The navigation process is as follows:
1. Start at AAA,
2. Go R to CCC,
3. Go L to ZZZ,
4. and we've reached the destination.

## Part 1

In Part 1, we need to count the number of steps to reach ZZZ.

### Solution Overview

Part 1 was straightforward,  I created a struct to hold the L and R values and hashed them using the key.  Then I just needed to step through the map to get the number of steps.


## Part 2

In Part 2, the problem becomes more complex. Now, we have ghosts starting at every node that ends with an 'A', and they all need to land on a node ending with a 'Z'. The catch is that they all need to be on a 'Z' node simultaneously.


### Solution Overview

My initial approach was to brute force the solution. I created a slice of current nodes, one for each ghost starting point, and stepped through the map as in Part 1. However, this approach was not feasible due to the large number of possible outcomes.

Upon further investigation, I noticed a pattern in how often the ghosts were landing on the 'Z' nodes. It seemed to cycle. I ended up finding the cycle length for each of the ghosts and then calculated the [Least Common Multiple (LCM)](https://en.wikipedia.org/wiki/Least_common_multiple).

The LCM formula is as follows:

$\text{LCM}(a, b) = \frac{|a \cdot b|}{\text{GCD}(a, b)}$

I then calculated the LCM for all the ghosts. The answer was in the range of $1.4×10^{13}$, which explains why the brute force approach was not viable.

## Challenges and Learnings

This problem was particularly interesting as it required a deeper level of investigation and the application of the LCM concept, which I hadn't used in a while.
