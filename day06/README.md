# Advent of Code 2023: Day 6: Wait For It

Visit the problem statement at [Advent of Code 2023 Day 6](https://adventofcode.com/2023/day/6).

## Problem Statement

We are racing boats!  Each race has a time component that tells us how long the race lasts and a field saying what the record distance is.  To race the boat you need to hold down the button and for each unit of time it gets held down your boat goes faster. The time pressing the button counts towards the total time of the race.

## Part 1

For parts 1 we need to figure out how many ways we can beat the record for each race.  The answer is the number of ways for each race multiplied together.

### Solution Overview

I used a two pointer approach on this one.  I found the min amount of time we can hold down the button and beat the target and then I found the max time.  The answer is $maxTime - minTime + 1$

## Part 2

Part two just combines the all the games into one massive game.  I think it was just so you couldn't brute force it.  

### Solution Overview

The two pointer approach worked on part 2 so I just need to change the parsing a bit.


## Challenges and Learnings

This one was pretty easy and after solving it iteratively I found out you can solve it using an quadratic equation but didn't implement the solution fully.

1. The equation for distance is: $d = v \cdot t$
2. The velocity equals the amount of time you hold down the button: $n$
3. Plug that in gives you $d=n(t-n)$
4. That gives you $d=tn - n^2$
5. Which gives us $-n^2 + tn - d = 0$




