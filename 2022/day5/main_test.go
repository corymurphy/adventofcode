package main

import (
	"fmt"
	"testing"

	"github.com/corymurphy/adventofcode/shared"
)

func Test_Part1(t *testing.T) {
	expected := "CMZ"

	input := shared.ReadInput("input_test")

	actual := part1(input)
	if expected != actual {
		t.Errorf("expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}

func Test_Part2(t *testing.T) {
	expected := "MCD"
	input := shared.ReadInput("input_test")
	actual := part2(input)
	if expected != actual {
		t.Errorf("expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}

func Test_Part1_Completed(t *testing.T) {
	expected := "PSNRGBTFT"

	input := shared.ReadInput("input")

	actual := part1(input)
	if expected != actual {
		t.Errorf("expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}

func Test_Part2_Completed(t *testing.T) {
	expected := "BNTZFPMMW"

	input := shared.ReadInput("input")

	actual := part2(input)
	if expected != actual {
		t.Errorf("expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}
