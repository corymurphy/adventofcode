package main

import (
	"fmt"
	"math"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input_sample_medium2")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

type Direction int

const (
	Unknown Direction = iota
	Up
	Down
	Left
	Right
)

func (d Direction) Reverse() Direction {
	switch d {
	case Unknown:
		return Unknown
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	}
	panic("invalid direction")
}

type Garden struct {
	Regions []Region
	Map     Map
}

type Map []string

type Region struct {
	Id    Plot
	Plots map[Position]bool
	Edges map[Edge]bool
}

type Edge struct {
	Position  Position
	Direction Direction
}

func NewEdge(pos Position, dir Direction) (edge Edge) {
	return Edge{
		Position:  pos,
		Direction: dir,
	}
}

type Plot struct {
	Position Position
	Type     string
}

type Position struct {
	Y int
	X int
}

func (p Position) Move(direction Direction, count int) Position {
	if direction == Up {
		return Position{X: p.X, Y: p.Y - 1}
	}

	if direction == Down {
		return Position{X: p.X, Y: p.Y + 1}
	}

	if direction == Left {
		return Position{X: p.X - 1, Y: p.Y}
	}

	if direction == Right {
		return Position{X: p.X + 1, Y: p.Y}
	}

	return Position{X: p.X, Y: p.Y}
}

func (g *Garden) IsRegionKnown(pos Position) bool {

	for _, region := range g.Regions {
		if _, contains := region.Plots[pos]; contains {
			return true
		}
	}
	return false
}

func (m *Map) Contains(pos Position) bool {

	if pos.Y < 0 {
		return false
	}

	if pos.X < 0 {
		return false
	}

	if pos.Y >= len(*m) {
		return false
	}

	if pos.X >= len((*m)[pos.Y]) {
		return false
	}

	return true
}

func DiscoverRegion(garden Map, pos Position, region *Region) {

	if region.Plots[pos] {
		return
	}
	region.Plots[pos] = true

	for _, next := range []Edge{
		NewEdge(pos.Move(Up, 1), Up),
		NewEdge(pos.Move(Down, 1), Down),
		NewEdge(pos.Move(Left, 1), Left),
		NewEdge(pos.Move(Right, 1), Right),
	} {
		if !garden.Contains(next.Position) {
			region.Edges[NewEdge(pos, next.Direction)] = true
			continue
		}

		if region.Id.Type != string(garden[next.Position.Y][next.Position.X]) {
			region.Edges[NewEdge(pos, next.Direction)] = true
			continue
		}

		DiscoverRegion(garden, next.Position, region)
	}
}

func (g *Garden) PerimeterCost() (cost int) {
	for _, region := range g.Regions {
		cost = cost + (len(region.Edges) * len(region.Plots))
	}
	return cost
}

func (g *Garden) SideCost(r *Region) (cost int) {

	// left
	for y := 0; y < len(g.Map); y++ {
		for x := 0; x < len(g.Map[y]); x++ {
			if r.Edges[NewEdge(Position{X: x, Y: y}, Left)] && !r.Edges[NewEdge(Position{X: x, Y: y - 1}, Left)] {
				cost++
				x++
				for ; x < len(g.Map[y]); x++ {
					if (!r.Edges[NewEdge(Position{X: x, Y: y}, Left)] && !r.Edges[NewEdge(Position{X: x, Y: y - 1}, Left)]) {
						break
					}
				}
			}
		}
	}

	// right
	for y := len(g.Map) - 1; y >= 0; y-- {
		for x := 0; x < len(g.Map[y]); x++ {
			if r.Edges[NewEdge(Position{X: x, Y: y}, Right)] && !r.Edges[NewEdge(Position{X: x, Y: y + 1}, Right)] {
				cost++
				x++
				for ; x < len(g.Map[y]); x++ {
					if (!r.Edges[NewEdge(Position{X: x, Y: y}, Right)] && !r.Edges[NewEdge(Position{X: x, Y: y + 1}, Right)]) {
						break
					}
				}
			}
		}
	}

	// up
	for x := 0; x < len(g.Map[0]); x++ {
		for y := 0; y < len(g.Map); y++ {
			if r.Edges[NewEdge(Position{X: x, Y: y}, Up)] && !r.Edges[NewEdge(Position{X: x - 1, Y: y}, Up)] {
				cost++
				y++
				for ; y < len(g.Map); y++ {
					if (!r.Edges[NewEdge(Position{X: x, Y: y}, Up)] && !r.Edges[NewEdge(Position{X: x - 1, Y: y}, Up)]) {
						break
					}
				}
			}
		}
	}

	// down
	for x := len(g.Map[0]) - 1; x >= 0; x-- {
		for y := 0; y < len(g.Map); y++ {
			if r.Edges[NewEdge(Position{X: x, Y: y}, Down)] && !r.Edges[NewEdge(Position{X: x + 1, Y: y}, Down)] {
				cost++
				y++
				for ; y < len(g.Map); y++ {
					if (!r.Edges[NewEdge(Position{X: x, Y: y}, Down)] && !r.Edges[NewEdge(Position{X: x + 1, Y: y}, Down)]) {
						break
					}
				}
			}
		}
	}

	return cost
}

func (g *Garden) Explore() {
	for y, row := range g.Map {
		for x := range row {

			pos := Position{X: x, Y: y}

			if g.IsRegionKnown(pos) {
				continue
			}

			region := Region{
				Plots: make(map[Position]bool),
				Id: Plot{
					Position: pos,
					Type:     string(g.Map[pos.Y][pos.X]),
				},
				Edges: make(map[Edge]bool),
			}

			DiscoverRegion(g.Map, pos, &region)

			g.Regions = append(g.Regions, region)
		}
	}
}

func part1(input []string) (answer int) {

	garden := Garden{
		Map: input,
	}
	garden.Explore()

	answer = garden.PerimeterCost()

	return answer
}

func part2(input []string) (answer int) {

	garden := Garden{
		Map: input,
	}
	garden.Explore()

	for _, region := range garden.Regions {
		answer += garden.SideCost(&region) * len(region.Plots)
	}

	return answer
}

func (r *Region) GetFrame() (frame [][]string) {
	yMax := 0
	yMin := math.MaxInt
	xMax := 0
	xMin := math.MaxInt

	for plot := range r.Plots {
		if plot.Y < yMin {
			yMin = plot.Y
		}

		if plot.X < xMin {
			xMin = plot.X
		}

		if plot.X > xMax {
			xMax = plot.X
		}

		if plot.Y > yMax {
			yMax = plot.Y
		}
	}

	height := yMax - yMin
	width := xMax - xMin

	frame = [][]string{}
	for y := range height + 1 {
		frame = append(frame, []string{})
		for range width + 1 {
			frame[y] = append(frame[y], ".")
		}
	}

	return frame
}

func (r *Region) Print() {

	yMin := math.MaxInt
	xMin := math.MaxInt

	for plot := range r.Plots {
		if plot.Y < yMin {
			yMin = plot.Y
		}

		if plot.X < xMin {
			xMin = plot.X
		}
	}

	frame := r.GetFrame()

	for plot := range r.Plots {
		frame[plot.Y-yMin][plot.X-xMin] = r.Id.Type
	}

	for _, row := range frame {
		for _, val := range row {
			if val != "." {
				fmt.Print(val)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()

}
