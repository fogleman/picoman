package picoman

import (
	"strings"
)

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
	if m.Position != m.Target {
		return false
	}
	for _, a := range m.Agents {
		if a.Active {
			return false
		}
	}
	return true
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

	// move player
	if m.Rock {
		p := m.Position.Add(d.Offset())
		for i := range m.Agents {
			a := &m.Agents[i]
			if !a.Active {
				continue
			}
			if a.Type == Rock {
				continue
			}
			d := p.Sub(a.Position).Abs()
			if d.X <= 1 && d.Y <= 1 {
				a.Target = p
				if a.Position != a.Target {
					a.Heading = m.Grid.DirectionTo(a.Position, a.Target)
				}
			}
		}
	} else {
		m.Position = m.Position.Add(d.Offset())
	}

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

	// move agents
	for i := range m.Agents {
		a := &m.Agents[i]
		if !a.Active {
			continue
		}
		if a.Position != a.Target {
			a.Heading = m.Grid.DirectionTo(a.Position, a.Target)
			a.Position = m.Grid.NextTo(a.Position, a.Target)
			if a.Position != a.Target {
				a.Heading = m.Grid.DirectionTo(a.Position, a.Target)
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
			continue
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
