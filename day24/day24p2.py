import sympy

with open("input.txt") as f:
    lines = f.readlines()

hailstones = []
for line in lines:
    split_line = line.replace("@", ",").split(",")
    int_line = map(int, split_line)
    hailstones.append(tuple(int_line))

x, y, z, dx, dy, dz = sympy.symbols("x, y, z, dx, dy, dz")

equations = []

for i, (xn, yn, zn, dxn, dyn, dzn) in enumerate(hailstones[:3]):
    equations.append((x - xn) * (dyn - dy) - (y - yn) * (dxn - dx))
    equations.append((y - yn) * (dzn - dz) - (z - zn) * (dyn - dy))


solutions = sympy.solve(equations)

print("Found", len(solutions), "solutions")
print("Solutions", solutions)

#find the integer solutions
int_solutions = []
for sol in solutions:
    if all(x % 1 == 0 for x in sol.values()):
        int_solutions.append(sol)

print("Found", len(int_solutions), " int solutions")
print("Int Solutions", int_solutions)

answer = int_solutions[0]
print(answer[x] + answer[y] + answer[z])
