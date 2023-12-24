package main

import (
	"adventofcode2023/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"github.com/gammazero/deque"
	mapset "github.com/deckarep/golang-set/v2"
)
type Point = utils.Point

func main() {
	text := utils.InputToString();

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}


func SolvePart1(input string) int {
	blocks := readBlocks(input)

	blocks = sinkBlocks(blocks)

	supportedBy, supports := createSupportedAndSupportsMaps(blocks)

	canBeDisintegrated := 0
	for _, block := range blocks {
		if len(supports[block.id]) == 0 {
			canBeDisintegrated++
			continue
		}
		//if all the blocks are supported by another block, then it can be disintegrated
		allTrue := true
		for _, supportedByCurrentBlock := range supports[block.id] {
			if len(supportedBy[supportedByCurrentBlock]) <= 1 {
				allTrue = false
			}
		}
		if allTrue {
			canBeDisintegrated++
		}
	}
	return canBeDisintegrated
}

func SolvePart2(input string) int {
	idCounter = 0

	blocks := readBlocks(input)
	blocks = sinkBlocks(blocks)
	supportedBy, supports := createSupportedAndSupportsMaps(blocks)
	// fmt.Println("Supports")
	// for id, support := range supports {
	// 	fmt.Printf("%s: %v\n", id, support)
	// }

	total := 0

	//For each block count the number of blocks that would fall if it was removed
	for _, block := range blocks {
		falling := mapset.NewSet[string]()
		falling.Add(block.id)

		q := deque.New[string]()
		for _, supportedByCurrentBlock := range supports[block.id] {
			q.PushBack(supportedByCurrentBlock)
		}

		for q.Len() > 0 {
			upper := q.PopFront()

			supportingBlocks := supportedBy[upper]

			allFalling := true
    		for _, block := range supportingBlocks {
        		if !falling.Contains(block) {
            		allFalling = false
            		break
				}
        	}

			if allFalling {
				falling.Add(upper)
				for _, supportedByCurrentBlock := range supports[upper] {
					q.PushBack(supportedByCurrentBlock)
				}
			}
		}
		// fmt.Printf("destroying: %s would cause %d blocks to fall\n", block.id, falling.Cardinality() - 1)
		total += falling.Cardinality() - 1
	}

	return total
}

func readBlocks(input string) []block {
	lines := strings.Split(input, "\n")
	blocks := make([]block, 0)
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		blocks = append(blocks, CreateBlockFromString(line))
	}
	return blocks
}

func sinkBlocks(blocks []block) []block {
	sortByZ := func(a, b int) bool {
		return blocks[a].MinZ() < blocks[b].MinZ()
	}

	//Sort the blocks by the min z value
	sort.Slice(blocks, sortByZ)

	//Sink all the blocks to the ground
	for i, upper := range blocks {
		minZ := 1
		//It's as low as it can go
		if upper.z1 <= minZ {
			continue
		}

		for j := i - 1; j >= 0; j-- {
			lower := blocks[j]
			overlap := upper.GetXYPoints().Intersect(lower.GetXYPoints())
			if overlap.Cardinality() > 0 {
				minZ = utils.Max(lower.z2 + 1, minZ)
			}
		}
		blocks[i].drop(minZ)
	}

	//Sort them again incase any shifted positions after sinking
	sort.Slice(blocks, sortByZ)

	return blocks
}

func createSupportedAndSupportsMaps(blocks []block) (map[string][]string, map[string][]string) {
	supportedBy := make(map[string][]string)
	supports := make(map[string][]string)

	for i, upper := range blocks {
		supportedBy[upper.id] = make([]string, 0)
		supports[upper.id] = make([]string, 0)

		for j := i - 1; j >= 0; j-- {
			lower := blocks[j]
			overlap := upper.GetXYPoints().Intersect(lower.GetXYPoints()).Cardinality() > 0

			if overlap && lower.z2 + 1 == upper.z1 {
				supportedBy[upper.id] = append(supportedBy[upper.id], lower.id)
				supports[lower.id] = append(supports[lower.id], upper.id)
			}
		}
	}
	return supportedBy, supports
}


type block struct {
	id string
	x1, y1, z1, x2, y2, z2 int
}
func (b block) String() string {
	return fmt.Sprintf("%s(%d,%d,%d~%d,%d,%d)", b.id, b.x1, b.y1, b.z1, b.x2, b.y2, b.z2)
}
func (b block) Size() int {
	return (b.x2 - b.x1 + 1) * (b.y2 - b.y1 + 1) * (b.z2 - b.z1 + 1)
}
func (b block) GetXYPoints() mapset.Set[Point] {
	points := mapset.NewSet[Point]()
	for x := b.x1; x <= b.x2; x++ {
		for y := b.y1; y <= b.y2; y++ {
			points.Add(Point{X:x, Y: y})
		}
	}
	return points
}
func (b block) MinZ() int {
	return b.z1
}
func (b *block) drop(z int) {
	delta := z - b.z1
	b.z1 += delta
	b.z2 += delta
}

func CreateBlockFromString(s string) block {
	parts := strings.Split(s, "~")
	firstPoint := strings.Split(parts[0], ",")
	secondPoint := strings.Split(parts[1], ",")
	x1, _ := strconv.Atoi(firstPoint[0])
	y1, _ := strconv.Atoi(firstPoint[1])
	z1, _ := strconv.Atoi(firstPoint[2])
	x2, _ := strconv.Atoi(secondPoint[0])
	y2, _ := strconv.Atoi(secondPoint[1])
	z2, _ := strconv.Atoi(secondPoint[2])

	//Make sure the first point is the smallest
	if (x1 > x2) {
		x1, x2 = x2, x1
	}
	if (y1 > y2) {
		y1, y2 = y2, y1
	}
	if (z1 > z2) {
		z1, z2 = z2, z1
	}

	id := getNextId()
	b := block{id, x1, y1, z1, x2, y2, z2}

	return b
}

var idCounter = 0
func getNextId() string {
	idCounter++
    id := ""
    num := idCounter
    for num > 0 {
        num-- // Adjust base-26 to be 0-25 instead of 1-26
        digit := num % 26
        id = string(rune('A')+rune(digit)) + id
        num /= 26
    }
    return id
}
