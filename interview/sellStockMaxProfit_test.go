package interview

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := [...]struct {
		desc  string
		input []int
		exp   int
	}{
		{"base case", []int{7, 1, 5, 3, 6, 4}, 5},
		{"case two", []int{7, 6, 4, 3, 1}, 0},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if actual := maxProfit(tst.input); actual != tst.exp {
				t.Errorf("actual %d != expected %d", actual, tst.exp)
			}
		})
	}
}
