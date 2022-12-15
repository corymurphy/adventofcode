package main

type Sand struct {
	start   *Vector
	finish  *Vector
	current *Vector
	graph   *Graph
}

func NewSand(start *Vector, graph *Graph) *Sand {
	return &Sand{
		start:   start,
		current: NewVector(start.X, start.Y),
		graph:   graph,
	}
}

func (s *Sand) Drop() {

	finished := false

	for !finished {

		// if bottom of graph

		if s.current.Y >= len(*s.graph) {
			finished = true
		}

		// try down
		if !s.graph.IsBlocked(*NewVector(s.current.X, s.current.Y+1)) {
			s.current.Y = s.current.Y + 1
			continue
		}

		// try left
		// if !s.graph.IsBlocked(*NewVector(s.current.X-1, s.current.Y)) &&
		if !s.graph.IsBlocked(*NewVector(s.current.X-1, s.current.Y+1)) {
			s.current.Y = s.current.Y + 1
			s.current.X = s.current.X - 1
			continue
		}

		// try right
		// if !s.graph.IsBlocked(*NewVector(s.current.X+1, s.current.Y)) &&
		if !s.graph.IsBlocked(*NewVector(s.current.X+1, s.current.Y+1)) {
			s.current.Y = s.current.Y + 1
			s.current.X = s.current.X + 1
			continue
		}
		finished = true
	}

	s.finish = s.current
	s.graph.SetSand(*s.finish)
}
