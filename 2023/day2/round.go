package main

import (
	"strings"

	"github.com/corymurphy/adventofcode/shared"
)

type Round struct {
	revealed map[string]int
}

var cubes []string = []string{
	"red",
	"blue",
	"green",
}

func getCubeCount(roundData string, cubeIndex int) int {

	start := 0

	for i := cubeIndex - 2; i > -1 && string(roundData[i]) != " "; i-- {
		start = i
	}
	return shared.ToInt(string(roundData[start : cubeIndex-1]))
}

func ParseRounds(roundsData string) *[]Round {
	rounds := []Round{}

	for _, roundData := range strings.Split(roundsData, "; ") {
		round := Round{
			revealed: map[string]int{},
		}

		for _, cube := range cubes {
			cubeIndex := strings.Index(roundData, cube)

			if cubeIndex < 0 {
				continue
			}
			round.revealed[cube] = getCubeCount(roundData, cubeIndex)
		}

		rounds = append(rounds, round)
	}

	return &rounds
}
