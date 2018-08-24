package picoman

import "math/rand"

const (
	MaxCells  = 64
	MaxAgents = 16
)

var (
	ZobristPlayer [MaxCells]uint64
	ZobristAgent  [MaxAgents][MaxCells][4]uint64
)

func init() {
	for i := range ZobristPlayer {
		ZobristPlayer[i] = rand.Uint64()
	}

	for i := 0; i < MaxAgents; i++ {
		for j := 0; j < MaxCells; j++ {
			for k := 0; k < 4; k++ {
				ZobristAgent[i][j][k] = rand.Uint64()
			}
		}
	}
}
