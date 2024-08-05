package interview

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		input []string
		want  string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"ab", "abc", "a"}, "a"},
	}

	for idx, tst := range cases {
		t.Run(fmt.Sprintf("case-%d", idx), func(t *testing.T) {
			if ans := LongestCommonPrefix(tst.input); ans != tst.want {
				t.Errorf("%s, want: %s", ans, tst.want)
			}
		})
	}
}
