package main

import (
	"fmt"
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

func part1(inputs []string) (answer int64) {

	iop := len(inputs) - 1
	l := len(inputs[0])

	for x := 0; x < l; {

		next := x + 1
		for next < l {
			if inputs[iop][next] == ' ' {
				next++
				continue
			}
			break
		}

		fmt.Print()

		op := inputs[iop][x]

		a := int64(0)

		for y := 0; y < iop; y++ {
			sval := inputs[y][x:next]
			val := shared.MustInt64(strings.TrimSpace(sval))

			if y == 0 {
				a = val
				continue
			}

			if op == '*' {
				a = a * val
			} else if op == '+' {
				a = a + val
			} else {
				panic("shouldn't get here")
			}

		}

		answer = answer + a
		x = next

		if x >= l {
			break
		}

	}
	return answer
}

func part2(inputs []string) (answer int64) {

	iop := len(inputs) - 1
	l := len(inputs[0])

	for x := 0; x < l; {

		next := x + 1
		for next < l {
			if inputs[iop][next] == ' ' {
				next++
				continue
			}
			break
		}

		op := inputs[iop][x]

		a := int64(0)

		for x < next {
			sval := ""
			for y := 0; y < iop; y++ {
				sval = sval + string(inputs[y][x])
			}

			sval = strings.TrimSpace(sval)

			if sval == "" {
				break
			}

			val := shared.MustInt64(sval)

			if a == 0 {
				a = val
				x++
				continue
			}

			if op == '*' {
				a = a * val
			} else if op == '+' {
				a = a + val
			} else {
				panic("shouldn't get here")
			}
			x++
		}

		answer = answer + a
		x = next

		if x >= l {
			break
		}

	}
	return answer
}
