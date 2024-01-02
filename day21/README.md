# Advent of Code 2023: Day 21: Step Counter

Visit the problem statement at [Advent of Code 2023 Day 21](https://adventofcode.com/2023/day/21).

## Problem Statement

The problem for Day 21 involves finding how many positions can be reached in N steps. We are given a map and the starting position is in the middle of the map. Here's the example map:

```
...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
```

## Part 1

In part 1 we need to find out how many positions can be reached in 64 steps. If you look a different set of positions can be reached depending on if it's an even or odd number of steps. For example in 2 steps you can reach your starting point, but there is now way to do that in 3.

### Solution Overview

I used a BFS to find the positions you can get to in 64 steps. I keep track of only the even positions in the seen set.

To get the total positions I just return count of how many unique positions we've visited on even steps.

## Part 2

Part 2, things get a bit harry. We now have 26,501,365 steps to take and the grid map extends infinitely in all directions. Yikes!

### Solution Overview

The first step is to make it so we can go in any direction and step off the single grid. I did this by creating a function to map the current point x,y to the original grid using the modular operator.

```go
func GetCharFromPoint(grid [][]rune, p Point) rune {
	m := MapPoint(p, len(grid), len(grid[0]))
	return grid[m.Y][m.X]
}

func MapPoint(p Point, R int, C int) Point {
	return Point{
		X: (p.X%C + C) % C,
		Y: (p.Y%R + R) % R,
	}
}
```

Then I recorded various steps vs positions reached. It turns out it's quadratic! Ok now I needed to figure out how to find the equation. Turns out Newtons' Interpolation method can provide a 2nd degree equation given 3 points of data.

For the x values I used the number of steps to cross the grid (the length of the column) with an offset of the middle of the grid.

Turns out the grid is 131x131 and the starting position is at the midpoint of 65.

If you take $\frac{26501365 - 65}{131} = 202300$

Here are the x and y values for my puzzle input.
| X   | Positions |
| --- | --------- |
| 0   | 3814      |
| 1   | 33952     |
| 2   | 94138     |

I used some code to find the coefficients of a newton polynomial using those values.  Then I solved the polynomial at $x = 202300$

## Challenges and Learnings

This one was tough.  I learned fairly quickly that it was a polynomial but I had never encounter a way to determine a polynomial from the points. Then I didn't release it wasn't in the standard form $P(x) =  ax^2 + bx + c$, it was in the newton form $P(x) = f[x_0] + f[x_0, x_1](x - x_0) + f[x_0, x_1, x_2](x - x_0)(x - x_1)$ so I needed another method to evaluate the polynomial given the coefficients.
