package main

import (
	"fmt"
	"strings"

	shared "github.com/corymurphy/adventofcode/shared"
)

var priority map[string]int = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
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
	for _, rucksack := range input {
		firstCompartment := rucksack[0 : len(rucksack)/2]
		secondCompartment := rucksack[len(rucksack)/2:]

		for j := 0; j < len(firstCompartment); j++ {
			item := firstCompartment[j]
			if strings.Contains(secondCompartment, string(item)) {
				total = total + priority[string(item)]
				break
			}
		}

	}
	return total
}

func part2(input []string) int {
	total := 0

	for i := 0; i < len(input)-2; i = i + 3 {
		firstRucksack := input[i]
		secondRucksack := input[i+1]
		thirdRucksack := input[i+2]

		for j := 0; j < len(firstRucksack); j++ {
			item := firstRucksack[j]

			if strings.Contains(secondRucksack, string(item)) &&
				strings.Contains(thirdRucksack, string(item)) {
				total = total + priority[string(item)]
				break
			}
		}
	}

	return total
}
