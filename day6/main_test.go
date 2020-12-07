package main

import (
	"fmt"
	"testing"
)

func Test_BoardingPass_SeatId(t *testing.T) {

	expected1 := 11
	expected2 := 6
	actual1, actual2 := YesSum([]string{"abc", "", "a", "b", "c", "", "ab", "ac", "", "a", "a", "a", "a", "", "b", ""})

	if expected1 != actual1 {
		t.Errorf("YesSum1 %s, got %s", fmt.Sprint(expected1), fmt.Sprint(actual1))
	}
	if expected2 != actual2 {
		t.Errorf("YesSum2 %s, got %s", fmt.Sprint(expected2), fmt.Sprint(actual2))
	}

	expected := 1
	actual := NewDeclaration([]string{"a", "a", "a", "a"}).YesCount1()

	if expected != actual {
		t.Errorf("aaaa YesCount1 %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}

	expected = 3
	actual = NewDeclaration([]string{"a", "b", "c", "a"}).YesCount1()

	if expected != actual {
		t.Errorf("abca YesCount1 %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}

	expected = 3
	actual = NewDeclaration([]string{"abc"}).YesCount2()

	if expected != actual {
		t.Errorf("abc YesCount2 %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}

	expected = 1
	actual = NewDeclaration([]string{"ab", "ac"}).YesCount2()

	if expected != actual {
		t.Errorf("ab ac YesCount2 %s, got %s", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}
