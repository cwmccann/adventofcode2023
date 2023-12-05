package main

import (
	"testing"
	"reflect"
)

var input =
`
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

func TestMapFunc(t *testing.T) {
	testMap := map[int]int{
		10: 10,
		97: 97,
		98: 50,
		99: 51,
		100: 100,
	}

	map1 := CreateMapFunc(50, 98, 2)

	for k, v := range testMap {
		if map1(k) != v {
			t.Errorf("Expected %d, got %d", v, map1(k))
		}
	}
}

func TestMapValue(t *testing.T) {
	testMap := map[int]int{
		10: 10,
		53: 55,
		54: 56,
		97: 99,
		98: 50,
		99: 51,
	}

	map1 := CreateMapFunc(50, 98, 2)
	map2 := CreateMapFunc(52, 50, 48)
	mapFuncs := []func(int) int{map1, map2}

	for k, v := range testMap {
		if MapValue(k, mapFuncs) != v {
			t.Errorf("Expected %d, got %d", v, MapValue(k, mapFuncs))
		}
	}
}

func TestGetSeeds(t *testing.T) {
	expected := []int{79, 14, 55, 13}
	seeds := GetSeeds(input)
	if !reflect.DeepEqual(seeds, expected) {
        t.Errorf("Expected %v, got %v", expected, seeds)
    }
}

func TestGetMap(t *testing.T) {
	testMap := map[int]int{
		10: 10,
		53: 55,
		54: 56,
		97: 99,
		98: 50,
		99: 51,
	}

	mapping, err := GetMap(input, "seed-to-soil")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	for k, v := range testMap {
		if mapping(k) != v {
			t.Errorf("Expected %d, got %d", v, mapping(k))
		}
	}

}

func TestGeneratePart2Seeds(t *testing.T) {
	seeds := []int{10, 2}
	expected := []int{10, 11}
	i := 0
	for seed := range GeneratePart2Seeds(seeds) {
		if seed != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], seed)
		}
		i++
	}
}

func TestGeneratePart2SeedsMultiple(t *testing.T) {
	seeds := []int{10, 2, 20, 5}
	expected := []int{10, 11, 20, 21, 22, 23, 24}
	i := 0
	for seed := range GeneratePart2Seeds(seeds) {
		if seed != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], seed)
		}
		i++
	}
}

func TestGeneratePart2SeedsFromInput(t *testing.T) {
	seeds := GetSeeds(input)
	actual := []int{}
	for seed := range GeneratePart2Seeds(seeds) {
		actual = append(actual, seed)
	}
	if len(actual) != 27 {
		t.Errorf("Expected 27, got %d", len(actual))
	}
}


func TestDay01Part1(t *testing.T) {
	solution := SolvePart1(input);
	if solution != 35 {
		t.Errorf("Solution incorrect: %d", solution)
	}
}

func TestDay01Part2(t *testing.T) {
	solution := SolvePart2(input);
	if solution != 46 {
		t.Errorf("Solution incorrect: %d", solution)
	}
}



