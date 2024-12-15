package main

import (
	"fmt"
	"math"
	"strings"

	shared "github.com/corymurphy/adventofcode/shared"
	"gonum.org/v1/gonum/mat"
)

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func isRound(val float64) bool {
	return math.Trunc(val) == val
}

func part1(input []string) (answer int) {

	for i := 0; i < len(input); i = i + 4 {
		a := map[string]float64{
			"X": float64(shared.ToInt(strings.TrimRight(strings.Split(strings.Split(input[i], " ")[2], "+")[1], ","))),
			"Y": float64(shared.ToInt(strings.Split(strings.Split(input[i], " ")[3], "+")[1])),
		}
		b := map[string]float64{
			"X": float64(shared.ToInt(strings.TrimRight(strings.Split(strings.Split(input[i+1], " ")[2], "+")[1], ","))),
			"Y": float64(shared.ToInt(strings.Split(strings.Split(input[i+1], " ")[3], "+")[1])),
		}

		prize := map[string]float64{
			"X": float64(shared.ToInt(strings.TrimRight(strings.Split(strings.Split(input[i+2], " ")[1], "=")[1], ","))),
			"Y": float64(shared.ToInt(strings.Split(strings.Split(input[i+2], " ")[2], "=")[1])),
		}

		matrix := mat.NewDense(2, 2, []float64{a["X"], b["X"], a["Y"], b["Y"]})
		vector := mat.NewVecDense(2, []float64{prize["X"], prize["Y"]})

		var x mat.VecDense
		if err := x.SolveVec(matrix, vector); err != nil {
			fmt.Println(err)
			return
		}
		aPressed := roundFloat(x.RawVector().Data[0], 2)
		bPressed := roundFloat(x.RawVector().Data[1], 2)

		if a["X"]*aPressed+b["X"]*bPressed != prize["X"] {
			continue
		}

		if a["Y"]*aPressed+b["Y"]*bPressed != prize["Y"] {
			continue
		}

		if isRound(aPressed) && isRound(bPressed) {
			answer = answer + (int(3*aPressed) + int(bPressed))
		}
	}

	return answer
}

func part2(input []string) (answer int) {

	for i := 0; i < len(input); i = i + 4 {
		a := map[string]float64{
			"X": float64(shared.ToInt(strings.TrimRight(strings.Split(strings.Split(input[i], " ")[2], "+")[1], ","))),
			"Y": float64(shared.ToInt(strings.Split(strings.Split(input[i], " ")[3], "+")[1])),
		}
		b := map[string]float64{
			"X": float64(shared.ToInt(strings.TrimRight(strings.Split(strings.Split(input[i+1], " ")[2], "+")[1], ","))),
			"Y": float64(shared.ToInt(strings.Split(strings.Split(input[i+1], " ")[3], "+")[1])),
		}

		prize := map[string]float64{
			"X": float64(shared.ToInt(strings.TrimRight(strings.Split(strings.Split(input[i+2], " ")[1], "=")[1], ","))),
			"Y": float64(shared.ToInt(strings.Split(strings.Split(input[i+2], " ")[2], "=")[1])),
		}

		matrix := mat.NewDense(2, 2, []float64{a["X"], b["X"], a["Y"], b["Y"]})
		vector := mat.NewVecDense(2, []float64{10000000000000 + prize["X"], 10000000000000 + prize["Y"]})

		var x mat.VecDense
		if err := x.SolveVec(matrix, vector); err != nil {
			fmt.Println(err)
			return
		}
		aPressed := roundFloat(x.RawVector().Data[0], 2)
		bPressed := roundFloat(x.RawVector().Data[1], 2)

		if a["X"]*aPressed+b["X"]*bPressed != 10000000000000+prize["X"] {
			continue
		}

		if a["Y"]*aPressed+b["Y"]*bPressed != 10000000000000+prize["Y"] {
			continue
		}

		if isRound(aPressed) && isRound(bPressed) {
			answer = answer + (int(3*aPressed) + int(bPressed))
		}
	}

	return answer
}
