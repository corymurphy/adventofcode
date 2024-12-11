package main

import (
	"fmt"
	"strconv"
	"strings"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) (answer int) {

	stones := CreateMapFromStones(strings.Split(input[0], " "))

	for range 25 {

		newStones := NewStones()
		for stone := range stones {

			c := stones[stone]
			if stone == 0 {
				newStones.Add(1, c)
			} else if len(strconv.Itoa(stone))%2 == 0 {

				s := strconv.Itoa(stone)

				first := shared.ToInt(s[:len(s)/2])
				second := shared.ToInt(s[len(s)/2:])
				newStones.Add(first, c)
				newStones.Add(second, c)

			} else {
				newStones.Add(stone*2024, c)
			}
		}
		stones = newStones
	}
	answer = stones.Count()

	return answer
}

type Stones map[int]int

func CreateMapFromStones(input []string) (stones Stones) {

	stones = make(map[int]int)

	for _, s := range input {

		stone := shared.ToInt(s)

		if c, contains := stones[stone]; contains {
			stones[stone] = c + 1
		} else {
			stones[stone] = 1
		}
	}
	return stones
}

func NewStones() (s Stones) {
	return make(map[int]int)
}

func (s *Stones) Count() (count int) {

	for _, val := range *s {
		count += val
	}
	return count
}

func (s *Stones) Add(stone int, count int) {
	if current, contains := (*s)[stone]; contains {
		(*s)[stone] = current + count
	} else {
		(*s)[stone] = count
	}
}

func part2(input []string) (answer int) {

	stones := CreateMapFromStones(strings.Split(input[0], " "))

	for range 75 {

		newStones := NewStones()
		for stone := range stones {

			c := stones[stone]
			if stone == 0 {
				newStones.Add(1, c)
			} else if len(strconv.Itoa(stone))%2 == 0 {

				s := strconv.Itoa(stone)

				first := shared.ToInt(s[:len(s)/2])
				second := shared.ToInt(s[len(s)/2:])
				newStones.Add(first, c)
				newStones.Add(second, c)

			} else {
				newStones.Add(stone*2024, c)
			}
		}
		stones = newStones
	}
	answer = stones.Count()
	return answer
}
