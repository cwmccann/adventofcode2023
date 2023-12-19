package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestRangeSubtraction(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
        name     string
        r1       Range
        r2       Range
        expected []Range
    }{
        {
            name:     "Test 1",
            r1:       NewRange(0, 10),
            r2:       NewRange(2, 8),
            expected: []Range{NewRange(0, 1), NewRange(9, 10)},
        },
        {
            name:     "Test 2",
            r1:       NewRange(0, 10),
            r2:       NewRange(0, 10),
            expected: []Range{},
        },
        {
            name:     "Test 3",
            r1:       NewRange(0, 10),
            r2:       NewRange(5, 15),
            expected: []Range{NewRange(0, 4)},
        },
		{
			name:     "Test 4",
			r1:       NewRange(0, 10),
			r2:       NewRange(0, 5),
			expected: []Range{NewRange(6, 10)},
		},
		{
			name:     "Test 5",
			r1:       NewRange(0, 10),
			r2:       NewRange(15, 20),
			expected: []Range{NewRange(0, 10)},
		},
		{
			name:     "Test 6",
			r1:       NewRange(0, 10),
			r2:       NewRange(0, 0),
			expected: []Range{NewRange(0, 10)},
		},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equal(tt.expected, tt.r1.Subtract(tt.r2))
        })
    }
}

func TestIsValid(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name     string
		r        Range
		expected bool
	}{
		{
			name:     "Valid",
			r:       NewRange(0, 10),
			expected: true,
		},
		{
			name:     "End before start",
			r:       NewRange(10, 0),
			expected: false,
		},
		{
			name:     "Empty Range",
			r:       NewRange(10, 10),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(tt.expected, tt.r.IsValid())
		})
	}
}

