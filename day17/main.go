package main

import (
	"adventofcode2023/utils"
	"container/heap"
	"fmt"
	"strings"
)

func main() {
	text := utils.InputToString()

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}

func SolvePart1(input string) int {
	grid := StringsToInt2D(strings.Split(input, "\n"))
	return findCost(grid, 1, 3)
}

func SolvePart2(input string) int {
	grid := StringsToInt2D(strings.Split(input, "\n"))
	return findCost(grid, 4, 10)
}

func findCost(grid [][]int, minLength int, maxLength int) int {
	R := len(grid)
	C := len(grid[0])

	// Priority queue for Dijkstra's Algorithm
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{row: 0, col: 0, dRow: 0, dCol: 0, dCount: 0, cost: 0})
	seen := make(map[string]bool)

	i := 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		i++

		if i%100000 == 0 {
			fmt.Printf("i = %d, len(heap) = %d, len(seen) = %d\n", i, pq.Len(), len(seen))
		}

		//We found the end, return the cost
		if item.row == R-1 && item.col == C-1 && item.dCount >= minLength {
			fmt.Printf("Found the end in i = %d, cost = %d\n", i, item.cost)
			return item.cost
		}

		// If we have already seen this node, skip it
		hash := item.Hash()
		if seen[hash] {
			continue
		}
		seen[hash] = true

		//See if we can keep going in the same direction
		if item.dCount < maxLength {

			//Check for the starting position
			if item.dRow != 0 || item.dCol != 0 {
				newRow := item.row + item.dRow
				newCol := item.col + item.dCol

				//Make sure we are not out of bounds
				if inBounds(grid, newRow, newCol) {
					newCost := item.cost + grid[newRow][newCol]
					heap.Push(&pq, &Item{
						row:    newRow,
						col:    newCol,
						dRow:   item.dRow,
						dCol:   item.dCol,
						dCount: item.dCount + 1,
						cost:   newCost,
					})
				}
			}
		}

		type Move struct{ dRow, dCol int }
		moves := []Move{{dRow: 1, dCol: 0}, {dRow: -1, dCol: 0}, {dRow: 0, dCol: 1}, {dRow: 0, dCol: -1}}

		for _, move := range moves {
			newRow := item.row + move.dRow
			newCol := item.col + move.dCol

			//Make sure we are not out of bounds
			if !inBounds(grid, newRow, newCol) {
				continue
			}
			//make sure we are not going forward since it's handle above
			if item.dRow == move.dRow && item.dCol == move.dCol {
				continue
			}
			//Not allowed to go backwards
			if item.dRow == -move.dRow && item.dCol == -move.dCol {
				continue
			}

			//Need to travel the min length
			//Also need to handle the case where we are at the start
			if (item.dCount != 0 && item.dCount < minLength) {
			  	continue
			}

			newCost := item.cost + grid[newRow][newCol]
			heap.Push(&pq, &Item{
				row:    newRow,
				col:    newCol,
				dRow:   move.dRow,
				dCol:   move.dCol,
				dCount: 1,
				cost:   newCost,
			})
		}
	}

	//Didn't find a path
	return -1
}

func inBounds(grid [][]int, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}

func StringsToInt2D(lines []string) [][]int {
	int2D := make([][]int, 0)
	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}
		int2D = append(int2D, StringToDigits(line))
	}
	return int2D
}

func StringToDigits(s string) []int {
	digits := make([]int, len(s))
	for i, c := range s {
		digits[i] = int(c - '0')
	}
	return digits
}

type Item struct {
	row, col   int // the position
	dRow, dCol int //the direction going
	dCount     int // the number of times we have gone in this direction
	cost       int // The cost of the node
	index      int // The index of the item in the heap
}

func (item *Item) String() string {
	return fmt.Sprintf("Item{row: %d, col: %d, dRow: %d, dCol: %d, dCount: %d, cost: %d, index: %d}", item.row, item.col, item.dRow, item.dCol, item.dCount, item.cost, item.index)
}
//Used to detect duplicates so don't include cost or index.
func (Item *Item) Hash() string {
	return fmt.Sprintf("%d,%d,%d,%d,%d", Item.row, Item.col, Item.dRow, Item.dCol, Item.dCount)
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // clean up memory
	item.index = -1 // it's not in the heap anymore
	*pq = old[0 : n-1]
	return item
}
