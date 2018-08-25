package picoman

func Search(m *Model) []Direction {
	for n := 0; n <= m.ActiveAgents(); n++ {
		memo := NewMemo()
		previousMemoSize := 0
		for i := 1; ; i++ {
			path := make([]Direction, i)
			if search(m, memo, path, 0, i, n) {
				return path
			}
			if memo.Size() == previousMemoSize {
				break
			}
			previousMemoSize = memo.Size()
		}
	}
	return nil
}

func search(m *Model, memo *Memo, path []Direction, depth, maxDepth, maxAgents int) bool {
	height := maxDepth - depth

	if m.Win() {
		return m.ActiveAgents() <= maxAgents
	}

	if height == 0 {
		return false
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
			solved := search(mc, memo, path, depth+1, maxDepth, maxAgents)
			if solved {
				memo.Set(key, height-1)
				path[depth] = d
				return true
			}
		}
	}
	return false
}
