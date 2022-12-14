package main

import "strings"

type RockStructure struct {
	Vectors []*Vector
}

func NewRockStructure(input string) *RockStructure {
	rawVectors := strings.Split(input, " -> ")
	vectors := []*Vector{}
	for _, v := range rawVectors {
		vectors = append(vectors, ParseNewVector(v))
	}
	return &RockStructure{
		Vectors: vectors,
	}
}
