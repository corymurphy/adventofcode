package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func ReadInput(path string) []string {
	lines, err := readLines(path)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return lines
}

func toInt(input string) int {
	value, _ := strconv.ParseInt(input, 10, 32)
	return int(value)
}
