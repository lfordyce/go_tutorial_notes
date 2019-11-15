package decorator

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sort"
	"strings"
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

	i := &sync.Map{}
	a = WrapCache(i)(a)
	a.Add(5, 6)

	value, ok := i.Load("x=5y=6")
	if ok {
		fmt.Println(value)
	}

	//fmt.Println(Do(a))
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

func TestCompose(t *testing.T) {
	raw := "\n\n\nHello Golang!!!\n\n\n"
	trim := strings.TrimSpace
	trimExclamation := func(s string) string {
		return strings.Trim(s, "!")
	}
	toLower := strings.ToLower

	s := compose(toLower, trimExclamation, trim)(raw)
	fmt.Println(s)
}

func TestComposeUrl(t *testing.T) {
	raw := "halo://prdtll3ep005f:8080"
	parse, _ := url.Parse(raw)

	i := composeUrl(
		changeScheme("http"),
		addPathToUrl(fmt.Sprintf("ui/liveencoder/%s/stats", "some-ID")),
	)(*parse)

	fmt.Println(i.String())
}

func TestAppendDecorator(t *testing.T) {
	s := "Hello, playground"

	var fn StringManipulator = ident

	fmt.Println(fn(s))

	fn = ToBase64(ToLower(fn))
	fmt.Println(fn(s))

}
