package picoman

type Memo struct {
	data map[uint64]int
	hits uint64
}

func NewMemo() *Memo {
	data := make(map[uint64]int)
	return &Memo{data, 0}
}

func (memo *Memo) Size() int {
	return len(memo.data)
}

func (memo *Memo) Hits() uint64 {
	return memo.hits
}

func (memo *Memo) Add(key uint64, depth int) bool {
	memo.hits++
	if before, ok := memo.data[key]; ok && before >= depth {
		return false
	}
	memo.data[key] = depth
	return true
}

func (memo *Memo) Set(key uint64, depth int) {
	memo.data[key] = depth
}
