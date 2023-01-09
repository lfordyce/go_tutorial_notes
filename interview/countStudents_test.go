package interview

import "testing"

func TestCountStudents(t *testing.T) {
	cases := [...]struct {
		desc       string
		students   []int
		sandwiches []int
		exp        int
	}{
		{"example 1", []int{1, 1, 0, 0}, []int{0, 1, 0, 1}, 0},
		{"example 2", []int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}, 3},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if actual := countStudents(tst.students, tst.sandwiches); actual != tst.exp {
				t.Errorf("expected (%d) != actual (%d)", tst.exp, actual)
			}
		})
	}
}
