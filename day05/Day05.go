package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"regexp"
	"strconv"
	"bufio"
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
	seeds := GetSeeds(input)
	seedToSoilMap, _ := GetMap(input, "seed-to-soil")
	soilToFertilizerMap, _ := GetMap(input, "soil-to-fertilizer")
	fertilizerToWaterMap, _ := GetMap(input, "fertilizer-to-water")
	waterToLightMap, _ := GetMap(input, "water-to-light")
	lightToTemperatureMap, _ := GetMap(input, "light-to-temperature")
	temperatureToHumidityMap, _ := GetMap(input, "temperature-to-humidity")
	humidityToLocationMap, _ := GetMap(input, "humidity-to-location")

	mappings := make([]int, len(seeds))
	for i, seed := range seeds {
		soil := seedToSoilMap(seed)
		fertilizer := soilToFertilizerMap(soil)
		water := fertilizerToWaterMap(fertilizer)
		light := waterToLightMap(water)
		temp := lightToTemperatureMap(light)
		humidity := temperatureToHumidityMap(temp)
		location := humidityToLocationMap(humidity)
		mappings[i] = location
	}

	return min(mappings)
}


func SolvePart2(input string) int {
	seedToSoilMap, _ := GetMap(input, "seed-to-soil")
	soilToFertilizerMap, _ := GetMap(input, "soil-to-fertilizer")
	fertilizerToWaterMap, _ := GetMap(input, "fertilizer-to-water")
	waterToLightMap, _ := GetMap(input, "water-to-light")
	lightToTemperatureMap, _ := GetMap(input, "light-to-temperature")
	temperatureToHumidityMap, _ := GetMap(input, "temperature-to-humidity")
	humidityToLocationMap, _ := GetMap(input, "humidity-to-location")

	mappings := make([]int, 0)
	cache := make(map[int]int)

	seeds := GetSeeds(input)
	for seed := range GeneratePart2Seeds(seeds) {
		if _, ok := cache[seed]; ok {
			//Already calculated so we don't need to add it to the list
		} else {
			soil := seedToSoilMap(seed)
			fertilizer := soilToFertilizerMap(soil)
			water := fertilizerToWaterMap(fertilizer)
			light := waterToLightMap(water)
			temp := lightToTemperatureMap(light)
			humidity := temperatureToHumidityMap(temp)
			location := humidityToLocationMap(humidity)
			mappings = append(mappings, location)
			cache[seed] = location
		}
	}
	return min(mappings)
}

// MapValue maps a value to another value using a list of mapping functions
func MapValue(x int, mappingFuncs []func(int) int) int {
	for _, mappingFunc := range mappingFuncs {
		newX := mappingFunc(x)
		//Don't want to double map them
		if newX != x {
			return newX
		}
	}
	return x
}

func CompositeMapFunc(mappingFuncs []func(int) int) func(int) int {
	return func (x int) int {
		return MapValue(x, mappingFuncs)
	}
}

// CreateMapFunc creates a new function that maps a value to another value.
// if the value falls outside the range of the source, the value is returned
func CreateMapFunc(destStart int, srcStart int, len int) func(int) int {
	return func (x int) int {
		if x < srcStart || x >= srcStart + len {
			return x
		}
		return destStart + (x - srcStart)
	}
}

func GetSeeds(input string) []int {
    re := regexp.MustCompile(`seeds: ([\d\s]+)`)
    match := re.FindStringSubmatch(input)
    if len(match) > 1 {
        stringSeeds := strings.Fields(match[1])
		intSeeds := make([]int, len(stringSeeds))
		for i, seed := range stringSeeds {
			intSeeds[i], _ = strconv.Atoi(seed)
		}
		return intSeeds
    }
    log.Panic("No seeds found")
	return nil
}

func GeneratePart2Seeds(seeds []int) <- chan int {
	if (len(seeds) % 2 != 0) {
		log.Panic("Invalid number of seeds")
	}

	starts := make([]int, len(seeds) / 2)
	lengths := make([]int, len(seeds) / 2)
	for i := 0; i < len(seeds); i += 2 {
		starts[i / 2] = seeds[i]
		lengths[i / 2] = seeds[i + 1]
	}
	count := 0
	ch:= make(chan int)
    go func() {
        for x := range starts {
            for i := 0; i < lengths[x]; i++ {
				if count % 100000 == 0 {
					fmt.Printf("Generated %d seeds\n", count)
				}
				count += 1
                ch <- starts[x] + i
            }
        }
        close(ch)
    }()
    return ch
}

func GetMap(input string, mapName string) (func(int) int, error) {
    scanner := bufio.NewScanner(strings.NewReader(input))
    reading := false
    var result []func(int) int

    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, mapName) {
            reading = true
            continue
        }
        if reading && strings.TrimSpace(line) == "" {
            break
        }
        if reading {
            parts := strings.Fields(line)
            if len(parts) != 3 {
                return nil, fmt.Errorf("invalid line: %s", line)
            }
            row := make([]int, len(parts))
            for i, part := range parts {
                num, err := strconv.Atoi(part)
                if err != nil {
                    return nil, err
                }
                row[i] = num
            }
            result = append(result, CreateMapFunc(row[0], row[1], row[2]))
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return CompositeMapFunc(result), nil
}

func min(slice []int) int {
    min := slice[0]
    for _, value := range slice {
        if value < min {
            min = value
        }
    }
    return min
}
