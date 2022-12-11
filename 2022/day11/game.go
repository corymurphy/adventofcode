package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

type Game struct {
	Monkeys         []*Monkey
	WorryManagement int
}

func parseInput(input []string) []*Monkey {

	monkeys := []*Monkey{}

	id := 0

	for i, line := range input {

		if strings.Contains(line, "Monkey ") {
			mod, op := parseOperation(input[i+2])
			monkey := Monkey{
				Items:          parseStartingItems(input[i+1]),
				Inspected:      0,
				WorryModifier:  mod,
				WorryOperation: op,
				ThrowTest:      parseThrowTest(input[i+3]),
				IfTrue:         parseIfTrue(input[i+4]),
				IfFalse:        parseIfFalse(input[i+5]),
				Id:             id,
			}

			monkeys = append(monkeys, &monkey)
			id++
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

func NewGame(input []string, worryManagement int) *Game {
	return &Game{
		Monkeys:         parseInput(input),
		WorryManagement: worryManagement,
	}
}

func (g *Game) Print() {
	for i, monkey := range g.Monkeys {
		fmt.Printf("monkey %d items %s\n", i, monkey.ItemsToString())
	}
}

func (g *Game) PlayRound(round int) {

	// fmt.Printf("starting round %d\n", round)
	for _, monkey := range g.Monkeys {
		g.TakeTurn(monkey)
	}
}

func (g *Game) TakeTurn(monkey *Monkey) {
	item, err := monkey.Items.Dequeue()

	i := 0
	for err == nil {

		monkey.Inspected = monkey.Inspected + 1

		if monkey.WorryOperation == Multiply {
			item = item * monkey.WorryModifier
		} else if monkey.WorryOperation == Add {
			item = item + monkey.WorryModifier
		} else if monkey.WorryOperation == Square {
			item = item * item
		}

		item = g.ManageWorry(item, monkey)

		throwTo := -1
		if item%monkey.ThrowTest == 0 {
			throwTo = monkey.IfTrue
		} else {
			throwTo = monkey.IfFalse
		}

		// fmt.Printf("throwing %d to %d\n", item, throwTo)

		g.Monkeys[throwTo].Items.Enqueue(item)

		item, err = monkey.Items.Dequeue()

		i++
		if i == 999 {
			panic("we're iterating wayyyy too many times") // hack lol
		}

	}
}

func (g *Game) Play(rounds int) int {
	for i := 1; i <= rounds; i++ {
		g.PlayRound(i)

		if i%1000 == 0 {
			results := []int{}
			for _, monkey := range g.Monkeys {
				results = append(results, monkey.Inspected)
			}
			sort.Ints(results)
			fmt.Println(results)
		}
	}

	results := []int{}
	for _, monkey := range g.Monkeys {
		results = append(results, monkey.Inspected)
	}
	sort.Ints(results)

	return results[len(results)-1] * results[len(results)-2]
}

func (g *Game) Supermod() int {
	mod := 1
	for _, monkey := range g.Monkeys {
		mod = mod * monkey.ThrowTest
	}
	return mod
}
func (g *Game) ManageWorry(initial int, monkey *Monkey) int {
	if g.WorryManagement == 3 {
		return initial / g.WorryManagement
	}
	return initial % g.Supermod()
}
