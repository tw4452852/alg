package boyermoore

import (
	"bytes"
	"fmt"
	"io"
)

// state log the progress when searching
type state struct {
	table    *moveTable
	source   string        // the source text
	r        *bytes.Reader // the source text reader
	search   []byte        // the search word
	err      error         // error when searching
	suffixes []string      // good suffixes when searching
	offset   int64         // current offset of reader
}

func newState(source, search string) *state {
	return &state{
		table:  newMoveTable(search),
		source: source,
		r:      bytes.NewReader([]byte(source)),
		search: []byte(search),
	}
}

// move offset of source until reach the end
// if there is error happened, return it
func (s *state) find() (found bool, offset int64, err error) {
	// if error already happened, can't move ahead
	if s.err != nil {
		err = s.err
		return
	}
	b := make([]byte, 1)

	var i int
	for i = len(s.search) - 1; i >= 0; i-- {
		_, err = s.r.ReadAt(b, s.offset+int64(i))
		if err != nil {
			s.err = err
			return
		}
		if b[0] != s.search[i] {
			break
		}
		if i == 0 {
			found = true
			break
		}
		// the good character must at the begining
		if b[0] == s.search[0] {
			// if we good byte, add it to good suffix
			if s.suffixes == nil {
				s.suffixes = make([]string, 1)
			}
			s.suffixes = append(s.suffixes, string(s.search[i:]))
		}

		b = b[0:1]
	}
	if found {
		offset = s.offset
	}
	// if we not found, move the source
	s.move(i, b[0])
	return
}

// move the source to prepare next finding
func (s *state) move(badOffset int, badCharacter byte) {
	moveStep := 0
	badStep := s.table.badStep(rune(badCharacter), badOffset)
	goodStep := -1
	if s.suffixes != nil {
		// find the max good step
		for _, suffix := range s.suffixes {
			s := s.table.goodStep(suffix)
			if s > goodStep {
				goodStep = s
			}
		}
		// clean up old stuff
		s.suffixes = s.suffixes[:0]
	}
	switch {
	case badOffset == -1:
		moveStep = len(s.search)
	case badStep > goodStep:
		moveStep = badStep
	default:
		moveStep = goodStep
	}
	if false {
		fmt.Printf("badOffset:%d, badCharacter:%c, bad:%d, good:%d\n",
			badOffset, badCharacter, badStep, goodStep)
		fmt.Printf("%v move %d\n", s, moveStep)
	}
	s.offset, s.err = s.r.Seek(int64(moveStep), 1)
}

// just for debug
func (s *state) String() string {
	return fmt.Sprintf("source(%q), search(%q), offset(%d), err(%v), suffixes(%v)",
		s.source, string(s.search), s.offset, s.err, s.suffixes)
}

// return the first offset in the source text
// if not found or some error happened, return -1
func SearchOne(source, search string) int64 {
	s := newState(source, search)
	for {
		found, offset, err := s.find()
		if found {
			return offset
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
			offsets = append(offsets, offset)
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
