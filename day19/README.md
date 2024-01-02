# Advent of Code 2023: Day 19: Aplenty

Visit the problem statement at [Advent of Code 2023 Day 19](https://adventofcode.com/2023/day/19).

## Problem Statement

We have a bunch of parts with some fields on it.  For example `{x=787,m=2655,a=1222,s=2876}` and we have a set of workflows the parts need to go through to see if they will be accepted or not.  Here's an example workflow: `px{a<2006:qkq,m>2090:A,rfg}`

Let's break this down.
- The workflow label is `px`
- If the part field `a < 2006` proceed to workflow `qkq`
- If the part field `m > 2090` The part is Accepted `A`
- Otherwise proceed to workflow `rfg`

All the workflows have a similar structure.  If the part meets a condition the either proceed to another workflow or are Accepted `A` or rejected `R`.  Each workflow also has a default next step.

## Part 1

In part1 we need to count how many parts get accepted and for each of them add up the `xmas` fields and return the total.

### Solution Overview

Part was a lot of careful parsing and building a set of data structure to support all the features.

Here's by object definitions
```go
type Part struct {
	x, m ,a, s int
}

type Rule struct {
	condition RuleCondition
	action string
}

type RuleCondition struct {
	field, operator string
	value int
}

type Workflow struct {
	rules []Rule
}
```

I also had a map of workflows keyed by the label.

After that it was just a matter of processing the parts and stashing the accepted parts into a slice.

## Part 2

Part 2 changes the parts have ranges of values (1-4000) instead of a single value.  Now I need to return how many parts would be accepted.

### Solution Overview

I couldn't brute force this one since we would have $4000^4$ different values to test.  So we had to work with ranges again.

To that I created a recursive function `CountAccepted` that takes the workflow map, a map of Ranges that contained the fields `xmas` and a Starting Range of `1-4000`

The function does the following:
1. If the workflow is `A` we return the the total which is the each range multiplied together.
2. If the workflow is a `R` return 0.  All the ranges are rejected.
3. Then I apply the rule for the current workflow
   1. No field means we are at the default case to return the results from that
   2. The I apply the operator to the field and split the range into two parts:
      1. the range that matches the rule `in`
      2. the range that doesn't match `out`
      3. if the `in` range is not empty move to the next workflow (recursive call)
      4. if the `out` range is not empty keep applying the rules in the current workflow

## Challenges and Learnings

As I've said before, ranges hurt my head.  This took a while to get all bit worked out.
