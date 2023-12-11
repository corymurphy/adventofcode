package main

import (
	"fmt"
	"strconv"
	"strings"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func part1(input []string) int {

	var calibration int = 0

	for _, calibrationKey := range input {

		calibrationKeyValue := []int{}
		for _, rawElement := range calibrationKey {
			element, err := strconv.ParseInt(string(rawElement), 10, 32)

			if err != nil {
				continue
			}
			calibrationKeyValue = append(calibrationKeyValue, int(element))

		}

		// yikes, but we're behind, send it
		calibration = calibration + shared.ToInt(fmt.Sprint(calibrationKeyValue[0])+fmt.Sprint(calibrationKeyValue[len(calibrationKeyValue)-1]))
	}
	return calibration
}

func part2(input []string) int {

	alphabetized := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	numerical := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var calibration int = 0

	for _, calibrationPoint := range input {

		calibrationPointValue := map[int]string{}

		for _, alpha := range alphabetized {
			setOccurrences(calibrationPointValue, calibrationPoint, 0, alpha)
		}

		keys := []int{}
		for key := range calibrationPointValue {
			keys = append(keys, key)
		}

		// fmt.Println(keys)
		max := getMax(keys)

		values := make([]int, max+1)

		for key, value := range calibrationPointValue {
			values[key] = numerical[value]
		}

		// fmt.Println(calibrationPointValue)
		// fmt.Println(values)

		first := getFirst(values)
		last := getLast(values)

		calculatedCalibrationPoint := shared.ToInt(fmt.Sprint(first) + fmt.Sprint(last))
		calibration = calibration + calculatedCalibrationPoint
		// fmt.Println(calculatedCalibrationPoint)
		// fmt.Println()
	}
	return calibration
}

func getLast(items []int) int {
	for i := len(items) - 1; i >= 0; i-- {
		if items[i] > 0 {
			return items[i]
		}
	}
	return 0
}

func getFirst(items []int) int {
	for _, value := range items {
		if value > 0 {
			return value
		}
	}
	return 0
}

func getMax(items []int) int {
	max := 0
	for _, item := range items {
		if item > max {
			max = item
		}
	}
	return max
}

func setOccurrences(calibrationPointValue map[int]string, calibrationPoint string, start int, alpha string) {

	for i := 0; i <= strings.Count(calibrationPoint, alpha); i++ {
		loc := strings.Index(calibrationPoint[start:], alpha)

		if loc < 0 {
			return
		}

		calibrationPointValue[start+loc] = alpha
		start = start + loc + 1
	}
}
