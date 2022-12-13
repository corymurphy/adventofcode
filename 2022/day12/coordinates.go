package main

type Vector struct {
	X int
	Y int
}

func NewCoordinates(x int, y int) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}

// func NewCoordinatesList(count int) []*Coordinates {
// 	cords := []*Coordinates{}
// 	for i := 0; i < count; i++ {
// 		cords = append(cords, NewCoordinates())
// 	}
// 	return cords
// }
