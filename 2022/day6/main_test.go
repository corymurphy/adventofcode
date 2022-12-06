package main

import (
	"testing"

	"github.com/corymurphy/adventofcode/shared"
)

func Test_Part1(t *testing.T) {
	shared.AssertEqual(t, 7, part1([]string{"mjqjpqmgbljsphdztnvjfqwrcgsmlb"}))
	shared.AssertEqual(t, 5, part1([]string{"bvwbjplbgvbhsrlpgdmjqwftvncz"}))
	shared.AssertEqual(t, 6, part1([]string{"nppdvjthqldpwncqszvftbrmjlhg"}))
	shared.AssertEqual(t, 10, part1([]string{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}))
	shared.AssertEqual(t, 11, part1([]string{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}))
}

func Test_Part2(t *testing.T) {
	shared.AssertEqual(t, 19, part2([]string{"mjqjpqmgbljsphdztnvjfqwrcgsmlb"}))
	shared.AssertEqual(t, 23, part2([]string{"bvwbjplbgvbhsrlpgdmjqwftvncz"}))
	shared.AssertEqual(t, 23, part2([]string{"nppdvjthqldpwncqszvftbrmjlhg"}))
	shared.AssertEqual(t, 29, part2([]string{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}))
	shared.AssertEqual(t, 26, part2([]string{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}))
}

func Test_Part1_Completed(t *testing.T) {
	expected := 1647

	input := shared.ReadInput("input")

	actual := part1(input)
	shared.AssertEqual(t, expected, actual)
}

func Test_Part2_Completed(t *testing.T) {
	expected := 2447

	input := shared.ReadInput("input")

	actual := part2(input)
	shared.AssertEqual(t, expected, actual)
}
