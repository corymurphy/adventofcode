package main

import "strings"

type Round struct {
	opponent string
	you      string
	expected string
	rock     map[string]int
	paper    map[string]int
	scissors map[string]int
	part2    map[string]map[string]int
}

func NewRound(input string) *Round {
	plays := strings.Split(input, " ")
	return &Round{
		opponent: plays[0],
		you:      plays[1],
		expected: plays[1],
		rock: map[string]int{
			should_lose: 3,
			should_win:  8,
			should_draw: 4,
		},
		paper: map[string]int{
			should_lose: 1,
			should_win:  9,
			should_draw: 5,
		},
		scissors: map[string]int{
			should_lose: 2,
			should_win:  7,
			should_draw: 6,
		},
		part2: map[string]map[string]int{
			rock: map[string]int{
				should_lose: rock_bonus + lose,
				should_win:  rock_bonus + win,
				should_draw: rock_bonus + draw,
			},
			paper: map[string]int{
				should_lose: paper_bonus + lose,
				should_win:  paper_bonus + win,
				should_draw: paper_bonus + draw,
			},
			scissors: map[string]int{
				should_lose: scissors_bonus + lose,
				should_win:  scissors_bonus + win,
				should_draw: scissors_bonus + draw,
			},
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
	// return r.part2[r.opponent][r.expected]
}
