# Advent of Code 2023: Day 20: Pulse Propagation

Visit the problem statement at [Advent of Code 2023 Day 20](https://adventofcode.com/2023/day/20).

## Problem Statement

Today is another workflow like problem.  We have a series of modules:
- Flip-flop modules (prefix %):
  - are either on or off, initially it's off
  - if it receives a high pulse it does nothing
  - if it receives a low pulse it toggles the state and
    - if it was off it sends a high pulse
    - if it was on it sends a low pulse

- Conjunction modules (prefix &):
  - remembers the state of all the components that feed it pulses
  - defaults to remembering a low pulse
  - when it receives a pulse it updates the remembered state and:
    - if all the remembered states are high sends a low pulse
    - otherwise a high pulse
- Broadcaster module:
  - forwards the pulse to all connected components
- Button:
  - sends a low pulse to the broadcaster component

Pulses are always processed in order they are received.

## Part 1

For part 1, we need to count the number of high and low pulses sent if we press the button 1000 times.

### Solution Overview

First is parsing all the components and connecting them up.  I created a map of module objects that contain:
```go
type Module struct {
	name    string // name for debugging
	mType   string // type of module (%,&,broadcaster)
	state   int //the state for the flip flop
	outputs []string //where it connects to
}
```

I also created a `map[string]map[string]int` to track the conjunction inputs.

Then using a queue I processed each signal sent starting with the button press.  Every time a low or high pulse is sent I updated a counter.

## Part 2

In Part 2, we need to determine the number of button presses required to send a low signal to the `rx` module.

### Solution Overview

Initially, I let the solution for Part 1 run overnight, but it didn't yield an answer. This led me to dig deeper into the problem.

Upon examining the input data, I noticed that the `rx` module is fed by a single conjunction module, which I refer to as the `rxFeeder`. This module sends a low signal when all its inputs last sent a high signal.

The `rxFeeder` is, in turn, fed by several other conjunction modules, which I call `rxFeederInputs`. After letting the Part 1 solution run for a while and observing the modules connected to the `rxFeeder`, I found that each `rxFeederInput` sends a high signal on a set cycle.

I tracked when the `rxFeederInputs` sent a high signal and recorded the button press number. When all of them were set, I multiplied all the values together, similar to calculating the Least Common Multiple (LCM). However, I found that submitting the answer without calculating the LCM worked.

## Challenges and Learnings

This problem was challenging. Implementing the workflow was tricky, and I spent a significant amount of time getting the conjunction module to work correctly. The cycle detection wasn't too difficult once I had a fresh perspective on the problem. This problem reminded me of [Day 14](../day14/README.md) where we dealt with cycles, and [Day 08](../day08/README.md) where we worked with LCM.
