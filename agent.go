package picoman

type AgentType int

const (
	Stationary AgentType = iota
	Rock
	// Pacing
	// Rotating
)

type Agent struct {
	Type     AgentType
	Active   bool
	Position Point
	Heading  Direction
	Target   Point
}

func NewStationaryAgent(p Point, d Direction) *Agent {
	a := Agent{}
	a.Type = Stationary
	a.Active = true
	a.Position = p
	a.Target = p
	a.Heading = d
	return &a
}

func NewRockAgent(p Point) *Agent {
	a := Agent{}
	a.Type = Rock
	a.Active = true
	a.Position = p
	a.Target = p
	return &a
}
