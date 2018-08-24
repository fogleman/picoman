package picoman

import (
	"math"
	"strings"
)

const Inf = math.MaxInt32

type Grid struct {
	W     int
	H     int
	Walls []int
	Dist  [][]int
	Next  [][]int
}

func NewGrid(w, h int) *Grid {
	walls := make([]int, w*h)
	return &Grid{w, h, walls, nil, nil}
}

func (g *Grid) Finalize() {
	g.floydWarshall()
}

func (g *Grid) AddWall(p Point, d Direction) {
	q := p.Add(d.Offset())
	i := p.Y*g.W + p.X
	j := q.Y*g.W + q.X
	g.Walls[i] |= int(d)
	if j >= 0 && j < len(g.Walls) {
		g.Walls[j] |= int(d.Opposite())
	}
}

func (g *Grid) CanMove(p Point, d Direction) bool {
	if d == W && p.X == 0 {
		return false
	}
	if d == E && p.X == g.W-1 {
		return false
	}
	if d == N && p.Y == 0 {
		return false
	}
	if d == S && p.Y == g.H-1 {
		return false
	}
	i := p.Y*g.W + p.X
	return g.Walls[i]&int(d) == 0
}

func (g *Grid) Distance(src, dst Point) int {
	i := src.Index(g.W)
	j := dst.Index(g.W)
	return g.Dist[i][j]
}

func (g *Grid) floydWarshall() {
	n := g.W * g.H

	dist := make([][]int, n)
	next := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		next[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dist[i][j] = Inf
			next[i][j] = -1
		}
	}

	for i := 0; i < n; i++ {
		p := Point{i % g.W, i / g.W}
		dist[i][i] = 0
		next[i][i] = i
		for _, d := range Directions {
			if g.CanMove(p, d) {
				j := p.Add(d.Offset()).Index(g.W)
				dist[i][j] = 1
				next[i][j] = j
			}
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					next[i][j] = next[i][k]
				}
			}
		}
	}

	g.Dist = dist
	g.Next = next
}

func (g *Grid) String() string {
	var rows []string
	for y := 0; y < g.H; y++ {
		if y != 0 {
			var row []string
			for x := 0; x < g.W; x++ {
				if g.CanMove(Point{x, y}, N) {
					row = append(row, "| ")
				} else {
					row = append(row, "  ")
				}
			}
			rows = append(rows, strings.Join(row, ""))
		}
		var row []string
		for x := 0; x < g.W; x++ {
			if x != 0 {
				if g.CanMove(Point{x, y}, W) {
					row = append(row, "-")
				} else {
					row = append(row, " ")
				}
			}
			row = append(row, "o")
		}
		rows = append(rows, strings.Join(row, ""))
	}
	return strings.Join(rows, "\n")
}
