# Advent of code 2023

I'm learning Go and thought to do the advent of code this year with a new language.

Let's see how far I get!  

**UPDATE**: I finished

## What is advent of Code

Advent of Code is an annual online event since 2015, that takes place from December 1st to December 25th. Each day, a new programming challenge is released for participants to solve. The challenges are themed around holiday-related stories and puzzles, and they become progressively more difficult as the event progresses.

Each day's puzzle is divided into two parts. The first part presents a problem that you need to solve. After you submit the correct answer for the first part, the second part of the puzzle is unlocked. The second part is a continuation or a twist on the first part, often requiring you to adapt your solution from the first part to meet new conditions or solve a more complex problem. This two-part structure allows for a progressive increase in difficulty and complexity, and often the second part of the puzzle provides a more challenging twist to the initial problem.

The event is designed to promote problem-solving and programming skills, and it's open to anyone, regardless of their skill level or experience. Participants can use any programming language to solve the challenges.

The problems cover a wide range of topics and often require participants to apply algorithms and data structures, making it a fun and educational event for learning and improving programming skills.


## Impressions

Participating in the Advent of Code has been a fun and challenging experience. The puzzles were definitely a stretch for me and have pushed me to think critically and creatively about problem-solving strategies.  I have also learned a bunch of new algorithms and formulas needed to solve the days problem.

The two-part structure of the puzzles has been particularly interesting. The twist in the second part often requires rethinking and adapting the solution from the first part, which has been a great exercise in flexible problem-solving.

Using Go for these challenges has also been enlightening. While the language's simplicity and efficiency are great, the lack of some built-in functions in the standard library was initially surprising. However, implementing these functions myself and creating a bank of utility functions has been good, reinforcing my understanding of fundamental programming concepts.

Overall, the Advent of Code has been a fantastic opportunity to learn Go and improve my programming skills and problem-solving abilities. I look forward to participating in future events and continuing to learn and grow as a developer.


## Project Structure

This project is organized by day, with each day's challenge having its own directory. The directory name corresponds to the day of the Advent of Code event (for example, `day01`, `day02`, etc.). 

Each directory contains the following files:

1. `main.go`: This is the Go script that contains the solution for that day's challenge. It includes both Part 1 and Part 2 of the challenge.  Note in the early days I has it numbered `day01.go`.  That changed when I moved to a common template.

2. `main_test.go`: This is the Go script that contains unit tests for the solution. It tests the solution using the provided example inputs and outputs.

3. `README.md`: This file provides a brief overview of the challenge for that day and discusses the approach taken to solve it.

Please note that the actual puzzle inputs (`input.txt`) are not included in this repository, in accordance with the Advent of Code's request not to share puzzle inputs.

In addition to the daily directories, there is a `utils` directory that contains utility functions that are used across multiple challenges. These functions are in a separate Go script and are imported as needed.

## Running the Code

To run the code for a specific day's challenge, navigate to the corresponding directory and use the `go run` command. This command runs the `main.go` script, which contains the solution for that day's challenge.

```bash
cd day01
go run .
```

To run the unit tests for a specific day's challenge, use the go test command in the corresponding directory. This command runs the main_test.go script, which contains the unit tests for the solution.

```bash
cd day01
go test
```

Please note that the actual puzzle inputs (input.txt) are not included in this repository, in accordance with the Advent of Code's request not to share puzzle inputs. You will need to download the puzzle input from the Advent of Code website and place it in the correct directory to run the code against the actual input.

### Using Aliases

For convenience, I have set up a couple of aliases to run the code and tests more quickly:

- `goc`: This alias is equivalent to `go run main.go`. It runs the `main.go` script, which contains the solution for that day's challenge.

- `got`: This alias is equivalent to `go test`. It runs the `main_test.go` script, which contains the unit tests for the solution.

To use these aliases, navigate to the directory for the day's challenge and enter the alias:

```bash
cd day01
goc  # to run the code
got  # to run the tests
```




## Solution Notes

I have notes for each days solutions on how I completed them.  Currently all but day 24 is written in Go.

<div style="display: flex; justify-content: space-between; width: 600px; margin: auto;">
<div style="width: 300px;">

- [Day 01: Trebuchet?!](day01/README.md)
- [Day 02: Cube Conundrum](day02/README.md)
- [Day 03: Gear Ratios](day03/README.md)
- [Day 04: Scratchcards](day04/README.md)
- [Day 05: Almanac Workflow](day05/README.md)
- [Day 06: Wait For It](day06/README.md)
- [Day 07: Camel Cards](day07/README.md)
- [Day 08: Haunted Wasteland](day08/README.md)
- [Day 09: Mirage Maintenance](day09/README.md)
- [Day 10: Pipe Maze](day10/README.md)
- [Day 11: Cosmic Expansion](day11/README.md)
- [Day 12: Hot Springs](day12/README.md)
- [Day 13: Point of Incidence](day13/README.md)

</div>
<div style="display: flex; justify-content: space-between;">
<div style="width: 300px;">

- [Day 14: Parabolic Reflector Dish](day14/README.md)
- [Day 15: Lens Library](day15/README.md)
- [Day 16: The Floor Will Be Lava](day16/README.md)
- [Day 17: Clumsy Crucible](day17/README.md)
- [Day 18: Lavaduct Lagoon](day18/README.md)
- [Day 19: Aplenty](day19/README.md)
- [Day 20: Pulse Propagation](day20/README.md)
- [Day 21: Step Counter](day21/README.md)
- [Day 22: Sand Slabs](day22/README.md)
- [Day 23: A Long Walk](day23/README.md)
- [Day 24: Never Tell Me The Odds](day24/README.md)
- [Day 25: Snowverload](day25/README.md)

</div>
</div>
