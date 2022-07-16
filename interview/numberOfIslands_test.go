package interview

import "testing"

func TestNumberOfIslands(t *testing.T) {
	cases := [...]struct {
		desc  string
		input [][]byte
		exp   int
	}{
		{"example 1", [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}, 3},
		{"example 2", [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}, 1},
	}
	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if actual := numIslands(tst.input); actual != tst.exp {
				t.Errorf("actual (%d) != expected (%d)", actual, tst.exp)
			}
		})
	}
}
