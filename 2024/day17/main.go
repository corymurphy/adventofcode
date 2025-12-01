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

func (cpu *Computer) Run() {
	for cpu.Counter < len(cpu.Program) {
		opcode := cpu.Program[cpu.Counter]
		operand := cpu.Program[cpu.Counter+1]

		// fmt.Println("opcode:", opcode, ";", "operand:", operand)

		// yep we're going it this way for now stay mad

		// adv
		if opcode == 0 {

			if operand >= 0 && operand <= 3 {
				cpu.A = int(float64(cpu.A) / math.Pow(2, float64(operand)))
			}

			if operand == 4 {
				cpu.A = int(float64(cpu.A) / math.Pow(2, float64(cpu.A)))
			}

			if operand == 5 {
				cpu.A = int(float64(cpu.A) / math.Pow(2, float64(cpu.B)))
			}

			if operand == 6 {
				cpu.A = int(float64(cpu.A) / math.Pow(2, float64(cpu.C)))
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
				cpu.B = int(float64(cpu.A) / math.Pow(2, float64(operand)))
			}

			if operand == 4 {
				cpu.B = int(float64(cpu.A) / math.Pow(2, float64(cpu.A)))
			}

			if operand == 5 {
				cpu.B = int(float64(cpu.A) / math.Pow(2, float64(cpu.B)))
			}

			if operand == 6 {
				cpu.B = int(float64(cpu.A) / math.Pow(2, float64(cpu.C)))
			}

			if operand == 7 {
				panic("reserved")
			}
		}

		if opcode == 7 {
			if operand >= 0 && operand <= 3 {
				cpu.C = int(float64(cpu.A) / math.Pow(2, float64(operand)))
			}

			if operand == 4 {
				cpu.C = int(float64(cpu.A) / math.Pow(2, float64(cpu.A)))
			}

			if operand == 5 {
				cpu.C = int(float64(cpu.A) / math.Pow(2, float64(cpu.B)))
			}

			if operand == 6 {
				cpu.C = int(float64(cpu.A) / math.Pow(2, float64(cpu.C)))
			}

			if operand == 7 {
				panic("reserved")
			}
		}

		cpu.Counter += 2
	}
}

func (cpu *Computer) Run2() {
	for cpu.Counter < len(cpu.Program) {
		opcode := cpu.Program[cpu.Counter]
		operand := cpu.Program[cpu.Counter+1]

		// fmt.Println("opcode:", opcode, ";", "operand:", operand)

		// yep we're going it this way for now stay mad

		// adv
		if opcode == 0 {

			if operand >= 0 && operand <= 3 {
				cpu.A = int(float64(cpu.A) / math.Pow(2, float64(operand)))
			}

			if operand == 4 {
				cpu.A = int(float64(cpu.A) / math.Pow(2, float64(cpu.A)))
			}

			if operand == 5 {
				cpu.A = int(float64(cpu.A) / math.Pow(2, float64(cpu.B)))
			}

			if operand == 6 {
				cpu.A = int(float64(cpu.A) / math.Pow(2, float64(cpu.C)))
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
				cpu.B = cpu.A % 8
			}

			if operand == 5 {
				cpu.B = cpu.B % 8
			}

			if operand == 6 {
				cpu.B = cpu.C % 8
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
			result := 0
			if operand >= 0 && operand <= 3 {
				result = operand % 8
			}

			if operand == 4 {
				result = cpu.A % 8
			}

			if operand == 5 {
				result = cpu.B % 8
			}

			if operand == 6 {
				result = cpu.C % 8
			}

			if operand == 7 {
				panic("reserved")
			}

			if result != cpu.Program[len(cpu.Out)] {
				return
			}

			cpu.Out = append(cpu.Out, result)

			// sl[len(sl)-1]

		}

		// bdv
		if opcode == 6 {
			if operand >= 0 && operand <= 3 {
				cpu.B = int(float64(cpu.A) / math.Pow(2, float64(operand)))
			}

			if operand == 4 {
				cpu.B = int(float64(cpu.A) / math.Pow(2, float64(cpu.A)))
			}

			if operand == 5 {
				cpu.B = int(float64(cpu.A) / math.Pow(2, float64(cpu.B)))
			}

			if operand == 6 {
				cpu.B = int(float64(cpu.A) / math.Pow(2, float64(cpu.C)))
			}

			if operand == 7 {
				panic("reserved")
			}
		}

		if opcode == 7 {
			if operand >= 0 && operand <= 3 {
				cpu.C = int(float64(cpu.A) / math.Pow(2, float64(operand)))
			}

			if operand == 4 {
				cpu.C = int(float64(cpu.A) / math.Pow(2, float64(cpu.A)))
			}

			if operand == 5 {
				cpu.C = int(float64(cpu.A) / math.Pow(2, float64(cpu.B)))
			}

			if operand == 6 {
				cpu.C = int(float64(cpu.A) / math.Pow(2, float64(cpu.C)))
			}

			if operand == 7 {
				panic("reserved")
			}
		}

		cpu.Counter += 2
	}
}

func main() {
	// part1 := part1(shared.ReadInput("input_sample"))
	part2 := part2(shared.ReadInput("input"))

	// fmt.Printf("\nPart 1 answer: %s\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) (answer string) {

	cpu := NewComputer(input)
	cpu.Run()

	answer = cpu.OutString()

	return answer
}

func ToInt(input string) int {
	value, _ := strconv.ParseInt(input, 10, 64)
	return int(value)
}

func (cpu *Computer) OutString() (output string) {
	output = strconv.Itoa(cpu.Out[0])
	for _, val := range cpu.Out[1:] {
		output = fmt.Sprintf("%s,%d", output, val)
	}
	return output
}

func (cpu *Computer) Reset(a int) {
	cpu.A = a
	cpu.B = 0
	cpu.C = 0
	cpu.Counter = 0
	cpu.Out = []int{}
}

func (cpu *Computer) ProgramEqualsOut() bool {

	if len(cpu.Program) != len(cpu.Out) {
		return false
	}

	for i, prog := range cpu.Program {
		if cpu.Out[i] != prog {
			return false
		}
	}
	return true
}

func part2(input []string) (answer int) {

	cpu := NewComputer(input)

	// 37826000000
	for a := 37826000000; ; a++ {

		cpu.Reset(a)
		cpu.Run2()

		// if a%10000000 == 0 {
		// 	fmt.Println(a)
		// }

		if cpu.ProgramEqualsOut() {
			answer = a
			break
		}
	}

	return answer
}
