package interview

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	cases := [...]struct {
		desc  string
		input int
		exp   string
	}{
		{
			desc:  "example 1",
			input: 6,
			exp:   "VI",
		},
		{
			desc:  "example 2",
			input: 58,
			exp:   "LVIII",
		},
		{
			desc:  "example 3",
			input: 1994,
			exp:   "MCMXCIV",
		},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if actual := intToRomanRev(tst.input); actual != tst.exp {
				t.Errorf("actual (%s) != expected (%s)", actual, tst.exp)
			}
		})
	}
}

func TestRomanToInt(t *testing.T) {
	cases := [...]struct {
		desc  string
		input string
		exp   int
	}{
		{"example 1", "MCMXCIV", 1994},
		{"example 2", "LVIII", 58},
		{"example 3", "VI", 6},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if actual := romanToInt(tst.input); actual != tst.exp {
				t.Errorf("actual (%d) != expected (%d)", actual, tst.exp)
			}
		})
	}
}

func TestFoo(t *testing.T) {
	s := "III"
	cut := s[2:3]
	fmt.Println(cut)
}
