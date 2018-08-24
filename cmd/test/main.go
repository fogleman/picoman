package main

import (
	"fmt"

	. "github.com/fogleman/picoman"
)

func Level4() *Model {
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

	return m
}

func Level5() *Model {
	g := NewGridFromDescription([]string{
		"o o o o",
		"  |    ",
		"o o-o-o",
		"      |",
		"o-o o-o",
		"|   | |",
		"o-o-o-o",
	})

	m := NewModel(g, Point{2, 3}, Point{1, 0})
	m.AddAgent(NewStationaryAgent(Point{2, 1}, E))
	m.AddAgent(NewRockAgent(Point{1, 2}))
	return m
}

func Level6() *Model {
	g := NewGridFromDescription([]string{
		"o o o o",
		"  |   |",
		"o-o-o-o",
		"|     |",
		"o-o-o o",
		"|   | |",
		"o-o-o o",
	})

	m := NewModel(g, Point{2, 3}, Point{3, 0})
	m.AddAgent(NewStationaryAgent(Point{1, 0}, S))
	m.AddAgent(NewStationaryAgent(Point{1, 1}, W))
	m.AddAgent(NewStationaryAgent(Point{3, 1}, W))
	m.AddAgent(NewStationaryAgent(Point{3, 3}, S))
	m.AddAgent(NewRockAgent(Point{0, 2}))
	m.AddAgent(NewRockAgent(Point{2, 2}))
	return m
}

func Level7() *Model {
	g := NewGridFromDescription([]string{
		"o-o-o-o-o",
		"|       |",
		"o-o-o-o-o",
		"  |   |  ",
		"o o-o-o o",
		"    |    ",
		"o o o o o",
	})

	m := NewModel(g, Point{2, 3}, Point{2, 0})
	m.AddAgent(NewRockAgent(Point{1, 2}))
	m.AddAgent(NewRockAgent(Point{2, 2}))
	m.AddAgent(NewRockAgent(Point{3, 2}))
	m.AddAgent(NewStationaryAgent(Point{0, 0}, E))
	m.AddAgent(NewStationaryAgent(Point{1, 0}, W))
	m.AddAgent(NewStationaryAgent(Point{3, 0}, W))
	m.AddAgent(NewStationaryAgent(Point{4, 0}, W))
	m.AddAgent(NewStationaryAgent(Point{1, 1}, S))
	m.AddAgent(NewStationaryAgent(Point{3, 1}, S))
	return m
}

func main() {
	m := Level7()
	fmt.Println(m)
	Search(m)
}
