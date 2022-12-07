package main

import (
	"fmt"
)

type Node struct {
	Children map[string]*Node
	Parent   map[string]*Node
	Name     string
	Type     NodeType
	size     int
	isRoot   bool
	Depth    int
}

func nodeSize(node Node, size int) int {
	if node.Type == File {
		return size + node.size
	}

	for _, child := range node.Children {
		size = nodeSize(*child, size)
	}
	return size
}

func (n *Node) Size() int {
	return nodeSize(*n, 0)
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

func GetCleanableDirectory(node Node, spaceNeeded int, nearest int) int {

	if node.Type != Directory {
		return nearest
	}

	if node.Size() > spaceNeeded && node.Size() < nearest {
		nearest = node.Size()
	}

	for _, child := range node.Children {
		if child.Type != Directory {
			continue
		}
		nearest = GetCleanableDirectory(*child, spaceNeeded, nearest)
	}
	return nearest
}
