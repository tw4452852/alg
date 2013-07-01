package kmp

import (
	"reflect"
	"testing"
)

func TestPMT(t *testing.T) {
	type c struct {
		s  string
		vs []int
	}
	cs := []c{
		c{"abcdabd", []int{0, 0, 0, 0, 1, 2, 0}},
		c{"", []int{}},
	}
	for i, c := range cs {
		got := newPmt(c.s)
		if !reflect.DeepEqual(c.vs, []int(got)) {
			t.Errorf("the %d case failed: expect %v, got %v\n",
				i, c.vs, []int(got))
		}
	}
}
