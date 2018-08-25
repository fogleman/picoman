package picoman

import "strings"

type Model struct {
	Grid     *Grid
	Position Point
	Target   Point
	Agents   []Agent
	Rock     bool
}

func NewModel(grid *Grid, position, target Point) *Model {
	return &Model{grid, position, target, nil, false}
}

func (m *Model) Copy() *Model {
	agents := make([]Agent, len(m.Agents))
	copy(agents, m.Agents)
	return &Model{m.Grid, m.Position, m.Target, agents, m.Rock}
}

func (m *Model) ZobristHash() uint64 {
	var h uint64
	h ^= ZobristPlayer[m.Position.Index(m.Grid.W)]
	for i, a := range m.Agents {
		if a.Active {
			h ^= ZobristAgent[i][a.Position.Index(m.Grid.W)][a.Heading.Index()]
		}
	}
	if m.Rock {
		h |= ZobristRock
	}
	return h
}

func (m *Model) Win() bool {
	return m.Position == m.Target
}

func (m *Model) ActiveAgents() int {
	activeCount := 0
	for _, a := range m.Agents {
		if a.Type == Rock {
			continue
		}
		if a.Active {
			activeCount++
		}
	}
	return activeCount
}

func (m *Model) AddAgent(a *Agent) {
	m.Agents = append(m.Agents, *a)
}

func (m *Model) CanMove(d Direction) bool {
	if m.Rock {
		return m.Grid.CanMoveRock(m.Position, d)
	}
	return m.Grid.CanMove(m.Position, d)
}

func (m *Model) DoMove(d Direction) bool {
	nextRock := false
	skipMove := make([]bool, len(m.Agents))

	// move player
	if m.Rock {
		p := m.Position.Add(d.Offset())
		for i := range m.Agents {
			a := &m.Agents[i]
			if !a.Active {
				continue
			}
			if a.Fixed {
				continue
			}
			d := p.Sub(a.Position).Abs()
			if d.X <= 1 && d.Y <= 1 {
				a.Target = p
				if a.Position != a.Target {
					a.Heading = m.Grid.DirectionTo(a.Position, a.Target)
				}
				skipMove[i] = true
			}
		}
	} else {
		m.Position = m.Position.Add(d.Offset())
	}

	plant := false
	for _, a := range m.Agents {
		if a.Type == Plant && a.Position == m.Position {
			plant = true
			break
		}
	}

	if !plant {
		// kill agents
		for i := range m.Agents {
			a := &m.Agents[i]
			if !a.Active {
				continue
			}
			if a.Position == m.Position {
				a.Active = false
				if a.Type == Rock {
					nextRock = true
				}
			}
		}

		// kill player
		for _, a := range m.Agents {
			if !a.Active {
				continue
			}
			if !m.Grid.CanMove(a.Position, a.Heading) {
				continue
			}
			p := a.Position.Add(a.Heading.Offset())
			if p == m.Position {
				return false
			}
		}
	}

	// move agents
	if !nextRock {
		for i := range m.Agents {
			a := &m.Agents[i]
			if !a.Active {
				continue
			}
			if a.Fixed {
				continue
			}
			if skipMove[i] {
				continue
			}
			if a.Position != a.Target {
				a.Heading = m.Grid.DirectionTo(a.Position, a.Target)
				a.Position = m.Grid.NextTo(a.Position, a.Target)
				if a.Position != a.Target {
					a.Heading = m.Grid.DirectionTo(a.Position, a.Target)
				}
			} else if a.Type == Pacing {
				if !m.Grid.CanMove(a.Position, a.Heading) {
					a.Heading = a.Heading.Opposite()
				}
				a.Position = a.Position.Add(a.Heading.Offset())
				if !m.Grid.CanMove(a.Position, a.Heading) {
					a.Heading = a.Heading.Opposite()
				}
				a.Target = a.Position
			}
		}
	}

	m.Rock = nextRock
	return true
}

func (m *Model) String() string {
	g := m.Grid
	var cells [][]string
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
			cells = append(cells, row)
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
		cells = append(cells, row)
	}

	p := m.Position
	cells[p.Y*2][p.X*2] = "*"

	for _, a := range m.Agents {
		if !a.Active {
			continue
		}
		p := a.Position
		if a.Type == Rock {
			cells[p.Y*2][p.X*2] = "R"
		}
		if a.Type == Waypoint {
			cells[p.Y*2][p.X*2] = "W"
		}
		if a.Type == Plant {
			cells[p.Y*2][p.X*2] = "P"
		}
		switch a.Heading {
		case N:
			cells[p.Y*2][p.X*2] = "^"
		case E:
			cells[p.Y*2][p.X*2] = ">"
		case S:
			cells[p.Y*2][p.X*2] = "v"
		case W:
			cells[p.Y*2][p.X*2] = "<"
		}

	}

	rows := make([]string, len(cells))
	for i := range rows {
		rows[i] = strings.Join(cells[i], "")
	}
	return strings.Join(rows, "\n")
}
