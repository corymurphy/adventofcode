package main

type Coordinate struct {
	X int
	Y int
}

func NewCoordinates(y int, x int) Coordinate {
	return Coordinate{
		X: x,
		Y: y,
	}
}
