package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

type Direction int

const (
	Unknown Direction = iota
	Up
	Down
	Left
	Right
)

func GetDirection(input rune) Direction {
	switch input {
	case '<':
		return Left
	case '>':
		return Right
	case '^':
		return Up
	case 'v':
		return Down
	}
	panic("invalid direction")
}

func (d Direction) Rune() rune {
	switch d {
	case Up:
		return '^'
	case Down:
		return 'v'
	case Left:
		return '<'
	case Right:
		return '>'
	}
	panic("invalid direction")
}

func (d Direction) String() string {
	switch d {
	case Up:
		return "up"
	case Down:
		return "down"
	case Left:
		return "left"
	case Right:
		return "right"
	}
	panic("invalid direction")
}

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

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

type Warehouse [][]rune
type Directions []Direction

type Position struct {
	X int
	Y int
}

func (p Position) Move(direction Direction) Position {
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

func (w *Warehouse) Print() {
	fmt.Println()
	for _, row := range *w {
		for _, val := range row {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
}

func (d *Directions) Print() {
	for _, dir := range *d {
		fmt.Println(dir.String())
	}
}

func NewWarehouse(input []string) (warehouse Warehouse) {

	warehouse = [][]rune{}
	for _, row := range input {
		x := []rune{}
		for _, val := range row {
			x = append(x, val)
		}
		warehouse = append(warehouse, x)
	}
	return warehouse
}

func NewDirections(input []string) (directions Directions) {

	directions = []Direction{}
	for _, row := range input {
		for _, val := range row {
			directions = append(directions, GetDirection(val))
		}
	}
	return
}

func Initialize(input []string) (directions Directions, warehouse Warehouse) {

	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			return NewDirections(input[i+1:]), NewWarehouse(input[:i])
		}
	}
	panic("this shouldn't happen")
}

func GetStart(w *Warehouse) (start Position) {

	for y, row := range *w {
		for x, val := range row {
			if val == '@' {
				return Position{
					X: x,
					Y: y,
				}
			}
		}
	}
	panic("shouldn't get here")
}

func (w *Warehouse) Contains(pos Position) bool {

	if pos.Y < 0 {
		return false
	}

	if pos.X < 0 {
		return false
	}

	if pos.Y >= len(*w) {
		return false
	}

	if pos.X >= len((*w)[pos.Y]) {
		return false
	}

	return true
}

func (w *Warehouse) PrintSimulation() {
	for y, row := range *w {
		for x, val := range row {

			// out := val
			// if val == 0 {
			// 	out = "."
			// }

			fmt.Printf("\033[%d;%dH%s", y+5, x+2, string(val))
		}
	}
	fmt.Println()
}

func (w *Warehouse) Move(p Position, dir Direction, item rune) (next Position) {
	// is position in warehosue?

	next = p.Move(dir)

	if (*w)[next.Y][next.X] == '#' {
		return p
	}

	if (*w)[next.Y][next.X] == 'O' {
		w.Move(next, dir, 'O')
	}

	if (*w)[next.Y][next.X] == 'O' {
		return p
	}

	(*w)[p.Y][p.X] = '.'
	(*w)[next.Y][next.X] = item

	return next

}

func Simulate(robot Position, w *Warehouse, d Directions) {

	for _, dir := range d {

		robot = w.Move(robot, dir, '@')

		// w.PrintSimulation()
		// time.Sleep(250 * time.Millisecond)
	}

	// w.PrintSimulation()
}

func (w *Warehouse) BoxGPS() (answer int) {
	for y, row := range *w {
		for x, val := range row {
			if val == 'O' {
				answer = answer + ((y * 100) + x)
			}
		}
	}
	return answer
}

func part1(input []string) (answer int) {

	d, w := Initialize(input)
	Simulate(GetStart(&w), &w, d)
	answer = w.BoxGPS()
	return answer
}

func part2(input []string) (answer int) {

	return answer
}
