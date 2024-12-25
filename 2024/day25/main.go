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

type Lock [][]string
type Key [][]string

type Schematics struct {
	Keys  []Key
	Locks []Lock
}

func GetSchematics(input []string) (schematics Schematics) {

	schematics = Schematics{
		Locks: []Lock{},
		Keys:  []Key{},
	}

	for i := 0; i < len(input); i++ {
		is := input[i : i+7]

		schematic := [][]string{}
		for _, row := range is {
			x := []string{}
			for _, val := range row {
				x = append(x, string(val))
			}
			schematic = append(schematic, x)
		}

		if input[i] == "#####" {
			schematics.Locks = append(schematics.Locks, schematic)
		}

		if input[i] == "....." {
			schematics.Keys = append(schematics.Keys, schematic)
		}

		i += 7
	}

	return schematics
}

func (k *Key) Heights() (heights []int) {
	heights = []int{0, 0, 0, 0, 0}

	h := 0
	for y := len(*k) - 1; y >= 0; y-- {
		for x, val := range (*k)[y] {
			if val == "#" {
				heights[x] = h
			}
		}
		h++
	}
	return heights
}

func (l *Lock) Heights() (heights []int) {

	heights = []int{0, 0, 0, 0, 0}
	for y, row := range *l {
		for x, val := range row {
			if val == "#" {
				heights[x] = y
			}
		}
	}
	return heights
}

func part1(input []string) (answer int) {

	schematics := GetSchematics(input)

	// for _, key := range schematics.Keys {
	// 	fmt.Println(key.Heights())
	// }

	// for _, lock := range schematics.Locks {
	// 	fmt.Println(lock.Heights())
	// }

	fits := 0

	for _, lock := range schematics.Locks {
		for _, key := range schematics.Keys {

			kh := key.Heights()
			lh := lock.Heights()

			for x := 0; x < len(kh); x++ {
				if kh[x]+lh[x] >= 6 {
					// fmt.Println("lock", lh, "and key", kh, "overlap on column", x)
					break
				}
				if x == len(kh)-1 {
					fits++
				}
			}
		}
	}

	// fmt.Println(fits)

	// answer = len(schematics.Locks)
	// answer += len(schematics.Keys)
	answer = fits
	return answer
}

func part2(input []string) (answer int) {

	return answer
}
