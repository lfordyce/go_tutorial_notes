package composite

import (
	"fmt"
	"strings"
)

type Scoreable interface {
	fmt.Stringer
	Score() int
}

type Country struct {
	name  string
	score int
}

func (c Country) Score() int {
	return c.score
}

func (c Country) String() string {
	return fmt.Sprintf("%s [%d]", c.name, c.score)
}

type Team struct {
	team1, team2 Scoreable
	score        int
}

func (t Team) Score() int {
	return t.score
}

func (t Team) String() string {
	return fmt.Sprintf("(%s + %s)", t.team1.String(), t.team2.String())
}

type TeamHeap []Scoreable

func (th TeamHeap) Len() int {
	return len(th)
}

func (th TeamHeap) Less(i, j int) bool {
	return th[i].Score() < th[j].Score()
}

func (th TeamHeap) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
}

func (th *TeamHeap) Push(t interface{}) {
	*th = append(*th, t.(Scoreable))
}

func (th *TeamHeap) Pop() interface{} {
	old := *th
	n := len(old)
	t := old[n-1]
	*th = old[0 : n-1]
	return t
}

func printScoreable(s Scoreable, indent int) {
	switch s.(type) {
	case Country:
		c := s.(Country)
		fmt.Printf("%s%s [%d]\n", strings.Repeat("  ", indent), c.name, c.score)

	case Team:
		t := s.(Team)
		fmt.Printf("%sTeam [%d]:\n", strings.Repeat("  ", indent), t.score)
		printScoreable(t.team1.(Scoreable), indent+1)
		printScoreable(t.team2.(Scoreable), indent+1)
	}
}
