package main

import (
	"fmt"
	"strconv"
	"strings"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input")
	input2 := shared.ReadInput("input2")

	part1 := part1(input)
	part2 := part2(input2)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) int64 {

	races := *ParseRaces(input)
	var sum int64 = 1

	for _, race := range races {
		strategies := *GetStrategies(race, 1, race.Time)
		sum = sum * getWinning(&strategies)
	}

	return sum
}

func part2(input []string) int64 {
	times := strings.Fields(input[0])
	records := strings.Fields(input[1])

	race := Race{
		Time:           ToInt(times[1]),
		RecordDistance: ToInt(records[1]),
	}

	strategies := *GetStrategies(race, 1, race.Time)
	return getWinning(&strategies)
}

func ToInt(input string) int64 {
	value, _ := strconv.ParseInt(input, 10, 64)
	return value
}
