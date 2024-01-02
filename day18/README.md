# Advent of Code 2023: Day 18: Lavaduct Lagoon

Visit the problem statement at [Advent of Code 2023 Day 18](https://adventofcode.com/2023/day/18).

## Problem Statement

We need to dig out an area to store lava.  We have a dig plan that give instructions (Right, Left, Up and Down) to dig.  The plan gives use a contained polygon and we need to also dig out the interior. 

Here is some example input:
```
R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
```

## Part 1

In part 1 we just need to find the area of the polygon that is created by the dig plan.  The color code is ignored.

### Solution Overview

For part 1, I implemented it using a flood fill algorithm like in [day 10](../day10/README.md).

This was slightly tricky since I didn't know how large of a grid to make.  I did it by executing the instruction once to figure out the min and max for the columns and rows then draw it out using the offset to 0,0 as the starting point.

## Part 2

In part 2, the dig plan instructions are actually in the color code.
The last character is direction to dig: `0` means `R`, `1` means `D`, `2` means `L`, and `3` means `U`.

The first 5 characters is the distance to dig.

For example:
```
#70c710 = R 461937
#0dc571 = D 56407
#5713f0 = R 356671
#d2c081 = D 863240
#59c680 = R 367720
#411b91 = D 266681
#8ceee2 = L 577262
#caa173 = U 829975
#1b58a2 = L 112010
#caa171 = D 829975
#7807d2 = L 491645
#a77fa3 = U 686074
#015232 = L 5411
#7a21e3 = U 500254
```

### Solution Overview

Flood fill won't work since the numbers are way too big.  Searching around the internet gave me the [Shoelace Formula](https://en.wikipedia.org/wiki/Shoelace_formula).  This gives the formula for calculating the area of any polygon that does not intersect itself.  Very cool.

- $A = \frac{1}{2} \left| \sum_{i=1}^{n-1} x_i y_{i+1} + x_n y_1 - \sum_{i=1}^{n-1} y_i x_{i+1} - y_n x_1 \right|$
 
Just applying this formula doesn't work in our case because we are working in a grid.  The formula will give you the area as if it's in the center of each grid.  There are a few ways to handle this.  I used Pick's Theorem.
- $A = i + \frac{b}{2} - 1$
- $A$ = the area of the polygon
- $i$ = the number of interior points
- $b$ = is number of boundary points

Rearranging the equation to get: $i = A -\frac{b}{2} + 1$

We can calculate the $b$ since it's the perimeter of the polygon, for $A$ we can use the Shoelace formula.

That will give us the number of interior points so we need to add back the perimeter.

Therefore:
- $A_{grid} = A_{shoelace} -\frac{b}{2} + 1 + b$
- $A_{grid} = A_{shoelace} + \frac{b}{2} + 1$


## Challenges and Learnings

I really liked this one, I learned some neat formulas and it took a while to really understand what was happening.

