package main

import (
	"2021/shared"
	"testing"
)

func TestPart1(t *testing.T) {

	input := shared.ReadInput("input")
	expected := 0
	actual := part1(input)

	shared.AssertEqual(t, expected, actual)
}
