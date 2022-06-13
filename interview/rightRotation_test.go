package interview

import (
	"fmt"
	"testing"
)

func TestRotateRight(t *testing.T) {
	cases := [...]struct {
		desc     string
		input    []int
		target   int
		expected []int
	}{
		{"base case", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
	}

	for id, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			linkedList := ArrayToListNode(tst.input)
			actual := rotateRightAlt(linkedList, tst.target)
			a := ListNodeToArray(actual)
			fmt.Printf("【output】:%v\n", a)
			if !equalArr(a, tst.expected) {
				t.Errorf("(%d), (%s) expected: %v != actual: %v", id, tst.desc, tst.expected, a)
			}
		})
	}
}

func equalArr(a []int, b []int) bool {
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
