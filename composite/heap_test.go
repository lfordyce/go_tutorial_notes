package composite

import (
	"container/heap"
	"testing"
)

func TestTeam_Score(t *testing.T) {
	teams := &TeamHeap{}

	heap.Init(teams)
	heap.Push(teams, Country{"Canada", 7})
	heap.Push(teams, Country{"US", 2})
	heap.Push(teams, Country{"Germany", 9})
	heap.Push(teams, Country{"Korea", 3})
	heap.Push(teams, Country{"Sweden", 7})

	for teams.Len() > 1 {
		t1 := heap.Pop(teams).(Scoreable)
		t2 := heap.Pop(teams).(Scoreable)
		heap.Push(teams, Team{t1, t2, t1.Score() + t2.Score()})
	}

	for teams.Len() > 0 {
		s := heap.Pop(teams)
		printScoreable(s.(Scoreable), 0)
	}
}
