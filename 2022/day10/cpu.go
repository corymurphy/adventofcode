package main

import "fmt"

func shouldDraw(cycle int, register int) bool {
	pixel := cycle - 1

	upper := register + 1
	lower := register - 1
	if pixel >= lower && pixel <= upper {
		return true
	}
	return false
}

func (c *CPU) Draw(cycle int, register int) {

	offset := (cycle / 40) * 40

	if shouldDraw(cycle, register+offset) {
		c.display[cycle-1] = 1
	}

}

func (c *CPU) Execute(program Program) int {
	signal := 0

	cycle := 1

	for i := 0; i < len(program); {
		op := program[i]

		// fmt.Printf("cycle: %d, register: %d\n", cycle, c.register)

		c.Draw(cycle, c.register)

		// if shouldDraw(cycle, c.register) {
		// 	c.display[cycle-1] = 1
		// 	// c.Draw(cycle, c.register)

		// 	// fmt.Printf("cycle: %d, register: %d\n", cycle, c.register)
		// }

		if isSignificantCycle(cycle) {
			signal = signal + (cycle * c.register)
		}

		op.remaining = op.remaining - 1

		if op.remaining <= 0 {
			if op.instruction == Addx {
				c.register = c.register + op.data
			}
			i++
		}

		cycle++
	}

	return signal
}

type CPU struct {
	register int
	display  []int
	// cycle             int
	// significantCycles []int
}

func SignificantCycles() []int {
	return []int{20, 60, 100, 140, 180, 220}
}

func isSignificantCycle(cycle int) bool {
	significant := false
	for _, significantCycle := range SignificantCycles() {
		if cycle == significantCycle {
			return true
		}
	}
	return significant
}

func NewCpu() *CPU {
	return &CPU{
		display:  make([]int, 240),
		register: 1,
	}
}

func (c *CPU) Display() {

	for i, pixel := range c.display {
		newLine := i%40 == 0

		if newLine {
			fmt.Println("")
		}

		if pixel == 1 {
			fmt.Printf("#")
		} else {
			fmt.Print(".")
		}
		// fmt.Printf()
	}
	fmt.Println("")
}

// type CRT []int

// func (c *CRT) Draw
