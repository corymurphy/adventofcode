package main

import (
	"testing"

	shared "github.com/corymurphy/adventofcode/shared"
)

func Test_Part1(t *testing.T) {
	expected := 26
	input := shared.ReadInput("input_test")
	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

// func Test_Sensor_Area(t *testing.T) {
// 	loc := NewCoordinates(2, 2)
// 	beacon := NewCoordinates(4, 4)
// 	Sensor := NewSensor(loc, beacon)

// 	expectedCoordinates := []Coordinate{
// 		NewCoordinates(0, 4),
// 		NewCoordinates(1, 3),
// 		NewCoordinates(2, 2),
// 		NewCoordinates(0, -4),
// 		NewCoordinates(1, -3),
// 		NewCoordinates(2, -2),
// 		NewCoordinates(0, 4),
// 		NewCoordinates(-1, 3),
// 		NewCoordinates(-2, 2),
// 		NewCoordinates(0, -4),
// 		NewCoordinates(1, -3),
// 		NewCoordinates(2, -2),
// 	}

// 	area := Sensor.Area()
// 	fmt.Println(area)
// 	shared.AssertEqual(t, len(expectedCoordinates), len(area))

// 	for i, expected := range expectedCoordinates {
// 		shared.AssertEqual(t, expected.X, area[i].X)
// 		shared.AssertEqual(t, expected.Y, area[i].Y)
// 	}

// }

// func Test_Part2(t *testing.T) {
// 	expected := 93
// 	input := shared.ReadInput("input_test")
// 	actual := part2(input)
// 	shared.AssertEqual(t, expected, actual)
// }

// func Test_Part1_Completed(t *testing.T) {
// 	expected := 793
// 	input := shared.ReadInput("input")
// 	actual := part1(input)
// 	shared.AssertEqual(t, expected, actual)
// }

// func Test_Part2_Completed(t *testing.T) {
// 	expected := 24165
// 	input := shared.ReadInput("input")
// 	actual := part2(input)
// 	shared.AssertEqual(t, expected, actual)
// }
