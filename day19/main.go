package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strconv"
	"strings"
)
type Range = utils.Range

func main() {
	text := utils.InputToString();

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}

type Part struct {
	x, m ,a, s int
}

type Rule struct {
	condition RuleCondition
	action string
}

type RuleCondition struct {
	field, operator string
	value int
}

type Workflow struct {
	rules []Rule
}

func SolvePart1(input string) int {
	sections := strings.Split(input, "\n\n")

	rulesStr := strings.Split(sections[0], "\n")
	partsStr := strings.Split(sections[1], "\n")

	workflows := parseWorkflows(rulesStr)
	parts := parseParts(partsStr)

	accepted := make([]Part, 0)

	for _, part := range parts {
		//All parts start in the "in" workflow
		workflow := workflows["in"]
		done := false
		for !done {
			// fmt.Println("Applying part to workflow", part, workflow.apply(part))
			result := workflow.apply(part)

			if result == "A" {
				accepted = append(accepted, part)
				done = true
			}
			if result == "R" {
				//rejected = append(rejected, part)
				done = true
			}
			workflow = workflows[result]
		}
	}

	//fmt.Println("Accepted:", accepted)

	total := 0
	for _, part := range accepted {
		total += part.x
		total += part.m
		total += part.a
		total += part.s
	}
	return total
}

func (w Workflow) apply(part Part) string {
	for _, rule := range w.rules {
		if rule.condition.field == "" {
			return rule.action
		}
		var partValue int

		switch rule.condition.field {
		case "x":
			partValue = part.x
		case "m":
			partValue = part.m
		case "a":
			partValue = part.a
		case "s":
			partValue = part.s
		default:
			panic("Invalid field: " + rule.condition.field)
		}

		if rule.condition.operator == "<" && partValue < rule.condition.value {
			return rule.action
		}
		if rule.condition.operator == ">" && partValue > rule.condition.value {
			return rule.action
		}
	}
	panic("No rule applied!")
}

func parseRuleCondition(s string) RuleCondition {
	var operator string
	var index int

	if strings.Contains(s, "<") {
		operator = "<"
		index = strings.Index(s, "<")
	} else if strings.Contains(s, ">") {
		operator = ">"
		index = strings.Index(s, ">")
	} else {
		panic("Invalid rule: " + s)
	}
	field := s[:index]
	value, _ := strconv.Atoi(s[index + 1:])

	return RuleCondition{field, operator, value}
}

func parseWorkflows(rulesStr []string) map[string]Workflow {
	workflows := make(map[string]Workflow)

	for _, ruleStr := range rulesStr {
		if strings.Trim(ruleStr, " ") == "" {
			continue
		}
		keyIndex := strings.Index(ruleStr, "{")
		key := ruleStr[:keyIndex]

		ruleStr = utils.RemoveAll(ruleStr[keyIndex:], "{}")
		parts := strings.Split(ruleStr, ",")
		rules := make([]Rule, 0)

		for _, part := range parts {
			if !strings.Contains(part, ":") {
				rules = append(rules, Rule{RuleCondition{}, part})
			} else {
				tokens := strings.Split(part, ":")
				ruleCondition := parseRuleCondition(tokens[0])
				action := tokens[1]
				rules = append(rules, Rule{ruleCondition, action})
			}
		}

		workflows[key] = Workflow{rules: rules}
	}

	return workflows
}

func parseParts(partsStr []string) []Part {
	parts := make([]Part, 0)

	for _, partStr := range partsStr {
		if strings.Trim(partStr, " ") == "" {
			continue
		}
		part := Part{}
		fmt.Sscanf(partStr, "{x=%d,m=%d,a=%d,s=%d}", &part.x, &part.m, &part.a, &part.s)
		parts = append(parts, part)
	}

	return parts
}

func Day19Range() Range {
	return Range{Start: 1, End: 4000}
}

type RangeMap map[string]Range
func NewPartRange() RangeMap {
	return RangeMap {
		"x": Day19Range(),
		"m": Day19Range(),
		"a": Day19Range(),
		"s": Day19Range(),
	}
}
func (r RangeMap) Clone() RangeMap {
	return RangeMap {
		"x": r["x"].Clone(),
		"m": r["m"].Clone(),
		"a": r["a"].Clone(),
		"s": r["s"].Clone(),
	}
}
func (r RangeMap) String() string {
	return fmt.Sprintf("{x=%v,m=%v,a=%v,s=%v}", r["x"], r["m"], r["a"], r["s"])
}

func SolvePart2(input string) int {
	fmt.Println("Solving Part 2")
	sections := strings.Split(input, "\n\n")
	rulesStr := strings.Split(sections[0], "\n")
	workflows := parseWorkflows(rulesStr)

	//accepted := make([]PartRange, 0)

	count := CountAccepted(workflows, NewPartRange(), "in")
	fmt.Println("Count:", count)
	return count
}

func CountAccepted(workflows map[string]Workflow, rangeMap RangeMap, workflow string) int {
	if workflow == "A" {
		fmt.Println("Accepted:", rangeMap)
		total := 1
		for _, v := range rangeMap {
			total *= (v.End - v.Start + 1)
		}
		return total
	}
	if workflow == "R" {
		return 0
	}

	total := 0
	w := workflows[workflow]

	for _, rule := range w.rules {
		//fmt.Println("Applying rule:", rule, " to part range:", rangeMap)

		field := rule.condition.field
		opStr := rule.condition.operator

		//Fallback rule
		if field == "" {
			total += CountAccepted(workflows, rangeMap, rule.action)
			continue
		}
		r := rangeMap[field]

		var op Range
		if opStr == "<" {
			op = lessThan(rule.condition.value)
		} else if opStr == ">" {
			op = greaterThan(rule.condition.value)
		} else {
			panic("Invalid operator: " + rule.condition.operator)
		}

		in := r.Intersection(op)
		out := r.Subtract(op)

		if in.IsValid() {
			//Copy the out range and apply the next workflow
			newMap := rangeMap.Clone()
			newMap[field] = in.Clone()
			total += CountAccepted(workflows, newMap, rule.action)
		}

		if len(out) == 1 && out[0].IsValid() {
			rangeMap[field] = out[0]
			rangeMap = rangeMap.Clone()
		}
	}
	return total
}

func lessThan(v int) Range {
	return Range{Start: 0, End: v - 1}
}

func greaterThan(v int) Range {
	return Range{Start: v + 1, End: 4000}
}

