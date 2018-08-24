package picoman

import (
	"fmt"
)

func Search(m *Model) {
	memo := NewMemo()
	for i := 1; ; i++ {
		path := make([]Direction, i)
		if search(m, memo, path, 0, i) {
			fmt.Println(i, memo.Size(), memo.Hits())
			fmt.Println(path)
			break
		}
	}
}

func search(m *Model, memo *Memo, path []Direction, depth, maxDepth int) bool {
	height := maxDepth - depth
	if height == 0 {
		return m.Win()
	}

	key := m.ZobristHash()
	if !memo.Add(key, height) {
		return false
	}

	for _, d := range Directions {
		if !m.CanMove(d) {
			continue
		}
		mc := m.Copy()
		if mc.DoMove(d) {
			solved := search(mc, memo, path, depth+1, maxDepth)
			if solved {
				memo.Set(key, height-1)
				path[depth] = d
				return true
			}
		}
	}
	return false
}
