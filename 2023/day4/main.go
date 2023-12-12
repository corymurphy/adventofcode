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

func part1(input []string) int {

	cards := *ParseCards(input)

	sum := 0
	for _, card := range cards {
		sum = sum + card.Points()
	}
	return sum
}

func part2(input []string) int {

	cards := *ParseCards(input)

	for current, card := range cards {
		for j := 0; j < card.instances; j++ {
			for i := current + 1; i <= len(card.matching)+current && i < len(cards); i++ {
				cards[i].AddInstance()
			}
		}
	}

	instances := 0
	for _, card := range cards {
		instances = instances + card.instances
	}

	return instances
}
