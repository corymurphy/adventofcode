package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strings"
	// "strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	part1 := part1(lines)
	part2 := part2(lines)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(lines []string) int32 {

	position := NewPosition()

	for _, line := range lines {

		command := NewCommand(line)
		fmt.Printf("%s %d\n", command.Direction, command.Units)

		position.MovePart1(command)
	}

	return position.LocationPart1()
}

func part2(input []string) int32 {
	position := NewPosition()

	for _, line := range input {

		command := NewCommand(line)
		fmt.Printf("%s %d\n", command.Direction, command.Units)

		position.MovePart2(command)
	}

	return position.LocationPart2()
}
