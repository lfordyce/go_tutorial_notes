package binary_tree

import (
	"fmt"
	"testing"
)

func TestTree_Traverse(t *testing.T) {

	values := []string{"d", "b", "c", "e", "e"}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}

	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i], data[i])
		if err != nil {
			t.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}

	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
}
