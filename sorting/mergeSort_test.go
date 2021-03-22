package sorting

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/pingcap/check"
)

var _ = check.Suite(&sortTestSuite{})

func TestT(t *testing.T) {
	check.TestingT(t)
}

func prepare(src []int64) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		src[i] = rand.Int63()
	}
}

type sortTestSuite struct{}

func (s *sortTestSuite) TestMergeSort(c *check.C) {
	lens := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1024, 1 << 13, 1 << 17, 1 << 19, 1 << 20}
	for i := range lens {
		src := make([]int64, lens[i])
		expect := make([]int64, lens[i])
		prepare(src)
		copy(expect, src)
		MergeSortAlt(src)
		sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })
		for i := 0; i < len(src); i++ {
			c.Assert(src[i], check.Equals, expect[i])
		}
	}
}

func (s *sortTestSuite) TestMergeSortShort(c *check.C) {
	src := []int64{9, 3, 6, 8, 13, 5, 6}

	expect := make([]int64, len(src))
	copy(expect, src)
	MergeSortAlt(src)
	sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })
	for i := 0; i < len(src); i++ {
		c.Assert(src[i], check.Equals, expect[i])
	}
}


func TestMergeSort(t *testing.T) {
	slice := generateSlice(50)
	fmt.Println("\n --- unsorted --- \n\n", slice)
	sorted := MergeSort(slice)
	fmt.Println("\n --- sorted --- \n\n", sorted)

	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] > sorted[i+1] {
			t.Error("Merge sort failed")
		}
	}
}

func TestMergeSortOffLength(t *testing.T) {
	numbers := [7]int{9, 3, 6, 8, 13, 5, 6}
	slice := numbers[:]
	sorted := MergeSort(slice)

	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] > sorted[i+1] {
			t.Error("Merge sort failed")
		}
	}
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return slice
}

func TestFooBar(t *testing.T) {
	values := []string{"a", "b", "c"}
	var funcs []func()

	for _, val := range values {
		val := val
		funcs = append(funcs, func() {
			fmt.Println(val)
		})
	}

	for _, fn := range funcs {
		fn()
	}

	var copies []*string
	for _, val := range values {
		copies = append(copies, &val)
	}
}