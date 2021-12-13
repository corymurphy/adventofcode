package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	input := readInput("input")

	part1 := computeFuel(input, fuelCostPart1)
	part2 := computeFuel(input, fuelCostPart2)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func diff(a int, b int) int {
	if a > b {
		return a - b
	} else if b > a {
		return b - a
	} else {
		return 0
	}
}

func initializeCrabLocations(input []string) []int {
	var locations []int
	for _, strLocation := range strings.Split(input[0], ",") {

		location := toInt(strLocation)

		if location >= len(locations) {
			diff := location - len(locations)
			for i := 0; i <= diff; i++ {
				locations = append(locations, 0)
			}
		}
		locations[location] = locations[location] + 1
	}
	return locations
}

func computeFuel(input []string, fuelCost func(int, int) int) int {

	locations := initializeCrabLocations(input)
	leastFuel := math.MaxInt
	optimalTarget := 0

	for target := 0; target < len(locations); target++ {

		fuel := 0
		for location, crabs := range locations {
			dist := diff(location, target)
			fuel = fuel + fuelCost(dist, crabs)
		}

		if fuel > 0 && fuel < leastFuel {
			leastFuel = fuel
			optimalTarget = target
		}

	}

	fmt.Printf("\nleast fuel: %d; optimal target: %d\n", leastFuel, optimalTarget)

	return leastFuel
}

func fuelCostPart1(dist int, crabs int) int {
	return crabs * dist
}

func fuelCostPart2(dist int, crabs int) int {
	if dist == 0 {
		return 0
	}

	if dist == 1 {
		return 1 * crabs
	}
	cost := dist
	for i := 0; i < dist; i++ {
		cost = cost + i
	}
	return cost * crabs
}
