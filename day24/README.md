# Advent of Code 2023: Day 24: Never Tell Me The Odds

https://adventofcode.com/2023/day/24

## Problem Statement

The snow machine isn't working, it's forming hail instead!

## Part 1

For part 1 we are given the position and velocity of each hail in the format `19, 13, 30 @ -2,  1, -2` where each number is `x, y, z @ dx, dy, dz`.

The problem is to find out which particles have an intersecting path ignoring the z position and it must have a positive time component.  An interesting thing to point out, is this is the first AOC day this year that required floating point numbers.

### Solution Overview

This is a standard linear algebra problem (which I'm a but rusty at).

Given $(x_1, y_1), (x_2, y_2)$ and the direction vectors $(dx_1, dy_1), (dx_2, dy_2)$ you can set up a system of equations to solve.
1. Setup the equations
   - $x_1 + t_1 \cdot dx_1 = x_2 + t_2 \cdot dx_2$
   - $y_1 + t_1 \cdot dy_1 = y_2 + t_2 \cdot dy_2$
2. Rearrange to $Ax = b$
   - $t_1 \cdot dx_1 - t_2 \cdot dx_2 = x_2 - x_1$
   - $t_1 \cdot dy_1 - t_2 \cdot dy_2 = y_2 - y_1$
3. Transfer it to matrix form
   - $\begin{pmatrix}
dx_1 & -dx_2 \\
dy_1 & -dy_2
\end{pmatrix}
\begin{pmatrix}
t_1 \\
t_2
\end{pmatrix}
=
\begin{pmatrix}
x_2 - x_1 \\
y_2 - y_1
\end{pmatrix}$
4. Solve the matrix equation, $A^{-1}b=t$.  I ended up using a Go library [gonum](https://www.gonum.org/) for this.
5. That gives me the values for $(t_1, t_2)$ or an error if the detriment is zero which means that the lines are parallel.
6. I did this for each line combination in the input and also added a check to ensure the $t$ values were positive (forward in time), and that the lines intersected within the given range.

## Part 2

Things get harder for part 2.  What we need to find is an equation of a particle (rock) that will intersects all the hail particles, in 3D space.

Given a number of hailstones with a position of $(x_n, y_n, z_n) and direction of (dx_n, dy_n, dz_n)$ we need to find a rock with a position of $(x, y, z)$ and a direction of $(dx, dy, dz)$ that will intersect all the hailstones.

In this case we also need to account for $t$.

1. Setup the equations
   - $x_n + t \cdot dx_n = x + t \cdot dx$
   - $y_n + t \cdot dy_n = y + t \cdot dy$
   - $z_n + t \cdot dz_n = z + t \cdot dz$

2. Rearrange to $t =$
   - $t = \frac{x - x_n}{dx_n - dx}$
   - $t = \frac{y - y_n}{dy_n - dy}$
   - $t = \frac{z - z_n}{dz_n - dz}$

3. We can then set these equal to each other and solve for $x, y, z$.
   - $\frac{x - x_n}{dx_n - dx} = \frac{y - y_n}{dy_n - dy} = \frac{z - z_n}{dz_n - dz}$

4. Now we can pick 2 of the equations and solve for $x, y, z$.
   - $\frac{x - x_n}{dx_n - dx} = \frac{y - y_n}{dy_n - dy}$
   - $\frac{x - x_n}{dx_n - dx} = \frac{z - z_n}{dz_n - dz}$
   - $\frac{y - y_n}{dy_n - dy} = \frac{z - z_n}{dz_n - dz}$

5. Rearrange so they equal 0
   - $(x - x_n)(dy_n - dy) - (y - y_n)(dx_n - dx) = 0$
   - $(x - x_n)(dz_n - dz) - (z - z_n)(dx_n - dx) = 0$
   - $(y - y_n)(dz_n - dz) - (z - z_n)(dy_n - dy) = 0$

6. we have 6 unknowns and 3 equations, so we need to have multiple values for $n$ to solve for $x, y, z, dx, dy, dz$.

7. At this point I got tired and used sympy to solve the equations for me.
``` python
#create the symbols

x, y, z, dx, dy, dz = sympy.symbols("x, y, z, dx, dy, dz")
equations = []

for i, (xn, yn, zn, dxn, dyn, dzn) in enumerate(hailstones[:3]):
    equations.append((x - xn) * (dyn - dy) - (y - yn) * (dxn - dx))
    equations.append((y - yn) * (dzn - dz) - (z - zn) * (dyn - dy))

solutions = sympy.solve(equations)
```

8. The only remaining step is to search for an integer solution in the range of the solutions.

I know there are a few ways to solve this without using sympy, and I'm planning to come back to fix this.


## Challenges and Learnings

This one was interesting because of the math.  Learning to use sympy was also pretty cool.  At some point I hope to come back and do the solution in go.
