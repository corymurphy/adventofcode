package main

import (
	"fmt"
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

type Computer struct {
	A       int
	B       int
	C       int
	Program []int
}

func NewComputer(input []string) (computer Computer) {
	computer = Computer{
		A:       shared.ToInt(strings.Split(input[0], ": ")[1]),
		B:       shared.ToInt(strings.Split(input[1], ": ")[1]),
		C:       shared.ToInt(strings.Split(input[2], ": ")[1]),
		Program: []int{},
	}

	program := strings.Split(strings.Split(input[4], ": ")[1], ",")
	for _, opcode := range program {
		computer.Program = append(computer.Program, shared.ToInt(opcode))
	}

	return computer

	// shared.ToInt(strings.Split(input[4], ": ")[1])
}

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) (answer int) {

	computer := NewComputer(input)
	fmt.Println(computer)
	return answer
}

func part2(input []string) (answer int) {

	return answer
}
