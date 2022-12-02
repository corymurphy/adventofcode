package main

import (
	"fmt"
	"testing"
)

func Test_Part1(t *testing.T) {
	expected := 24000
	input := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
		"",
	}

	actual := mostCarriedCalories(input)
	if expected != actual {
		t.Errorf("expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}

func Test_Part2(t *testing.T) {
	expected := 45000
	input := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
		"",
	}
	actual := topThreeCarriedCalories(input)
	if expected != actual {
		t.Errorf("expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}
