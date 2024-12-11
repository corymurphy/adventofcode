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

func part1(input []string) (answer int) {

	trailMap := NewTrailMap(input)
	// trailMap.PrintTrailHeads()

	trailheads := trailMap.GetTrailHeads()

	for _, trailhead := range trailheads {
		discovered := Discovered{}
		TrailScore(trailMap, trailhead, &discovered)
		// break
	}

	answer = score

	return answer
}

func part2(input []string) (answer int) {

	return answer
}

var score int = 0

type TrailMap [][]int

type Discovered []Coordinate

func (d *Discovered) Contains(coordinate Coordinate) bool {
	for _, this := range *d {
		if this.X == coordinate.X && this.Y == coordinate.Y {
			return true
		}
	}
	return false
}

// type  map[Coordinate]Coordinate

func TrailScore(trailMap TrailMap, start Coordinate, discovered *Discovered) {

	// right
	if len(trailMap[0]) > start.X+1 && trailMap[start.Y][start.X+1]-trailMap[start.Y][start.X] == 1 {
		fmt.Print(trailMap[start.Y][start.X+1])
		if trailMap[start.Y][start.X+1] == 9 {
			if !discovered.Contains(Coordinate{X: start.X + 1, Y: start.Y}) {
				*discovered = append(*discovered, Coordinate{X: start.X + 1, Y: start.Y})
				score++
				// return 0
				// return 0
				// fmt.Println()
			}
			fmt.Println()
		} else {
			TrailScore(trailMap, Coordinate{X: start.X + 1, Y: start.Y}, discovered)
		}
	}

	// left
	if start.X > 0 && trailMap[start.Y][start.X-1]-trailMap[start.Y][start.X] == 1 {
		fmt.Print(trailMap[start.Y][start.X-1])
		if trailMap[start.Y][start.X-1] == 9 {
			if !discovered.Contains(Coordinate{X: start.X - 1, Y: start.Y}) {
				*discovered = append(*discovered, Coordinate{X: start.X - 1, Y: start.Y})
				score++
				// return 0

			}
			fmt.Println()
		} else {
			TrailScore(trailMap, Coordinate{X: start.X - 1, Y: start.Y}, discovered)
		}
	}

	// up

	if start.Y > 0 && trailMap[start.Y-1][start.X]-trailMap[start.Y][start.X] == 1 {
		fmt.Print(trailMap[start.Y-1][start.X])
		if trailMap[start.Y-1][start.X] == 9 {
			if !discovered.Contains(Coordinate{X: start.X, Y: start.Y - 1}) {
				*discovered = append(*discovered, Coordinate{X: start.X, Y: start.Y - 1})
				score++
				// return 0
				// fmt.Println()
			}
			fmt.Println()
		} else {
			TrailScore(trailMap, Coordinate{X: start.X, Y: start.Y - 1}, discovered)
		}

	}

	// down
	if len(trailMap) > start.Y+1 && trailMap[start.Y+1][start.X]-trailMap[start.Y][start.X] == 1 {
		fmt.Print(trailMap[start.Y+1][start.X])
		if trailMap[start.Y+1][start.X] == 9 {

			if !discovered.Contains(Coordinate{X: start.X, Y: start.Y + 1}) {
				*discovered = append(*discovered, Coordinate{X: start.X, Y: start.Y + 1})
				score++
				// return 0
				// fmt.Println()
			}
			fmt.Println()
		} else {
			TrailScore(trailMap, Coordinate{X: start.X, Y: start.Y + 1}, discovered)
		}

	}
	fmt.Println()

	// return 0
}

type Coordinate struct {
	X int
	Y int
}

// func (m *TrailMap) IsTrailHead()
func (m *TrailMap) GetTrailHeads() (trailheads []Coordinate) {
	for y, row := range *m {
		for x, val := range row {
			if val == 0 {
				trailheads = append(trailheads, Coordinate{X: x, Y: y})
			}
		}
	}
	return trailheads
}

func NewTrailMap(input []string) (trailMap TrailMap) {

	for _, row := range input {
		horizontal := []int{}
		for _, val := range row {
			horizontal = append(horizontal, shared.ToInt(string(val)))
		}
		trailMap = append(trailMap, horizontal)
	}

	return trailMap
}

func (m *TrailMap) PrintTrailHeads() {
	fmt.Println()
	for _, row := range *m {
		for _, val := range row {
			if val == 0 || val == 9 {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
