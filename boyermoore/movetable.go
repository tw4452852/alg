package boyermoore

import (
	"bytes"
)

// precount a moveable table used when searching
type moveTable struct {
	source []byte   // source string
	bad    [256]int // bad character array
	good   []int    // good suffix moveable table
}

// generate table according to the origin string
func newMoveTable(str string) *moveTable {
	s := []byte(str)
	good := make([]int, len(s))

	for i := len(s) - 1; i > 0; i-- {
		suffix := s[i:]
		prefix := s[:len(suffix)]
		preIndex := len(prefix) - 1
		if bytes.Compare(suffix, prefix) == 0 {
			good[i] = len(s) - 1 - preIndex
		} else {
			good[i] = -1
		}
	}

	t := &moveTable{
		source: s,
		good:   good,
	}

	for i, b := range s {
		t.bad[b] = i
	}
	return t
}

// return good suffix move step count
// if not a good suffix, return -1
func (m *moveTable) goodStep(position int) int {
	if 0 <= position && position < len(m.source) {
		return m.good[position]
	}
	return -1
}

// return bad character move step count
func (m *moveTable) badStep(badCharacter byte, badPosition int) int {
	if bytes.IndexByte(m.source[:badPosition], badCharacter) == -1 {
		return badPosition + 1
	}
	return badPosition - m.bad[badCharacter]
}

// return the maximun of good and bad step
func (m *moveTable) maxStep(badCharacter byte, badPosition int) int {
	bad := m.badStep(badCharacter, badPosition)

	var good int
	for i := badPosition + 1; i < len(m.source); i++ {
		step := m.goodStep(i)
		if step > good {
			good = step
		}
	}

	if bad > good {
		return bad
	}
	return good
}
