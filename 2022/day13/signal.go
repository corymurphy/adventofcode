package main

import "fmt"

type Signal []*PacketsPair

func NewSignal(input []string) *Signal {

	signal := Signal{}
	pairIndex := 1
	for i := 0; i < len(input)-1; i = i + 3 {
		pairIndex++
		pair := NewPacketsPair(input[i:i+2], pairIndex)
		signal = append(signal, pair)
	}

	return &signal
}

func (s *Signal) Add(pair *PacketsPair) {
	// (*s).append
}

func (s *Signal) Print() {
	for _, pair := range *s {
		fmt.Println(pair.setA)
		fmt.Println(pair.setB)
		fmt.Println("--------------")
	}
}
