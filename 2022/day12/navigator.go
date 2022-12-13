package main

import "fmt"

type Navigator struct {
	start     *Node
	visited   *map[Vector]bool
	territory Map
}

// setups up our tree
func (n *Navigator) Explore() {

	n.setVisited(n.start.vector)
}

func (n *Navigator) setVisited(vector Vector) {
	(*n.visited)[vector] = true
}

func isNavigable(from *Node, to *Node) bool {

	if Height[from.height]-1 == Height[to.height] ||
		Height[from.height] == Height[to.height] ||
		Height[from.height]+1 == Height[to.height] {
		return true
	}
	return false
}

func (n *Navigator) hasVisited(vector Vector) bool {
	if _, exists := (*n.visited)[vector]; exists {
		return true
	} else {
		return false
	}
}

func (n *Navigator) Up(node *Node) (*Node, error) {
	if node.vector.Y == 0 {
		return &Node{}, fmt.Errorf("at top of territory, no node up")
	}

	up := Vector{X: node.vector.X, Y: node.vector.Y - 1}

	if n.hasVisited(up) {
		return &Node{}, fmt.Errorf("already visited")
	}

	next := &Node{
		vector: up,
		height: n.territory.Height(up),
		next:   []*Node{},
	}

	if isNavigable(node, next) {
		return next, nil
	} else {
		return &Node{}, fmt.Errorf("node is not navigable")
	}
}

func (n *Navigator) Down(node *Node) (*Node, error) {
	if node.vector.Y == len(n.territory)-1 {
		return &Node{}, fmt.Errorf("at bottom of territory, no node up")
	}

	up := Vector{X: node.vector.X, Y: node.vector.Y - 1}

	if n.hasVisited(up) {
		return &Node{}, fmt.Errorf("already visited")
	}

	next := &Node{
		vector: up,
		height: n.territory.Height(up),
		next:   []*Node{},
	}

	if isNavigable(node, next) {
		return next, nil
	} else {
		return &Node{}, fmt.Errorf("node is not navigable")
	}
}

func (n *Navigator) Left(node *Node) (*Node, error) {
	if node.vector.X == 0 {
		return &Node{}, fmt.Errorf("at left edge of territory, no node left")
	}

	up := Vector{X: node.vector.X, Y: node.vector.Y - 1}

	if n.hasVisited(up) {
		return &Node{}, fmt.Errorf("already visited")
	}

	next := &Node{
		vector: up,
		height: n.territory.Height(up),
		next:   []*Node{},
	}

	if isNavigable(node, next) {
		return next, nil
	} else {
		return &Node{}, fmt.Errorf("node is not navigable")
	}
}

func (n *Navigator) Right(node *Node) (*Node, error) {
	if node.vector.X == len(n.territory[0])-1 {
		return &Node{}, fmt.Errorf("at right edge of territory")
	}

	up := Vector{X: node.vector.X, Y: node.vector.Y - 1}

	if n.hasVisited(up) {
		return &Node{}, fmt.Errorf("already visited")
	}

	next := &Node{
		vector: up,
		height: n.territory.Height(up),
		next:   []*Node{},
	}

	if isNavigable(node, next) {
		return next, nil
	} else {
		return &Node{}, fmt.Errorf("node is not navigable")
	}
}

func (n *Navigator) FindNext(node *Node) []*Node {

}

func NewNavigator(input []string) *Navigator {

	territory := NewTerritoryMap(input)
	start := Vector{X: 0, Y: 0}

	navigator := &Navigator{
		territory: territory,
		start: &Node{
			vector: start,
			height: territory.Height(start),
		},
		visited: map[Vector]bool{},
	}

	navigator.Explore()

	return navigator
}

func (n *Navigator) Shortest() int {
	return 0
}
