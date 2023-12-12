package main

import (
	"fmt"
	"slices"

	shared "github.com/corymurphy/adventofcode/shared"
)

var numbers []rune = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
var GEAR rune = '*'

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) int {

	symbols := findSymbols(input)
	partNumbers := []int{}
	sumPartNumbers := 0

	for columnIndex := 0; columnIndex < len(input); columnIndex++ {
		for rowIndex := 0; rowIndex < len(input[columnIndex]); rowIndex++ {

			element := rune(input[columnIndex][rowIndex])

			if !slices.Contains(numbers, element) {
				continue
			}

			hasAdjacentSymbol := hasAdjacentSymbol(input, symbols, columnIndex, rowIndex)

			if !hasAdjacentSymbol {
				continue
			}

			partNumber, advance := getPartNumber(input[columnIndex], rowIndex)
			rowIndex = advance

			partNumbers = append(partNumbers, partNumber)
			sumPartNumbers = sumPartNumbers + partNumber
		}
	}

	return sumPartNumbers
}

func part2(input []string) int {
	sumGearRatio := 0

	for columnIndex := 0; columnIndex < len(input); columnIndex++ {
		for rowIndex := 0; rowIndex < len(input[columnIndex]); rowIndex++ {

			element := rune(input[columnIndex][rowIndex])

			if element != GEAR {
				continue
			}

			gears := getAdjacentInt(input, columnIndex, rowIndex)

			if len(gears) < 2 {
				continue
			}

			gearRatio := 0

			for gear := range gears {
				if gearRatio == 0 {
					gearRatio = gear.value
					continue
				}

				gearRatio = gearRatio * gear.value
			}

			sumGearRatio = sumGearRatio + gearRatio
		}
	}

	return sumGearRatio
}

func getPartNumber(row string, start int) (int, int) {
	number := ""
	end := start

	// leading
	for i := start - 1; i >= 0 && slices.Contains(numbers, rune(row[i])); i-- {
		number = string(row[i]) + number
	}

	// trailing
	for i := start; i < len(row) && slices.Contains(numbers, rune(row[i])); i++ {
		number = number + string(row[i])
		end = i
	}

	return shared.ToInt(number), end
}

func getGear(column int, row string, foundAt int) Coordinate {
	number := ""
	end := foundAt
	start := foundAt

	// leading
	for i := foundAt - 1; i >= 0 && slices.Contains(numbers, rune(row[i])); i-- {
		number = string(row[i]) + number
		start = i
	}

	// trailing
	for i := foundAt; i < len(row) && slices.Contains(numbers, rune(row[i])); i++ {
		number = number + string(row[i])
		end = i
	}

	return Coordinate{column: column, start: start, end: end, value: shared.ToInt(number)}
}

func hasAdjacentSymbol(input []string, symbols []rune, columnIndex int, rowIndex int) bool {

	// left
	if rowIndex > 0 {
		left := input[columnIndex][rowIndex-1]
		if slices.Contains(symbols, rune(left)) {
			return true
		}
	}

	// top left
	if rowIndex > 0 && columnIndex > 0 {
		topLeft := input[columnIndex-1][rowIndex-1]
		if slices.Contains(symbols, rune(topLeft)) {
			return true
		}
	}

	// top
	if columnIndex > 0 {
		top := input[columnIndex-1][rowIndex]
		if slices.Contains(symbols, rune(top)) {
			return true
		}
	}

	// top right
	if columnIndex > 0 && rowIndex+1 < len(input[columnIndex]) {
		topRight := input[columnIndex-1][rowIndex+1]
		if slices.Contains(symbols, rune(topRight)) {
			return true
		}
	}

	// right
	if rowIndex+1 < len(input[columnIndex]) {
		right := input[columnIndex][rowIndex+1]
		if slices.Contains(symbols, rune(right)) {
			return true
		}
	}

	// bottom right
	if columnIndex+1 < len(input) && rowIndex+1 < len(input[columnIndex]) {
		bottomRight := input[columnIndex+1][rowIndex+1]
		if slices.Contains(symbols, rune(bottomRight)) {
			return true
		}
	}

	// bottom
	if columnIndex+1 < len(input) {
		bottom := input[columnIndex+1][rowIndex]
		if slices.Contains(symbols, rune(bottom)) {
			return true
		}
	}

	// bottom left
	if columnIndex+1 < len(input) && rowIndex > 0 {
		bottomleft := input[columnIndex+1][rowIndex-1]
		if slices.Contains(symbols, rune(bottomleft)) {
			return true
		}
	}

	return false
}

func getAdjacentInt(input []string, columnIndex int, rowIndex int) map[Coordinate]bool {

	found := map[Coordinate]bool{}

	// left
	if rowIndex > 0 {
		left := input[columnIndex][rowIndex-1]
		if slices.Contains(numbers, rune(left)) {
			found[getGear(columnIndex, input[columnIndex], rowIndex-1)] = true
		}
	}

	// top left
	if rowIndex > 0 && columnIndex > 0 {
		topLeft := input[columnIndex-1][rowIndex-1]
		if slices.Contains(numbers, rune(topLeft)) {
			found[getGear(columnIndex-1, input[columnIndex-1], rowIndex-1)] = true
		}
	}

	// top
	if columnIndex > 0 {
		top := input[columnIndex-1][rowIndex]
		if slices.Contains(numbers, rune(top)) {
			found[getGear(columnIndex-1, input[columnIndex-1], rowIndex)] = true
		}
	}

	// top right
	if columnIndex > 0 && rowIndex+1 < len(input[columnIndex]) {
		topRight := input[columnIndex-1][rowIndex+1]
		if slices.Contains(numbers, rune(topRight)) {
			found[getGear(columnIndex-1, input[columnIndex-1], rowIndex+1)] = true
		}
	}

	// right
	if rowIndex+1 < len(input[columnIndex]) {
		right := input[columnIndex][rowIndex+1]
		if slices.Contains(numbers, rune(right)) {
			found[getGear(columnIndex, input[columnIndex], rowIndex+1)] = true
		}
	}

	// bottom right
	if columnIndex+1 < len(input) && rowIndex+1 < len(input[columnIndex]) {
		bottomRight := input[columnIndex+1][rowIndex+1]
		if slices.Contains(numbers, rune(bottomRight)) {
			found[getGear(columnIndex+1, input[columnIndex+1], rowIndex+1)] = true
		}
	}

	// bottom
	if columnIndex+1 < len(input) {
		bottom := input[columnIndex+1][rowIndex]
		if slices.Contains(numbers, rune(bottom)) {
			found[getGear(columnIndex+1, input[columnIndex+1], rowIndex)] = true
		}
	}

	// bottom left
	if columnIndex+1 < len(input) && rowIndex > 0 {
		bottomLeft := input[columnIndex+1][rowIndex-1]
		if slices.Contains(numbers, rune(bottomLeft)) {
			found[getGear(columnIndex+1, input[columnIndex+1], rowIndex-1)] = true
		}
	}

	return found
}

func findSymbols(input []string) []rune {
	symbols := []rune{}

	for _, column := range input {
		for _, element := range column {
			if element != '.' && !slices.Contains(numbers, element) {
				if !slices.Contains(symbols, element) {
					symbols = append(symbols, element)
				}

			}
		}
	}
	return symbols
}
