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

	for y := c.Limits.MinY(); y <= c.Limits.maxY+2; y++ {
		// for y := c.Limits.MinY(); y <= 510; y++ {
		fmt.Println("")
		for x := c.Limits.MinX() - 10; x <= c.Limits.maxX+10; x++ {
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

	// fmt.Println(c.Limits.maxY)
	last := NewVector(start.X, start.Y)
	sands := 0
	for !c.InAbyss(*last) {

		// if last.Y == start.Y && last.X == start.X {
		// 	break
		// }

		sand := NewSand(start, c.Graph)
		sand.Drop()

		last = NewVector(sand.finish.X, sand.finish.Y)
		if c.InAbyss(*last) {
			break
		}
		sands++
	}

	// c.Draw()

	return sands
}

func (c *Cave) StartSand2(start *Vector) int {
	(*c.Graph)[start.Y][start.X] = "+"
	c.Limits.Analyze([]*Vector{start})

	// fmt.Println(c.Limits.maxY)
	last := NewVector(start.X, start.Y+1)
	sands := 0
	for !c.FullToStart(*start, *last) {

		// if last.Y == start.Y && last.X == start.X {
		// 	break
		// }

		sand := NewSand(start, c.Graph)
		sand.Drop()

		last = NewVector(sand.finish.X, sand.finish.Y)
		if c.FullToStart(*start, *last) {
			break
		}
		sands++
		// fmt.Println(sands)
	}

	c.Draw()

	return sands
}

func (c *Cave) FullToStart(start Vector, v Vector) bool {
	return start.X == v.X && start.Y == v.Y
}

func NewGraph(limits *Limits) *Graph {

	graph := Graph{}

	for y := 0; y <= limits.maxY+2; y++ {
		row := []string{}
		for x := 0; x <= limits.maxX+10000; x++ {
			row = append(row, ".")
		}
		graph = append((graph), row)
	}
	return &graph
}
func (c *Cave) PlotFloor() {
	y := c.Limits.maxY + 2

	for x := range (*c.Graph)[y] {
		(*c.Graph)[y][x] = "#"
	}
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

	c.PlotFloor()
}

func (c *Cave) PlotRockLine(start Vector, finish Vector) {

	// you know what, this is bad, and I feel bad

	// dif vertical

	(*c.Graph)[start.Y][start.X] = "#"
	(*c.Graph)[finish.Y][finish.X] = "#"

	yDiff := start.Y - finish.Y

	if yDiff != 0 {

		if yDiff < 0 { // go down
			for y := start.Y; y < finish.Y; y++ {
				(*c.Graph)[y][start.X] = "#"
			}
		} else { // go up
			for y := start.Y; y > finish.Y; y-- {
				(*c.Graph)[y][start.X] = "#"
			}
		}

	}

	xDiff := start.X - finish.X

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
