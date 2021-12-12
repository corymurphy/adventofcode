package main

import (
	"fmt"
	// "strings"
)

func main() {
	input := readInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) int {
	diagReport := NewDiagnosticReport(input)
	return diagReport.PowerConsumption()
}

func part2(input []string) int {
	diagReport := NewDiagnosticReport(input)
	return diagReport.LifeSupportRating()
}
