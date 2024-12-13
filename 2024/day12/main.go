package main

import (
	"fmt"

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
	// input   []string
	Map Map
}

type Map []string

type Region struct {
	Id    Plot
	Plots map[Position]int
	Edges int
	Side  int
	// Sides map[int, int]
	RightSides map[Position]int
	LeftSides  map[Position]int
	DownSides  map[Position]int
	UpSides    map[Position]int
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

// func MapRegion(input []string, garden Garden, start Position) (Region )

// func (g *Garden) DiscoverRegion(start)

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

// func Search(input Map, region *Region, pos Position) {

// 	// if
// }

func (r *Region) Add(pos Position) {
	if val, contains := r.Plots[pos]; contains {
		r.Plots[pos] = val + 1
	} else {
		r.Plots[pos] = 1
	}
}

func (r *Region) AddSide(pos Position, dir Direction) {

	if dir == Left {
		// for _, y := range r.LeftSides {
		// 	fmt.Println(y)
		// 	if shared.Abs(pos.Y-y) == 1 {
		// 		return
		// 	}
		// }

		// for side, _ := range r.LeftSides {
		// 	if shared.Abs(side.X-pos.X) == 1 {
		// 		return
		// 	}
		// }

		// if val, contains := r.LeftSides[pos]; contains {
		// 	r.LeftSides[pos] = r.LeftSides[pos] + val
		// } else {
		// 	r.LeftSides[pos] = 1
		// }

		r.LeftSides[pos] = 1
	}

	if dir == Right {

		r.RightSides[pos] = 1
		// for side, _ := range r.RightSides {
		// 	if shared.Abs(side.X-pos.X) == 1 {
		// 		return
		// 	}
		// }

		// if val, contains := r.RightSides[pos]; contains {
		// 	r.RightSides[pos] = r.RightSides[pos] + val
		// } else {
		// 	r.RightSides[pos] = 1
		// }

		// r.RightSides[pos.X] = pos.Y
		// if pos.x is sides with right
	}

	if dir == Up {
		r.UpSides[pos] = 1
	}

	if dir == Down {
		r.DownSides[pos] = 1
	}
}

func (r *Region) Discovered(pos Position) bool {
	_, contains := r.Plots[pos]
	return contains
}

func DiscoverRegion(input Map, start Position, region *Region) {

	// fmt.Println(start)
	// a := input[start.Y][start.X]

	// region = Region{
	// 	Plots: make(map[Position]Plot),
	// 	Id: Plot{
	// 		Position: start,
	// 		Type:     string(input[start.Y][start.X]),
	// 	},
	// }

	// rev
	if region.Id.Type != string(input[start.Y][start.X]) {
		region.Edges++
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

	// fmt.Println(start)

	if input.Contains(up) {
		DiscoverRegion(input, up, region)
	} else {
		region.Edges++

	}

	if input.Contains(down) {
		DiscoverRegion(input, down, region)
	} else {
		region.Edges++

	}

	if input.Contains(left) {
		DiscoverRegion(input, left, region)
	} else {
		region.Edges++

	}

	if input.Contains(right) {
		DiscoverRegion(input, right, region)
	} else {
		region.Edges++
	}

}

func DiscoverRegion2(input Map, start Position, region *Region, dir Direction, previousEdges []Direction) (edge bool) {

	// fmt.Println(start)
	// a := input[start.Y][start.X]

	// region = Region{
	// 	Plots: make(map[Position]Plot),
	// 	Id: Plot{
	// 		Position: start,
	// 		Type:     string(input[start.Y][start.X]),
	// 	},
	// }

	if region.Id.Type != string(input[start.Y][start.X]) {
		// if prev != Unknown {
		region.Edges++
		// }

		// region.AddSide(start, dir)
		return true
	}

	if region.Discovered(start) {
		return false
	}

	region.Add(start)

	// rev := dir.Rev()
	// prev := start.Move(rev, 1)
	up := start.Move(Up, 1)
	down := start.Move(Down, 1)
	left := start.Move(Left, 1)
	right := start.Move(Right, 1)

	// fmt.Println(start)

	if input.Contains(up) {
		// hasEdge := DiscoverRegion2(input, up, region, Up, previousEdges)
	} else {
		region.Edges++
	}

	if input.Contains(down) {
		// hasEdge := DiscoverRegion2(input, down, region, Down, previousEdges)
	} else {
		region.Edges++
	}

	if input.Contains(left) {
		// hasEdge := DiscoverRegion2(input, left, region, Left, previousEdges)
	} else {
		region.Edges++
	}

	if input.Contains(right) {
		hasEdge := DiscoverRegion2(input, right, region, Right, previousEdges)

		if hasEdge {

		}
	} else {
		region.Edges++
	}

}

func (g *Garden) PerimeterCost() (cost int) {
	for _, region := range g.Regions {
		cost = cost + (region.Edges * len(region.Plots))
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
				LeftSides:  map[Position]int{},
				RightSides: map[Position]int{},
				DownSides:  map[Position]int{},
				UpSides:    map[Position]int{},
			}

			DiscoverRegion2(input, pos, &region, Unknown, []Direction{})

			garden.Regions = append(garden.Regions, region)
		}
	}

	for _, region := range garden.Regions {
		// answer = answer
		answer = answer + (region.Edges * len(region.Plots))
		// answer = answer + ((len(region.DownSides) + len(region.UpSides) + len(region.LeftSides) + len(region.RightSides)) * len(region.Plots))

		// fmt.Println(region.Id.Type, region.LeftSides, region.RightSides, region.DownSides, region.UpSides)
		// fmt.Println(region.Id.Type)
		// fmt.Println(region)

		// region.CountSides()
	}

	return answer
}

func (r *Region) CountSides() (count int) {

	fmt.Print(r.Id.Type)

	fmt.Println(count)

	return 0
}
