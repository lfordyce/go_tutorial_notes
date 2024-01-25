package generics

import (
	"strconv"
	"testing"
)

// Setter2 is a type constraint that requires that the type
// implement a Set method that sets the value from a string,
// and also requires that the type be a pointer to its type parameter.
type Setter2[B any] interface {
	Set(string)
	*B // non-interface type constraint element
}

// FromStrings2 takes a slice of strings and returns a slice of T,
// calling the Set method to set each returned value.
//
// We use two different type parameters so that we can return
// a slice of type T but call methods on *T aka PT.
// The Setter2 constraint ensures that PT is a pointer to T.
func FromStrings2[T any, PT Setter2[T]](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		// The type of &result[i] is *T which is in the type set
		// of Setter2, so we can convert it to PT.
		p := PT(&result[i])
		// PT has a Set method.
		p.Set(v)
	}
	return result
}

// Settable is an integer type that can be set from a string.
type Settable int

// Set sets the value of *p from a string.
func (p *Settable) Set(s string) {
	i, _ := strconv.Atoi(s) // real code should not ignore the error
	*p = Settable(i)
}

func TestFromStrings2(t *testing.T) {
	nums := FromStrings2[Settable, *Settable]([]string{"1", "2"})
	t.Log(nums)
}
