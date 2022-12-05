package main

import (
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

func CountStacks(input string) int {
	return len(strings.Fields(input))
}

func InitializeStacks(input []string) map[int]*shared.Stack {

	stacks := map[int]*shared.Stack{}

	stackCount := CountStacks(input[len(input)-1])

	for i := 1; i <= stackCount; i++ {
		stacks[i] = shared.NewStack()
	}

	for i := len(input) - 2; i >= 0; i-- {

		stackCounter := 1
		for j := 1; j <= len(input[i]); j = j + 4 {

			crate := string(input[i][j])
			if crate != " " {
				stack := stacks[stackCounter]
				stack.Push(crate)
				stacks[stackCounter] = stack
			}
			stackCounter++
		}

	}
	return stacks
}
