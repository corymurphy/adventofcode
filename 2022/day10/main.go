package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

// const (
// 	significantCycles = []int{20}
// )

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) int {
	program := NewProgram(input)
	cpu := NewCpu()
	return cpu.Execute(program)
}

func part2(input []string) int {
	return 0
}
