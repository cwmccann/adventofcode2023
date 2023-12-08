package main

import ( "testing" )

var input =
`
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

func TestMakeRound(t *testing.T) {
	round := MakeRound("1 green, 3 red, 6 blue")
	if round.blue != 6 {
		t.Errorf("Blue incorrect: %d", round.blue)
	}
	if round.red != 3 {
		t.Errorf("Red incorrect: %d", round.red)
	}
	if round.green != 1 {
		t.Errorf("Green incorrect: %d", round.green)
	}
}

func TestMakeGame(t *testing.T) {
	game := MakeGame("Game 13: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	if game.id != 13 {
		t.Errorf("Game id incorrect: %d", game.id)
	}
	if len(game.rounds) != 3 {
		t.Errorf("Game rounds incorrect: %d", len(game.rounds))
	}
}

func TestPart1(t *testing.T) {
	solution := SolvePart1(input);
	if solution != 8 {
		t.Errorf("Solution incorrect: %d", solution)
	}
}

func TestPart2(t *testing.T) {
	solution := SolvePart2(input);
	if solution != 2286 {
		t.Errorf("Solution incorrect: %d", solution)
	}
}
