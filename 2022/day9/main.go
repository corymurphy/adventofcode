package main

import (
	"fmt"
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

func NewGrid(size int) [][]int {
	// grid := [size][size]int{}
	grid := make([][]int, size)
	for i := len(grid) - 1; i >= 0; i-- {
		row := make([]int, size)
		if i == 0 {
			row[0] = 0
		}
		grid[i] = row
	}
	// for i := 0; i <= size; i++ {
	// 	row := []int{}
	// 	for j := 0; j <= size; j++ {
	// 		if j == 0 && i == 0 {
	// 			row = append(row, 1)
	// 		} else {
	// 			row = append(row, 0)
	// 		}

	// 	}
	// 	grid = append(grid, row)
	// }

	return grid
}

func PrintGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		fmt.Println("")
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d", grid[i][j])
		}
		// fmt.Println("")
	}
	fmt.Println("")
	fmt.Println("")
}

func PrintGridBackwards(grid [][]int) {
	// fmt.Println("")
	for i := len(grid) - 1; i >= 0; i-- {
		fmt.Println("")
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= 1 {
				fmt.Printf("%s", "#")
			} else {
				fmt.Printf("%s", ".")
			}
			// fmt.Printf("%d", grid[i][j])
		}
		// fmt.Println("")

	}
	fmt.Println("")
	fmt.Println("")
}

func GetBridgeSize(input []string) int {
	size := 0

	for _, row := range input {
		distance := shared.ToInt(strings.Split(row, " ")[1])
		if distance > size {
			size = distance
		}
	}

	return size

}

func SimulateBridge(input []string, grid [][]int) [][]int {

	head := NewCoordinates()
	tail := NewCoordinates()

	for _, item := range input {
		instruction := NewInstruction(item)
		Move(&grid, instruction, head, tail)
	}

	return grid
}

func CountVisited(grid [][]int) int {
	visited := 0
	for _, col := range grid {
		for _, loc := range col {
			if loc > 0 {
				visited++
			}
		}
	}
	return visited
}

func SimulateBridgeTails(input []string, grid [][]int, tailCount int) [][]int {

	head := NewCoordinates()
	// tail := NewCoordinates()
	tails := NewCoordinatesList(tailCount)

	for _, item := range input {
		instruction := NewInstruction(item)
		MoveDynamic(&grid, instruction, head, tails)
	}

	return grid
}

func part1(input []string) int {
	grid := NewGrid(GetBridgeSize(input) * 20)
	grid = SimulateBridge(input, grid)
	return CountVisited(grid)
}

func part2(input []string) int {
	grid := NewGrid(GetBridgeSize(input) * 20)
	grid = SimulateBridgeTails(input, grid, 9)
	return CountVisited(grid)
}
