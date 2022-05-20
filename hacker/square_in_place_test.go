package hacker

import (
	"fmt"
	"testing"
)

func TestSquareInPlace(t *testing.T) {
	x := 1.5
	// TODO: replace the ? placeholder
	SquareInPlace(&x)
	fmt.Print(x)
	if x != 2.25 {
		t.Error("square failed")
	}
}
