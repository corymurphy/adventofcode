package main

import (
	"fmt"
	"strconv"
	"strings"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input_sample2")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

type Calibration struct {
	Value      uint64
	Components []uint64
}

func ToUint64(input string) (result uint64) {

	result, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		panic(err)
	}

	return result
}

func GetCalibrations(input []string) (calibrations []Calibration) {
	for _, row := range input {

		components := []uint64{}
		for _, component := range strings.Split(strings.Split(row, ": ")[1], " ") {
			components = append(components, ToUint64(component))
		}

		calibrations = append(calibrations,
			Calibration{
				Value: ToUint64(strings.Split(row, ": ")[0]),

				Components: components,
			},
		)
	}
	return calibrations
}

func generateBinaryPermutations(a string, b string, n int) []string {
	result := []string{}

	var generate func(current string, length int)
	generate = func(current string, length int) {
		if length == 0 {
			result = append(result, current)
			return
		}

		generate(current+a, length-1)
		generate(current+b, length-1)
	}

	generate("", n)
	return result
}

func generateBinaryPermutations3(a string, b string, c string, n int) []string {
	result := []string{}

	var generate func(current string, length int)
	generate = func(current string, length int) {
		if length == 0 {
			result = append(result, current)
			return
		}

		generate(current+a, length-1)
		generate(current+b, length-1)
		generate(current+c, length-1)
	}

	generate("", n)
	return result
}

func part1(input []string) (answer uint64) {

	calibrations := GetCalibrations(input)

	for _, calibration := range calibrations {
		// fmt.Println(calibration.Value)
		// operators := GetPermutations(len(calibration.Components) - 1)

		perms := generateBinaryPermutations("*", "+", len(calibration.Components)-1)

		operators := [][]string{}

		for _, generated := range perms {
			ops := []string{}
			for _, op := range generated {
				ops = append(ops, string(op))
				// fmt.Print(string(op))
			}
			// fmt.Println()
			operators = append(operators, ops)
		}

		// fmt.Println(perms)

		for _, operator := range operators {

			var total uint64 = calibration.Components[0]

			// fmt.Print(calibration.Value, ":", " ", total)

			for i := 1; i < len(calibration.Components); i++ {
				if operator[i-1] == "*" {
					total = total * calibration.Components[i]

				}

				if operator[i-1] == "+" {
					total = total + calibration.Components[i]
				}
				// fmt.Print(" ", operator[i-1], " ", calibration.Components[i])
			}

			// fmt.Print(" ----> ", total)

			// fmt.Println()

			if total == calibration.Value {
				answer = answer + total
				break
			}
		}
	}
	return answer
}

func part2(input []string) (answer uint64) {

	calibrations := GetCalibrations(input)

	for _, calibration := range calibrations {
		// fmt.Println(calibration.Value)
		// operators := GetPermutations(len(calibration.Components) - 1)

		operators := [][]string{}

		// for _, generated := range generateBinaryPermutations3("*", "+", len(calibration.Components)-1) {
		// 	ops := []string{}
		// 	for _, op := range generated {
		// 		ops = append(ops, string(op))
		// 		fmt.Print(string(op))
		// 	}
		// 	fmt.Println()
		// 	operators = append(operators, ops)
		// }

		for _, generated := range generateBinaryPermutations3("||", "*", "+", len(calibration.Components)-1) {
			ops := []string{}
			for i := 0; i < len(generated); i++ {
				if generated[i] == '|' && generated[i+1] == '|' {
					ops = append(ops, "||")
					i++
				} else {
					ops = append(ops, string(generated[i]))
				}
			}
			// fmt.Println()
			operators = append(operators, ops)
		}

		// for _, generated := range generateBinaryPermutations("||", "+", len(calibration.Components)-1) {
		// 	ops := []string{}

		// 	for i := 0; i < len(generated); i++ {
		// 		if generated[i] == '|' && generated[i+1] == '|' {
		// 			ops = append(ops, "||")
		// 			i++
		// 		} else {
		// 			ops = append(ops, string(generated[i]))
		// 		}
		// 	}
		// 	operators = append(operators, ops)
		// }

		// fmt.Println(perms)

		// fmt.Println(operators)

		for _, operator := range operators {

			var total uint64 = calibration.Components[0]

			// fmt.Print(calibration.Value, ":", " ", total)

			for i := 1; i < len(calibration.Components); i++ {

				if operator[i-1] == "*" {
					total = total * calibration.Components[i]
					continue
				}

				if operator[i-1] == "+" {
					total = total + calibration.Components[i]
					continue
				}

				if operator[i-1] == "||" {

					// if t == 4 {
					// 	fmt.Println(total)
					// }

					total = ToUint64(strconv.FormatUint(total, 10) + strconv.FormatUint(calibration.Components[i], 10))

					// fmt.Println(total)
					// total = total + calibration.Components[i]
					// fmt.Println(total)
					continue

				}

				// panic("you can't be here bro")

				// fmt.Println(total)

				// fmt.Print(" ", operator[i-1], " ", calibration.Components[i])
			}

			// fmt.Print(" ----> ", total)

			// fmt.Println()

			if total == calibration.Value {
				// fmt.Println("here")
				answer = answer + total
				break
			}
		}
	}
	return answer
}
