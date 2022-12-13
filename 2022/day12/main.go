package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

// don't look, this is hacky
func part1(input []string) int {
	navigator := NewNavigator(input)

	// fmt.Println("starting shortest search")
	// navigator.DrawVisited()
	return navigator.Shortest()
}

func part2(input []string) int {
	// navigator := NewNavigator(input)
	AllShortest(input)
	// fmt.println()
	return 0
}

type Map [][]string

func NewVisitedMap(input []string) *Map {

	vMap := make(Map, len(input))
	for i, row := range input {
		vMap[i] = make([]string, len(row))
		for j, _ := range row {
			vMap[i][j] = "."
		}
	}
	return &vMap
}

func (m *Map) Height(v Vector) string {
	return (*m)[v.Y][v.X]
}

func NewTerritoryMap(input []string) Map {

	heightmap := make(Map, len(input))

	for i := len(input) - 1; i >= 0; i-- {

		heightmap[i] = make([]string, len(input[i]))

		for j, loc := range input[i] {
			heightmap[i][j] = string(loc)
		}
	}

	return heightmap
}

func NewHeightMap(input []string) *Map {

	heightmap := make(Map, len(input))

	for i := len(input) - 1; i >= 0; i-- {

		heightmap[i] = make([]string, len(input[i]))

		for j, loc := range input[i] {
			heightmap[i][j] = string(loc)
		}
	}

	return &heightmap
}
