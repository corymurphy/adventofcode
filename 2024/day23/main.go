package main

import (
	"fmt"
	"slices"
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

type Connection struct {
	A string
	B string
}

type Connections []string

func GetNetwork(input []string) (network map[string]Connections) {
	network = make(map[string]Connections)

	for _, line := range input {

		a := strings.Split(line, "-")[0]
		b := strings.Split(line, "-")[1]

		if connections, contains := network[a]; contains {
			network[a] = append(connections, b)
		} else {
			network[a] = Connections{b}
		}

		if connections, contains := network[b]; contains {
			network[b] = append(connections, a)
		} else {
			network[b] = Connections{a}
		}
	}
	return network
}

func (s1 Connections) Contains(s2 Connections) (con Connections) {

	con = Connections{}

	for _, el1 := range s1 {
		for _, el2 := range s2 {
			if el1 == el2 {
				con = append(con, el1)
			}
		}
	}
	return con
}

func part1(input []string) (answer int) {

	network := GetNetwork(input)

	inter := make(map[string]int)

	for cpu1, con1 := range network {

		for _, cpu2 := range con1 {

			con2 := network[cpu2]

			for _, common := range con1.Contains(con2) {

				connected := []string{cpu1, cpu2, common}
				slices.Sort(connected)
				inter[strings.Join(connected, ":")] = 0

			}
		}
	}

	for conns, _ := range inter {
		for _, computer := range strings.Split(conns, ":") {
			if strings.HasPrefix(computer, "t") {
				answer++
				break
			}
		}
	}
	return answer
}

// func FindConnected

func part2(input []string) (answer int) {
	// network := GetNetwork(input)

	return answer
}
