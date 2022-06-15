package interview

import "testing"

func TestCourseSchedule(t *testing.T) {
	cases := [...]struct {
		desc       string
		input      [][]int
		numCourses int
		exp        bool
	}{
		{"example 1", [][]int{{1, 0}}, 2, true},
		{"example 2", [][]int{{1, 0}, {0, 1}}, 2, false},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if actual := canFinish(tst.numCourses, tst.input); actual != tst.exp {
				t.Errorf("actual %v != expected %v", actual, tst.exp)
			}
		})
	}
}
