package main

import (
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

type Instruction int

const (
	Noop    Instruction = 0
	Addx    Instruction = 1
	Unknown Instruction = -1
)

func (i Instruction) String() string {
	switch i {
	case Noop:
		return "noop"
	case Addx:
		return "addx"
	default:
		return "Unknown"
	}
}

type Operation struct {
	data        int
	instruction Instruction
	remaining   int
}

func parseInstruction(input string) Instruction {
	switch input {
	case "noop":
		return Noop
	case "addx":
		return Addx
	default:
		return Unknown
	}
}

func NewOperation(line string) *Operation {
	instruction := parseInstruction(strings.Split(line, " ")[0])
	data := 0
	remaining := 1
	if instruction == Addx {
		data = shared.ToInt(strings.Split(line, " ")[1])
		remaining = 2
	}
	return &Operation{
		data:        data,
		instruction: instruction,
		remaining:   remaining,
	}
}

type Program []*Operation

func NewProgram(input []string) Program {
	program := []*Operation{}
	for _, line := range input {
		program = append(program, NewOperation(line))
	}
	return program
}
