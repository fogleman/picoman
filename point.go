package picoman

type Point struct {
	X, Y int
}

func (p Point) Index(w int) int {
	return p.Y*w + p.X
}

func (a Point) Add(b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}
