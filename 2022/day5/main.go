package main

import (
	"fmt"
	"strings"

	shared "github.com/corymurphy/adventofcode/shared"
)

type Dock struct {
	Stacks   map[int]*CmStack
	Commands Commands
}

func CountStacks(input string) int {
	return len(strings.Fields(input))
}

func InitializeStacks(input []string) map[int]*CmStack {

	stacks := map[int]*CmStack{}

	stackCount := CountStacks(input[len(input)-1])

	for i := 1; i <= stackCount; i++ {
		stacks[i] = NewStack()
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

func (d *Dock) TopCrates() string {
	result := ""
	for i := 1; i <= len(d.Stacks); i++ {
		stack := d.Stacks[i]
		item, err := stack.Peek()
		if err != nil || item != "" {
			result = result + item
		}
	}
	return result
}

func (d *Dock) Rearrange() {
	for _, command := range d.Commands {
		for i := 1; i <= command.Move; i++ {
			from := command.From
			to := command.To
			value, _ := d.Stacks[from].Pop()
			d.Stacks[to].Push(value)
		}
	}
}

func (d *Dock) RearrangePart2() {
	for _, command := range d.Commands {
		staging := NewStack()

		// this is lazy, i could improve the logic here
		for i := 1; i <= command.Move; i++ {
			value, _ := d.Stacks[command.From].Pop()
			staging.Push(value)
		}

		for i := 1; i <= command.Move; i++ {
			to := command.To
			value, _ := staging.Pop()
			d.Stacks[to].Push(value)
		}
	}
}

func NewDock(input []string) *Dock {

	commandSeparatorIndex := 0
	for i, row := range input {
		if row == "" {
			commandSeparatorIndex = i
		}
	}

	stacks := input[0:commandSeparatorIndex]
	commands := input[commandSeparatorIndex:]

	return &Dock{
		Stacks:   InitializeStacks(stacks),
		Commands: *NewCommands(commands),
	}
}

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %s\n\n", part1)
	fmt.Printf("\nPart 2 answer: %s\n\n", part2)
}

func part1(input []string) string {
	dock := NewDock(input)
	dock.Rearrange()
	return dock.TopCrates()
}

func part2(input []string) string {
	dock := NewDock(input)
	dock.RearrangePart2()
	return dock.TopCrates()
}
