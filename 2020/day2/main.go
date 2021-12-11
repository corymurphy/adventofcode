package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	valid := 0
	for _, line := range lines {
		p := Parse1(line)
		if p.Valid() {
			valid++
		}
	}
	fmt.Printf("part 1: %s\n", fmt.Sprint(valid))
}

func part2() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	valid := 0
	for _, line := range lines {
		p := Parse2(line)
		if p.Valid() {
			valid++
		}
	}
	fmt.Printf("part 2: %s\n", fmt.Sprint(valid))
}

func main() {
	part1()
	part2()
}

type Password1 struct {
	Max      int64
	Min      int64
	Required string
	Value    string
}

type Password2 struct {
	IndexA   int64
	IndexB   int64
	Required string
	Value    string
}

func (p Password1) ToString() string {
	return fmt.Sprintf("min: %s; max: %s; req: %s;val: %s; occ: %s", fmt.Sprint(p.Min), fmt.Sprint(p.Max), p.Required, p.Value, fmt.Sprint(p.Occurrences()))
}

func (p Password1) Valid() bool {
	if p.Occurrences() >= p.Min && p.Occurrences() <= p.Max {
		return true
	} else {
		return false
	}
}

func (p Password2) Valid() bool {

	if string(p.Value[p.IndexA]) == p.Required && string(p.Value[p.IndexB]) != p.Required {
		return true
	}

	if string(p.Value[p.IndexA]) != p.Required && string(p.Value[p.IndexB]) == p.Required {
		return true
	}

	return false

}

func (p Password1) Occurrences() int64 {
	return int64(strings.Count(p.Value, p.Required))
}

func Parse1(input string) Password1 {

	inputSplit := strings.Split(input, " ")
	validRange := strings.Split(inputSplit[0], "-")

	return Password1{
		Max:      parseInt(validRange[1]),
		Min:      parseInt(validRange[0]),
		Required: strings.Replace(inputSplit[1], ":", "", -1),
		Value:    inputSplit[2],
	}
}

func Parse2(input string) Password2 {

	inputSplit := strings.Split(input, " ")
	validRange := strings.Split(inputSplit[0], "-")

	return Password2{
		IndexA:   parseInt(validRange[0]) - 1,
		IndexB:   parseInt(validRange[1]) - 1,
		Required: strings.Replace(inputSplit[1], ":", "", -1),
		Value:    inputSplit[2],
	}
}

func parseInt(num string) int64 {
	if val, err := strconv.ParseInt(num, 10, 64); err != nil {
		log.Fatalf("error parsing int: %s", err)
		return 0
	} else {
		return val
	}
}

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
