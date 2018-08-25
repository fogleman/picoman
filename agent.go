package picoman

type AgentType int

const (
	Stationary AgentType = iota
	Pacing
	Rock
	Waypoint
	Plant
)

type Agent struct {
	Type     AgentType
	Active   bool
	Fixed    bool
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

func NewPacingAgent(p Point, d Direction) *Agent {
	a := Agent{}
	a.Type = Pacing
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
	a.Fixed = true
	return &a
}

func NewWaypointAgent(p Point) *Agent {
	a := Agent{}
	a.Type = Waypoint
	a.Active = true
	a.Position = p
	a.Target = p
	a.Fixed = true
	return &a
}

func NewPlantAgent(p Point) *Agent {
	a := Agent{}
	a.Type = Plant
	a.Active = false
	a.Position = p
	a.Target = p
	a.Fixed = true
	return &a
}
