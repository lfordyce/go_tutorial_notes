package interview

import (
	"fmt"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	cases := []struct {
		input string
		exp   int
	}{
		{"(()", 2},
		{")()())", 4},
		{"()(())", 6},
		{"()(())))", 6},
		{"()(()", 2},
	}

	for idx, tst := range cases {
		t.Run(fmt.Sprintf("case:%d", idx), func(t *testing.T) {
			if ans := longestValidParenthese(tst.input); ans != tst.exp {
				t.Errorf("%d, want: %d, got: %d", idx, tst.exp, ans)
			}
		})
	}
}
