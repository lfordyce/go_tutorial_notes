package composite

import (
	"encoding/json"
	"fmt"
)

type BNode struct {
	ID       int
	Name     string
	ParentID int
	Depth    int
	Path     string
	Child    *BNode
}

func MakeTree(indata string) error {
	nodes := make([]BNode, 0)

	if err := json.Unmarshal([]byte(indata), &nodes); err != nil {
		return err
	}

	m := make(map[int]*BNode)
	for i := range nodes {
		m[nodes[i].ID] = &nodes[i]
	}

	for i, n := range nodes {
		if m[n.ParentID] != nil {
			m[n.ParentID].Child = &nodes[i]
		}
	}

	outdata, err := json.Marshal(m[1])
	if err != nil {
		panic(err)
	}

	fmt.Println(string(outdata))
	return nil
}
