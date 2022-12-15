package main

import (
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

const (
	mod int = 0
)

type Vector struct {
	X int
	Y int
}

func NewVector(x int, y int) *Vector {
	return &Vector{
		X: x + mod,
		Y: y,
	}
}

func ParseNewVector(input string) *Vector {
	vector := strings.Split(input, ",")
	return &Vector{
		X: shared.ToInt(vector[0]),
		Y: shared.ToInt(vector[1]),
	}
}
