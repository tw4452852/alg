package kmp

import (
	"fmt"
	"io"
	"strings"
)

// state are the context when searching
type state struct {
	source string          // source text
	search string          // search word
	r      *strings.Reader // source reader according to source
	t      pmt             // patrial match table
	e      error           // striky error
	offset int64           // current offset of reader
}

func newState(source, search string) *state {
	return &state{
		source: source,
		search: search,
		r:      strings.NewReader(source),
		t:      newPmt(search),
	}
}

// move offset of source until reach the end
// if there is error happened, return it
func (s *state) find() (found bool, offset int64, err error) {
	if s.e != nil {
		err = s.e
		return
	}
	b := make([]byte, 1)
	var matchCount int
	for i, c := range s.search {
		_, err = s.r.ReadAt(b, s.offset+int64(i))
		if err != nil {
			s.e = err
			return
		}
		if c != rune(b[0]) {
			break
		}
		matchCount++
		if i == len(s.search)-1 {
			found = true
			break
		}
		b = b[0:1]
	}
	if found {
		offset = s.offset
	}
	s.move(matchCount)
	return
}

// move the source to prepare next finding
func (s *state) move(matchCount int) {
	// matchCount is the position of the partrial character
	var moveStep int64
	if matchCount == 0 {
		moveStep = 1
	} else {
		moveStep = int64(matchCount - s.t.get(matchCount))
	}
	if false {
		fmt.Printf("matchCount %d\n", matchCount)
		fmt.Printf("%v move %d\n", s, moveStep)
	}
	s.offset, s.e = s.r.Seek(moveStep, 1)
}

// just for debug
func (s *state) String() string {
	return fmt.Sprintf("source(%q), search(%q), offset(%d), err(%v)",
		s.source, s.search, s.offset, s.e)
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
	panic("not reachable")
}
