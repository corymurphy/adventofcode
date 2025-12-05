package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)

}

// ..xx.xx@x.
// x@@.@.@.@@
// @@@@@.x.@@
// @.@@@@..@.
// x@.@@@@.@x
// .@@@@@@@.@
// .@.@.@.@@@
// x.@@@.@@@@
// .@@@@@@@@.
// x.x.@@@.x.

func part1(inputs []string) (answer int) {

	fmt.Println()
	// for y, row := range inputs {

	// warehouse := []string{}
	for y := 0; y < len(inputs); y++ {

		w := ""
		row := inputs[y]
		for x, val := range row {

			if val != '@' && val != 'x' {
				w = w + "."
				continue
			}

			adj := 0

			// left
			if x > 0 && row[x-1] == '@' {
				adj++
			}
			// left up
			if y > 0 && x > 0 && inputs[y-1][x-1] == '@' {
				adj++
			}
			// up
			if y > 0 && inputs[y-1][x] == '@' {
				adj++
			}

			// right up
			if y > 0 && x < len(row)-1 && inputs[y-1][x+1] == '@' {
				adj++
			}

			// right
			if x < len(row)-1 && row[x+1] == '@' {
				adj++
			}

			// right down
			if y < len(inputs)-1 && x < len(row)-1 && inputs[y+1][x+1] == '@' {
				adj++
			}

			// down
			if y < len(inputs)-1 && inputs[y+1][x] == '@' {
				adj++
			}

			// left down
			if y < len(inputs)-1 && x > 0 && inputs[y+1][x-1] == '@' {
				adj++
			}

			if adj >= 4 {
				w = w + string(val)
				continue
			}

			answer++
			w = w + "x"

		}

	}
	return answer
}

func part2(inputs []string) (answer int) {

	warehouse := [][]rune{}

	for _, row := range inputs {
		wr := []rune{}
		for _, val := range row {
			wr = append(wr, val)
		}
		warehouse = append(warehouse, wr)
	}

	for {
		removed, warehouse := SetMoveable(warehouse)

		if removed == 0 {
			break
		}

		for y := 0; y < len(warehouse); y++ {
			for x := 0; x < len(warehouse[y]); x++ {
				if warehouse[y][x] == 'x' {
					warehouse[y][x] = '.'
				}
			}
		}

		// fmt.Println(removed)
		answer = answer + removed

		// for _, wr := range warehouse {
		// 	for _, r := range wr {
		// 		fmt.Print(string(r))
		// 	}
		// 	fmt.Println()
		// }

	}

	// warehouse

	return answer
}

func SetMoveable(w [][]rune) (int, [][]rune) {

	moved := 0

	for y := 0; y < len(w); y++ {

		// w := ""
		row := w[y]
		// wr := []rune{}
		for x, val := range row {

			if val != '@' && val != 'x' {
				// w = w + "."
				w[y][x] = '.'
				continue
			}

			adj := 0

			// left
			if x > 0 && (row[x-1] == '@' || row[x-1] == 'x') {
				adj++
			}
			// left up
			if y > 0 && x > 0 && (w[y-1][x-1] == '@' || w[y-1][x-1] == 'x') {
				adj++
			}
			// up
			if y > 0 && (w[y-1][x] == '@' || w[y-1][x] == 'x') {
				adj++
			}

			// right up
			if y > 0 && x < len(row)-1 && (w[y-1][x+1] == '@' || w[y-1][x+1] == 'x') {
				adj++
			}

			// right
			if x < len(row)-1 && (row[x+1] == '@' || row[x+1] == 'x') {
				adj++
			}

			// right down
			if y < len(w)-1 && x < len(row)-1 && (w[y+1][x+1] == '@' || w[y+1][x+1] == 'x') {
				adj++
			}

			// down
			if y < len(w)-1 && (w[y+1][x] == '@' || w[y+1][x] == 'x') {
				adj++
			}

			// left down
			if y < len(w)-1 && x > 0 && (w[y+1][x-1] == '@' || w[y+1][x-1] == 'x') {
				adj++
			}

			if adj >= 4 {
				// w = w + string(val)

				continue
			}

			moved++
			// w = w + "x"
			w[y][x] = 'x'

		}
		// wr = append(wr, )

	}
	return moved, w
}

// func Is
// func (w *Warehouse) PrintSimulation() {
// 	for y, row := range *w {
// 		for x, val := range row {

// 			// out := val
// 			// if val == 0 {
// 			// 	out = "."
// 			// }

// 			fmt.Printf("\033[%d;%dH%s", y+5, x+2, string(val))
// 		}
// 	}
// 	fmt.Println()
// }
