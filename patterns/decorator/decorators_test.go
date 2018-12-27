package decorator

import (
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
	"testing"
)

func TestSortedFunc(t *testing.T) {
	stable := TimedSortFunc(sort.Stable)
	unStable := TimedSortFunc(sort.Sort)

	randomHeroList1 := RandomHeroSort(1000, 5000)
	randomHeroList2 := RandomHeroSort(1000, 5000)

	fmt.Printf("Unstable sorting funciton:n")
	unStable(randomHeroList1)
	fmt.Printf("Stable sorting function:n")
	stable(randomHeroList2)
}

func TestLogging(t *testing.T) {

	var fn LikeFunc = LikePost
	// foo := fn(123, 789)
	// fmt.Println(foo)
	fn = DecoratedLike(fn)
	fmt.Println(fn(123, 432))
	fmt.Println(fn(654321, 12345678765432))
}

func TestAdderDecorator(t *testing.T) {

	a := AdderFunc(
		func(x, y int) int {
			return x + y
		},
	)
	fmt.Println(Do(a))
}

func TestAdderMiddleware(t *testing.T) {
	var a Adder = AdderFunc(
		func(x, y int) int {
			return x + y
		},
	)

	a = WrapLogger(log.New(os.Stdout, "test ", 1))(a)

	fmt.Println(Do(a))
}

func Do(adder Adder) int {
	return adder.Add(1, 2)
}

func TestAdderCach(t *testing.T) {
	var a Adder = AdderFunc(
		func(x, y int) int {
			return x + y
		},
	)

	a = WrapCache(&sync.Map{})(a)

	fmt.Println(Do(a))
}

func TestAdderChainMiddleware(t *testing.T) {
	var a Adder = AdderFunc(
		func(x, y int) int {
			return x + y
		},
	)
	a = Chain(
		WrapLogger(log.New(os.Stdout, "test", 1)),
		WrapCache(&sync.Map{}),
	)(a)

	a.Add(10, 20)
}
