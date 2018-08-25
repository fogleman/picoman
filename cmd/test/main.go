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

func Level8() *Model {
	g := NewGridFromDescription([]string{
		"o-o-o-o  ",
		"  | | |  ",
		"  o-o o  ",
		"  | | |  ",
		"o-o o-o-o",
		"  | | |  ",
		"  o o o  ",
		"  |   |  ",
		"  o-o-o  ",
	})

	m := NewModel(g, Point{0, 2}, Point{4, 2})
	m.AddAgent(NewPacingAgent(Point{1, 1}, N))
	m.AddAgent(NewPacingAgent(Point{2, 1}, N))
	m.AddAgent(NewPacingAgent(Point{3, 3}, S))
	m.AddAgent(NewStationaryAgent(Point{0, 0}, W))
	return m
}

func Level9() *Model {
	g := NewGridFromDescription([]string{
		"    o-o",
		"    | |",
		"o-o o o",
		"  | |  ",
		"o-o-o  ",
		"| | |  ",
		"o-o-o  ",
	})

	m := NewModel(g, Point{2, 3}, Point{3, 1})
	m.AddAgent(NewStationaryAgent(Point{3, 0}, W))
	m.AddAgent(NewRockAgent(Point{0, 2}))
	m.AddAgent(NewRockAgent(Point{1, 1}))
	m.AddAgent(NewWaypointAgent(Point{0, 1}))
	return m
}

func Level10() *Model {
	g := NewGridFromDescription([]string{
		"    o    ",
		"    |    ",
		"o o o-o o",
		"| |   | |",
		"o o-o-o o",
		"|   |   |",
		"o-o-o-o-o",
	})

	m := NewModel(g, Point{2, 3}, Point{2, 0})
	m.AddAgent(NewRockAgent(Point{0, 1}))
	m.AddAgent(NewRockAgent(Point{4, 3}))
	m.AddAgent(NewPacingAgent(Point{3, 2}, W))
	m.AddAgent(NewPacingAgent(Point{2, 1}, N))
	m.AddAgent(NewWaypointAgent(Point{1, 1}))
	return m
}

func Level11() *Model {
	g := NewGridFromDescription([]string{
		"  o-o-o  ",
		"  |      ",
		"o-o-o-o-o",
		"  | | | |",
		"o-o-o-o-o",
		"  | | |  ",
		"o-o-o-o  ",
	})

	m := NewModel(g, Point{0, 3}, Point{4, 1})
	m.AddAgent(NewPlantAgent(Point{2, 2}))
	m.AddAgent(NewPacingAgent(Point{2, 1}, W))
	m.AddAgent(NewPacingAgent(Point{1, 2}, W))
	m.AddAgent(NewPacingAgent(Point{3, 2}, E))
	m.AddAgent(NewWaypointAgent(Point{3, 0}))
	return m
}

func Level12() *Model {
	g := NewGridFromDescription([]string{
		"        o",
		"        |",
		"o-o-o-o-o",
		"|     | |",
		"o o-o-o o",
		"| | |   |",
		"o o-o   o",
	})

	m := NewModel(g, Point{2, 3}, Point{0, 3})
	m.AddAgent(NewPlantAgent(Point{1, 1}))
	m.AddAgent(NewPlantAgent(Point{2, 2}))
	m.AddAgent(NewPacingAgent(Point{4, 1}, W))
	m.AddAgent(NewPacingAgent(Point{1, 2}, E))
	m.AddAgent(NewWaypointAgent(Point{4, 0}))
	return m
}

func Level13() *Model {
	g := NewGridFromDescription([]string{
		"  o-o-o o",
		"  | | | |",
		"o-o o o o",
		"| | | | |",
		"o-o-o o-o",
		"| | | |  ",
		"o o-o-o  ",
	})

	m := NewModel(g, Point{0, 3}, Point{4, 0})
	m.AddAgent(NewPlantAgent(Point{2, 0}))
	m.AddAgent(NewPlantAgent(Point{3, 3}))
	m.AddAgent(NewPacingAgent(Point{1, 0}, S))
	m.AddAgent(NewPacingAgent(Point{2, 3}, N))
	m.AddAgent(NewPacingAgent(Point{4, 1}, N))
	m.AddAgent(NewRockAgent(Point{3, 1}))
	m.AddAgent(NewWaypointAgent(Point{2, 2}))
	return m
}

func Level14() *Model {
	g := NewGridFromDescription([]string{
		"  o-o-o  ",
		"  | |    ",
		"  o-o-o-o",
		"    |    ",
		"o-o-o-o-o",
		"|     | |",
		"o-o-o-o o",
		"| | |   |",
		"o-o-o-o o",
	})

	m := NewModel(g, Point{3, 4}, Point{0, 2})
	m.AddAgent(NewPacingAgent(Point{3, 1}, W))
	m.AddAgent(NewPacingAgent(Point{3, 2}, E))
	m.AddAgent(NewPacingAgent(Point{2, 3}, E))
	m.AddAgent(NewStationaryAgent(Point{1, 3}, S))
	m.AddAgent(NewWaypointAgent(Point{1, 0}))
	return m
}

func Level15() *Model {
	g := NewGridFromDescription([]string{
		"o-o-o-o  ",
		"| |   |  ",
		"o o-o o-o",
		"| | | | |",
		"o-o o o o",
		"| | | | |",
		"o-o-o-o-o",
		"  | |    ",
		"o-o-o-o  ",
	})

	m := NewModel(g, Point{4, 2}, Point{0, 4})
	m.AddAgent(NewPacingAgent(Point{0, 2}, S))
	m.AddAgent(NewPacingAgent(Point{3, 1}, S))
	m.AddAgent(NewPacingAgent(Point{1, 3}, N))
	m.AddAgent(NewStationaryAgent(Point{2, 1}, W))
	m.AddAgent(NewStationaryAgent(Point{0, 3}, E))
	m.AddAgent(NewStationaryAgent(Point{2, 4}, W))
	m.AddAgent(NewPlantAgent(Point{2, 3}))
	return m
}

func main() {
	levels := []*Model{
		Level4(),
		Level5(),
		Level6(),
		Level7(),
		Level8(),
		Level9(),
		Level10(),
		Level11(),
		Level12(),
		Level13(),
		Level14(),
		Level15(),
	}

	for _, level := range levels {
		m := level
		// fmt.Println(m)
		// fmt.Println()

		moves := Search(m)
		fmt.Println(len(moves), moves)
		// fmt.Println()

		// for _, move := range moves {
		// 	m.DoMove(move)
		// }

		// fmt.Println(m)
		// fmt.Println()
	}

	// moves := []Direction{E, E, N, S}
	// for _, move := range moves {
	// 	m.DoMove(move)
	// 	fmt.Println(move)
	// 	fmt.Println(m)
	// 	fmt.Println()
	// }
}
