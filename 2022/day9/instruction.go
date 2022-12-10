package main

import (
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

type Direction int

const (
	Unknown Direction = 0
	Up      Direction = 1
	Down    Direction = 2
	Left    Direction = 3
	Right   Direction = 4
)

type Instruction struct {
	Distance  int
	Direction Direction
}

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		return "Unknown"
	}
}

func parseDirection(input string) Direction {
	switch input {
	case "U":
		return Up
	case "D":
		return Down
	case "L":
		return Left
	case "R":
		return Right
	default:
		return Unknown
	}
}

func NewInstruction(input string) Instruction {
	instruction := strings.Split(input, " ")
	return Instruction{
		Distance:  shared.ToInt(instruction[1]),
		Direction: parseDirection(instruction[0]),
	}
}
