package main

import (
	"fmt"

	. "github.com/fogleman/picoman"
)

func main() {
	g := NewGrid(5, 4)
	g.AddWall(Point{2, 0}, S)
	g.AddWall(Point{2, 1}, E)
	g.AddWall(Point{2, 3}, N)
	g.AddWall(Point{4, 0}, S)
	g.AddWall(Point{4, 0}, W)
	g.AddWall(Point{4, 3}, N)
	g.AddWall(Point{4, 3}, W)
	g.AddWall(Point{4, 2}, W)
	g.Finalize()

	m := NewModel(g, Point{0, 0}, Point{4, 2})
	m.AddAgent(NewStationaryAgent(Point{3, 0}, S))
	m.AddAgent(NewStationaryAgent(Point{1, 1}, N))
	m.AddAgent(NewStationaryAgent(Point{2, 1}, W))
	m.AddAgent(NewStationaryAgent(Point{1, 2}, N))
	m.AddAgent(NewStationaryAgent(Point{2, 2}, W))
	m.AddAgent(NewStationaryAgent(Point{3, 2}, N))
	m.AddAgent(NewStationaryAgent(Point{0, 3}, E))

	fmt.Println(m)
	Search(m)
}
