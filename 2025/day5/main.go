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

func NewFreshRange(input string) FreshRange {
	from := MustInt64(strings.Split(input, "-")[0])
	to := MustInt64(strings.Split(input, "-")[1])
	fr := FreshRange{
		From: from,
		To:   to,
	}
	return fr
}

type FreshRanges []FreshRange
type FreshRange struct {
	From int64
	To   int64
}

func (f *FreshRange) FreshCount() (count int64) {
	count = f.To - f.From + 1
	return count
}

func (f *FreshRanges) IsFresh(a int64) bool {
	for _, fr := range *f {
		if a >= fr.From && a <= fr.To {
			return true
		}
	}
	return false
}

func part1(inputs []string) (answer int) {

	freshes := FreshRanges{}
	availableStart := 0
	for i, ir := range inputs {
		if ir == "" {
			availableStart = i + 1
			break
		}
		freshes = append(freshes, NewFreshRange(ir))
	}

	for _, input := range inputs[availableStart:] {
		available := MustInt64(input)
		if freshes.IsFresh(available) {
			answer++
		}
	}

	return answer
}

func part2(inputs []string) (answer int64) {

	freshes := FreshRanges{}
	for _, ir := range inputs {
		if ir == "" {
			break
		}
		freshes = append(freshes, NewFreshRange(ir))
	}

	optimized := map[FreshRange]bool{}
	for _, fr := range freshes {
		optimal := FreshRange{From: fr.From, To: fr.To}
		for {
			from := optimal.From
			to := optimal.To
			for _, compare := range freshes {
				if compare.From <= optimal.From && compare.To >= optimal.From {
					optimal.From = compare.From
				}
				if compare.To >= optimal.To && compare.From <= optimal.To {
					optimal.To = compare.To
				}
			}
			if optimal.From == from && optimal.To == to {
				break
			}
		}
		optimized[optimal] = true
	}

	for o := range optimized {
		answer = answer + o.FreshCount()
	}

	return answer
}

func MustInt64(input string) int64 {
	val, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		panic("unable to convert")
	}
	return val
}
