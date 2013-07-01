package kmp

// patial match table
type pmt []int // suffix and prefix common part length

func newPmt(s string) pmt {
	vs := make([]int, len(s))
	for i := range vs {
		vs[i] = commonLength(s[:i+1])
	}
	return vs
}

// count the suffix and prefix common part length
func commonLength(s string) (length int) {
	ss := suffix(s)
	ps := prefix(s)

	if ss == nil || ps == nil {
		return
	}
	for _, s := range ss {
		for _, p := range ps {
			if s == p {
				return len(s)
			}
		}
	}
	return
}

func suffix(s string) []string {
	if len(s) <= 1 {
		return nil
	}
	ss := make([]string, len(s)-1)
	for i := range ss {
		ss[i] = s[i+1:]
	}
	return ss
}

func prefix(s string) []string {
	if len(s) <= 1 {
		return nil
	}
	ps := make([]string, len(s)-1)
	for i := range ps {
		ps[i] = s[:i+1]
	}
	return ps
}

func (t pmt) get(i int) int {
	if i <= 0 && i < len(t) {
		return t[i]
	}
	return 0
}
