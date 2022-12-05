package main

import (
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

type Commands []Command

func NewCommands(input []string) *Commands {

	commands := Commands{}
	for _, row := range input {
		if row == "" {
			continue
		}
		commands = append(commands, *NewCommand(row))
	}
	return &commands
}

type Command struct {
	Move int
	From int
	To   int
}

func NewCommand(input string) *Command {
	args := strings.Split(input, " ")
	return &Command{
		Move: shared.ToInt(args[1]),
		From: shared.ToInt(args[3]),
		To:   shared.ToInt(args[5]),
	}
}
