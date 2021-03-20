package concurrency

import (
	"fmt"
	"testing"
)

func TestNewIterator(t *testing.T) {
	data := []string{"Sphinx of black quartz, judge my vow",
		"The sky is blue and the water too",
		"Cozy lummox gives smart squid who asks for job pen",
		"Jackdaws love my big sphinx of quartz",
		"The quick onyx goblin jumps over the lazy dwarf"}
	histogram := make(map[string]int)
	iter := NewIterator(data) // returns handle to data channel

	for value, ok := iter.Next(); ok; value, ok = iter.Next() {
		histogram[value.(string)]++
		fmt.Println(value, ok)
	}
	fmt.Println(iter.Error())
	fmt.Println(histogram)
}
