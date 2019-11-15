package composite

type ANode struct {
	Name     string  `json:"name"`
	Children []ANode `json:"children,omitempty"`
}

func AddToTree(root []ANode, names []string) []ANode {
	if len(names) > 0 {
		var i int
		for i = 0; i < len(root); i++ {
			if root[i].Name == names[0] {
				break
			}
		}
		if i == len(root) {
			root = append(root, ANode{Name: names[0]})
		}
		root[i].Children = AddToTree(root[i].Children, names[1:])
	}
	return root
}
