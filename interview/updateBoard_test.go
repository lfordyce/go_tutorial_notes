package interview

import (
	"fmt"
	"testing"
)

func TestUpdateBoard(t *testing.T) {
	cases := []struct {
		desc   string
		input  [][]byte
		clicks []int
	}{
		{"example 1", [][]byte{
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'M', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
		}, []int{3, 0}},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			out := updateBoard(tst.input, tst.clicks)
			fmt.Println(out)
		})
	}
}
