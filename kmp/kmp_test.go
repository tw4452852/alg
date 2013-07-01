package kmp

import (
	"strings"
	"testing"
)

type test struct {
	source  string
	search  string
	offsets []int64
}

func TestSearchOne(t *testing.T) {
	cases := []test{
		test{"here is a simple example", "example", []int64{17}},
		test{"hello tw", "world", []int64{-1}},
		test{"hello tw hello tw", "tw", []int64{6}},
	}
	for _, c := range cases {
		offsets := []int64{SearchOne(c.source, c.search)}
		if !checkOffsets(c, offsets) {
			t.Errorf("searchOne %q in %q failed: expect %v, got %v\n",
				c.search, c.source, c.offsets, offsets)
		}
	}
}

func TestSearchAll(t *testing.T) {
	cases := []test{
		test{"hello world", "world", []int64{6}},
		test{"hello tw", "world", nil},
		test{"hello tw tw", "tw", []int64{6, 9}},
		test{"hello twtw", "tw", []int64{6, 8}},
		test{"hello wwww", "w", []int64{6, 7, 8, 9}},
	}
	for _, c := range cases {
		offsets := SearchAll(c.source, c.search)
		if !checkOffsets(c, offsets) {
			t.Errorf("searchAll %q in %q failed: expect %v, got %v\n",
				c.search, c.source, c.offsets, offsets)
		}
	}
}

func checkOffsets(expect test, offsets []int64) (passed bool) {
	if len(expect.offsets) != len(offsets) {
		return false
	}
	for ei, eo := range expect.offsets {
		if eo != offsets[ei] {
			return false
		}
	}
	return true
}

const (
	source = "this is a simple example, this is a simple example."
	search = "example"
)

func BenchmarkStandard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Index(source, search)
	}
}

func BenchmarkSearchOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchOne(source, search)
	}
}

func BenchmarkSearchAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchAll(source, search)
	}
}
