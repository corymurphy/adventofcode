package main

import "fmt"

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

func (n *Node) Print() {
	if n.Type == File {
		fmt.Printf("%s- %s (%s, size=%d)\n", newRepeatingString(n.Depth, " "), n.Name, n.Type.String(), n.size)
	} else {
		fmt.Printf("%s- %s (%s, size=%d)\n", newRepeatingString(n.Depth, " "), n.Name, n.Type.String(), n.Size())
	}

	for _, value := range n.Children {
		value.Print()
	}
}

func SumDirectories(node Node, sizeAtMost int, sum int) int {

	if node.Type == Directory && node.Size() <= sizeAtMost {
		sum = sum + node.Size()
	}

	for _, child := range node.Children {
		sum = SumDirectories(*child, sizeAtMost, sum)
	}
	return sum
}
