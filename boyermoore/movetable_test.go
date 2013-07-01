package boyermoore

import (
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
				source: "example",
				good: map[string]int{
					"e": 6,
				},
			},
		},
		tableCase{
			s: "",
			t: &moveTable{
				source: "",
				good:   map[string]int{},
			},
		},
		tableCase{
			s: "eeiee",
			t: &moveTable{
				source: "eeiee",
				good: map[string]int{
					"e":  4,
					"ee": 3,
				},
			},
		},
	}

	for _, c := range cases {
		rt := newMoveTable(c.s)
		checkMoveTable(c.t, rt, t)
	}
}

func checkMoveTable(expect, got *moveTable, t *testing.T) {
	if expect.source != got.source {
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
		suffix       string
		suffixStep   int
		badCharacter rune
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
				stepExpect{"e", 6, 's', 6, 7},
				stepExpect{"le", -1, 'p', 6, 2},
				stepExpect{"ple", -1, 'e', 4, 1},
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
