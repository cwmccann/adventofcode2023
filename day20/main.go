package main

import (
	"adventofcode2023/utils"
	"fmt"
	"github.com/gammazero/deque"
	"strings"
)

func main() {
	text := utils.InputToString()

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}

type Module struct {
	name    string
	mType   string
	state   int
	outputs []string
}

type ModuleWithInput struct {
	module Module
	input  int
	from   string
}

func SolvePart1(input string) int {
	moduleMap := parseInput(input)

	//Map of maps for the conjunction inputs
	conjunctionMap := make(map[string]map[string]int)
	//init the conjunction map
	for name, module := range moduleMap {
		if module.mType == "&" {
			conjunctionMap[name] = make(map[string]int)
		}
	}
	//populate the conjunction map
	for name, module := range moduleMap {
		for _, output := range module.outputs {
			if conjunctionMap[output] != nil {
				conjunctionMap[output][name] = 0
			}
		}
	}

	lowPulses := 0 //For the button push
	highPulses := 0

	q := deque.New[ModuleWithInput]()

	for i := 0; i < 1000; i++ {
		//Push the "button"
		q.PushBack(ModuleWithInput{module: moduleMap["broadcaster"], input: 0})
		lowPulses += 1
		//fmt.Printf("button -low-> broadcaster\n")

		for q.Len() > 0 {
			moduleWithInput := q.PopFront()
			m := moduleWithInput.module
			from := moduleWithInput.from
			name := m.name
			in := moduleWithInput.input
			out := in

			if m.mType == "broadcaster" {
				//Do nothing
			} else if m.mType == "%" {
				//Flip-flop modules (prefix %) are either on or off; they are initially off.
				//If a flip-flop module receives a high pulse, it is ignored and nothing happens.
				//However, if a flip-flop module receives a low pulse, it flips between on and off.
				//If it was off, it turns on and sends a high pulse. If it was on, it turns off and sends a low pulse.
				if in == 1 {
					continue
				}

				if m.state == 0 {
					m.state = 1
					out = 1
				} else {
					m.state = 0
					out = 0
				}
			} else if m.mType == "&" {
				//Conjunction modules (prefix &) remember the type of the most recent pulse received from each of their
				//connected input modules; they initially default to remembering a low pulse for each input.
				//When a pulse is received, the conjunction module first updates its memory for that input.
				//Then, if it remembers high pulses for all inputs, it sends a low pulse; otherwise, it sends a high pulse.

				//Get the map for this conjunction
				conjunctionMap[name][from] = in
				//Check if all inputs are high
				allHigh := true
				for _, value := range conjunctionMap[name] {
					if value == 0 {
						allHigh = false
						break
					}
				}

				if allHigh {
					out = 0
				} else {
					out = 1
				}
			}

			//Stash the state back into the map
			moduleMap[name] = m

			for _, output := range m.outputs {
				// outString := "low"
				// if out == 1 {
				// 	outString = "high"
				// }
				//fmt.Printf("%s -%s-> %s\n", name, outString, output)
				q.PushBack(ModuleWithInput{module: moduleMap[output], input: out, from: name})
				if out == 0 {
					lowPulses += 1
				} else {
					highPulses += 1
				}
			}
		}
	}

	return lowPulses * highPulses
}

func parseInput(input string) map[string]Module {
	modules := make(map[string]Module)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		name := strings.Trim(parts[0], " ")

		mType := ""
		if strings.Contains(name, "%") {
			mType = "%"
			name = strings.Trim(name, "%")
		} else if strings.Contains(name, "&") {
			mType = "&"
			name = strings.Trim(name, "&")
		} else {
			mType = "broadcaster"
		}

		out := strings.Trim(parts[1], " ")
		outputs := strings.Split(out, ", ")

		module := Module{
			name:    name,
			mType:   mType,
			state:   0,
			outputs: outputs,
		}
		modules[name] = module
	}
	return modules
}

func SolvePart2(input string) int {
	moduleMap := parseInput(input)

	//Map of maps for the conjunction inputs
	conjunctionMap := make(map[string]map[string]int)
	//init the conjunction map
	for name, module := range moduleMap {
		if module.mType == "&" {
			conjunctionMap[name] = make(map[string]int)
		}
	}
	//populate the conjunction map
	for name, module := range moduleMap {
		for _, output := range module.outputs {
			if conjunctionMap[output] != nil {
				conjunctionMap[output][name] = 0
			}
		}
	}

	//Find the element that feeds rx
	rxFeeder := findRXFeeder(moduleMap)

	//assert that the rxFeeder is a conjunction
	if moduleMap[rxFeeder].mType != "&" {
		panic("rxFeeder is not a conjunction")
	}

	//Find the inputs to the rxFeeder
	rxFeederInputs := make(map[string]int)
	for key := range conjunctionMap[rxFeeder] {
		rxFeederInputs[key] = -1
		if moduleMap[key].mType != "&" {
			panic("rxFeederInput is not a conjunction")
		}
	}

	i := 0
	q := deque.New[ModuleWithInput]()
	ans := 0

	for ans == 0 {
		//Push the "button"
		q.PushBack(ModuleWithInput{module: moduleMap["broadcaster"], input: 0})
		i += 1

		for q.Len() > 0 && ans == 0 {
			moduleWithInput := q.PopFront()
			m := moduleWithInput.module
			from := moduleWithInput.from
			name := m.name
			in := moduleWithInput.input
			out := in

			if m.mType == "broadcaster" {
				//Do nothing
			} else if m.mType == "%" {
				//Flip-flop modules (prefix %) are either on or off; they are initially off.
				//If a flip-flop module receives a high pulse, it is ignored and nothing happens.
				//However, if a flip-flop module receives a low pulse, it flips between on and off.
				//If it was off, it turns on and sends a high pulse. If it was on, it turns off and sends a low pulse.
				if in == 1 {
					continue
				}

				if m.state == 0 {
					m.state = 1
					out = 1
				} else {
					m.state = 0
					out = 0
				}
			} else if m.mType == "&" {
				//Conjunction modules (prefix &) remember the type of the most recent pulse received from each of their
				//connected input modules; they initially default to remembering a low pulse for each input.
				//When a pulse is received, the conjunction module first updates its memory for that input.
				//Then, if it remembers high pulses for all inputs, it sends a low pulse; otherwise, it sends a high pulse.

				//Set the state
				conjunctionMap[name][from] = in

				if AllHigh(conjunctionMap[name]) {
					out = 0
				} else {
					out = 1
				}
			}

			//Stash the state back into the map
			moduleMap[name] = m

			//Track when the feeders output a high pulse
			if out == 1 && rxFeederInputs[name] == -1 {
				rxFeederInputs[name] = i

				//Check if we have all the inputs
				allDone := true;
				for _, value := range rxFeederInputs {
					if value == -1 {
						allDone = false
						break
					}
				}

				if allDone {
					ans = 1
					for _, value := range rxFeederInputs {
						ans *= value
					}
					break;
				}
			}

			//Push the outputs onto the queue
			for _, output := range m.outputs {
				q.PushBack(ModuleWithInput{module: moduleMap[output], input: out, from: name})
			}
		}
	}
	return ans
}

func AllHigh(states map[string]int) bool {
	for _, value := range states {
		if value == 0 {
			return false
		}
	}
	return true
}

func AnyHigh(states map[string]int) bool {
	for _, value := range states {
		if value == 1 {
			return true
		}
	}
	return false
}


func findRXFeeder(moduleMap map[string]Module) string {
	for name, module := range moduleMap {
		for _, output := range module.outputs {
			if output == "rx" {
				return name
			}
		}
	}
	return ""
}



