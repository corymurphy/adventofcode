package main

import (
	"fmt"
	"sort"

	shared "github.com/corymurphy/adventofcode/shared"
)

func topThreeCarriedCalories(input []string) int {

	calories := []int{}

	elf := 0

	for i, snack := range input {

		if len(snack) == 0 || i == len(input)-1 {
			calories = append(calories, elf)
			elf = 0
			continue
		}

		snackCalories := shared.ToInt(snack)

		elf = elf + snackCalories
	}

	sort.Ints(calories)
	length := len(calories)
	return calories[length-1] + calories[length-2] + calories[length-3]

}

func mostCarriedCalories(input []string) int {
	most := 0

	elf := 0

	for _, snack := range input {

		if len(snack) == 0 {
			elf = 0
			continue
		}

		snackCalories := shared.ToInt(snack)

		elf = elf + snackCalories

		if elf > most {
			most = elf
		}
	}

	return most
}

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) int {
	return mostCarriedCalories(input)
}

func part2(input []string) int {
	return topThreeCarriedCalories(input)
}
