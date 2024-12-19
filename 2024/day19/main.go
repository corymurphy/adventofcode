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

	pmax := 0
	patterns := strings.Split(input[0], ", ")
	sort.Slice(patterns, func(i, j int) bool {
		if len(patterns[i]) > pmax {
			pmax = len(patterns[i])
		}
		if len(patterns[j]) > pmax {
			pmax = len(patterns[j])
		}
		return len(patterns[i]) > len(patterns[j])
	})

	designs := input[2:]

	possible := []string{}

	for _, design := range designs {

		// s := 0
		endFit := false

		for i := range 4 {

			di := 0
			fit := false
			// pmin := 1
			retries := 0

			for di < len(design) {

				fit = false

				fmt.Println(i, retries, di, pmax)

				for _, pattern := range patterns {
					pl := len(pattern)

					if pl+di > len(design) {
						continue
					}

					if design[di:pl+di] == pattern {
						endFit = len(design) == pl+di
						di += pl
						fit = true
						break
					}
				}

				if endFit {
					break
				}

				if fit {
					continue
				}

				if retries > pmax || di <= 0 {
					break
				}

				retries++

				di = di - i
			}
		}

		if endFit {
			possible = append(possible, design)
		}
	}

	answer = len(possible)

	return answer
}

func part2(input []string) (answer int) {

	return answer
}
