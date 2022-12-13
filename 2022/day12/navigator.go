package main

import (
	"fmt"
)

const (
	finish string = "E"
	start  string = "S"
)

type Navigator struct {
	start     *Node
	visited   *map[Vector]bool
	territory Map
	nodes     *Queue
	input     []string
}

// setups up our tree
func (n *Navigator) Explore() {

	n.setVisited(n.start.vector)
	n.nodes.Enqueue(n.start)
	for !n.nodes.Empty() {
		// node := nil
		if node, err := n.nodes.Dequeue(); err == nil {
			// fmt.Printf("exploring node at %v\n", node)
			finish := n.FindNext(node)
			if finish {
				// fmt.Printf("found finish at %v\n", node)
				return
			}
		} else {
			fmt.Println(err)
			panic("this shouldn't happen")
		}
	}

	// fmt.Println(n.visited)
}

func (n *Navigator) setVisited(vector Vector) {
	(*n.visited)[vector] = true
}

// allow two down for some reason
func isNavigable(from *Node, to *Node) bool {

	if Height[from.height]-1 == Height[to.height] ||
		Height[from.height] == Height[to.height] ||
		Height[from.height]+1 == Height[to.height] ||
		Height[from.height]-2 == Height[to.height] {
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
		finish: n.territory.Height(up) == finish,
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

	down := Vector{X: node.vector.X, Y: node.vector.Y + 1}

	if n.hasVisited(down) {
		return &Node{}, fmt.Errorf("already visited")
	}

	next := &Node{
		vector: down,
		height: n.territory.Height(down),
		next:   []*Node{},
		finish: n.territory.Height(down) == finish,
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

	up := Vector{X: node.vector.X - 1, Y: node.vector.Y}

	if n.hasVisited(up) {
		return &Node{}, fmt.Errorf("already visited")
	}

	next := &Node{
		vector: up,
		height: n.territory.Height(up),
		next:   []*Node{},
		finish: n.territory.Height(up) == finish,
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

	right := Vector{X: node.vector.X + 1, Y: node.vector.Y}

	if n.hasVisited(right) {
		return &Node{}, fmt.Errorf("already visited")
	}

	next := &Node{
		vector: right,
		height: n.territory.Height(right),
		next:   []*Node{},
		finish: n.territory.Height(right) == finish,
	}

	if isNavigable(node, next) {
		return next, nil
	} else {
		return &Node{}, fmt.Errorf("node is not navigable")
	}
}

// TODO: refactor to reduce duplication
func (n *Navigator) FindNext(node *Node) bool {

	n.setVisited(node.vector)
	nextNodes := []*Node{}
	foundFinish := false

	// fmt.Println("right")
	if next, err := n.Right(node); err == nil {
		n.nodes.Enqueue(next)
		// fmt.Println(n.nodes.Size())
		nextNodes = append(nextNodes, next)
		n.setVisited(next.vector)
		if next.height == finish {
			// next.finish = true
			foundFinish = true
			// return true
		}
	}

	// fmt.Println("left")
	if next, err := n.Left(node); err == nil {
		n.nodes.Enqueue(next)
		nextNodes = append(nextNodes, next)
		n.setVisited(next.vector)
		if next.height == finish {
			// next.finish = true
			foundFinish = true
			// return true
		}
	}

	// fmt.Println("up")
	if next, err := n.Up(node); err == nil {
		// fmt.Println("up")
		n.nodes.Enqueue(next)
		nextNodes = append(nextNodes, next)
		n.setVisited(next.vector)
		if next.finish {
			foundFinish = true
			// next.finish = true
			// return true
		}
	}

	// fmt.Println("down")
	if next, err := n.Down(node); err == nil {
		n.nodes.Enqueue(next)
		nextNodes = append(nextNodes, next)
		n.setVisited(next.vector)
		if next.height == finish {
			foundFinish = true
			// next.finish = true
			// return true
		}
	}

	node.next = nextNodes
	return foundFinish

}

func FindStarts(territory Map) []Vector {

	starts := []Vector{}
	for col, _ := range territory {
		for row, val := range territory[col] {
			if val == start || val == "a" {
				starts = append(starts, Vector{X: row, Y: col})
				// return Vector{X: row, Y: col}
			}
		}
	}

	return starts

	panic("unable to find start")
	// return Vector{}
}

func FindStart(territory Map) Vector {

	for col, _ := range territory {
		for row, val := range territory[col] {
			if val == start {
				return Vector{X: row, Y: col}
			}
		}
	}

	panic("unable to find start")
	// return Vector{}
}

func (n *Navigator) DrawVisited() {
	graph := NewVisitedMap(n.input)
	for v := range *n.visited {
		(*graph)[v.Y][v.X] = "x"
	}

	fmt.Println("")
	for col, _ := range *graph {
		fmt.Println("")
		for _, val := range (*graph)[col] {
			fmt.Printf("%s", val)
		}
	}
	fmt.Println("")
}

func NewNavigator(input []string) *Navigator {

	territory := NewTerritoryMap(input)
	start := FindStart(territory)

	navigator := &Navigator{
		territory: territory,
		start: &Node{
			vector: start,
			height: territory.Height(start),
		},
		visited: &map[Vector]bool{},
		nodes:   NewQueue(),
		input:   input,
	}

	navigator.Explore()

	return navigator
}

func AllShortest(input []string) {
	territory := NewTerritoryMap(input)

	for _, start := range FindStarts(territory) {
		navigator := &Navigator{
			territory: territory,
			start: &Node{
				vector: start,
				height: territory.Height(start),
			},
			visited: &map[Vector]bool{},
			nodes:   NewQueue(),
			input:   input,
		}

		navigator.Explore()
		navigator.Shortest()
	}
}

func (n *Navigator) Shortest() int {
	distance := 0
	n.Next(n.start, distance)
	return distance
}

func (n *Navigator) Next(node *Node, distance int) {

	distance++

	for _, next := range node.next {
		if next.height == finish {
			// fmt.Println("found finish!")
			fmt.Println(distance)
			// return distance
			// (*distances).append(*distances, distance)
			// return distance
		} else {
			n.Next(next, distance)
		}
	}

	// return
}
