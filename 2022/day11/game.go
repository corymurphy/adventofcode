package main

import (
	"fmt"
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

type Game struct {
	Monkeys []*Monkey
}

func parseInput(input []string) []*Monkey {

	monkeys := []*Monkey{}

	for i, line := range input {

		if strings.Contains(line, "Monkey ") {
			mod, op := parseOperation(input[i+2])
			monkey := Monkey{
				Items:          parseStartingItems(input[i+1]),
				Inspected:      0,
				WorryModifier:  mod,
				WorryOperation: op,
				ThrowTest:      parseThrowTest(input[3]),
				IfTrue:         parseIfTrue(input[4]),
				IfFalse:        parseIfFalse(input[5]),
			}

			monkeys = append(monkeys, &monkey)
		}
	}
	return monkeys
}

func parseThrowTest(input string) int {
	trimmed := strings.TrimPrefix(input, "  Test: divisible by ")
	return shared.ToInt(trimmed)
}

func parseIfTrue(input string) int {
	trimmed := strings.TrimPrefix(input, "    If true: throw to monkey ")
	return shared.ToInt(trimmed)
}

func parseIfFalse(input string) int {
	trimmed := strings.TrimPrefix(input, "    If false: throw to monkey ")
	return shared.ToInt(trimmed)
}

func parseOperation(input string) (int, Operation) {
	trimmed := strings.TrimPrefix(input, "  Operation: new = ")

	values := strings.Split(trimmed, " ")

	op := Unknown
	mod := 0

	if values[2] == "old" {
		return 0, Square
	}

	op = ParseOperation(values[1])
	mod = shared.ToInt(values[2])

	return mod, op
}

func parseStartingItems(input string) *shared.IntQueue {
	queue := shared.NewIntQueue()

	trimmedInput := strings.TrimPrefix(input, "  Starting items: ")

	for _, item := range strings.Split(trimmedInput, ", ") {
		queue.Enqueue(shared.ToInt(item))
	}

	return queue
}

func NewGame(input []string) *Game {
	return &Game{
		Monkeys: parseInput(input),
	}
}

func (g *Game) Print() {
	for i, monkey := range g.Monkeys {
		fmt.Printf("monkey %d starting items %s\n", i, monkey.ItemsToString())
	}
}
