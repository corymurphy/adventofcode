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

func play1(input []string) {
	heightMap := NewHeightMap(input)
	visitedMap := NewVisitedMap(input)
	fmt.Println(heightMap)
	fmt.Println(visitedMap)

	loc := NewCoordinates(5, 6)

	visited := map[Vector]bool{
		*loc:                  true,
		*NewCoordinates(7, 8): true,
	}

	fmt.Println(visited)

	newLoc := NewCoordinates(5, 6)

	val := visited[*newLoc]

	fmt.Println(val)
}

func part1(input []string) int {
	play1(input)
	return 0
}

func part2(input []string) int {
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

	return &heightmap
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
