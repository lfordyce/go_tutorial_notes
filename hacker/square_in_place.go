package hacker

import (
	"errors"
	"math"
)

// SquareInPlace replaces the value of v with its square value
// Note: this is a void function
// Note: square of two is four and not 1.41421356237...
// TODO: replace T with the appropriate type
func SquareInPlace(v *float64) {
	// TODO: implement the square in place for v
	//*v = *v * *v
	//*v *= *v
	*v = math.Pow(*v, 2)
}

func SquareInPlaceGeneric(value interface{}) {
	// TODO: implement the square in place for v
	switch v := value.(type) {
	case *float32:
		*v *= *v
	case *float64:
		*v *= *v
	default:
		panic(errors.New("invalid type"))
	}
}