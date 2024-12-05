package main

import (
	"fmt"
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

func parseRules(input []string) (rules map[int][]int) {
	rules = make(map[int][]int)
	for _, line := range input {

		if line == "" {
			break
		}

		rule := strings.Split(line, "|")
		preceeding := shared.ToInt(rule[0])
		succeeding := shared.ToInt(rule[1])

		if _, contains := rules[preceeding]; !contains {
			rules[preceeding] = []int{succeeding}
		} else {
			rules[preceeding] = append(rules[preceeding], succeeding)
		}

	}
	return rules
}

func parseUpdates(input []string) (updates [][]int) {
	updates = [][]int{}

	inUpdates := false

	for _, line := range input {

		if line == "" {
			inUpdates = true
			continue
		}

		if !inUpdates {
			continue
		}

		instructions := []int{}
		for _, instruction := range strings.Split(line, ",") {
			instructions = append(instructions, shared.ToInt(instruction))
		}
		updates = append(updates, instructions)

	}
	return updates
}

func isValid(preceedingValues []int, rules []int) bool {

	for _, preceeding := range preceedingValues {
		for _, rule := range rules {
			if preceeding == rule {
				return false
			}
		}
	}
	return true
}

func isValidUpdate(update []int, rules map[int][]int) (valid bool, invalidAt int) {

	for i, page := range update {

		preceeding := update[:i]
		if len(preceeding) == 0 {
			continue
		}

		rule := rules[page]

		valid = isValid(preceeding, rule)

		if !valid {
			return false, i
		}
	}

	return true, 0
}

func part1(input []string) (answer int) {

	rules := parseRules(input)
	updates := parseUpdates(input)

	for _, update := range updates {
		valid, _ := isValidUpdate(update, rules)
		if valid {
			answer = answer + update[len(update)/2]
		}
	}

	return answer
}

func part2(input []string) (answer int) {

	rules := parseRules(input)
	updates := parseUpdates(input)

	for _, update := range updates {
		valid, incorrectIndex := isValidUpdate(update, rules)
		if valid {
			continue
		}

		for {

			a := update[incorrectIndex]
			b := update[incorrectIndex-1]

			update[incorrectIndex-1] = a
			update[incorrectIndex] = b
			valid := false
			valid, incorrectIndex = isValidUpdate(update, rules)

			if valid {
				answer = answer + update[len(update)/2]
				break
			}
		}
	}
	return answer
}
