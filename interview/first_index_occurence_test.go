package interview

import (
	"fmt"
	"testing"
)

func TestFirstIndexOccrence(t *testing.T) {

	cases := []struct {
		haystack string
		needle   string
		idx      int
	}{
		{"hello", "ll", 2},
		{"abacbabc", "abc", 5},
		{"abacbabc", "abcd", -1},
		{"abacbabc", "", 0},
	}

	for idx, tst := range cases {
		t.Run(fmt.Sprintf("case:%d", idx), func(t *testing.T) {
			if ans := FirstIndexOccurence(tst.haystack, tst.needle); ans != tst.idx {
				t.Errorf("%d, want:%d, got:%d", idx, tst.idx, ans)
			}
		})
	}
}
