package main

import (
	"fmt"
	"strconv"
	"strings"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)

}

func part1(inputs []string) (answer int64) {

	ranges := []Range{}
	for _, input := range strings.Split(inputs[0], ",") {

		from, err := strconv.ParseInt(strings.Split(input, "-")[0], 10, 64)
		if err != nil {
			panic(err)
		}
		to, err := strconv.ParseInt(strings.Split(input, "-")[1], 10, 64)
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, Range{

			From: from,
			To:   to,
		})
	}

	for _, r := range ranges {
		pids := ProductIds(r)
		for _, pid := range pids {
			if IsValidPID1(pid) {
				answer = answer + pid
			}
		}
	}

	return answer

}

func IsValidPID1(input int64) bool {

	pid := strconv.FormatInt(input, 10)

	if len(pid)%2 != 0 {
		return false
	}
	if pid[0:len(pid)/2] == pid[len(pid)/2:] {
		return true
	}

	return false
}

func IsValidPID2(input int64) bool {

	pid := strconv.FormatInt(input, 10)
	for i := 1; i <= len(pid)/2; i++ {

		if len(pid)%i != 0 {
			continue
		}

		test := pid[:i]
		valid := false
		for j := 0; j <= len(pid) && j+i <= len(pid); j = j + i {
			if test != pid[j:j+i] {
				valid = false
				break
			}
			if test == pid[j:j+i] {
				valid = true
			}
		}
		if valid {
			return valid
		}
	}
	return false
}

func ProductIds(r Range) []int64 {

	ids := []int64{}
	for i := r.From; i <= r.To; i++ {
		ids = append(ids, i)
	}
	return ids
}

func part2(inputs []string) (answer int64) {
	ranges := []Range{}
	for _, input := range strings.Split(inputs[0], ",") {

		from, err := strconv.ParseInt(strings.Split(input, "-")[0], 10, 64)
		if err != nil {
			panic(err)
		}
		to, err := strconv.ParseInt(strings.Split(input, "-")[1], 10, 64)
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, Range{

			From: from,
			To:   to,
		})
	}
	for _, r := range ranges {
		pids := ProductIds(r)
		for _, pid := range pids {
			if IsValidPID2(pid) {
				answer = answer + pid
			}
		}
	}
	return answer
}

type Range struct {
	From int64
	To   int64
}
