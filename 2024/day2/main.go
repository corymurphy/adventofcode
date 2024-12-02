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

func part1(input []string) (answer int) {

	for _, reportSerial := range input {

		report := []int{}

		for _, level := range strings.Split(reportSerial, " ") {

			report = append(report, shared.ToInt(level))
		}

		if safe, _ := isReportSafe(report); safe {
			answer++
			continue
		}
	}

	return answer
}

func isReportSafe(report []int) (safe bool, i int) {

	decreasing := false
	increasing := false

	for i, level := range report {
		if i == 0 {
			continue
		}

		diff := level - report[i-1]

		if diff >= 1 && diff <= 3 && !increasing {
			decreasing = true
			continue
		}

		if diff <= -1 && diff >= -3 && !decreasing {
			increasing = true
			continue
		}

		return false, i
	}

	return true, i
}

func removeAt(items []int, removedAt int) (removed []int) {
	for i, item := range items {
		if i != removedAt {
			removed = append(removed, item)
		}
	}
	return removed
}

func isReportSafe2(report []int) (safe bool) {

	if safe, _ := isReportSafe(report); safe {
		return true
	}

	for i := range report {

		if safe, _ := isReportSafe(removeAt(report, i)); safe {
			return true
		}
	}

	return false
}

func part2(input []string) (answer int) {

	for _, reportSerial := range input {

		report := []int{}

		for _, level := range strings.Split(reportSerial, " ") {

			report = append(report, shared.ToInt(level))
		}

		if safe := isReportSafe2(report); safe {
			answer++
			continue
		}
	}

	return answer
}
