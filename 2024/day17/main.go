package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

type Computer struct {
	A       int
	B       int
	C       int
	Program []int
	Counter int
	Out     []int
}

func NewComputer(input []string) (computer Computer) {
	computer = Computer{
		A:       ToInt(strings.Split(input[0], ": ")[1]),
		B:       ToInt(strings.Split(input[1], ": ")[1]),
		C:       ToInt(strings.Split(input[2], ": ")[1]),
		Program: []int{},
		Out:     []int{},
	}

	program := strings.Split(strings.Split(input[4], ": ")[1], ",")
	for _, opcode := range program {
		computer.Program = append(computer.Program, ToInt(opcode))
	}

	return computer
}

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) (answer int) {

	cpu := NewComputer(input)
	// fmt.Println(cpu)

	for cpu.Counter < len(cpu.Program) {
		opcode := cpu.Program[cpu.Counter]
		operand := cpu.Program[cpu.Counter+1]

		// fmt.Println("opcode:", opcode, ";", "operand:", operand)

		// yep we're going it this way for now stay mad

		// adv
		if opcode == 0 {

			if operand >= 0 && operand <= 3 {
				cpu.A = cpu.A / int(math.Pow(2, float64(operand)))
			}

			if operand == 4 {
				cpu.A = cpu.A / int(math.Pow(2, float64(cpu.A)))
			}

			if operand == 5 {
				cpu.A = cpu.A / int(math.Pow(2, float64(cpu.B)))
			}

			if operand == 6 {
				cpu.A = cpu.A / int(math.Pow(2, float64(cpu.C)))
			}

			if operand == 7 {
				panic("reserved")
			}
		}

		// bxl
		if opcode == 1 {
			cpu.B = cpu.B ^ operand
		}

		// bst
		if opcode == 2 {
			if operand >= 0 && operand <= 3 {
				cpu.B = operand % 8
			}

			if operand == 4 {
				// cpu.B = int(math.Pow(2, float64(cpu.A))) % 8
				cpu.B = cpu.A % 8
			}

			if operand == 5 {
				// cpu.B = int(math.Pow(2, float64(cpu.B))) % 8
				cpu.B = cpu.B % 8
			}

			if operand == 6 {
				// cpu.B = int(math.Pow(2, float64(cpu.C))) % 8
				cpu.B = cpu.C % 8
			}

			if operand == 7 {
				panic("reserved")
			}
		}

		// jnz
		if opcode == 3 {
			if cpu.A != 0 {
				cpu.Counter = operand
				continue
			}
		}

		// bxc
		if opcode == 4 {
			cpu.B = cpu.B ^ cpu.C
		}

		// out
		if opcode == 5 {
			if operand >= 0 && operand <= 3 {
				cpu.Out = append(cpu.Out, operand%8)
			}

			if operand == 4 {
				cpu.Out = append(cpu.Out, cpu.A%8)
			}

			if operand == 5 {
				cpu.Out = append(cpu.Out, cpu.B%8)
			}

			if operand == 6 {
				cpu.Out = append(cpu.Out, cpu.C%8)
			}

			if operand == 7 {
				panic("reserved")
			}
		}

		// bdv
		if opcode == 6 {
			if operand >= 0 && operand <= 3 {
				cpu.B = cpu.A / int(math.Pow(2, float64(operand)))
			}

			if operand == 4 {
				cpu.B = cpu.A / int(math.Pow(2, float64(cpu.A)))
			}

			if operand == 5 {
				cpu.B = cpu.A / int(math.Pow(2, float64(cpu.B)))
			}

			if operand == 6 {
				cpu.B = cpu.A / int(math.Pow(2, float64(cpu.C)))
			}

			if operand == 7 {
				panic("reserved")
			}
		}

		if opcode == 7 {
			if operand >= 0 && operand <= 3 {
				cpu.C = cpu.A / int(math.Pow(2, float64(operand)))
			}

			if operand == 4 {
				cpu.C = cpu.A / int(math.Pow(2, float64(cpu.A)))
			}

			if operand == 5 {
				cpu.C = cpu.A / int(math.Pow(2, float64(cpu.B)))
			}

			if operand == 6 {
				cpu.C = cpu.A / int(math.Pow(2, float64(cpu.C)))
			}

			if operand == 7 {
				panic("reserved")
			}
		}

		cpu.Counter += 2
	}

	// fmt.Println(cpu.Out)

	output := ""
	for _, val := range cpu.Out {
		output = fmt.Sprintf("%s%d", output, val)
	}

	// fmt.Println(output)
	// fmt.Println(ToInt(output))

	answer = ToInt(output)

	return answer
}

func ToInt(input string) int {
	value, _ := strconv.ParseInt(input, 10, 64)
	return int(value)
}

func part2(input []string) (answer int) {

	// fmt.Println(math.Pow(2, 4))
	return answer
}
