package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestReverse(t *testing.T) {
    tests := []struct {
        name string
        input []int
        expected []int
    }{
        {
            name: "Test 1",
            input: []int{1, 2, 3, 4, 5},
            expected: []int{5, 4, 3, 2, 1},
        },
        {
            name: "Test 2",
            input: []int{5, 4, 3, 2, 1},
            expected: []int{1, 2, 3, 4, 5},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Reverse(tt.input)
            assert.Equal(t, tt.expected, result)
        })
    }
}
