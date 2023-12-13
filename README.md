# Advent of code 2023

I'm learning Go and thought to do the advent of code this year with a new language.

Let's see how far I get!

## Impressions

### After Day 12
So day 12 is history and I'm getting ready for day 13.  So far I quite like Go as a language but I'm slightly shocked to how little is included in the standard library. Since Aoc deals with lots of lists of integers it's been a bit painful to have to write so many functions I would just expect tpo be there.  For example if the you at my utils package (still a mess) there's functions for finding the min and max in a slice of ints, absolute value, and sum the list.  All these things I would expect to be included in the standard package.  I understand it's partially because Go didn't have support for generics until very recently.

Overall though I like the language, it's type safe, which I prefer for most cases.  It looks familiar to me and is quite readable.  I haven't worked with go routines much yet but they look very interesting.

## Solution Notes

### Day 12 Hot Springs

https://adventofcode.com/2023/day/12

This challenge really made me feel like I need to brush up on my recursion technique.  I solved it by creating a recursive method `countArrangements`.  It took in the condition record for the springs, the count of each spring, and the state of the recursion as a struct. 

The goal is to count how many way you can replace the `?` with either a `#` or a `.` to satisfy the grouping of `#`.  Check the problem details on the AoC website for more information on the problem.

The state holds:
  - `s` the current state of the string.  Not really needed until part 2.  
  - `i` indicating the position in the string we are considering
  - `gi` the index into which group we are processing
  - `current` the current number of `#`'s in the group

We first check if we are at the end of the string and have replaced all the `?`  if we have found a solution we return 1, otherwise we return 0 to indicate it's not a solution.

If we aren't at the end we consider the character in the string.  For the `.` or `#` characters update the state and keep going.
The more interesting part if when there is a `?`.  At that point try both characters and see if we find a solution.  This is a kinds of a naive approach and additional logic could be added to be smarter.  For example if we have processed all the groups trying a `#` won't help.  It worked though.

Part 2 wasn't too bad.  After "unfolding" the input string and groups, my old solution wouldn't run in time. To fix that I memoized the function by keeping a cache of the state and current answer.  At first it would work and I need to add a unique key for the two replacement functions.  I ended up just adding `newS` which had the replaced value in it.  This also made debugging a bit easier.




