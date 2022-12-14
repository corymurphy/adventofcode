package main

import "fmt"

type Cave struct {
	Rocks  []*RockStructure
	Graph  [][]string
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
	// fmt.Println(*limits)

	// fmt.Println(cave.Rocks)

	return cave
}

func (c *Cave) PlotRocks() {

	for _, rock := range c.Rocks {
		for i, v := range rock.Vectors {

			c.Graph[v.Y][v.X] = "#"

			if i > 0 {
				prev := rock.Vectors[i-1]
				c.PlotRockLine(prev, v)
			}

		}
	}
}

func (c *Cave) Draw() {

	// fmt.Println(len(c.Graph))
	// fmt.Println(len(c.Graph[0]))
	for y := c.Limits.MinY(); y <= c.Limits.maxY; y++ {
		fmt.Println("")
		for x := c.Limits.MinX(); x <= c.Limits.maxX; x++ {
			fmt.Printf("%s", c.Graph[y][x])
		}
	}
}

func (c *Cave) StartSand(start *Vector) int {
	c.Graph[start.Y][start.X] = "+"
	c.Limits.Analyze([]*Vector{start})
	c.Draw()

	return 0
}

func NewGraph(limits *Limits) [][]string {

	graph := [][]string{}

	for y := 0; y <= limits.maxX; y++ {
		row := []string{}
		for x := 0; x <= limits.maxX; x++ {
			row = append(row, ".")
		}
		graph = append(graph, row)
	}
	return graph
}

func (c *Cave) PlotRockLine(start *Vector, finish *Vector) {

	// you know what, this is bad, and I feel bad

	// dif vertical
	yDiff := start.Y - finish.Y
	if yDiff != 0 {

		if yDiff < 0 { // do down
			for y := start.Y; y < finish.Y; y++ {
				c.Graph[y][start.X] = "#"
			}
		} else { // go up
			for y := start.Y; y < finish.Y; y-- {
				c.Graph[y][start.X] = "#"
			}
		}

	}

	// dif horizontal
	xDiff := start.X - finish.X
	if xDiff != 0 {
		if xDiff < 0 { // go left
			for x := start.X; x < finish.X; x++ {
				c.Graph[start.Y][x] = "#"
			}
		} else { // go right
			for x := start.X; x > finish.X; x-- {
				c.Graph[start.Y][x] = "#"
			}
		}
	}
}
