package main

import ( "testing" )

var input =
`
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`
func TestIsGear(t *testing.T) {
	if !isGear('*') {
		t.Errorf("* is a gear, but isGear returned false")
	}
	if isGear('1') {
		t.Errorf("1 is not a gear, but isGear returned true")
	}
	if isGear('#') {
		t.Errorf("# is not a gear, but isGear returned true")
	}
}

func TestIsSymbol(t *testing.T) {
	if !isSymbol('*') {
		t.Errorf("* is a symbol, but isSymbol returned false")
	}
	if !isSymbol('+') {
		t.Errorf("+ is a symbol, but isSymbol returned false")
	}
	if !isSymbol('#') {
		t.Errorf("+ is a symbol, but isSymbol returned false")
	}
	if isSymbol('1') {
		t.Errorf("1 is not a symbol, but isSymbol returned true")
	}
	if isSymbol('.') {
		t.Errorf(". is not a symbol, but isSymbol returned true")
	}
}

func TestHasAdjacentSymbols(t *testing.T) {
	schematic := StringTo2DArray(input)
	//Top left corner 4
	if HasAdjacentSymbol(schematic, 0, 0, isSymbol).IsValid() {
		t.Errorf("Expected false, got true")
	}
	//It's also a gear
	if HasAdjacentSymbol(schematic, 0, 0, isGear).IsValid() {
		t.Errorf("Expected false, got true")
	}

	//Top right corner 7
	if !HasAdjacentSymbol(schematic, 0, 2, isSymbol).IsValid() {
		t.Errorf("Expected true, got false")
	}
	//Third Row 5
	if !HasAdjacentSymbol(schematic, 2, 3, isSymbol).IsValid() {
		t.Errorf("Expected true, got false")
	}
	//Third Row 6
	if !HasAdjacentSymbol(schematic, 2, 6, isSymbol).IsValid() {
		t.Errorf("Expected true, got false")
	}
	if HasAdjacentSymbol(schematic, 2, 6, isGear).IsValid() {
		char := schematic[2][6]
		t.Errorf("there is no gear here: %s", string(char))
	}
	//Third Row second 3
	if !HasAdjacentSymbol(schematic, 2, 7, isSymbol).IsValid() {
		t.Errorf("Expected true, got false")
	}
	//Third Row third 3
	if HasAdjacentSymbol(schematic, 2, 8, isSymbol).IsValid() {
		t.Errorf("Expected true, got false")
	}
}

func TestPart1(t *testing.T) {
	solution := SolvePart1(input);
	if solution != 4361 {
		t.Errorf("Solution incorrect: %d", solution)
	}
}

func TestPart2(t *testing.T) {
	solution := SolvePart2(input);
	if solution != 467835 {
		t.Errorf("Solution incorrect: %d", solution)
	}
}
