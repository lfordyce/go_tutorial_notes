package interview

import "testing"

//You are given an array representing a row of seats where seats[i] = 1 represents a person sitting in the ith seat, and seats[i] = 0 represents that the ith seat is empty (0-indexed).
//
//There is at least one empty seat, and at least one person sitting.
//
//Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized.
//
//Return that maximum distance to the closest person.

// EX 1
//Input: seats = [1,0,0,0,1,0,1]
//Output: 2
//Explanation:
//If Alex sits in the second open seat (i.e. seats[2]), then the closest person has distance 2.
//If Alex sits in any other open seat, the closest person has distance 1.
//Thus, the maximum distance to the closest person is 2.

// EX 2
//Input: seats = [1,0,0,0]
//Output: 3
//Explanation:
//If Alex sits in the last seat (i.e. seats[3]), the closest person is 3 seats away.
//This is the maximum distance possible, so the answer is 3.

//EX 3
//Input: seats = [0,1]
//Output: 1

func TestMaxArea(t *testing.T) {
	cases := [...]struct {
		desc  string
		input []int
		exp   int
	}{
		{"base case", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}

	for _, tst := range cases {
		t.Run(tst.desc, func(t *testing.T) {
			if actual := maxArea(tst.input); actual != tst.exp {
				t.Errorf("expected: %v != actual: %v", tst.exp, actual)
			}
		})
	}
}
