package main

import (
	"fmt"
	"sort"
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

		snackCalories := toInt(snack)

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

		snackCalories := toInt(snack)

		elf = elf + snackCalories

		if elf > most {
			most = elf
		}
	}

	return most
}

func main() {
	input := ReadInput("input")
	fmt.Println(mostCarriedCalories(input))
	fmt.Println(topThreeCarriedCalories(input))
}
