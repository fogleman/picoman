package picoman

type Direction int

const (
	N Direction = 1 << iota
	E
	S
	W
)

var Directions = []Direction{N, E, S, W}

func (d Direction) String() string {
	switch d {
	case N:
		return "N"
	case E:
		return "E"
	case S:
		return "S"
	case W:
		return "W"
	}
	return "?"
}

func (d Direction) Index() int {
	switch d {
	case N:
		return 0
	case E:
		return 1
	case S:
		return 2
	case W:
		return 3
	}
	return 0
}

func (d Direction) Opposite() Direction {
	switch d {
	case N:
		return S
	case E:
		return W
	case S:
		return N
	case W:
		return E
	}
	return N
}

func (d Direction) Offset() Point {
	switch d {
	case N:
		return Point{0, -1}
	case E:
		return Point{1, 0}
	case S:
		return Point{0, 1}
	case W:
		return Point{-1, 0}
	}
	return Point{}
}
