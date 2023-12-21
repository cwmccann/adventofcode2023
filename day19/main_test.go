package main

import (
	"adventofcode2023/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)
type TestCase = utils.TestCase

var input =
`
px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}
`

func TestWorkflowApplyPartRange(t *testing.T) {
    assert := assert.New(t)

    tests := []struct {
        name        string
        workflowDef []string
        expected    int
    } {
        {
            name:        "less than",
            workflowDef: []string{"in{s<1000:A,R}"},
            expected:    999*4000*4000*4000,
        },
        {
            name:        "greater than",
            workflowDef: []string{"in{s>1000:A,R}"},
            expected:    3000*4000*4000*4000,
        },
        {
            name:        "combined",
            workflowDef: []string{"in{s>1000:ppr,R}", "ppr{x<1000:A,R}"},
            expected:    3000*999*4000*4000,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            workflows := parseWorkflows(tt.workflowDef)
            partRange := NewPartRange()
            assert.Equal(tt.expected, CountAccepted(workflows, partRange, "in"), tt.name)
        })
    }
}

func TestPart1(t *testing.T) {
    tests := []utils.TestCase{
        {
            Name:     "Test 1",
            Input:    input,
            Expected: 19114,
        },
        // Add more test cases here
    }

    assert := assert.New(t)
    for _, tt := range tests {
        t.Run(tt.Name, func(t *testing.T) {
            assert.Equal(tt.Expected, SolvePart1(tt.Input), tt.Name)
        })
    }
}

func TestPart2(t *testing.T) {
    tests := []utils.TestCase{
        {
            Name:     "Test 1",
            Input:    input,
            Expected: 167409079868000,
        },
        // Add more test cases here
    }

    assert := assert.New(t)
    for _, tt := range tests {
        t.Run(tt.Name, func(t *testing.T) {
            assert.Equal(tt.Expected, SolvePart2(tt.Input), tt.Name)
        })
    }
}



