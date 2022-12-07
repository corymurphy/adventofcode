package main

type NodeType int64

const (
	Directory NodeType = 0
	File      NodeType = 1
)

func (n NodeType) String() string {
	switch n {
	case File:
		return "file"
	case Directory:
		return "dir"
	}
	return "unknown"
}
