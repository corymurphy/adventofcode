package main

type Operation int

const (
	Unknown  Operation = 0
	Add      Operation = 1
	Multiply Operation = 2
	Square   Operation = 3
)

func (o Operation) String() string {
	switch o {
	case Add:
		return "+"
	case Multiply:
		return "*"
	case Square:
		return "^"
	default:
		return "Unknown"
	}
}

func ParseOperation(input string) Operation {
	switch input {
	case "+":
		return Add
	case "*":
		return Multiply
	default:
		return Unknown
	}
}
