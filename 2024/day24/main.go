package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

// type Wires struct {
// }

type Instruction struct {
	Input1 string
	Input2 string
	Op     string
	Output string
}

type Wire struct {
	Key   string
	Value string
	// Value int
	// HasValue
}

type Monitor struct {
	Wires        []Wire
	Instructions []Instruction
}

func (m *Monitor) SetWire(new Wire) {

	for i, wire := range m.Wires {
		if wire.Key == new.Key {
			m.Wires[i].Value = new.Value
			return
		}
	}
	m.Wires = append(m.Wires, new)
}

func (m *Monitor) GetWire(key string) Wire {
	for _, wire := range m.Wires {
		if wire.Key == key {
			return wire
		}
	}
	return Wire{Key: key, Value: ""}
}

func NewMonitor(input []string) (monitor Monitor) {

	monitor = Monitor{
		Wires:        []Wire{},
		Instructions: []Instruction{},
	}

	instruct := false
	for _, line := range input {

		if instruct {

			monitor.Instructions = append(monitor.Instructions, Instruction{
				Input1: strings.Split(line, " ")[0],
				Op:     strings.Split(line, " ")[1],
				Input2: strings.Split(line, " ")[2],
				Output: strings.Split(line, " ")[4],
			})

			continue
		}

		if line == "" {
			instruct = true
			continue
		}

		wire := Wire{
			Key:   strings.Split(line, ": ")[0],
			Value: strings.Split(line, ": ")[1],
		}
		monitor.SetWire(wire)
	}

	return monitor
}

func (m *Monitor) Start() {

	queue := NewQueue()

	for _, instruction := range m.Instructions {
		queue.Enqueue(&instruction)
	}

	for !queue.Empty() {

		instruction, _ := queue.Dequeue()

		i1 := m.GetWire(instruction.Input1)
		i2 := m.GetWire(instruction.Input2)

		out := m.GetWire(instruction.Output)

		if i1.Value == "" || i2.Value == "" {
			queue.Enqueue(instruction)
			continue
		}

		if instruction.Op == "AND" {
			if i1.Value == "1" && i2.Value == "1" {
				out.Value = "1"
			}
			if i1.Value == "0" || i2.Value == "0" {
				out.Value = "0"
			}
		}

		if instruction.Op == "OR" {
			if i1.Value == "1" || i2.Value == "1" {
				out.Value = "1"
			}
			if i1.Value == "0" && i2.Value == "0" {
				out.Value = "0"
			}

		}

		if instruction.Op == "XOR" {
			if i1.Value != i2.Value {
				out.Value = "1"
			}
			if i1.Value == i2.Value {
				out.Value = "0"
			}
		}
		m.SetWire(out)
	}
}

func part1(input []string) (answer int) {

	monitor := NewMonitor(input)
	monitor.Start()

	sort.Slice(monitor.Wires, func(i, j int) bool {
		return monitor.Wires[i].Key < monitor.Wires[j].Key
	})

	serialized := ""

	for i := len(monitor.Wires) - 1; i > 0; i-- {
		wire := monitor.Wires[i]
		if strings.HasPrefix(wire.Key, "z") {
			serialized = serialized + wire.Value
		}
	}

	output, _ := strconv.ParseInt(serialized, 2, 64)

	answer = int(output)

	return answer
}

func part2(input []string) (answer int) {

	return answer
}
