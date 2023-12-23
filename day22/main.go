package main

import (
	"adventofcode2023/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"

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
	lines := strings.Split(input, "\n")
	blocks := make([]block, 0)
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		blocks = append(blocks, CreateBlockFromString(line))
	}

	//Sort the blocks by the min z value
	sort.Slice(blocks, func(a, b int) bool {
		return blocks[a].MinZ() < blocks[b].MinZ()
	})


	//Sink all the blocks to the ground
	for i, block1 := range blocks {
		fmt.Printf("block1 %d: %s\n", i, block1)

		minZ := 1
		//It's as low as it can go
		if block1.z1 >= minZ {
			continue
		}

		for j := i - 1; j >= 0; j-- {
			block2 := blocks[j]
			fmt.Printf("block2 %d: %s\n", j, block2)

			if block1.GetXYPoints().Intersect(block2.GetXYPoints()).Cardinality() > 0 {
				minZ = utils.Max(block1.z2 + 1, minZ)
			}
		}

		fmt.Printf("minZ: %d\n", minZ)
		blocks[i].drop(minZ)
	}

	fmt.Println("After sinking")
	for i, block := range blocks {
		fmt.Printf("block %d: %s\n", i, block)
	}

	// //Make a map of all the blocks and which ones overlap in the xy plane
	// blockMap := make(map[int][]block)
	// for i, block1 := range blocks {
	// 	for j, block2 := range blocks {
	// 		if i == j {
	// 			continue
	// 		}
	// 		if block1.GetXYPoints().Intersect(block2.GetXYPoints()).Cardinality() > 0 {
	// 			blockMap[i] = append(blockMap[i], block2)
	// 		}
	// 		//Todo add the reverse as well
	// 	}
	// 	//Sort the blocks by the min z value
	// 	sort.Slice(blockMap[i], func(a, b int) bool {
	// 		return blockMap[i][a].MinZ() < blockMap[i][b].MinZ()
	// 	})
	// }

	// for k, blocks := range blockMap {
	// 	fmt.Printf("block %d has %d overlapping blocks\n", k, len(blocks))
	// }

	// canRemove := make(map[int]bool)
	// //Check if there are any blocks above this block
	// for i, block1 := range blocks {
	// 	overlappingBlocks := blockMap[i]
	// 	hasOneAbove := false
	// 	for _, block2 := range overlappingBlocks {
	// 		if block1.z2 < block2.z1 {
	// 			//Now is is the only one supporting this block?
	// 		}
	// 	}




	return -1;
}

func SolvePart2(input string) int {
	return -1;
}

type block struct {
	x1, y1, z1, x2, y2, z2 int
}
func (b block) String() string {
	return fmt.Sprintf("%d,%d,%d~%d,%d,%d", b.x1, b.y1, b.z1, b.x2, b.y2, b.z2)
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
	delta := z - b.z2
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

	b := block{x1, y1, z1, x2, y2, z2}

	return b
}

