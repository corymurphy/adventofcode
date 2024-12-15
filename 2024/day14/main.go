package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

type Space [][]int
type Robots []Robot

type Robot struct {
	Position Position
	Velocity Velocity
}

type Position struct {
	X int
	Y int
}

type Velocity struct {
	X int
	Y int
}

func NewSpace(width int, height int) (s Space) {

	for y := 0; y < height; y++ {
		row := []int{}
		for x := 0; x < width; x++ {
			row = append(row, 0)
		}
		s = append(s, row)
	}
	return s
}

func NewRobots(input []string) (r Robots) {
	for _, line := range input {
		r = append(r, *NewRobot(line))
	}
	return r
}

func NewRobot(input string) (r *Robot) {

	p := strings.Split(strings.Split(input, " ")[0], "=")[1]
	v := strings.Split(strings.Split(input, " ")[1], "=")[1]

	return &Robot{
		Position: Position{
			X: shared.ToInt(strings.Split(p, ",")[0]),
			Y: shared.ToInt(strings.Split(p, ",")[1]),
		},
		Velocity: Velocity{
			X: shared.ToInt(strings.Split(v, ",")[0]),
			Y: shared.ToInt(strings.Split(v, ",")[1]),
		},
	}
}

func (s Space) Print() {
	fmt.Println()
	for _, row := range s {
		for _, val := range row {
			if val == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(val)
			}
		}
		fmt.Println()
	}
}

func (s Space) PrintSimulation() {
	for y, row := range s {
		for x, val := range row {

			out := strconv.Itoa(val)
			if val == 0 {
				out = "."
			}

			fmt.Printf("\033[%d;%dH%s", y+5, x+2, out)
		}
	}
	fmt.Println()
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (r *Robot) Move(limits Position) {

	x := r.Position.X + r.Velocity.X
	y := r.Position.Y + r.Velocity.Y

	if x < 0 {
		x = mod(x, limits.X)
	}
	if x >= limits.X {
		x = x % limits.X
	}
	if y < 0 {
		y = mod(y, limits.Y)
	}
	if y >= limits.Y {
		y = y % limits.Y
	}

	r.Position.X = x
	r.Position.Y = y
}

func (s *Space) Remove(at *Position) {
	current := (*s)[at.Y][at.X]

	if current == 0 {
		return
	}

	(*s)[at.Y][at.X] = current - 1
}

func (s *Space) Add(at *Position) {
	(*s)[at.Y][at.X] = (*s)[at.Y][at.X] + 1
}

func (rs *Robots) Simulate(space Space, seconds int) {

	limits := Position{
		X: len(space[0]),
		Y: len(space),
	}

	center := Position{
		X: len(space[0]) / 2,
		Y: len(space) / 2,
	}
	count := len(*rs)

	avg := math.MaxInt64
	a := 0
	for i := 0; i < seconds; i++ {

		sum := 0
		for r := range *rs {

			// sum = ( - )

			robot := (*rs)[r]
			space.Remove(&robot.Position)
			robot.Move(limits)
			space.Add(&robot.Position)

			sum = sum + shared.Abs(center.X-robot.Position.X) + shared.Abs(center.Y-robot.Position.Y)

			(*rs)[r] = robot
		}
		na := sum / count
		if na < avg {
			avg = na

		}

		// if na == 32 {
		// 	a = i
		// 	// break
		// }

		// fmt.Printf("\033[%d;%dH%s", 1, 0, strconv.Itoa(i))
		// if i == 8052 {
		// 	space.PrintSimulation()
		// }
		// space.PrintSimulation()
		// time.Sleep(10 * time.Millisecond)
	}

	fmt.Println(avg)
	fmt.Println(a)
}

func (s *Space) SafetyFactor() (answer int) {

	// top left
	// fmt.Println()
	qs := 0
	for _, row := range (*s)[:(len(*s)-1)/2] {
		for _, val := range row[:(len(row)-1)/2] {
			qs += val
			// if val == 0 {
			// 	fmt.Print(".")
			// } else {
			// 	fmt.Print(val)
			// }
			// fmt.Print(val)
		}
		// fmt.Println()
	}
	// fmt.Println()

	answer = qs

	// bottom right
	// fmt.Println()
	qs = 0
	for _, row := range (*s)[(len(*s)+1)/2:] {
		for _, val := range row[:(len(row)-1)/2] {
			qs += val
			// if val == 0 {
			// 	fmt.Print(".")
			// } else {
			// 	fmt.Print(val)
			// }
			// fmt.Print(val)
		}
		// fmt.Println()
	}
	// fmt.Println()

	answer = answer * qs

	// top right
	// fmt.Println()
	qs = 0
	for _, row := range (*s)[:(len(*s)-1)/2] {
		for _, val := range row[(len(row)+1)/2:] {
			qs += val
			// if val == 0 {
			// 	fmt.Print(".")
			// } else {
			// 	fmt.Print(val)
			// }
			// fmt.Print(val)
		}
		// fmt.Println()
	}
	// fmt.Println()

	answer = answer * qs

	// bottom right
	// fmt.Println()
	qs = 0
	for _, row := range (*s)[(len(*s)+1)/2:] {
		for _, val := range row[(len(row)+1)/2:] {
			qs += val
			// if val == 0 {
			// 	fmt.Print(".")
			// } else {
			// 	fmt.Print(val)
			// }
		}
		// fmt.Println()
	}
	// fmt.Println()

	answer = answer * qs

	return answer
}

func part1(input []string) (answer int) {

	// robots := NewRobots(input)

	// space := NewSpace(101, 103)

	// robots.Simulate(space, 100)

	// answer = space.SafetyFactor()

	return answer
}

func part2(input []string) (answer int) {

	robots := NewRobots(input)

	space := NewSpace(101, 103)

	robots.Simulate(space, 15000)

	// answer = space.SafetyFactor()

	return answer
}
