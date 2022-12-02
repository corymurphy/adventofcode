package main

import "strings"

type Round struct {
	opponent string
	you      string
	expected string
	rock     map[string]int
	paper    map[string]int
	scissors map[string]int
}

func NewRound(input string) *Round {
	plays := strings.Split(input, " ")
	return &Round{
		opponent: plays[0],
		you:      plays[1],
		expected: plays[1],
		rock: map[string]int{
			lose: 3,
			win:  8,
			draw: 4,
		},
		paper: map[string]int{
			lose: 1,
			win:  9,
			draw: 5,
		},
		scissors: map[string]int{
			lose: 2,
			win:  7,
			draw: 6,
		},
	}
}

func (r *Round) Play() int {

	if r.opponent == rock && r.you == rock_response {
		return 3 + rock_bonus
	} else if r.opponent == rock && r.you == paper_response {
		return 6 + paper_bonus
	} else if r.opponent == rock && r.you == scissors_response {
		return 0 + scissors_bonus
	} else if r.opponent == paper && r.you == rock_response {
		return 0 + rock_bonus
	} else if r.opponent == paper && r.you == paper_response {
		return 3 + paper_bonus
	} else if r.opponent == paper && r.you == scissors_response {
		return 6 + scissors_bonus
	} else if r.opponent == scissors && r.you == rock_response {
		return 6 + rock_bonus
	} else if r.opponent == scissors && r.you == paper_response {
		return 0 + paper_bonus
	} else if r.opponent == scissors && r.you == scissors_response {
		return 3 + scissors_bonus
	} else {
		return 0
	}

}

// i think I could do a tuple here, just lazy
func (r *Round) Play2() int {
	if r.opponent == rock {
		return r.rock[r.expected]
	} else if r.opponent == scissors {
		return r.scissors[r.expected]
	} else if r.opponent == paper {
		return r.paper[r.expected]
	}
	return 0
}
