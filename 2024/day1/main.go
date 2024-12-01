package main

import (
	"fmt"
	"sort"
	"strings"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) (answer int) {

	leftLocations := []int{}
	rightLocations := []int{}

	for _, row := range input {

		splitRow := strings.Split(row, "   ")
		left := shared.ToInt(splitRow[0])
		right := shared.ToInt(splitRow[1])

		leftLocations = append(leftLocations, left)
		rightLocations = append(rightLocations, right)

	}

	sort.Ints(leftLocations)
	sort.Ints(rightLocations)

	for i := range leftLocations {
		left := leftLocations[i]
		right := rightLocations[i]
		diff := left - right

		if diff < 0 {
			diff = diff * -1
		}

		answer = answer + diff

		// fmt.Printf("%d 		%d %d	%d\n", i, left, right, diff)
	}
	return answer
}

func part2(input []string) (answer int) {

	leftLocations := []int{}
	rightLocations := []int{}

	for _, row := range input {

		splitRow := strings.Split(row, "   ")
		left := shared.ToInt(splitRow[0])
		right := shared.ToInt(splitRow[1])

		leftLocations = append(leftLocations, left)
		rightLocations = append(rightLocations, right)

	}

	for _, left := range leftLocations {

		occurances := 0
		for _, right := range rightLocations {
			if left == right {
				occurances++
			}
		}
		occurances = left * occurances

		fmt.Println(occurances)
		answer = answer + occurances
	}

	return answer
}

// func InsertSorted(s []int, e int) []int {
// 	s = append(s, 0)
// 	i := sort.Search(len(s), func(i int) bool { return s[i] > e })
// 	copy(s[i+1:], s[i:])
// 	s[i] = e
// 	return s
// }
