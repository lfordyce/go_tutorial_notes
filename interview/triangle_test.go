package interview

import "testing"

func TestTriangle(t *testing.T) {
	cases := [...]struct {
		desc  string
		input [][]int
		exp   int
	}{
		{"simple", [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if actual := minimumTotal(tst.input); tst.exp != actual {
				t.Errorf("expected: %v != actual: %v", tst.exp, actual)
			}
		})
	}
}
