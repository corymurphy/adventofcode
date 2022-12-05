package main

import (
	"github.com/corymurphy/adventofcode/shared"
)

type Dock struct {
	Stacks   map[int]*shared.Stack
	Commands Commands
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
			value, _ := d.Stacks[command.From].Pop()
			d.Stacks[command.To].Push(value)
		}
	}
}

func (d *Dock) RearrangePart2() {
	for _, command := range d.Commands {
		staging := shared.NewStack()

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
			break
		}
	}

	stacks := input[0:commandSeparatorIndex]
	commands := input[commandSeparatorIndex:]

	return &Dock{
		Stacks:   InitializeStacks(stacks),
		Commands: *NewCommands(commands),
	}
}
