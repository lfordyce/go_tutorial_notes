package interview

import (
	"fmt"
	"testing"
)

//Range expression                          1st value          2nd value
//
//array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
//string          s  string type            index    i  int    see below  rune
//map             m  map[K]V                key      k  K      m[k]       V
//channel         c  chan E, <-chan E       element  e  E
//
// The reason for this is that range copies the values from the slice you're iterating over.
// So, range uses a[i] as its second value for arrays/slices, which effectively means that the value is copied,
// making the original value untouchable.
func TestRangeBehavior(t *testing.T) {
	x := make([]int, 3)
	x[0], x[1], x[2] = 1, 2, 3
	for i, val := range x {
		fmt.Println(&x[i], "vs", &val)
	}
}

type Attribute struct {
	Key, Val string
}
type Node struct {
	Attr []Attribute
}

func TestSliceModificationIteration(t *testing.T) {
	n := Node{[]Attribute{
		{"foo", ""},
		{"href", ""},
		{"bar", ""},
	}}

	for i := range n.Attr {
		attr := &n.Attr[i]
		if attr.Key == "href" {
			attr.Val = "something"
		}
	}

	for _, v := range n.Attr {
		fmt.Printf("%#v\n", v)
	}
}

type NodeAlt struct {
	Attr []*Attribute
}

func TestSliceModificationIterationAlt(t *testing.T) {
	n := NodeAlt{[]*Attribute{
		&Attribute{"foo", ""},
		&Attribute{"href", ""},
		&Attribute{"bar", ""},
	}}

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			attr.Val = "something"
		}
	}

	for _, v := range n.Attr {
		fmt.Printf("%#v\n", v)
	}
}

type myStruct struct {
	Name  string
	Count int
}

//when you for…range through a collection, the object returned is a copy of the original held in the collection (“value semantics”).
//So, the elem variable is a copy that you are assigning the count to.
func TestIterationUpdateExample(t *testing.T) {
	mod := 4
	cnt := 16
	chartRecords := make([]myStruct, 0)
	for i := 0; i <= cnt; i++ {
		n := myStruct{Count: i, Name: fmt.Sprintf("Joe%2d", i)} //Load some data
		chartRecords = append(chartRecords, n)
	}

	fmt.Printf("======NOW MODIFY VALUES THIS WAY========\r\n")
	i := 0
	for idx := range chartRecords {
		mm := modMe(mod, i)
		chartRecords[idx].Count = mm
		fmt.Printf("No: %2d | Count: %2d | Name = %s\r\n", i, chartRecords[idx].Count, chartRecords[idx].Name) //Print out this elem.Count element in the range
		i = i + 1
	}

	fmt.Printf("======CHECK AGAIN AND VALUES ARE AS DESIRED========\r\n") //Now lets loop through the same range
	i = 0
	for _, elem := range chartRecords {
		fmt.Printf("No: %2d | Count: %2d | Name = %s\r\n", i, elem.Count, elem.Name) //Print out this elem.Count element in the range
		i = i + 1
	}
}

func modMe(mod int, value int) int {
	return value % mod
}

// Filtering without allocating
// This trick uses the fact that a slice shares the same backing array and capacity as the original,
// so the storage is reused for the filtered slice. Of course, the original contents are modified.
func TestDeleteWhileIterating(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	isValid := func(x int) bool { return x%3 == 0 }

	temp := s[:0]
	for _, x := range s {
		if isValid(x) {
			temp = append(temp, x)
		}
	}
	s = temp
	fmt.Println(s)
}
