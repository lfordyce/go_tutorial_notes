package avl

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type intKey int

func (k intKey) Less(k2 Key) bool { return k < k2.(intKey) }
func (k intKey) Eq(k2 Key) bool   { return k == k2.(intKey) }

func dump(tree *Node) {
	b, err := json.MarshalIndent(tree, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func TestAvlTree(t *testing.T) {
	var tree *Node
	fmt.Println("Empty tree:")
	dump(tree)

	fmt.Println("\nInsert test:")

	Insert(&tree, intKey(4))
	Insert(&tree, intKey(5))
	Insert(&tree, intKey(1))
	Insert(&tree, intKey(2))
	Insert(&tree, intKey(6))

	dump(tree)

}
