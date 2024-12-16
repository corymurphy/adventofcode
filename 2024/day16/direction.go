package main

type Direction int

const (
	Unknown Direction = iota
	Up
	Down
	Left
	Right
)

func (d Direction) TurnLeft() Direction {
	switch d {
	case Up:
		return Left
	case Down:
		return Right
	case Left:
		return Down
	case Right:
		return Up
	}
	panic("invalid direction")
}

func (d Direction) TurnRight() Direction {
	switch d {
	case Up:
		return Right
	case Down:
		return Left
	case Left:
		return Up
	case Right:
		return Down
	}
	panic("invalid direction")
}

func GetDirection(input rune) Direction {
	switch input {
	case '<':
		return Left
	case '>':
		return Right
	case '^':
		return Up
	case 'v':
		return Down
	}
	panic("invalid direction")
}

func (d Direction) Rune() rune {
	switch d {
	case Up:
		return '^'
	case Down:
		return 'v'
	case Left:
		return '<'
	case Right:
		return '>'
	}
	panic("invalid direction")
}

func (d Direction) String() string {
	switch d {
	case Up:
		return "up"
	case Down:
		return "down"
	case Left:
		return "left"
	case Right:
		return "right"
	}
	panic("invalid direction")
}

func (d Direction) Horizontal() bool {
	return d == Right || d == Left
}

func (d Direction) Vertical() bool {
	return d == Up || d == Down
}

func (d Direction) Reverse() Direction {
	switch d {
	case Unknown:
		return Unknown
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	}
	panic("invalid direction")
}
