package main

func (c *CPU) Execute(program Program) int {
	signal := 0

	cycle := 1

	for i := 0; i < len(program); {
		op := program[i]

		// fmt.Printf("cycle: %d, register: %d\n", cycle, c.register)

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
		register: 1,
	}
}
