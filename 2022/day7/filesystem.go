package main

import (
	"math"
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

type FileSystem struct {
	Root *Node
	Cwd  *Node
}

func NewFileSystem(input []string) *FileSystem {
	fs := &FileSystem{}
	fs.init(input)
	return fs
}

func newRepeatingString(length int, value string) string {

	result := ""
	for i := 0; i <= length; i++ {
		result = result + value
	}
	return result
}

func (f *FileSystem) Print() {
	f.ChangeDirectory([]string{".."})
	f.Root.Print()
}

func (f *FileSystem) TotalSize() int {
	return f.Root.Size()
}

// func (f *FileSystem) Print2() {
// 	f.ChangeDirectory([]string{".."})
// 	f.Root.Print
// }

func (f *FileSystem) SumDirectories(sizeAtMost int) int {
	return SumDirectories(*f.Root, sizeAtMost, 0)
}

func (f *FileSystem) ProcessCommand(line string, input []string, index int) int {
	args := strings.Split(line, " ")
	command := args[1]
	switch command {
	case ChangeDirectory.String():
		f.ChangeDirectory(args[2:])
		return 0
	case ListDirectory.String():
		return f.ListDirectory(args[2:], input, index)
	}
	return 0
}

func (f *FileSystem) GetCleanableDirectory(spaceNeeded int) int {
	return GetCleanableDirectory(*f.Root, spaceNeeded, math.MaxInt)
}

func (f *FileSystem) ChangeDirectory(args []string) {

	directory := args[0]

	if f.Cwd == nil && directory == "/" {
		node := &Node{
			Name:     directory,
			Type:     Directory,
			Children: map[string]*Node{},
			Parent:   map[string]*Node{},
			isRoot:   true,
			Depth:    0,
		}

		f.Cwd = node
		f.Root = node
		f.Cwd.Parent[".."] = node
		return
	}

	if directory == ".." {
		f.Cwd = f.Cwd.Parent[".."]
		return
	}

	f.Cwd = f.Cwd.Children[directory]
}

func getStdout(input []string, index int) ([]string, int) {

	stdout := []string{}
	advance := 0
	for i := index; i < len(input); i++ {
		advance++
		if isCommand(input[i]) || input[i] == "" {
			return stdout, advance - 1
		} else {
			stdout = append(stdout, input[i])
		}
	}
	return stdout, advance
}

func (f *FileSystem) ListDirectory(args []string, input []string, index int) int {
	result, advance := getStdout(input, index+1)

	for _, line := range result {
		name := strings.Split(line, " ")[1]
		if strings.HasPrefix(line, "dir") {
			f.Cwd.Children[name] = &Node{
				Children: map[string]*Node{},
				Parent:   map[string]*Node{"..": f.Cwd},
				Name:     name,
				Type:     Directory,
				Depth:    f.Cwd.Depth + 2,
			}

		} else {
			size := shared.ToInt(strings.Split(line, " ")[0])
			f.Cwd.Children[name] = &Node{
				Children: map[string]*Node{},
				Parent:   map[string]*Node{"..": f.Cwd},
				Name:     name,
				Type:     File,
				size:     size,
				Depth:    f.Cwd.Depth + 2,
			}
		}
	}

	return advance
}

func (f *FileSystem) init(input []string) error {

	for i := 0; i < len(input); i++ {
		line := input[i]
		if isCommand(line) {
			advanceIndex := f.ProcessCommand(line, input, i)
			i = i + advanceIndex
		}
	}
	return nil
}

func isCommand(input string) bool {
	return strings.HasPrefix(input, "$")
}
