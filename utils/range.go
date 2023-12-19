package utils

import (
	"fmt"
)

type Range struct {
    Start, End int
}
func (r Range) Contains(n int) bool {
    return n >= r.Start && n <= r.End
}
func NewRange(start, end int) Range {
    return Range{Start: start, End: end}
}
func (r Range) String() string {
	return fmt.Sprintf("[%d-%d]", r.Start, r.End)
}
func (r Range) IsValid() bool {
	return r.Start < r.End
}
func (r Range) Length() int {
	return r.End - r.Start + 1
}
func (r Range) Equals(other Range) bool {
	return r.Start == other.Start && r.End == other.End
}
func (r Range) Clone() Range {
	return NewRange(r.Start, r.End)
}
func (r Range) Intersects(other Range) bool {
    return r.Start < other.End && r.End > other.Start
}
func (r Range) Intersection(other Range) Range {
    if !r.Intersects(other) {
        return NewRange(0, 0)
    }
    return Range{
        Start: max(r.Start, other.Start),
        End:   min(r.End, other.End),
    }
}
func (r Range) Subtract(other Range) []Range {
	ranges := make([]Range, 0)
	if !r.Intersects(other) {
		return []Range{NewRange(r.Start, r.End)}
	}
	if r.Start < other.Start {
		ranges = append(ranges, NewRange(r.Start, other.Start-1))
	}
	if r.End > other.End {
		ranges = append(ranges, NewRange(other.End+1, r.End))
	}
	return ranges
}
