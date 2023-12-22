package utils


type Point struct {
	X, Y int
}
func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}
func (p Point) Sub(other Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - other.Y}
}
func (p Point) Abs() Point {
	return Point{X: Abs(p.X), Y: Abs(p.Y)}
}
func (p Point) Manhattan() int {
	return Abs(p.X) + Abs(p.Y)
}
func (p Point) ManhattanTo(other Point) int {
	return p.Sub(other).Manhattan()
}
func (p Point) GetCardinalNeighbors() []Point {
	return []Point{
		{X: p.X, Y: p.Y - 1},
		{X: p.X, Y: p.Y + 1},
		{X: p.X - 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y},
	}
}

func (p Point) IsInGrid(rows int, cols int) bool {
	return p.X >= 0 && p.X < cols && p.Y >= 0 && p.Y < rows
}


