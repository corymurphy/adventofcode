package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

type Space [][]rune

type Position struct {
	X int
	Y int
}

func (p *Position) Up() (move Position) {
	return Position{
		Y: p.Y - 1,
		X: p.X,
	}
}

func (p *Position) Down() (move Position) {
	return Position{
		Y: p.Y + 1,
		X: p.X,
	}
}

func (p *Position) Left() (move Position) {
	return Position{
		X: p.X - 1,
		Y: p.Y,
	}
}

func (p *Position) Right() (move Position) {
	return Position{
		X: p.X + 1,
		Y: p.Y,
	}
}

type Route struct {
	Position   Position
	ReachedEnd bool
	Score      int
}

func NewSpace(input []string, bytes int) (space *Space) {
	space = &Space{}

	size := 0

	fallen := map[Position]int{}

	for _, line := range input {

		x := shared.ToInt(strings.Split(line, ",")[0])
		y := shared.ToInt(strings.Split(line, ",")[1])

		if x > size {
			size = x
		}
		if y > size {
			size = y
		}
	}

	for i := 0; i < bytes; i++ {
		line := input[i]

		x := shared.ToInt(strings.Split(line, ",")[0])
		y := shared.ToInt(strings.Split(line, ",")[1])

		if x > size {
			size = x
		}
		if y > size {
			size = y
		}
		pos := Position{
			X: x,
			Y: y,
		}

		fallen[pos] = 0
	}

	for y := range size + 1 {
		row := []rune{}
		for x := range size + 1 {
			pos := Position{X: x, Y: y}
			if _, contains := fallen[pos]; contains {
				row = append(row, '#')
				continue
			}
			row = append(row, '.')

		}
		*space = append(*space, row)
	}
	return space
}

func (s *Space) Print() {
	fmt.Println()
	for _, row := range *s {
		for _, val := range row {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
}

func (s *Space) IsMoveValid(move Position) bool {
	if move.Y < 0 || move.X < 0 {
		return false
	}

	if move.Y >= len(*s) {
		return false
	}

	if move.X >= len((*s)[move.Y]) {
		return false
	}

	if (*s)[move.Y][move.X] == '.' {
		return true
	}

	return false
}

func (s *Space) PrintSimulation() {
	for y, row := range *s {
		for x, val := range row {
			fmt.Printf("\033[%d;%dH%s", y+5, x+2, string(val))
		}
	}
	fmt.Println()
}

func Simulate(start *Route, end Position, space *Space) (routes []*Route) {
	routes = []*Route{start}

	queue := NewQueue()
	queue.Enqueue(start)
	for !queue.Empty() {

		fmt.Println(queue.Size())
		route, err := queue.Dequeue()

		space.PrintSimulation()

		if err != nil {
			panic(err)
		}

		if (*space)[route.Position.Y][route.Position.X] != '.' {
			continue
		}

		(*space)[route.Position.Y][route.Position.X] = 'O'

		if end.X == route.Position.X && end.Y == route.Position.Y {
			route.ReachedEnd = true
			break
		}

		route.Score++

		for i, next := range []Position{
			route.Position.Up(),
			route.Position.Down(),
			route.Position.Left(),
			route.Position.Right(),
		} {

			if !space.IsMoveValid(next) {
				continue
			}

			if i == 0 {
				route.Position = next
				queue.Enqueue(route)
				continue
			}

			forked := &Route{
				Position: next,
				Score:    route.Score,
			}
			routes = append(routes, forked)
			queue.Enqueue(forked)
		}
	}

	return routes
}

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) (answer int) {

	space := NewSpace(input, 12)

	start := &Route{
		Position: Position{X: 0, Y: 0},
	}
	end := Position{
		X: len(*space) - 1,
		Y: len(*space) - 1,
	}

	routes := Simulate(start, end, space)

	shortest := math.MaxInt
	for _, route := range routes {
		if !route.ReachedEnd {
			continue
		}
		if route.Score < shortest {
			shortest = route.Score
		}
	}

	answer = shortest

	return answer
}

func part2(input []string) (answer int) {

	bytes := len(input)
	exitable := false

	for !exitable {
		space := NewSpace(input, bytes)

		bytes--

		start := &Route{
			Position: Position{X: 0, Y: 0},
		}
		end := Position{
			X: len(*space) - 1,
			Y: len(*space) - 1,
		}

		routes := Simulate(start, end, space)

		for _, route := range routes {
			if route.ReachedEnd {
				exitable = true
			}
		}

	}

	fmt.Println(input[bytes+1])

	answer = bytes

	return answer
}
