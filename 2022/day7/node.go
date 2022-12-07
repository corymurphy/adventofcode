package main

type Node struct {
	Children map[string]*Node
	Parent   map[string]*Node
	Name     string
	Type     NodeType
	size     int
	isRoot   bool
	Depth    int
}

func (n *Node) Size() int {
	if n.Type == File {
		// fmt.Println(n.Depth)
		return n.size

	} else {
		size := 0
		for _, child := range n.Children {
			size = size + child.Size()
		}
		return size
	}
}

// func (n *Node) Append(node Node) {
// 	n.Children[node.Name] = node
// }
