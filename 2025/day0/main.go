package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)

}

func part1(inputs []string) (answer int) {

	rotations := []Rotation{}
	for _, input := range inputs {
		if input[0] == 'L' {
			rotations = append(rotations, Rotation{Direction: Left, Count: shared.ToInt(input[1:])})
			continue
		}

		if input[0] == 'R' {
			rotations = append(rotations, Rotation{Direction: Right, Count: shared.ToInt(input[1:])})
			continue
		}
	}

	padLock := NewPadLock(50)
	for _, rotation := range rotations {
		padLock.Rotate(rotation)
	}

	answer = padLock.WasZero
	return answer
}

func part2(inputs []string) (answer int) {

	rotations := []Rotation{}
	for _, input := range inputs {
		if input[0] == 'L' {
			rotations = append(rotations, Rotation{Direction: Left, Count: shared.ToInt(input[1:])})
			continue
		}

		if input[0] == 'R' {
			rotations = append(rotations, Rotation{Direction: Right, Count: shared.ToInt(input[1:])})
			continue
		}
	}

	padLock := NewPadLock(50)
	for _, rotation := range rotations {
		padLock.Rotate2(rotation)
	}
	answer = padLock.WasZero
	return answer
}

type Direction int

const (
	Left Direction = iota
	Right
)

type Rotation struct {
	Direction Direction
	Count     int
}

type PadLock struct {
	WasZero  int
	Position int
}

func NewPadLock(position int) PadLock {
	return PadLock{
		WasZero:  0,
		Position: position,
	}
}

func (p *PadLock) Rotate2(r Rotation) {

	n := -1
	c := r.Count
	o := 100

	if c >= 100 {
		p.WasZero = p.WasZero + (c / 100)
		o = (c / 100) * 100
		c = c - o
	}

	if r.Direction == Right {
		n = p.Position + c

		if n > 100 {
			n = n - 100
			p.WasZero = p.WasZero + 1
		}
		if n == 100 {
			n = n - 100
		}
	}

	if r.Direction == Left {

		initial := p.Position
		n = p.Position - c

		if n < 0 {
			n = n + 100
			if initial != 0 {
				p.WasZero = p.WasZero + 1
			}
		}
	}

	if n == 0 {
		p.WasZero = p.WasZero + 1
	}

	p.Position = n

}

func (p *PadLock) Rotate(r Rotation) {

	n := 0
	c := r.Count
	o := 100

	if c >= 100 {
		o = (c / 100) * 100
		c = c - o
	}

	if r.Direction == Right {
		n = p.Position + c

		if n >= 100 {
			n = n - 100
		}
	}

	if r.Direction == Left {
		n = p.Position - c
		if n < 0 {
			n = n + 100
		}
	}

	if n == 0 {
		p.WasZero = p.WasZero + 1
	}

	p.Position = n

}
