package boyermoore

import (
	"reflect"
	"testing"
)

func TestMoveTable(t *testing.T) {
	type tableCase struct {
		s string
		t *moveTable
	}

	cases := []tableCase{
		tableCase{
			s: "example",
			t: &moveTable{
				source: []byte("example"),
				good:   []int{0, -1, -1, -1, -1, -1, 6},
			},
		},
		tableCase{
			s: "",
			t: &moveTable{
				source: []byte(""),
				good:   []int{},
			},
		},
		tableCase{
			s: "eeiee",
			t: &moveTable{
				source: []byte("eeiee"),
				good:   []int{0, -1, -1, 3, 4},
			},
		},
	}

	for _, c := range cases {
		rt := newMoveTable(c.s)
		checkMoveTable(c.t, rt, t)
	}
}

func checkMoveTable(expect, got *moveTable, t *testing.T) {
	if !reflect.DeepEqual(expect.source, got.source) {
		t.Errorf("source check failed: expect %q, got %q\n",
			expect.source, got.source)
	}
	// good table
	if len(expect.good) != len(got.good) {
		t.Errorf("good table len check failed: expect %d, got %d\n",
			len(expect.good), len(got.good))
		return
	}
	for ei, es := range expect.good {
		if es != got.good[ei] {
			t.Errorf("good suffix(%v) check failed: expect %v, got %v\n",
				ei, es, got.good[ei])
			return
		}
	}
}

func TestStep(t *testing.T) {
	type stepExpect struct {
		suffix       int
		suffixStep   int
		badCharacter byte
		badPosition  int
		badStep      int
	}

	type stepCase struct {
		s       string
		expects []stepExpect
	}

	cases := []stepCase{
		stepCase{
			s: "exaeple",
			expects: []stepExpect{
				stepExpect{6, 6, 's', 6, 7},
				stepExpect{5, -1, 'p', 6, 2},
			},
		},
	}

	for _, c := range cases {
		m := newMoveTable(c.s)
		for _, expect := range c.expects {
			gs := m.goodStep(expect.suffix)
			bs := m.badStep(expect.badCharacter, expect.badPosition)
			if gs != expect.suffixStep {
				t.Errorf("good step check failed: expect %d, got %d\n",
					expect.suffixStep, gs)
			}
			if bs != expect.badStep {
				t.Errorf("bad step check failed: expect %d, got %d\n",
					expect.badStep, bs)
			}
		}
	}
}
