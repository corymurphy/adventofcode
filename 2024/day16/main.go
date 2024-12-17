package main

import (
	"fmt"
	"math"

	"github.com/corymurphy/adventofcode/shared"

	"github.com/google/uuid"
)

type Maze [][]rune

type Route struct {
	Position   Position
	ReachedEnd bool
	Score      int
	Visited    []Position
	Id         string
	SourceId   string
}

type Position struct {
	X         int
	Y         int
	Direction Direction
}

type Game struct {
	Maze    Maze
	Visited []Visited
	Routes  []*Route
	Start   Position
}

func main() {
	input := shared.ReadInput("input_sample")

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
	maze, start := NewMaze(input)
	return &Game{
		Maze:    maze,
		Start:   start,
		Visited: []Visited{{Position: start, Cost: 0}},
	}
}

func NewMaze(input []string) (maze Maze, start Position) {

	maze = [][]rune{}

	for y, row := range input {
		r := []rune{}
		for x, val := range row {
			r = append(r, val)

			if val == 'S' {
				start = Position{X: x, Y: y, Direction: Right}
			}
		}
		maze = append(maze, r)
	}
	return maze, start
}

type Visited struct {
	Position Position
	Cost     int
	// From     Direction
}

func (g *Game) IsMoveValid(move Position, cost int) bool {
	if move.Y < 0 || move.X < 0 {
		return false
	}

	if move.Y >= len(g.Maze) {
		return false
	}

	if move.X >= len((g.Maze)[move.Y]) {
		return false
	}

	// , contains := g.VisitedFrom[move]

	// if contains {
	// 	return false
	// }

	for _, v := range g.Visited {
		if v.Position.X == move.X && v.Position.Y == move.Y && v.Position.Direction == move.Direction && v.Cost <= cost {
			return false
		}
	}

	// if g.Maze[move.Y][move.X] == '#' {
	// 	return false
	// }

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

func (g *Game) ExploreBfs() {

	queue := NewQueue()

	for _, route := range g.Routes {
		queue.Enqueue(route)
	}

	for !queue.Empty() {

		route, err := queue.Dequeue()

		g.Visited = append(g.Visited, Visited{Position: route.Position, Cost: route.Score})

		route.Visited = append(route.Visited, route.Position)
		if err != nil {
			panic(err)
		}

		paths := 0

		current := Route{
			Score:    route.Score,
			Position: Position{X: route.Position.X, Y: route.Position.Y, Direction: route.Position.Direction},
			// Visited:  route.Visited,
			Id:       route.Id,
			SourceId: route.SourceId,
		}

		// fmt.Println(copy(current.Visited, route.Visited))

		next := g.Next(route.Position, route.Position.Direction)
		if g.IsMoveValid(next, current.Score) {
			paths++
			route.Score++

			// route.V[route.Position] = 0

			if g.Maze[next.Y][next.X] == 'E' {
				route.ReachedEnd = true
				// route.Visited = append(route.Visited, next)

				continue
			} else {

				route.Position = next
				queue.Enqueue(route)
			}
		}
		for _, next := range []Position{
			g.Next(current.Position, current.Position.Direction.TurnLeft()),
			g.Next(current.Position, current.Position.Direction.TurnRight()),
		} {

			if !g.IsMoveValid(next, current.Score) {
				continue
			}
			paths++

			if paths == 1 {
				route.Position.Direction = next.Direction
				route.Score += 1000
				queue.Enqueue(route)
			}

			if paths > 1 {

				forked := &Route{
					Position: current.Position,
					Score:    current.Score + 1000,
					// Visited:  route.Visited,
					Id:       uuid.New().String(),
					SourceId: current.Id,
					// Visited:
					// V:        current.V,
				}

				fmt.Println(len(route.Visited))

				// fmt.Println(copy(forked.Visited, route.Visited))

				// copy(forked.Visited, current.Visited)

				g.Routes = append(g.Routes, forked)

				forked.Position.Direction = next.Direction

				queue.Enqueue(forked)
			}
		}
	}
}

func (g *Game) Compete() {

	route := Route{Position: g.Start, Visited: []Position{g.Start}, Id: "1"}
	g.Routes = []*Route{&route}
	g.ExploreBfs()
}

func part1(input []string) (answer int) {

	game := NewGame(input)

	game.Compete()

	answer = math.MaxInt

	for _, route := range game.Routes {
		if route.ReachedEnd {
			if route.Score < answer {
				answer = route.Score
			}
		}
	}

	return answer
}

func (g *Game) Source(id string) Route {
	for _, source := range g.Routes {
		if source.Id == id {
			return *source
		}
	}
	return Route{Id: "0", Visited: []Position{}}
}

func part2(input []string) (answer int) {
	game := NewGame(input)

	game.Compete()

	cheapest := math.MaxInt

	for _, route := range game.Routes {
		if route.ReachedEnd {
			if route.Score < cheapest {
				cheapest = route.Score
			}
		}
	}

	var winning map[Position]int = make(map[Position]int)
	for _, route := range game.Routes {
		// fmt.Println(route.SourceId)
		if route.ReachedEnd && route.Score == cheapest {

			// fmt.Println(route.SourceId)

			// fmt.Println(len(route.Visited))

			for _, pos := range route.Visited {
				pos.Direction = Unknown
				winning[pos] = 0
			}

			// sid := route.SourceId

			// source := game.Source(route.SourceId)
			// sid := source.SourceId

			// for sid != "0" && sid != "" {

			// 	source = game.Source(sid)
			// 	sid = source.SourceId

			// 	for _, pos := range source.Visited {
			// 		pos.Direction = Unknown
			// 		winning[pos] = 0
			// 	}

			// }

			// for _, pos := range game.Source(route).Visited {
			// 	pos.Direction = Unknown
			// 	winning[pos] = 0
			// }

			// fmt.Println(game.Source(route))

			// for {

			// }

			// fmt.Println(route)
		}
	}

	for pos := range winning {
		game.Maze[pos.Y][pos.X] = 'O'
	}

	game.Maze.Print()

	answer = len(winning)

	return answer
}
