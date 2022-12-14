package main

import (
	"fmt"

	"github.com/corymurphy/adventofcode/shared"
)

type PacketsPair struct {
	raw   []string
	setA  string
	setB  string
	index int
}

func NewPacketsPair(input []string, index int) *PacketsPair {
	return &PacketsPair{
		raw:   input,
		setA:  input[0],
		setB:  input[1],
		index: index,
	}
}

func (p *PacketsPair) longestSetLength() int {
	if len(p.setA) > len(p.setB) {
		return len(p.setA)
	} else {
		return len(p.setB)
	}
}

func isSequenceStart(input rune) bool {

	if input == '[' {
		return true
	}
	return false
}

func isSequenceEnd(input rune) bool {
	if input == ']' {
		return true
	}
	return false
}

func isInt(input rune) bool {
	ints := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}

	for _, i := range ints {
		if i == input {
			return true
		}
	}
	return false
}

func (p *PacketsPair) InOrder() bool {
	inOrder := false

	// fmt.Println(p.longestSetLength())

	a := p.setA[1 : len(p.setA)-1]
	b := p.setB[1 : len(p.setB)-1]

	left := getSequences(a)
	right := getSequences(b)

	fmt.Println(left)
	fmt.Println(right)

	inOrder = compare(left, right)
	// aSequences := []interface{}

	// fmt.Println(a)
	// fmt.Println(b)

	// aSequence := NewSequence()
	// bSequence := NewSequence()

	// parseSequence(aSequence, a)
	// parseSequence(bSequence, b)
	// for i := 1; i < p.longestSetLength(); i++ {

	// 	// parseSequence()
	// 	if i > len(a) {

	// 	}

	// 	// fmt.Println(p.setA[i : p.longestSetLength()-1])

	// }

	return inOrder
}

// func ContainsSequence
func compare(left []string, right []string) bool {
	inOrder := true

	for i, l := range left {

		if i > len(right) {
			return false
		}

		r := right[i]

		// if both items are ints
		if areInts(l, r) {
			if shared.ToInt(l) > shared.ToInt(r) {
				return false
			}
		}

		if areMixed(l, r) {
			if compareMixed(l, r) {
				return false
			}
		}

		// if mixed

		// if i > len(right) {
		// 	return false
		// }
	}
	return inOrder
}

func areInts(l string, r string) bool {
	return len(l) == 1 && isInt(rune(l[0])) && len(r) == 1 && isInt(rune(r[1]))
}

func ContainsSequence(input string) bool {
	for _, val := range input {
		if isSequenceEnd(val) || isSequenceStart(val) {
			return true
		}
	}
	return false
}

func areMixed(l string, r string) bool {
	if len(l) == 1 && isInt(rune(l[0])) && ContainsSequence(r) {
		return true
	}

	if len(r) == 1 && isInt(rune(r[1])) && ContainsSequence(l) {
		return true
	}
	return false
}

func compareMixed(l string, r string) bool {
	if len(l) == 1 && isInt(rune(l[0])) {
		rSeq := getSequences(r)

		if shared.ToInt(l) > rSeq[0] {
			return false
		}
	}

	if len(r) == 1 && isInt(rune(r[1])) {
		lSeq := getSequences(l)

		if 
	}
}

func getSequences(input string) []string {
	// seqCount := 0

	sequences := []string{}

	for i := 0; i < len(input); i++ {
		char := rune(input[i])
		// fmt.Println(i)
		if isSequenceStart(char) {
			end := getSequenceEnd(input, i)
			seq := input[i:end]
			sequences = append(sequences, seq)
			i = end
		}
		if isInt(char) {
			sequences = append(sequences, string(char))
		}
	}
	// for i, char := range input {

	// }
	return sequences
}

func getSequenceEnd(input string, start int) int {
	seqCount := 1
	for i, char := range input[start:] {
		if isSequenceStart(char) {
			seqCount++
		}

		if isSequenceEnd(char) {
			seqCount--
		}

		// we've found the end
		if seqCount < 1 {
			return i
		}
	}
	return len(input)
}

// func parseSequence(seq *Sequence, input string) {

// 	// seqCount := 0
// 	hasChildren := false
// 	children := []*Sequence{}
// 	for i, char := range input {

// 		if isInt(char) {
// 			hasChildren = true
// 			(seq.children) = append(seq.children, &Sequence{
// 				value: []int{},
// 			})
// 		}
// 		// if isSequenceStart(char) {
// 		// 	end := getSequenceEnd(input, i)
// 		// 	parseSequence()
// 		// 	// seqCount++
// 		// }
// 	}

// 	if hasChildren {
// 		seq.vertex = false
// 	}

// }

func NewSequence() *Sequence {
	return &Sequence{
		vertex:   true,
		value:    []int{},
		children: []*Sequence{},
	}
}

type Sequence struct {
	vertex   bool
	value    []int
	children []*Sequence
}

// func parseSequence(input []string)
