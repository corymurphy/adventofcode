package main

import (
	"fmt"
	"strings"
	"strconv"
)

const DAYS = 80

type Simulator struct {
	LanternFish []LanternFish
	Day         int
	InitialInput []string
}

func NewSimulator(input []string) Simulator {


	var allLanternfish []LanternFish
	for _, ageStr := range strings.Split(input[0], ",") {
		lanternfish := NewLanternFish(toInt(ageStr))
		allLanternfish = append(allLanternfish, lanternfish)
	}

	return Simulator{
		LanternFish: allLanternfish,
		Day:         0,
		InitialInput: input,
	}
}

func (s *Simulator) LanternFishCount() int {
	return len(s.LanternFish)
}

func (s *Simulator) AgeFish() {
	spawned := 0
	for i := 0; i < len(s.LanternFish); i++ {
		fish := s.LanternFish[i]
		spawn := fish.IncrementAge()
		s.LanternFish[i] = fish
		if spawn {
			spawned++
		}
	}

	for i := 0; i < spawned; i++ {
		s.LanternFish = append(s.LanternFish, NewLanternFish(8))
	}

}

func (s *Simulator) Run() {
	s.LogDay()
	for i := 0; i < DAYS; i++ {
		s.AgeFish()
		s.LogDay()
	}
}

func (s *Simulator) LogDay() {
	line := ""
	for _, fish := range s.LanternFish {
		line = fmt.Sprintf("%s %d", line, fish.Age)
	}
	
}

// cheated on this by looking at examples
func (s *Simulator) RunPart2() int {
	ages := strings.Split(s.InitialInput[0], ",")
	calendar := [9]int{}

	for _, age := range ages {
		ageInt, err := strconv.Atoi(age)
		if err != nil {
			panic(err)
		}

		calendar[ageInt]++
	}

	for day := 0; day < 256; day++ {
		calendar = [9]int{
			calendar[1], calendar[2], calendar[3], calendar[4], calendar[5],
			calendar[6], calendar[7] + calendar[0], calendar[8], calendar[0],
		}
	}

	sum := 0
	for _, count := range calendar {
		sum += count
	}

	return sum

}
