package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

type Gaurd struct {
	X     int
	Y     int
	Value rune
}

type Location struct {
	Visits int
	Value  rune
}

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func getLabMap(input []string) (lab [][]Location, gaurd Gaurd) {

	for y, line := range input {
		row := []Location{}
		for x, pos := range line {
			row = append(row, Location{Value: pos, Visits: 0})

			if pos == '^' || pos == 'v' || pos == '>' || pos == '<' {
				gaurd = Gaurd{X: x, Y: y, Value: pos}
			}
		}
		lab = append(lab, row)
	}
	return lab, gaurd
}

func printLab(lab [][]Location) {
	fmt.Println("")
	for _, row := range lab {
		for _, pos := range row {
			fmt.Print(string(pos.Value))
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func printLabVisits(lab [][]Location) {
	fmt.Println("")
	for _, row := range lab {
		for _, pos := range row {

			if pos.Value == 'x' {
				fmt.Print(pos.Visits)
			} else {
				fmt.Print(string(pos.Value))
			}

		}
		fmt.Println("")
	}
	fmt.Println("")
}

func surveil(lab [][]Location, gaurd Gaurd) (loop bool) {
	for {
		if gaurd.Value == '^' {

			if gaurd.Y == 0 {
				lab[gaurd.Y][gaurd.X].Value = 'x'
				break
			}

			if lab[gaurd.Y-1][gaurd.X].Value == '#' {
				lab[gaurd.Y][gaurd.X].Value = '>'
				gaurd.Value = '>'
				continue
			}

			lab[gaurd.Y][gaurd.X].Value = 'x'
			gaurd.Y = gaurd.Y - 1
			lab[gaurd.Y][gaurd.X].Value = '^'
			lab[gaurd.Y][gaurd.X].Visits = lab[gaurd.Y][gaurd.X].Visits + 1
			if lab[gaurd.Y][gaurd.X].Visits >= 5 {
				return true
			}
			continue
		}

		if gaurd.Value == 'v' {

			if gaurd.Y+1 >= len(lab) {
				lab[gaurd.Y][gaurd.X].Value = 'x'
				break
			}

			if lab[gaurd.Y+1][gaurd.X].Value == '#' {
				lab[gaurd.Y][gaurd.X].Value = '<'
				gaurd.Value = '<'
				continue
			}

			lab[gaurd.Y][gaurd.X].Value = 'x'
			gaurd.Y = gaurd.Y + 1
			lab[gaurd.Y][gaurd.X].Value = 'v'
			lab[gaurd.Y][gaurd.X].Visits = lab[gaurd.Y][gaurd.X].Visits + 1
			if lab[gaurd.Y][gaurd.X].Visits >= 5 {
				return true
			}
			continue
		}

		if gaurd.Value == '>' {
			if gaurd.X+1 >= len(lab[gaurd.Y]) {
				lab[gaurd.Y][gaurd.X].Value = 'x'
				break
			}

			if lab[gaurd.Y][gaurd.X+1].Value == '#' {
				lab[gaurd.Y][gaurd.X].Value = 'v'
				gaurd.Value = 'v'
				continue
			}

			lab[gaurd.Y][gaurd.X].Value = 'x'
			gaurd.X = gaurd.X + 1
			lab[gaurd.Y][gaurd.X].Value = '>'
			lab[gaurd.Y][gaurd.X].Visits = lab[gaurd.Y][gaurd.X].Visits + 1
			if lab[gaurd.Y][gaurd.X].Visits >= 5 {
				return true
			}
			continue
		}

		if gaurd.Value == '<' {
			if gaurd.X == 0 {
				lab[gaurd.Y][gaurd.X].Value = 'x'
				break
			}

			if lab[gaurd.Y][gaurd.X-1].Value == '#' {
				lab[gaurd.Y][gaurd.X].Value = '^'
				gaurd.Value = '^'
				continue
			}

			lab[gaurd.Y][gaurd.X].Value = 'x'
			gaurd.X = gaurd.X - 1
			lab[gaurd.Y][gaurd.X].Value = '<'
			lab[gaurd.Y][gaurd.X].Visits = lab[gaurd.Y][gaurd.X].Visits + 1
			if lab[gaurd.Y][gaurd.X].Visits >= 5 {
				return true
			}
			continue
		}
	}
	return false
}

func part1(input []string) (answer int) {

	lab, gaurd := getLabMap(input)

	surveil(lab, gaurd)

	for _, row := range lab {
		for _, pos := range row {
			if pos.Value == 'x' {
				answer++
			}
		}
	}

	return answer
}

func part2(input []string) (answer int) {

	lab, gaurd := getLabMap(input)

	surveil(lab, gaurd)

	for y, row := range lab {
		for x, pos := range row {
			if pos.Value == 'x' {

				newLab, gaurd := getLabMap(input)

				if gaurd.X == x && gaurd.Y == y {
					continue
				}

				newLab[y][x].Value = '#'

				if surveil(newLab, gaurd) {
					answer++
				}
			}
		}
	}

	return answer
}
