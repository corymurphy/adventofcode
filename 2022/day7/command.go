package main

type Command int

const (
	ChangeDirectory Command = 0
	ListDirectory   Command = 1
)

func (c Command) String() string {
	switch c {
	case ChangeDirectory:
		return "cd"
	case ListDirectory:
		return "ls"
	}
	return "unknown"
}
