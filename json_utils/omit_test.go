package json_utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestOmitEmpty(t *testing.T) {
	d := Dog{
		Breed: "retriever",
	}

	b, _ := json.Marshal(d)
	fmt.Println(string(b))
}

func TestZeroNilDiference(t *testing.T) {
	n := 3
	d1 := Restaurant{
		NumberOfCustomers: &n,
	}

	b, _ := json.Marshal(d1)
	fmt.Println(string(b))

	d2 := Restaurant{}
	b2, _ := json.Marshal(d2)
	fmt.Println(string(b2))
}
