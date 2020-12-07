package main

import (
	"fmt"
	"testing"
)

func Test_BoardingPass_SeatId(t *testing.T) {

	expected := 70
	actual := NewBoardingPass("BFFFBBFRRR").Row()

	if expected != actual {
		t.Errorf("Row expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}

	expected = 7
	actual = NewBoardingPass("BFFFBBFRRR").Column()

	if expected != actual {
		t.Errorf("Column expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}

	expected = 567
	actual = NewBoardingPass("BFFFBBFRRR").SeatId()

	if expected != actual {
		t.Errorf("SeatId expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}

	expected = 102
	actual = NewBoardingPass("BBFFBBFRLL").Row()

	if expected != actual {
		t.Errorf("Row expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}

	expected = 4
	actual = NewBoardingPass("BBFFBBFRLL").Column()

	if expected != actual {
		t.Errorf("Column expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}

	expected = 820
	actual = NewBoardingPass("BBFFBBFRLL").SeatId()

	if expected != actual {
		t.Errorf("SeatId expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}

	expected = 955
	actual = part1()
	if expected != actual {
		t.Errorf("SeatId expected %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}
