package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

type Antenna struct {
	X         int
	Y         int
	Frequency rune
}

type Location struct {
	X         int
	Y         int
	Peers     []Location
	Value     rune
	Antinodes int
}

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func CreateCity(input []string) (city [][]Location) {

	for y, line := range input {
		row := []Location{}
		for x, pos := range line {
			row = append(row, Location{X: x, Y: y, Value: pos, Peers: []Location{}, Antinodes: 0})
		}
		city = append(city, row)
	}
	return city
}

func PresentCity(city [][]Location) {
	fmt.Println()
	for _, row := range city {
		for _, pos := range row {
			fmt.Print(string(pos.Value))
		}
		fmt.Println()
	}
	fmt.Println()
}

func PresentCityWithAntinodes(city [][]Location) {
	fmt.Println()
	for _, row := range city {
		for _, pos := range row {
			if pos.Antinodes > 0 && pos.Value == '.' {
				fmt.Print("#")
				continue
			}
			fmt.Print(string(pos.Value))
		}
		fmt.Println()
	}
	fmt.Println()
}

func PresentPeers(city [][]Location, start Location) {
	fmt.Println()
	for y, row := range city {
		for x, pos := range row {

			if start.X == x && start.Y == y {
				fmt.Print(string(pos.Value))
				continue
			}

			hasPeers := false
			for _, peer := range start.Peers {
				if peer.X == x && peer.Y == y {
					fmt.Print(string(peer.Value))
					hasPeers = true
				}
			}

			if !hasPeers {
				fmt.Print(string('.'))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func FindPeers(city [][]Location, start *Location) {
	for y, row := range city {
		for x, pos := range row {
			if pos.Value == start.Value && y != start.Y && x != start.X {
				start.Peers = append(start.Peers, pos)
			}
		}
	}
}

func FindAntinodes(city [][]Location, start Location) {

	for _, peer := range start.Peers {

		diffX := start.X - peer.X
		diffY := start.Y - peer.Y

		antinodeX := start.X + diffX
		antinodeY := start.Y + diffY

		if antinodeX >= len(city[0]) || antinodeX < 0 {
			continue
		}

		if antinodeY >= len(city) || antinodeY < 0 {
			continue
		}

		city[antinodeY][antinodeX].Antinodes = city[antinodeY][antinodeX].Antinodes + 1
	}
}

func FindAntinodes2(city [][]Location, start Location) {

	for _, peer := range start.Peers {

		diffX := start.X - peer.X
		diffY := start.Y - peer.Y

		antinodeX := start.X + diffX
		antinodeY := start.Y + diffY

		city[peer.Y][peer.X].Antinodes = city[peer.Y][peer.X].Antinodes + 1
		city[start.Y][start.X].Antinodes = city[start.Y][start.X].Antinodes + 1

		for {
			if antinodeX >= len(city[0]) || antinodeX < 0 {
				break
			}

			if antinodeY >= len(city) || antinodeY < 0 {
				break
			}

			city[antinodeY][antinodeX].Antinodes = city[antinodeY][antinodeX].Antinodes + 1

			antinodeX = antinodeX + diffX
			antinodeY = antinodeY + diffY
		}
	}
}

func CountAntinodes(city [][]Location) (count int) {
	for _, row := range city {
		for _, pos := range row {
			if pos.Antinodes > 0 {
				count++
			}
		}
	}
	return count
}

func part1(input []string) (answer int) {

	city := CreateCity(input)

	for _, row := range city {
		for _, pos := range row {

			if pos.Value == '.' {
				continue
			}

			FindPeers(city, &pos)
			FindAntinodes(city, pos)
		}
	}

	answer = CountAntinodes(city)

	return answer
}

func part2(input []string) (answer int) {

	city := CreateCity(input)

	for _, row := range city {
		for _, pos := range row {

			if pos.Value == '.' {
				continue
			}

			FindPeers(city, &pos)
			FindAntinodes2(city, pos)
		}
	}

	PresentCityWithAntinodes(city)
	answer = CountAntinodes(city)

	return answer
}
