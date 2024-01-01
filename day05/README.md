# Advent of Code 2023: Day 5: Almanac Workflow

Visit the problem statement at [Advent of Code 2023 Day 5](https://adventofcode.com/2023/day/5).

## Problem Statement

We have an almanac and need to figure out which seeds go to which location. It's a complicated workflow based on ranges.

## Part 1

We are given a number of section workflows. Each section has a list of mappings, and each mapping is made up of three numbers: destination range start, source range start, and length.

For example:
```
50 98 2
52 50 48
```
means the range on the left gets mapped to the range on the right:
```
98-99 -> 50-51
50-98 -> 52-97
```

There are multiple workflows to map seeds to the final location map.  We need to perform the workflow and find the lowest location number.

here are multiple workflows to map seeds to the final location map. We need to perform the workflow and find the lowest location number.

### Solution Overview

The first step is to parse the file. I split the file on `\n\n` to separate the sections. The first section is a list of seed values all the others sections are then processed to create a mapping of destination to source ranges. 

Each mapping is an struct with a apply method on it.  It takes in a number, applies the mapping and then returns the new number.

I then apply all these mappings to the seeds to determine their final locations.

Search of the final locations and get the lowest value.

## Part 2

Part 2 changes things up by making the seeds ranges as well!  So for part 1 this seed map

```
seeds: 79 14 55 13
```

means 4 seeds, with the values of 79, 14, 55, and 13.

In part 2, the same line means seeds 79-92 and 55-57 for a total of 27 seeds.

### Solution Overview

Part 2 starts with a some new parsing of the seed section to get the ranges and then proceeds with the same parsing for the workflow sections.

Next we need to apply a range of values to each range mapping.  When we do this there are up to 3 ranges of values produced.

If the input range overlaps with the mapping range, it can produce two ranges: the intersection and the leftover.
```
Input:           |-----|
Mapping:           |-----|
Intersection:      |---|
LeftOver:        |-|
```

If the input range contains the mapping range entirely, it can produce three ranges: two leftovers and the intersection.
```
Input:           |----------|
Mapping:           |-----|
Intersection:      |-----|
LeftOver:        |-|     |--|
```

Since each step of the workflow is like any other, I treated them like a list of items to work on. If a range got split on a mapping, I would add the mapped (intersection) values to the new list and any leftover section back into the work list. This approach handles the fall-through from different ranges. After each section is processed, I'd start again with the new mapped values.


## Challenges and Learnings

This problem presented several challenges, mainly that ranges hurt my brain.  Once I understood that a range can be split up into up to 3 ranges by another it wasn't too bad.  In a later day there was more ranges required so I moved a bunch of the logic into a util file.


