package main

import (
	"fmt"
	"math"

	"github.com/corymurphy/adventofcode/shared"
)

type Maze [][]rune
type Reindeer struct {
	Start Position
}

var debug bool = false

type Route struct {
	// Start      Position
	Position   Position
	ReachedEnd bool
	// Moves      int
	Score    int
	Id       int
	SourceId int
}

type Position struct {
	X         int
	Y         int
	Direction Direction
}

type Game struct {
	Reindeer     Reindeer
	Maze         Maze
	Visited      []Position
	Routes       []*Route
	ActiveRoutes int
}

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func (m Maze) Print() {
	fmt.Println()
	for _, row := range m {
		for _, val := range row {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
}

func NewGame(input []string) (game *Game) {
	maze, reindeer := NewMaze(input)
	return &Game{
		Maze:     maze,
		Reindeer: reindeer,
		Visited:  []Position{reindeer.Start},
	}
}

func NewMaze(input []string) (maze Maze, reindeer Reindeer) {

	maze = [][]rune{}
	reindeer = Reindeer{}

	for y, row := range input {
		r := []rune{}
		for x, val := range row {
			r = append(r, val)

			if val == 'S' {
				reindeer.Start = Position{X: x, Y: y, Direction: Right}
			}
		}
		maze = append(maze, r)
	}
	return maze, reindeer
}

func (g *Game) IsMoveValid(move Position) bool {
	if move.Y < 0 || move.X < 0 {
		return false
	}

	if move.Y >= len(g.Maze) {
		return false
	}

	if move.X >= len((g.Maze)[move.Y]) {
		return false
	}

	for _, v := range g.Visited {
		if v.X == move.X && v.Y == move.Y {
			return false
		}
	}

	if g.Maze[move.Y][move.X] == '.' || g.Maze[move.Y][move.X] == 'E' {
		return true
	}

	return false
}

func (g *Game) Next(pos Position, dir Direction) (next Position) {

	next = Position{Direction: dir, X: pos.X, Y: pos.Y}
	switch pos.Direction {
	case Up:
		switch dir {
		case Up:
			next.Y--
			return next
		case Left:
			next.X--
			return next
		case Right:
			next.X++
			return next
		}
		panic("invalid")
	case Down:
		switch dir {
		case Left:
			next.X++
			next.Direction = Right
			return next
		case Right:
			next.X--
			next.Direction = Left
			return next
		case Down:
			next.Y++
			return next
		}
		panic("invalid")
	case Left:
		switch dir {
		case Left:
			next.X--
			return next
		case Down:
			next.Y++
			return next
		case Up:
			next.Y--
			return next
		}
		panic("invalid")
	case Right:
		switch dir {
		case Right:
			next.X++
			return next
		case Down:
			next.Y++
			return next
		case Up:
			next.Y--
			return next
		}
		panic("invalid")
	}
	panic("invalid")
}

func (r *Route) ScoreMove(dir Direction) {
	if r.Position.Direction == dir {
		r.Score++
	} else {
		r.Score++
		r.Score += 1000
	}
}

func (g *Game) ExploreBfs() {

	queue := NewQueue()

	for _, route := range g.Routes {
		queue.Enqueue(route)
	}

	for !queue.Empty() {
		route, err := queue.Dequeue()

		if err != nil {
			panic(err) // shouldn't happen
		}

		valid := 0

		score := route.Score

		for _, next := range []Position{
			g.Next(route.Position, route.Position.Direction),
			g.Next(route.Position, route.Position.Direction.TurnLeft()),
			g.Next(route.Position, route.Position.Direction.TurnRight()),
		} {

			if !g.IsMoveValid(next) {
				continue
			}
			valid++

			if valid == 1 {

				route.ScoreMove(next.Direction)

				g.Maze[route.Position.Y][route.Position.X] = route.Position.Direction.Rune()

				if g.Maze[next.Y][next.X] == 'E' {
					route.ReachedEnd = true
					continue
				}
				route.Position = next
				queue.Enqueue(route)
			}

			if valid > 1 {
				forked := Route{
					Position: route.Position,
					Score:    score,
					Id:       len(g.Routes) + 1,
					SourceId: route.Id,
				}
				forked.ScoreMove(next.Direction)

				g.Routes = append(g.Routes, &forked)

				g.Maze[forked.Position.Y][forked.Position.X] = forked.Position.Direction.Rune()

				if g.Maze[next.Y][next.X] == 'E' {
					forked.ReachedEnd = true
					continue
				}
				forked.Position = next
				queue.Enqueue(&forked)
			}
		}
	}
}

func (g *Game) Compete() {

	route := Route{Position: g.Reindeer.Start, Id: 1}
	g.Routes = []*Route{&route}
	g.ExploreBfs()
}

func part1(input []string) (answer int) {

	game := NewGame(input)

	game.Compete()

	answer = math.MaxInt

	for _, route := range game.Routes {
		if route.ReachedEnd {
			fmt.Println(route)
			if route.Score < answer {
				answer = route.Score
			}
		}
	}

	game.Maze.Print()

	return answer
}

func part2(input []string) (answer int) {

	return answer
}
