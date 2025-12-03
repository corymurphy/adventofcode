package main

import (
	"fmt"
	"strconv"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)

}

func part1(inputs []string) (answer int) {

	for _, input := range inputs {
		bank := *NewBank(input)
		answer = answer + shared.ToInt(bank.GetJoltage1())
	}
	return answer
}

func GetLargest(input string) (largest int, index int) {

	for i, test := range input {
		if shared.ToInt(string(test)) > largest {
			largest = shared.ToInt(string(test))
			index = i
		}
	}
	return largest, index
}

func part2(inputs []string) (answer int64) {
	for _, input := range inputs {
		bank := *NewBank(input)
		answer = answer + bank.GetJoltage2(input)
	}

	return answer
}

type Joltage struct {
}
type Bank struct {
	Batteries []int
}

func (b *Bank) GetJoltage2(input string) int64 {

	result := ""
	for need := 12; need > 0; need-- {
		sub := input[:len(input)-(need)+1]
		largest, start := GetLargest(sub)
		result = result + strconv.Itoa(largest)
		input = input[start+1:]
	}
	return MustInt64(result)
}

func (b *Bank) GetJoltage1() string {

	possibilities := []int{}

	for i, first := range b.Batteries {

		if i >= len(b.Batteries)-1 {
			break
		}

		for _, second := range b.Batteries[i+1:] {
			possibilities = append(possibilities, shared.ToInt(
				strconv.Itoa(first)+strconv.Itoa(second),
			))
		}
	}

	highest := 0
	for _, possible := range possibilities {
		if possible > highest {
			highest = possible
		}
	}

	return strconv.Itoa(highest)
}

func NewBank(input string) *Bank {
	batteries := []int{}
	for _, b := range input {
		battery := shared.ToInt(string(b))
		batteries = append(batteries, battery)
	}
	return &Bank{
		Batteries: batteries,
	}
}

func MustInt64(input string) int64 {
	val, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		panic("unable to convert")
	}
	return val
}
