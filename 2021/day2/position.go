package main

type Position struct {
	Horizontal int32
	Depth      int32
	Aim		   int32
}

func NewPosition() Position {
	return Position{
		Horizontal: 0,
		Depth:      0,
		Aim: 0,
	}
}

func (p *Position) MovePart1(command Command) {
	switch command.Direction {
	case "forward":
		p.Horizontal = p.Horizontal + command.Units
	case "up":
		p.Depth = p.Depth - command.Units
	case "down":
		p.Depth = p.Depth + command.Units
	}
}

func (p *Position) MovePart2(command Command) {
	switch command.Direction {
	case "forward":
		p.Horizontal = p.Horizontal + command.Units
		p.Depth = p.Depth + (p.Aim * command.Units)
	case "up":
		p.Aim = p.Aim - command.Units
	case "down":
		p.Aim = p.Aim + command.Units
	}
}

func (p *Position) LocationPart1() int32 {
	return p.Horizontal * p.Depth
}

func (p *Position) LocationPart2() int32 {
	return p.Horizontal * p.Depth
}
