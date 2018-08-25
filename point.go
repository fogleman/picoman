package picoman

type Point struct {
	X, Y int
}

func PointFromIndex(i, w int) Point {
	x := i % w
	y := i / w
	return Point{x, y}
}

func (p Point) Index(w int) int {
	return p.Y*w + p.X
}

func (p Point) Abs() Point {
	x, y := p.X, p.Y
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return Point{x, y}
}

func (a Point) Add(b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}

func (a Point) Sub(b Point) Point {
	return Point{a.X - b.X, a.Y - b.Y}
}
