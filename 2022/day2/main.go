package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

const rock string = "A"
const paper string = "B"
const scissors string = "C"

const rock_response string = "X"
const paper_response string = "Y"
const scissors_response string = "Z"

const rock_bonus int = 1
const paper_bonus int = 2
const scissors_bonus int = 3

const should_lose string = "X"
const should_draw string = "Y"
const should_win string = "Z"

const lose int = 0
const draw int = 3
const win int = 6

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) int {

	total := 0
	for _, rawRound := range input {
		round := NewRound(rawRound)
		total = total + round.Play()
	}
	return total
}

func part2(input []string) int {
	total := 0
	for _, rawRound := range input {
		round := NewRound(rawRound)
		total = total + round.Play2()
	}
	return total
}
