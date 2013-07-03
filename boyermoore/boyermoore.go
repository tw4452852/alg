package boyermoore

import (
	"fmt"
	"io"
)

// state log the progress when searching
type state struct {
	table    *moveTable
	source   []byte   // the source text
	search   []byte   // the search word
	suffixes []string // good suffixes when searching
	offset   int      // current offset of reader
}

func newState(source, search string) *state {
	return &state{
		table:  newMoveTable(search),
		source: []byte(source),
		search: []byte(search),
	}
}

// move offset of source until reach the end
// if there is error happened, return it
func (s *state) find() (found bool, offset int, err error) {
	// out of range
	if s.offset+len(s.search) > len(s.source) {
		err = io.EOF
		return
	}

	var i int
	for i = len(s.search) - 1; i >= 0; i-- {
		if s.source[s.offset+i] != s.search[i] {
			break
		}
		if i == 0 {
			found = true
			continue
		}
	}
	if found {
		offset = s.offset
	}
	s.move(i, s.source[s.offset+i])
	return
}

// move the source to prepare next finding
func (s *state) move(badOffset int, badCharacter byte) {
	var moveStep int
	if badOffset == -1 {
		moveStep = len(s.search)
	} else {
		moveStep = s.table.maxStep(badCharacter, badOffset)
	}
	if false {
		fmt.Printf("badOffset:%d, badCharacter:%c\n",
			badOffset, badCharacter)
		fmt.Printf("%v move %d\n", s, moveStep)
	}
	s.offset += moveStep
}

// just for debug
func (s *state) String() string {
	return fmt.Sprintf("source(%q), search(%q), offset(%d), suffixes(%v)",
		s.source, s.search, s.offset, s.suffixes)
}

// return the first offset in the source text
// if not found or some error happened, return -1
func SearchOne(source, search string) int64 {
	s := newState(source, search)
	for {
		found, offset, err := s.find()
		if found {
			return int64(offset)
		}
		if err != nil {
			return -1
		}
	}
	panic("not reachable")
}

// return all the offset in the source text
// if not found or some error happened, return nil
func SearchAll(source, search string) []int64 {
	s := newState(source, search)
	offsets := make([]int64, 0)
	for {
		found, offset, err := s.find()
		if found {
			offsets = append(offsets, int64(offset))
			continue
		}
		if err != nil {
			if err == io.EOF {
				return offsets
			}
			return nil
		}
	}
	panic("not reach")
}
