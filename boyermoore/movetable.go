package boyermoore

import (
	"strings"
)

// precount a moveable table used when searching
type moveTable struct {
	source string         // source string
	good   map[string]int // good suffix moveable table
}

// generate table according to the origin string
func newMoveTable(s string) *moveTable {
	good := make(map[string]int)

	for i := len(s) - 1; i > 0; i-- {
		suffix := s[i:]
		prefix := s[:len(suffix)]
		preIndex := len(prefix) - 1
		if strings.EqualFold(suffix, prefix) {
			good[string(suffix)] = len(s) - 1 - preIndex
		}
	}

	return &moveTable{
		source: s,
		good:   good,
	}
}

// return good suffix move step count
// if not a good suffix, return -1
func (m *moveTable) goodStep(suffix string) int {
	if step, found := m.good[suffix]; found {
		return step
	}
	return -1
}

// return bad character move step count
func (m *moveTable) badStep(badCharacter rune, badPosition int) int {
	if badPosition > len(m.source)-1 {
		panic("bad position")
	}
	return badPosition - strings.LastIndex(m.source[:badPosition], string(badCharacter))
}
