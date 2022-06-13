package interview

import (
	"testing"
)

func TestTwoSumInputBST(t *testing.T) {
	cases := [...]struct {
		desc  string
		input []int
		k     int
		exp   bool
	}{
		{"given example", []int{5, 3, 6, 2, 4, NULL, 7}, 9, true},
		{"base case", []int{3, 9, 20, NULL, NULL, 17, 7}, 29, true},
		{"more complicated", []int{1, 2, 3, 4, NULL, NULL, 5}, 9, true},
		{"another one", []int{1, 2, 3, 4, NULL, NULL, 5}, 4, true},
	}

	for id, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			root := TreeNodeFromInts(tst.input)
			if result := findTarget(root, tst.k); tst.exp != result {
				t.Errorf("(%d) (%s) test case failed", id, tst.desc)
			}
		})
	}
}

func TestTwoSums(t *testing.T) {
	cases := [...]struct {
		desc   string
		input  []int
		target int
		exp    []int
	}{
		{"base case", []int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for id, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if result := twoSum(tst.input, tst.target); !IntArrayEqual(result, tst.exp) {
				t.Errorf("(%d) (%s) test case failed expected: %v != actual: %v", id, tst.desc, tst.exp, result)
			}
		})
	}
}

func IntArrayEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
