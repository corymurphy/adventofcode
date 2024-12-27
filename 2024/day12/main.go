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
	Plots map[Position]int
	// Edges int
	Edges map[Position]int
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

func (r *Region) Add(pos Position) {
	if val, contains := r.Plots[pos]; contains {
		r.Plots[pos] = val + 1
	} else {
		r.Plots[pos] = 1
	}
}

func (r *Region) Discovered(pos Position) bool {
	_, contains := r.Plots[pos]
	return contains
}

func DiscoverRegion(input Map, start Position, region *Region) {

	if region.Id.Type != string(input[start.Y][start.X]) {
		// region.Edges++
		region.Edges[start] = region.Edges[start] + 1
		return
	}

	if region.Discovered(start) {
		return
	}

	region.Add(start)

	up := start.Move(Up, 1)
	down := start.Move(Down, 1)
	left := start.Move(Left, 1)
	right := start.Move(Right, 1)

	if input.Contains(up) {
		DiscoverRegion(input, up, region)
	} else {
		// region.Edges++
		region.Edges[start] = region.Edges[start] + 1

	}

	if input.Contains(down) {
		DiscoverRegion(input, down, region)
	} else {
		// region.Edges++
		region.Edges[start] = region.Edges[start] + 1
	}

	if input.Contains(left) {
		DiscoverRegion(input, left, region)
	} else {
		// region.Edges++
		region.Edges[start] = region.Edges[start] + 1
	}

	if input.Contains(right) {
		DiscoverRegion(input, right, region)
	} else {
		region.Edges[start] = region.Edges[start] + 1
	}

}

func (g *Garden) PerimeterCost() (cost int) {
	for _, region := range g.Regions {
		for _, edges := range region.Edges {
			cost = cost + (edges * len(region.Plots))
		}
	}
	return cost
}

func part1(input []string) (answer int) {

	garden := Garden{
		Map: input,
	}

	for y, row := range input {
		for x := range row {

			pos := Position{X: x, Y: y}

			if garden.IsRegionKnown(pos) {
				continue
			}

			region := Region{
				Plots: make(map[Position]int),
				Id: Plot{
					Position: pos,
					Type:     string(input[pos.Y][pos.X]),
				},
				Edges: make(map[Position]int),
			}

			DiscoverRegion(input, pos, &region)

			garden.Regions = append(garden.Regions, region)
		}
	}

	answer = garden.PerimeterCost()

	return answer
}

func part2(input []string) (answer int) {

	garden := Garden{
		Map: input,
	}

	for y, row := range input {
		for x := range row {

			pos := Position{X: x, Y: y}

			if garden.IsRegionKnown(pos) {
				continue
			}

			region := Region{
				Plots: make(map[Position]int),
				Id: Plot{
					Position: pos,
					Type:     string(input[pos.Y][pos.X]),
				},
				Edges: make(map[Position]int),
			}

			DiscoverRegion(input, pos, &region)

			garden.Regions = append(garden.Regions, region)
		}
	}

	// for _, region := range garden.Regions {

	// }

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
