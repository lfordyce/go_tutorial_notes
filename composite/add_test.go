package composite

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestAddToTree(t *testing.T) {
	s := []string{
		"a/b/c",
		"a/b/g",
		"a/b/e",
		"a/d",
	}
	var tree []ANode
	for i := range s {
		tree = AddToTree(tree, strings.Split(s[i], "/"))
	}

	b, err := json.Marshal(tree)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
