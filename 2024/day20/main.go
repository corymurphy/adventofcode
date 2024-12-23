package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func (p *Position) Up() (move Position) {
	return Position{
		Y: p.Y - 1,
		X: p.X,
	}
}

func (p *Position) Down() (move Position) {
	return Position{
		Y: p.Y + 1,
		X: p.X,
	}
}

func (p *Position) Left() (move Position) {
	return Position{
		X: p.X - 1,
		Y: p.Y,
	}
}

func (p *Position) Right() (move Position) {
	return Position{
		X: p.X + 1,
		Y: p.Y,
	}
}

type Position struct {
	X int
	Y int
}

type Track [][]rune

func (t *Track) Print() {
	fmt.Println()
	for _, row := range *t {
		for _, val := range row {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
}

func NewTrack(input []string) (track *Track, start Position, end Position) {
	track = &Track{}
	start = Position{}
	end = Position{}

	for y, ys := range input {
		*track = append(*track, []rune{})
		for x, val := range ys {
			(*track)[y] = append((*track)[y], val)

			if val == 'S' {
				start.X = x
				start.Y = y
			}

			if val == 'E' {
				end.X = x
				end.Y = y
			}
		}
	}
	return track, start, end
}

type Route struct {
	Position Position
	Finished bool
	Elapsed  int
}

func (track *Track) Next(pos Position) (moves []Position) {

	moves = []Position{}

	for _, move := range []Position{
		pos.Down(),
		pos.Up(),
		pos.Left(),
		pos.Right(),
	} {

		if move.Y < 0 || move.X < 0 {
			continue
		}

		if move.Y >= len(*track) {
			continue
		}

		if move.X >= len((*track)[move.Y]) {
			continue
		}

		if (*track)[move.Y][move.X] == '.' || (*track)[move.Y][move.X] == 'E' {
			moves = append(moves, move)
			continue
		}
	}

	return moves
}

func Race(track *Track, start Position, end Position) (elapsed int) {

	rs := &Route{
		Position: start,
	}
	completed := []*Route{}

	queue := NewQueue()
	queue.Enqueue(rs)

	for !queue.Empty() {

		routes := []*Route{}
		for !queue.Empty() {
			route, err := queue.Dequeue()
			if err != nil {
				panic(err)
			}
			routes = append(routes, route)
		}

		for _, route := range routes {
			elapsed := route.Elapsed

			for i, move := range track.Next(route.Position) {

				if i == 0 {

					route.Elapsed++

					if (*track)[move.Y][move.X] == 'E' {
						route.Finished = true
						(*track)[route.Position.Y][route.Position.X] = 'o'
						route.Position = move
						completed = append(completed, route)
						break
					}

					(*track)[route.Position.Y][route.Position.X] = 'o'
					route.Position = move
					(*track)[route.Position.Y][route.Position.X] = 'x'
				}

				if i > 0 {

					forked := &Route{
						Position: move,
						Elapsed:  elapsed + 1,
					}
					queue.Enqueue(forked)
				}

				queue.Enqueue(route)
			}
		}

	}

	for _, r := range completed {
		// fmt.Println(r)
		if r.Finished {
			elapsed = r.Elapsed
		}
	}

	return elapsed
}

func (t *Track) GetCheats() (cheats []Position) {

	cheats = []Position{}
	for y, row := range *t {
		for x, val := range row {

			if val != '#' {
				continue
			}

			if y <= 0 || y >= len(*t)-1 {
				continue
			}

			if x <= 0 || x >= len((*t)[y])-1 {
				continue
			}

			pos := Position{X: x, Y: y}

			up := pos.Up()
			down := pos.Down()

			if (t.Value(up) == '.' || t.Value(up) == 'E' || t.Value(up) == 'S') &&
				(t.Value(down) == '.' || t.Value(down) == 'E' || t.Value(down) == 'S') {
				cheats = append(cheats, pos)
				continue
			}

			left := pos.Left()
			right := pos.Right()

			if (t.Value(left) == '.' || t.Value(left) == 'E' || t.Value(left) == 'S') &&
				(t.Value(right) == '.' || t.Value(right) == 'E' || t.Value(right) == 'S') {

				cheats = append(cheats, pos)
				continue
			}
		}
	}
	return cheats
}

func part1(input []string) (answer int) {

	track, start, end := NewTrack(input)

	normal := Race(track, start, end)

	track, _, _ = NewTrack(input)
	cheats := track.GetCheats()

	savings := map[int]int{}
	for _, cheat := range cheats {
		track, _, _ = NewTrack(input)
		(*track)[cheat.Y][cheat.X] = '.'
		elapsed := Race(track, start, end)

		saved := normal - elapsed
		savings[saved]++

		if saved >= 100 {
			answer++
		}

	}

	return answer
}

func (t *Track) Value(pos Position) (value rune) {
	return (*t)[pos.Y][pos.X]
}

func part2(input []string) (answer int) {

	return answer
}
