package main

import "fmt"

type Cave struct {
	Rocks  []*RockStructure
	Graph  *Graph
	Limits Limits
}

func NewCave(input []string) *Cave {

	limits := NewLimits()
	rocks := []*RockStructure{}

	for _, line := range input {
		rock := NewRockStructure(line)
		limits.Analyze(rock.Vectors)
		rocks = append(rocks, rock)
	}

	cave := &Cave{
		Rocks:  rocks,
		Graph:  NewGraph(limits),
		Limits: *limits,
	}

	cave.PlotRocks()
	return cave
}

func (c *Cave) Draw() {

	for y := c.Limits.MinY(); y <= c.Limits.maxY; y++ {
		// for y := c.Limits.MinY(); y <= 510; y++ {
		fmt.Println("")
		for x := c.Limits.MinX(); x <= c.Limits.maxX; x++ {
			fmt.Printf("%s", (*c.Graph)[y][x])
		}
	}
}

func (c *Cave) InAbyss(v Vector) bool {
	return v.Y >= c.Limits.maxY
	// return false
}

func (c *Cave) StartSand(start *Vector) int {
	(*c.Graph)[start.Y][start.X] = "+"
	c.Limits.Analyze([]*Vector{start})

	fmt.Println(c.Limits.maxY)
	last := NewVector(start.X, start.Y)
	sands := 0
	for !c.InAbyss(*last) {

		sand := NewSand(start, c.Graph)
		sand.Drop()

		last = NewVector(sand.finish.X, sand.finish.Y)
		if c.InAbyss(*last) {
			break
		}
		sands++
	}

	c.Draw()

	return sands
}

func NewGraph(limits *Limits) *Graph {

	graph := Graph{}

	for y := 0; y <= limits.maxX; y++ {
		row := []string{}
		for x := 0; x <= limits.maxX; x++ {
			row = append(row, ".")
		}
		graph = append((graph), row)
	}
	return &graph
}

func (c *Cave) PlotRocks() {

	for _, rock := range c.Rocks {

		for i, v := range rock.Vectors {

			if i > 0 {
				// prev :=
				c.PlotRockLine(*rock.Vectors[i-1], *v)
				// (*c.Graph)[prev.Y][prev.X] = "#"

				// if v.Y == 66 {
				// 	fmt.Println("here")
				// }
			}

			// (*c.Graph)[v.Y][v.X] = "#"

		}
	}
}

func (c *Cave) PlotRockLine(start Vector, finish Vector) {

	// you know what, this is bad, and I feel bad

	// dif vertical

	(*c.Graph)[start.Y][start.X] = "#"
	(*c.Graph)[finish.Y][finish.X] = "#"

	yDiff := start.Y - finish.Y

	if start.Y > 65 && start.Y < 78 {
		fmt.Printf("y diff %d\n", yDiff)
	}

	if yDiff != 0 {

		if yDiff < 0 { // go down
			for y := start.Y; y < finish.Y; y++ {
				(*c.Graph)[y][start.X] = "#"
			}
		} else { // go up
			if start.X == 466 || start.X == 462 {
				fmt.Println("here")
			}
			for y := start.Y; y > finish.Y; y-- {
				(*c.Graph)[y][start.X] = "#"
			}
		}

	}

	xDiff := start.X - finish.X

	if start.Y > 65 && start.Y < 78 {
		fmt.Printf("x diff %d\n", xDiff)
	}

	if xDiff != 0 {
		if xDiff < 0 { // go left
			for x := start.X; x < finish.X; x++ {
				(*c.Graph)[start.Y][x] = "#"
			}
		} else { // go right
			for x := start.X; x > finish.X; x-- {
				(*c.Graph)[start.Y][x] = "#"
			}
		}
	}
}
