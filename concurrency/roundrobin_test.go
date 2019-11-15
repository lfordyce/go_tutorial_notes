package concurrency

import (
	"fmt"
	"testing"
)

func TestRountRobin(t *testing.T) {
	execute()
}

type Attribute struct {
	Key, Val string
}

type Node struct {
	Attr []*Attribute
}

func TestIterateAndEdit(t *testing.T) {
	n := Node{[]*Attribute{
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
		fmt.Printf("%#v\n", *v)
	}
}
