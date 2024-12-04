package main

import (
	"fmt"
	"strconv"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func getParams(input string) (valid bool, first int, second int) {

	firstSearch := true
	firstParam := ""
	secondSearch := false
	secondParam := ""

	for i, current := range input {

		if i >= 7 {
			if isInt(firstParam) && isInt(secondParam) && input[i] == ')' {
				return true, shared.ToInt(firstParam), shared.ToInt(secondParam)
			}
			return false, first, second
		}

		if firstSearch && i >= 3 && !isInt(firstParam) {
			return false, first, second
		}

		if firstSearch && current == ',' {
			if isInt(firstParam) {
				secondSearch = true
				firstSearch = false
				first = shared.ToInt(firstParam)
				continue
			} else {
				return false, first, second
			}
		}

		if secondSearch && current == ')' {
			if isInt(secondParam) {
				second = shared.ToInt(secondParam)
				return true, first, second
			} else {
				return false, first, second
			}
		}

		if firstSearch {
			firstParam = firstParam + string(current)
		}

		if secondSearch {
			secondParam = secondParam + string(current)
		}

	}
	return false, first, second
}

func isInt(input string) (valid bool) {

	if _, err := strconv.Atoi(input); err == nil {
		return true
	}
	return false
}

func part1(input []string) (answer int) {

	for _, row := range input {

		for i := range row {

			if i >= len(row)-5 {
				continue
			}

			if row[i:i+4] != "mul(" {
				continue
			}

			valid, first, second := getParams(row[i+4:])

			if valid {
				answer = answer + (first * second)
			}
		}
	}

	return answer
}

func part2(input []string) (answer int) {

	enabled := true
	for _, row := range input {

		for i := range row {

			if i >= len(row)-5 {
				continue
			}
			if len(row[i:]) >= 7 && row[i:i+7] == "don't()" {
				enabled = false
			}

			if row[i:i+4] == "do()" {
				enabled = true
			}

			if row[i:i+4] != "mul(" {
				continue
			}

			valid, first, second := getParams(row[i+4:])

			if valid && enabled {
				answer = answer + (first * second)
			}
		}
	}

	return answer
}
