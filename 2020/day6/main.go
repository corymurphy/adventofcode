package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Declarations []Declaration

type Declaration struct {
	answersInput []string
}

func NewDeclaration(answers []string) Declaration {
	return Declaration{answersInput: answers}
}

func (d Declaration) ParseAnswers() []string {
	parsedAnswers := []string{}

	for _, a := range d.answersInput {
		parsedAnswers = append(parsedAnswers, strings.Split(a, "")...)
	}

	return parsedAnswers
}

func (d Declaration) YesCount1() int {
	answers := make(map[string]bool)

	yesCount := 0
	for _, answer := range d.ParseAnswers() {
		if _, hit := answers[answer]; !hit {
			yesCount++
		}
		answers[answer] = true
	}

	return yesCount
}

func (d Declaration) YesCount2() int {

	answers := make(map[string]int)

	yesCount := 0

	for _, answer := range d.ParseAnswers() {
		if count, hit := answers[answer]; hit {
			answers[answer] = count + 1
		} else {
			answers[answer] = 1
		}
	}

	for _, count := range answers {
		if len(d.answersInput) == count {
			yesCount++
		}
	}

	return yesCount
}

func main() {
	part1()
}

func part1() {
	yesSum1, yesSum2 := YesSum(readLines("input"))
	fmt.Printf("yesSum1: %s; yesSum2: %s \n", fmt.Sprint(yesSum1), fmt.Sprint(yesSum2))
}

func YesSum(lines []string) (int, int) {

	streamStart := true
	startIndex := 0
	endIndex := 0

	yesSum1 := 0
	yesSum2 := 0

	for i, line := range lines {

		if streamStart {
			startIndex = i
		}

		if line == "" {
			endIndex = i
			yesSum1 = yesSum1 + NewDeclaration(lines[startIndex:endIndex]).YesCount1()
			yesSum2 = yesSum2 + NewDeclaration(lines[startIndex:endIndex]).YesCount2()
			streamStart = true
		} else {
			streamStart = false
		}

	}

	return yesSum1, yesSum2
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
