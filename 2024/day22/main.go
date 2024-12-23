package main

import (
	"fmt"
	"strconv"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input_sample2")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func ToInt(input string) int {
	value, _ := strconv.ParseInt(input, 10, 64)
	return int(value)
}

func part1(input []string) (answer int) {

	for _, l := range input {
		secret := ToInt(l)

		for range 2000 {
			secret = ((secret * 64) ^ secret) % 16777216

			secret = ((secret / 32) ^ secret) % 16777216

			secret = ((secret * 2048) ^ secret) % 16777216
		}
		answer += secret
	}

	return answer
}

type Sequence []int

func (s *Sequence) Add(new int) {

	if len(*s) < 4 {
		*s = append(*s, new)
		return
	}

	for i := len(*s) - 1; i > 0; i-- {
		(*s)[i] = (*s)[i-1]
	}

	(*s)[0] = new
}

func (s *Sequence) Print() {
	fmt.Println(s)
}

func part2(input []string) (answer int) {

	for _, l := range []string{"123"} {
		secret := ToInt(l)

		// s := NewQueue()

		sequence := Sequence{}

		previous := secret
		fmt.Printf("%s: %d\n", l, secret%10)
		for range 10 {

			secret = ((secret * 64) ^ secret) % 16777216

			secret = ((secret / 32) ^ secret) % 16777216

			secret = ((secret * 2048) ^ secret) % 16777216

			// change :=

			sequence.Add(secret%10 - previous%10)
			fmt.Printf("%d: %d (%d)\n", secret, secret%10, secret%10-previous%10)
			previous = secret

			// fmt.Println()

			if secret%10 == 6 && len(sequence) >= 4 {
				sequence.Print()
			}

			// if secret == 6 && queue.Size() >= 4 {

			// 	// for queue.Empty()
			// 	fmt.Println(queue.Dequeue(), queue.Dequeue(), queue.Dequeue(), queue.Dequeue())
			// }
			// continue

			// if i != 0 {
			// 	fmt.Printf("%s: %d (%d)", l)
			// } else {
			// 	fmt.Printf("%s: %d", l, secret%10)
			// }

			// fmt.Println(secret % 10)

		}

		// sequence.Print()
		// queue.Print()
		// fmt.Println(l, ": ", secret, previous)
		fmt.Println()

		answer += secret
	}

	return answer
}
