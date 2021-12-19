package main

import (
	"2021/shared"
	"fmt"
	"strings"
	"sort"
	"strconv"
)

func main() {

	// input := shared.ReadInput("day8/input")
	input := shared.ReadInput("day8/example")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

// employeeSalary := map[string]int{
// 	"John": 1000
// 	"Sam": 2000
// 	}

//  size -> output value
func lengthEncodings() map[int]int {
	return map[int]int{
		2: 1,
		3: 7,
		4: 4,
		7: 8,
	}
}

func valueEncodings() map[string]string {
	return map[string]string{
		"bcdef": "5",
		"acdfg": "2",
		"abcdf": "3",
		"abcdef": "9",
		"bcdefg": "6",
		"abcdeg": "0",
	}
}

// s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
// sort.Strings(s)
// fmt.Println(s)

func stringToCharArray(intput string) []string {
	// str := "ab£"
    // chars := []rune(intput)
	chars :=  []string{}
    for i := 0; i < len(intput); i++ {
		chars = append(chars, string(intput[i]))
        // char := string(intput[i])
    }
	return chars
}

func charArrayToString(input []string) string {
	output := ""
	for _, val := range input {
		output = output + val
	}
	return output
}

func sortString(input string) string {
	inputArray := stringToCharArray(input)
	sort.Strings(inputArray)
	return charArrayToString(inputArray)
}

// func hasValidOutput(encodings map[int]int, output string) bool {

// }

func incrementCount(validSignalPatterns map[int]int, index int) map[int]int {

	// fmt.Println(index)
	if val, ok := validSignalPatterns[index]; ok {
		validSignalPatterns[index] = val + 1
	} else {
		validSignalPatterns[index] = 1
	}
	return validSignalPatterns
}


func part1(input []string) int {

	validSignalPatterns := map[int]int{}
	encodings := lengthEncodings()
	outputs := 0

	for _, row := range input {
		signals := strings.Split(row, " | ")
		signalPatterns := signals[0] 
		outputValues := signals[1]

		for _, pattern := range strings.Split(signalPatterns, " ") {
			
			if val, ok := encodings[len(pattern)]; ok {
				// fmt.Printf("%s %d %d\n", pattern, len(pattern), val)
				validSignalPatterns = incrementCount(validSignalPatterns, val)
			}
		}

		for _, outputSignal := range strings.Split(outputValues, " ") {
			
			outputValue := encodings[len(outputSignal)]
			if _, ok := validSignalPatterns[outputValue];  ok {
				// fmt.Printf("val: %d; index: %d\n", val, len(output))
				outputs++
			}
		}
		
	}


	return outputs
}

func part2(input []string) int {
	// validSignalPatterns := map[int]int{}
	// encodings := lengthEncodings()
	// outputs := 0
	output := 0

	for _, row := range input {
		signals := strings.Split(row, " | ")
		// signalPatterns := signals[0] 
		outputValues := signals[1]

		rowOutput := ""

		for _, outputSignal := range strings.Split(outputValues, " ") {


			if val, ok := lengthEncodings()[len(outputSignal)]; ok {
				// fmt.Println(string(val))
				rowOutput = rowOutput + strconv.Itoa(val)
			}

			if val, ok := valueEncodings()[sortString(outputSignal)]; ok {
				rowOutput = rowOutput + val
			}


		}

		fmt.Println(rowOutput)

		// for _, pattern := range strings.Split(signalPatterns, " ") {
			
		// 	if val, ok := encodings[len(pattern)]; ok {
		// 		// fmt.Printf("%s %d %d\n", pattern, len(pattern), val)
		// 		validSignalPatterns = incrementCount(validSignalPatterns, val)
		// 	}
		// }

		// for _, outputSignal := range strings.Split(outputValues, " ") {
			
		// 	outputValue := encodings[len(outputSignal)]
		// 	if _, ok := validSignalPatterns[outputValue];  ok {
		// 		// fmt.Printf("val: %d; index: %d\n", val, len(output))
		// 		outputs++
		// 	}
		// }
		
	}

	return output
}

