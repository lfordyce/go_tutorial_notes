package concurrency

import (
	"fmt"
	"math/big"
	"testing"
)

func TestMux(t *testing.T) {
	r := make([]chan big.Int, 10)
	for i := 0; i < 10; i++ {
		r[i] = fromTo(i*10, i*10+10)
	}
	all := Mux(r)
	for l := range all {
		fmt.Println(l)
	}
}
