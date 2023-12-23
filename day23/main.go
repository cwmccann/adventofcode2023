package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strings"
	"github.com/gammazero/deque"
	"time"

)

type Point = utils.Point

func main() {
	start := time.Now()
	text := utils.InputToString()

	solution := SolvePart1(text)
	elapsed := time.Since(start)

	fmt.Printf("Solution Part 1: %d took %s\n", solution, elapsed)

	start = time.Now()
	solution = SolvePart2(text)
	elapsed = time.Since(start)
	fmt.Printf("Solution Part 2: %d took %sa\n", solution, elapsed)
}

type Path struct {
	p Point
	path []Point
}

func SolvePart1(input string) int {
	lines := utils.TrimAndRemoveEmptyLines(strings.Split(input, "\n"))
	grid := utils.StringsToRune2D(lines)

	R := len(grid)
	C := len(grid[0])

	start := Point{X: 1, Y: 0}
	goal := Point{X: C - 2, Y: R - 1}

	q := deque.New[Path]()
	complete := make([]Path, 0)

	q.PushBack(Path{p: start, path: []Point{start}})
	for q.Len() > 0 {
		// fmt.Println("q len = ", q.Len())
		qi := q.PopFront()
		p := qi.p
		path := qi.path

		//printGrid(grid, path, p)

		if p == goal {
			complete = append(complete, qi)
			continue
		}

		usedPath := false
		for _, newP := range p.GetCardinalNeighbors() {
			if !newP.IsInGrid(R, C) {
				continue
			}

			if newP.IsInList(path) {
				continue
			}

			lastC := grid[p.Y][p.X]
			if (lastC == '>') && (newP.X < p.X) {
				continue
			}
			if (lastC == '<') && (newP.X > p.X) {
				continue
			}
			if (lastC == '^') && (newP.Y > p.Y) {
				continue
			}
			if (lastC == 'v') && (newP.Y < p.Y) {
				continue
			}

			c := grid[newP.Y][newP.X]

			if c == '#' {
				continue
			}

			//This was a pain to figure out
			//The paths between nodes was getting crossed and weird things were happening
			//Turns out the path needs to be copied when we hit a branch
			var newPath []Point
			if usedPath {
				newPath = make([]Point, len(path))
				copy(newPath, path)
				newPath = append(newPath, newP)
			} else {
				newPath = append(path, newP)
			}

			if utils.RuneInString(c, ".>v^<") {
				usedPath = true
				q.PushBack(Path{p: newP, path: newPath})
			}
		}
	}

	max := 0
	for _, p := range complete {
		// fmt.Printf("%d,%d\n", len(p.path), p.path)

		if len(p.path) > max {
			max = len(p.path)
		}
	}
	return max - 1
}

func SolvePart2(input string) int {
	lines := utils.TrimAndRemoveEmptyLines(strings.Split(input, "\n"))
	grid := utils.StringsToRune2D(lines)

	R := len(grid)
	C := len(grid[0])

	type Edge struct {
		p Point
		length int
	}

	//find all the vertices (spots with more than 2 neighbors)
	//Make it a map of point that contains a list of points
	//which will be the collapsed vertices it joins to
	vertices := make(map[Point][]Edge, 0)

	//Add the start and end points
	start := Point{X: 1, Y: 0}
	vertices[start] = make([]Edge, 0)
	goal := Point{X: C - 2, Y: R - 1}
	vertices[goal] = make([]Edge, 0)

	for y, row := range grid {
		for x, char := range row {
			if char == '.' {
				p := Point{X: x, Y: y}
				neighbors := 0
				for _, neighbor := range p.GetCardinalNeighbors() {
					if !neighbor.IsInGrid(R, C) {
						continue
					}

					c := grid[neighbor.Y][neighbor.X]
					if c != '#' {
						neighbors++
					}
				}
				if neighbors > 2 {
					//Add it to the list of vertices
					vertices[p] = make([]Edge, 0)
				}
			}
		}
	}
	// fmt.Println("Vertices: ", len(vertices))

	//Now find the edges and the length of the path between them
	//Use BFS to find the shortest path between any two vertices
	for v := range vertices {
		q := deque.New[Path]()
		q.PushBack(Path{p: v, path: []Point{v}})

		for q.Len() > 0 {
			// fmt.Println("q len = ", q.Len())
			qi := q.PopFront()
			p := qi.p
			path := qi.path

			if p != v {
				if _, ok := vertices[p]; ok {
					//We found a vertex
					//Add it to the list of edges
					// fmt.Printf("Found edge from %v to %v with length %d\n", v, p, len(path) - 1)
					vertices[v] = append(vertices[v], Edge{p: p, length: len(path) - 1})
					continue
				}
			}

			for _, newP := range p.GetCardinalNeighbors() {
				if !newP.IsInGrid(R, C) || newP.IsInList(path) {
					continue
				}

				c := grid[newP.Y][newP.X]
				if c == '#' {
					continue
				}

				//we just need to make a newPath when we are at the starting vertex
				//otherwise we can just append to the existing path
				var newPath []Point
				if p == v {
					newPath = make([]Point, len(path))
					copy(newPath, path)
					newPath = append(newPath, newP)
				} else {
					newPath = append(path, newP)
				}

				if utils.RuneInString(c, ".>v^<") {
					q.PushBack(Path{p: newP, path: newPath})
				}
			}
		}
	}

	//The graph is now a list of vertices and edges
	// fmt.Println("Vertices: ", len(vertices))
	// for v, edges := range vertices {
	// 	fmt.Println(v, edges)
	// }

	//do a dfs to find the longest path between the start and end
	//use a stack to keep track of the path
	//when we hit a vertex, we can add the length of the edge to the path
	//when we hit the end, we can compare the path length to the max

	type DfsPath struct {
		e Edge
		path []Point
	}

	max := 0
	stack := deque.New[DfsPath]();
	stack.PushFront(DfsPath {e: Edge{p: start, length: 0}, path: []Point{}})

	// fmt.Println("Starting DFS")
	for stack.Len() > 0 {
		//pop the top of the stack
		dfsPath := stack.PopFront()
		e := dfsPath.e

		//if we are at the end, compare the path length to the max
		if e.p == goal {
			if e.length > max {
				max = e.length
			}
			continue
		}

		//otherwise, add the edges to the stack
		for _, edge := range vertices[e.p] {
			if edge.p.IsInList(dfsPath.path) {
				continue
			}
			newPath := make([]Point, len(dfsPath.path))
			copy(newPath, dfsPath.path)
			newPath = append(newPath, edge.p)

			newEdge := Edge{p: edge.p, length: e.length + edge.length}
			stack.PushFront(DfsPath {e: newEdge, path: newPath})
		}
	}

	return max;
}

func printGrid(grid [][]rune, path []Point, cur Point) {
	// Create a copy of the grid
    gridCopy := make([][]rune, len(grid))
    for i := range grid {
        gridCopy[i] = make([]rune, len(grid[i]))
        copy(gridCopy[i], grid[i])
    }

    // Modify the copy, not the original grid
    for _, p := range path {
        gridCopy[p.Y][p.X] = 'O'
    }
	gridCopy[cur.Y][cur.X] = 'X'

    for _, line := range gridCopy {
        fmt.Println(string(line))
    }
    fmt.Println()
}



