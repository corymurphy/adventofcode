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

func countOccurrences(char rune, input string) int {
	count := 0
	for _, el := range input {
		if el == char {
			count++
		}
	}
	return count
}

func isUnique(input string) bool {

	for _, char := range input {
		if countOccurrences(char, input) > 1 {
			return false
		}
	}
	return true
}

func getMessageStart(signal string, evalLength int) int {
	for i := evalLength; i <= len(signal); i++ {
		if isUnique(signal[i-evalLength : i]) {
			return i
		}
	}
	return 0
}

func part1(input []string) int {
	return getMessageStart(input[0], 4)
}

func part2(input []string) int {
	return getMessageStart(input[0], 14)
}
