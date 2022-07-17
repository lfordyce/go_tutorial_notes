package interview

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	cases := [...]struct {
		desc  string
		input string
		exp   []string
	}{
		{"example 1", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"example 2", "236", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			out := letterCombinations(tst.input)
			fmt.Println(out)
		})
	}
}
