package main

import (
	"fmt"
	"strings"

	shared "github.com/corymurphy/adventofcode/shared"
)

type Assignment struct {
	Low  int
	High int
}

type AssignmentPair struct {
	AssignmentA Assignment
	AssignmentB Assignment
}

func NewAssignment(input string) *Assignment {
	assignment := strings.Split(input, "-")
	return &Assignment{
		Low:  shared.ToInt(assignment[0]),
		High: shared.ToInt(assignment[1]),
	}
}

func NewAssignmentPair(input string) *AssignmentPair {
	assignmentPair := strings.Split(input, ",")
	return &AssignmentPair{
		AssignmentA: *NewAssignment(assignmentPair[0]),
		AssignmentB: *NewAssignment(assignmentPair[1]),
	}
}

func (a *AssignmentPair) Contains() bool {
	overlaps := false

	if a.AssignmentA.Low >= a.AssignmentB.Low &&
		a.AssignmentA.High <= a.AssignmentB.High {
		overlaps = true
	}
	if a.AssignmentB.Low >= a.AssignmentA.Low &&
		a.AssignmentB.High <= a.AssignmentA.High {
		overlaps = true
	}
	return overlaps
}

func (a *AssignmentPair) Overlaps() bool {
	overlaps := false

	if a.AssignmentA.High >= a.AssignmentB.Low &&
		a.AssignmentA.High <= a.AssignmentB.High {
		return true
	}

	if a.AssignmentB.High >= a.AssignmentA.Low &&
		a.AssignmentB.High <= a.AssignmentA.High {
		return true
	}

	return overlaps
}

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) int {

	total := 0

	for _, row := range input {
		if NewAssignmentPair(row).Contains() {
			total++
		}
	}

	return total
}

func part2(input []string) int {
	total := 0
	for _, row := range input {
		if NewAssignmentPair(row).Overlaps() {
			total++
		}
	}
	return total
}
