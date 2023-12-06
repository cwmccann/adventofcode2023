package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	text := string(input)
	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}


func SolvePart1(input string) int {
	sectionsAsStrings := strings.Split(input, "\n\n")
	seeds := getSeeds(sectionsAsStrings[0])

	//Remove the seeds section
	sectionsAsStrings = sectionsAsStrings[1:]

	//Remove the first line withe the name of the map
	sections := parseSections(sectionsAsStrings)

	mapped := make([]int, len(seeds))

	for i, seed := range seeds {
		mapped[i] = seed

		for _, section := range sections {

			for _, mapping := range section {
				if mapping.Source.Contains(mapped[i]) {
					mapped[i] = mapping.Apply(mapped[i])
					break
				}
			}
		}
	}
	return minIntSlice(mapped)
}

func SolvePart2(input string) int {
	sectionsAsStrings := strings.Split(input, "\n\n")

	//Seeds are now ranges!
	seedValues := getSeeds(sectionsAsStrings[0])
	seeds := make([]Range, len(seedValues) / 2)
	for i := 0; i < len(seedValues); i += 2 {
		seeds[i / 2] = NewRange(seedValues[i], seedValues[i] + seedValues[i + 1])
	}

	//Remove the seeds section
	sectionsAsStrings = sectionsAsStrings[1:]

	//Remove the first line withe the name of the map
	sections := parseSections(sectionsAsStrings)

	for _, section := range sections {
		new := make([]Range, 0)
		for len(seeds) > 0 {
			//Pop the last item off the seeds to process
			seed := seeds[len(seeds)-1]
			seeds = seeds[:len(seeds)-1]

			found := false
			for _, mapping := range section {
				if mapping.Source.Intersects(seed) {
					found = true
					intersection, _ := mapping.Source.Intersection(seed)

					//Map the intersection of the range to the destination
					mappedStart := mapping.Apply(intersection.Start)
					mappedEnd := mapping.Apply(intersection.End)
					new = append(new, NewRange(mappedStart, mappedEnd))

					//Add any other parts of the range that weren't mapped to new seeds
					if (seed.Start < intersection.Start) {
						seeds = append(seeds, NewRange(seed.Start, intersection.Start - 1))
					}
					if (intersection.End < seed.End) {
						seeds = append(seeds, NewRange(intersection.End, seed.End))
					}

					break
				}

			}
			//Pass through, we didn't find any mapping
			if (!found) {
				new = append(new, seed)
			}
		}
		seeds = new
	}
	//fmt.Println("Final seeds", seeds)
	return minRange(seeds)
}

func parseSections(sectionsAsStrings []string) [][]Mapping {
	allMappings := make([][]Mapping, len(sectionsAsStrings))
	for i, section := range sectionsAsStrings {
		lines := strings.Split(section, "\n")

		lines = lines[1:]

		sectionMapping := make([]Mapping, len(lines))
		for j, line := range lines {
			if (strings.Trim(line, " ") == "") {
				continue
			}
			tokens := strings.Fields(line)
			dest, _ := strconv.Atoi(tokens[0])
			source, _ := strconv.Atoi(tokens[1])
			length, _ := strconv.Atoi(tokens[2])
			sectionMapping[j] = Mapping{Source: NewRange(source, source + length), Destination: NewRange(dest, dest + length)}
		}
		allMappings[i] = sectionMapping
	}
	return allMappings
}

func getSeeds(inputLine string) []int {
	tokens := strings.Fields(inputLine)

	//remove the first token `seeds:`
	tokens = tokens[1:]

	seeds := make([]int, len(tokens))
	for i, token := range tokens {
		seeds[i], _ = strconv.Atoi(token)
	}
	return seeds
}

func minIntSlice(slice []int) int {
    min := slice[0]
    for _, value := range slice[1:] {
        if value < min {
            min = value
        }
    }
    return min
}
func minRange(ranges []Range) int {
	min := ranges[0].Start
	for i, r := range ranges {
		if r.Start == 0 {
			fmt.Println("Found 0 at index", i)
		}
		if r.Start < min {
			min = r.Start
		}
	}
	return min
}

type Range struct {
    Start, End int
}
func (r Range) Contains(n int) bool {
    return n >= r.Start && n <= r.End
}
func NewRange(start, end int) Range {
    return Range{Start: start, End: end}
}
func (r Range) Intersects(other Range) bool {
    return r.Start < other.End && r.End > other.Start
}
func (r Range) Intersection(other Range) (Range, error) {
    if !r.Intersects(other) {
        return Range{}, fmt.Errorf("ranges do not intersect")
    }
    return Range{
        Start: max(r.Start, other.Start),
        End:   min(r.End, other.End),
    }, nil
}



type Mapping struct {
    Source      Range
    Destination Range
}
func (m Mapping) String() string {
    return fmt.Sprintf("Source: %+v, Destination: %+v", m.Source, m.Destination)
}
func (m Mapping) Apply(n int) int {
	if m.Source.Contains(n) {
		return n - m.Source.Start + m.Destination.Start
	}
	return n
}


func minInt(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func maxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func countDuplicates(ranges []Range) int {
    counts := make(map[Range]int)
    for _, r := range ranges {
        key := Range{r.Start, r.End}
        counts[key]++
    }

    duplicates := 0
    for _, count := range counts {
        if count > 1 {
            duplicates += count - 1
        }
    }
    return duplicates
}
