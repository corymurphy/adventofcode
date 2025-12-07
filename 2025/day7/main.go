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

func PrintSimulation(lab [][]rune) {
	// time.Sleep(1 * time.Millisecond)
	for y, row := range lab {
		for x, val := range row {
			fmt.Printf("\033[%d;%dH%s", y+5, x+2, string(val))
		}
	}
	fmt.Println()
}

func Print(lab [][]rune) {
	for _, row := range lab {
		for _, val := range row {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
}

func GetStart(inputs []string) (start Beam) {
	for x, val := range inputs[0] {
		if val == 'S' {
			return Beam{
				X: x,
				Y: 0,
			}
		}
	}
	return Beam{}
}

func GetLab(inputs []string) (lab [][]rune, start Beam) {
	lab = [][]rune{}

	for y, input := range inputs {

		row := []rune{}
		for x, val := range input {
			if val == 'S' {
				start = Beam{X: x, Y: y}
			}
			row = append(row, val)
		}
		lab = append(lab, row)
	}

	return lab, start
}

func part1(inputs []string) (answer int) {

	visited := map[Beam]bool{}
	lab, start := GetLab(inputs)
	split := 0

	beams := NewQueue()

	beams.Enqueue(start)
	for !beams.Empty() {
		beam, err := beams.Dequeue()
		if err != nil {
			panic("this shouldn't happen")
		}

		if beam.Y+1 >= len(inputs) {
			continue
		}

		beam.Y = beam.Y + 1

		if lab[beam.Y][beam.X] != '^' {
			lab[beam.Y][beam.X] = '|'
			visited[beam] = true
			beams.Enqueue(beam)
			continue
		}

		wasSplit := false

		beamLeft := Beam{X: beam.X - 1, Y: beam.Y}
		if _, exists := visited[beamLeft]; !exists {
			lab[beamLeft.Y][beamLeft.X] = '|'
			visited[beamLeft] = true
			beams.Enqueue(beamLeft)
			wasSplit = true
		}

		beamRight := Beam{X: beam.X + 1, Y: beam.Y}
		if _, exists := visited[beamRight]; !exists {
			lab[beamRight.Y][beamRight.X] = '|'
			visited[beamRight] = true
			beams.Enqueue(beamRight)
			wasSplit = true
		}

		if wasSplit {
			split++
		}
	}

	answer = split

	return answer
}

type Beam struct {
	X int
	Y int
}

func part2(inputs []string) (answer int64) {
	_, start := GetLab(inputs)
	answer = countTimelines(inputs, pos{r: 0, c: start.X})
	return answer
}

type pos struct{ r, c int }

func countTimelines(lab []string, start pos) int64 {
	memo := make(map[pos]int64)

	var dfs func(p pos) int64
	dfs = func(p pos) int64 {
		// Out‑of‑bounds → the beam has exited the diagram → one finished timeline.
		if p.r >= len(lab) || p.c < 0 || p.c >= len(lab[p.r]) {
			return 1
		}
		// If we have already solved this sub‑problem, reuse it.
		if v, ok := memo[p]; ok {
			return v
		}

		// Look at the cell *directly below* the current one.
		nextR := p.r + 1
		// Reached the bottom → finished timeline.
		if nextR >= len(lab) {
			memo[p] = 1
			return 1
		}
		cellBelow := lab[nextR][p.c]

		var result int64
		if cellBelow == '^' {
			// Splitter: the beam stops on the splitter and spawns left/right beams.
			left := pos{nextR, p.c - 1}
			right := pos{nextR, p.c + 1}
			result = dfs(left) + dfs(right)
		} else {
			// Empty space (or the start cell itself) → keep falling.
			result = dfs(pos{nextR, p.c})
		}
		memo[p] = result
		return result
	}

	// The problem states that the beam *starts moving downwards* from S,
	// i.e. the first step is the cell directly below S.
	return dfs(start)
}
