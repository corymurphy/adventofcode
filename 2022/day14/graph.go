package main

type Graph [][]string

func (g *Graph) IsRock(input string) bool {
	return input == "#"
}

func (g *Graph) IsSand(input string) bool {
	return input == "o"
}

func (g *Graph) IsBlocked(v Vector) bool {
	return g.IsOutOfBounds(v) || g.IsSand((*g)[v.Y][v.X]) || g.IsRock((*g)[v.Y][v.X])
}

// TODO: i think this is too far
func (g *Graph) IsOutOfBounds(v Vector) bool {
	return v.Y < 0 || v.X < 0 || v.Y >= len(*g) || v.X >= len((*g)[v.Y])
}

func (g *Graph) SetSand(v Vector) {
	(*g)[v.Y][v.X] = "o"
}
