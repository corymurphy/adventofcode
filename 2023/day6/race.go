package main

import (
	"strings"
)

type Race struct {
	Time           int64
	RecordDistance int64
}

func ParseRaces(input []string) *[]Race {
	races := []Race{}

	times := strings.Fields(input[0])
	records := strings.Fields(input[1])

	for i := 1; i < len(times); i++ {
		race := Race{
			Time:           ToInt(times[i]),
			RecordDistance: ToInt(records[i]),
		}
		races = append(races, race)
	}

	return &races
}
