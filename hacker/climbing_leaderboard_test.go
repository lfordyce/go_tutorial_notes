package hacker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ClimbingLeaderboard(t *testing.T) {
	cases := []struct {
		score    []int32
		alice    []int32
		expected []int32
	}{
		{
			[]int32{100, 90, 90, 80},
			[]int32{70, 80, 105},
			[]int32{4, 3, 1},
		},
		{
			[]int32{100, 100, 50, 40, 40, 20, 10},
			[]int32{5, 25, 50, 120},
			[]int32{6, 4, 2, 1},
		},
	}

	for idx, tst := range cases {
		t.Run(fmt.Sprintf("%d:", idx), func(t *testing.T) {
			result := climbinLeaderboard(tst.score, tst.alice)
			assert.Equal(t, tst.expected, result)
		})
	}
}

func Test_ClimbingLeaderboardAlt(t *testing.T) {
	cases := []struct {
		score    []int32
		alice    []int32
		expected []int32
	}{
		{
			[]int32{100, 90, 90, 80},
			[]int32{70, 80, 105},
			[]int32{4, 3, 1},
		},
		{
			[]int32{100, 100, 50, 40, 40, 20, 10},
			[]int32{5, 25, 50, 120},
			[]int32{6, 4, 2, 1},
		},
	}

	for idx, tst := range cases {
		t.Run(fmt.Sprintf("%d:", idx), func(t *testing.T) {
			result := climbingLeaderboardAlt(tst.score, tst.alice)
			assert.Equal(t, tst.expected, result)
		})
	}
}
