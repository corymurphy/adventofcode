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

func part1(input []string) int {
	landscape := parse(input)
	return countVisible(landscape)
}

func part2(input []string) int {
	landscape := parse(input)
	return getOptimalView(landscape)
}

func parse(input []string) [][]int {
	landscape := [][]int{}
	for _, row := range input {
		trees := []int{}
		for _, tree := range row {
			trees = append(trees, shared.ToInt(string(tree)))
		}
		landscape = append(landscape, trees)
	}
	return landscape
}

func getLargest(list []int) int {
	largest := 0
	for _, value := range list {
		if value > largest {
			largest = value
		}
	}
	return largest
}

func getOptimalView(landscape [][]int) int {
	scores := []int{}
	for colIndex := range landscape {
		for rowIndex := range landscape[colIndex] {

			if colIndex == 0 || rowIndex == 0 {
				continue
			}

			viewLeft := viewLeft(colIndex, rowIndex, landscape)
			viewRight := viewRight(colIndex, rowIndex, landscape)
			viewTop := viewTop(colIndex, rowIndex, landscape)
			viewBottom := viewBottom(colIndex, rowIndex, landscape)

			score := viewLeft * viewRight * viewTop * viewBottom

			scores = append(scores, score)
		}
	}
	return getLargest(scores)
}

func countVisible(landscape [][]int) int {
	visible := 0

	for colIndex := range landscape {
		for rowIndex := range landscape[colIndex] {
			if isVisible(colIndex, rowIndex, landscape) {
				visible++
			}
		}
	}
	return visible
}

func isVisible(startColumn int, startRow int, landscape [][]int) bool {

	visible := true

	// check top
	if startColumn == 0 {
		return true
	}

	// check left
	if startRow == 0 {
		return true
	}

	// check right
	if startRow == len(landscape[startColumn])-1 {
		return true
	}

	// check bottom
	if startColumn == len(landscape)-1 {
		return true
	}

	if !checkBottom(startColumn, startRow, landscape) &&
		!checkTop(startColumn, startRow, landscape) &&
		!visibleRight(startColumn, startRow, landscape) &&
		!visibleLeft(startColumn, startRow, landscape) {
		visible = false
	}

	return visible
}

func checkBottom(startColumn int, startRow int, landscape [][]int) bool {
	visible := true
	tree := landscape[startColumn][startRow]

	for i := startColumn + 1; i <= len(landscape)-1; i++ {
		height := landscape[i][startRow]

		if tree <= height {
			visible = false
			break
		}
	}
	return visible
}

func checkTop(startColumn int, startRow int, landscape [][]int) bool {
	visible := true
	tree := landscape[startColumn][startRow]
	for i := startColumn - 1; i >= 0; i-- {
		height := landscape[i][startRow]
		if tree <= height {
			visible = false
			break
		}
	}
	return visible
}

func visibleRight(startColumn int, startRow int, landscape [][]int) bool {
	visible := true
	tree := landscape[startColumn][startRow]
	for i := startRow + 1; i <= len(landscape)-1; i++ {
		height := landscape[startColumn][i]

		if tree <= height {
			visible = false
			break
		}
	}
	return visible
}

func visibleLeft(startColumn int, startRow int, landscape [][]int) bool {
	visible := true
	tree := landscape[startColumn][startRow]
	for i := startRow - 1; i >= 0; i-- {
		height := landscape[startColumn][i]
		if tree <= height {
			visible = false
			break
		}
	}
	return visible
}
