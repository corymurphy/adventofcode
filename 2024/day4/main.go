package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

type Coordinate[x int, y int] struct {
	X     int
	Y     int
	Value rune
}

// var (
// 	visualizer = make(map)
// )

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func horizontal(input []string, rowIndex int, colIndex int) (found int) {

	// if len(input[rowIndex]) < colIndex+3 {
	// 	return found
	// }

	if len(input[rowIndex]) > colIndex+3 &&
		input[rowIndex][colIndex+1] == 'M' &&
		input[rowIndex][colIndex+2] == 'A' &&
		input[rowIndex][colIndex+3] == 'S' {
		found = found + 1
	}

	if colIndex >= 3 &&
		input[rowIndex][colIndex-1] == 'M' &&
		input[rowIndex][colIndex-2] == 'A' &&
		input[rowIndex][colIndex-3] == 'S' {
		found = found + 1
	}

	return found
}

func vertical(input []string, rowIndex int, colIndex int) (found int) {

	// down
	if len(input) > rowIndex+3 {
		if input[rowIndex+1][colIndex] == 'M' &&
			input[rowIndex+2][colIndex] == 'A' &&
			input[rowIndex+3][colIndex] == 'S' {
			found = found + 1
		}
	}

	if rowIndex >= 3 {
		if input[rowIndex-1][colIndex] == 'M' &&
			input[rowIndex-2][colIndex] == 'A' &&
			input[rowIndex-3][colIndex] == 'S' {
			found = found + 1
		}
	}

	// up

	return found
}

func diagonal(input []string, rowIndex int, colIndex int) (found int) {

	// down right
	if colIndex+3 < len(input[rowIndex]) &&
		rowIndex+3 < len(input) &&
		input[rowIndex+1][colIndex+1] == 'M' &&
		input[rowIndex+2][colIndex+2] == 'A' &&
		input[rowIndex+3][colIndex+3] == 'S' {
		found = found + 1
	}

	// up right
	if colIndex+3 < len(input[rowIndex]) &&
		rowIndex >= 3 &&
		input[rowIndex-1][colIndex+1] == 'M' &&
		input[rowIndex-2][colIndex+2] == 'A' &&
		input[rowIndex-3][colIndex+3] == 'S' {
		found = found + 1
	}

	// down left
	if colIndex >= 3 &&
		rowIndex+3 < len(input) &&
		input[rowIndex+1][colIndex-1] == 'M' &&
		input[rowIndex+2][colIndex-2] == 'A' &&
		input[rowIndex+3][colIndex-3] == 'S' {
		found = found + 1
	}

	// up left
	if colIndex >= 3 &&
		rowIndex >= 3 &&
		input[rowIndex-1][colIndex-1] == 'M' &&
		input[rowIndex-2][colIndex-2] == 'A' &&
		input[rowIndex-3][colIndex-3] == 'S' {
		found = found + 1
	}

	return found
}

func part1(input []string) (answer int) {

	for rowIndex, row := range input {
		for colIndex, startValue := range row {

			if startValue == 'X' {

				answer = answer + horizontal(input, rowIndex, colIndex)
				answer = answer + vertical(input, rowIndex, colIndex)
				answer = answer + diagonal(input, rowIndex, colIndex)
			}

		}
	}

	return answer
}

func part2(input []string) (answer int) {

	for rowIndex, row := range input {

		for colIndex, startValue := range row {

			if rowIndex < 1 || rowIndex+1 >= len(input) {
				// fmt.Print(".")
				continue
			}

			if colIndex < 1 || colIndex+1 >= len(input[rowIndex]) {
				// fmt.Print(".")
				continue
			}

			if startValue != 'A' {
				// fmt.Print(".")
				continue
			}

			// fmt.Print("A")

			if input[rowIndex-1][colIndex-1] == 'M' &&
				input[rowIndex+1][colIndex-1] == 'M' &&
				input[rowIndex-1][colIndex+1] == 'S' &&
				input[rowIndex+1][colIndex+1] == 'S' {
				answer++
			}

			if input[rowIndex-1][colIndex-1] == 'S' &&
				input[rowIndex+1][colIndex-1] == 'S' &&
				input[rowIndex-1][colIndex+1] == 'M' &&
				input[rowIndex+1][colIndex+1] == 'M' {
				answer++
			}

			if input[rowIndex-1][colIndex-1] == 'M' &&
				input[rowIndex+1][colIndex-1] == 'S' &&
				input[rowIndex-1][colIndex+1] == 'M' &&
				input[rowIndex+1][colIndex+1] == 'S' {
				answer++
			}

			if input[rowIndex-1][colIndex-1] == 'S' &&
				input[rowIndex+1][colIndex-1] == 'M' &&
				input[rowIndex-1][colIndex+1] == 'S' &&
				input[rowIndex+1][colIndex+1] == 'M' {
				answer++
			}

		}
		// fmt.Print("\n")
	}

	return answer
}
